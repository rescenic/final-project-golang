// services/auth_service.go

package services

import (
	"errors"
	"gumuruh-clinic/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *AuthService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	var admin models.Admin

	// Find admin by email
	result := s.db.Where("email = ?", req.Email).First(&admin)
	if result.Error != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if account is active
	if !admin.Active {
		return nil, errors.New("account is inactive")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate token, passing "nama_lengkap"
	token, err := s.generateToken(admin.ID, admin.Email, admin.Role, admin.NamaLengkap)
	if err != nil {
		return nil, err
	}

	// Return the login response with the token
	return &models.LoginResponse{
		Token:  token,
		ID:     admin.ID,
		Email:  admin.Email,
		Name:   admin.NamaLengkap,
		Active: admin.Active,
		Role:   admin.Role,
	}, nil
}

func (s *AuthService) generateToken(userID uint, email, role string, namaLengkap string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	// Include "nama_lengkap" in the claims
	claims := jwt.MapClaims{
		"user_id":      userID,
		"email":        email,
		"role":         role,
		"nama_lengkap": namaLengkap,
		"exp":          expirationTime.Unix(),
		"iat":          time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
