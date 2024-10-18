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

	// Generate token and pass the role
	token, err := s.generateToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:  token,
		ID:     admin.ID,
		Email:  admin.Email,
		Name:   admin.NamaLengkap,
		Active: admin.Active,
		Role:   admin.Role,
	}, nil
}

func (s *AuthService) generateToken(userID uint, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.JWTClaim{
		UserID: userID,
		Email:  email,
		Role:   role, // Set the role dynamically
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
