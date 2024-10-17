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
	switch req.Role {
	case "admin":
		return s.loginAdmin(req)
	case "dokter":
		return s.loginDokter(req)
	case "pasien":
		return s.loginPasien(req)
	default:
		return nil, errors.New("invalid role")
	}
}

func (s *AuthService) loginAdmin(req *models.LoginRequest) (*models.LoginResponse, error) {
	var admin models.Admin
	if err := s.db.Where("email = ? AND active = ?", req.Email, true).First(&admin).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(admin.ID, admin.Email, "admin")
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		Role:  "admin",
		ID:    admin.ID,
		Email: admin.Email,
	}, nil
}

func (s *AuthService) loginDokter(req *models.LoginRequest) (*models.LoginResponse, error) {
	var dokter models.Dokter
	if err := s.db.Where("email = ? AND active = ?", req.Email, true).First(&dokter).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dokter.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(dokter.ID, dokter.Email, "dokter")
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		Role:  "dokter",
		ID:    dokter.ID,
		Email: dokter.Email,
	}, nil
}

func (s *AuthService) loginPasien(req *models.LoginRequest) (*models.LoginResponse, error) {
	var pasien models.Pasien
	if err := s.db.Where("email = ? AND active = ?", req.Email, true).First(&pasien).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pasien.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(pasien.ID, pasien.Email, "pasien")
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		Role:  "pasien",
		ID:    pasien.ID,
		Email: pasien.Email,
	}, nil
}

func (s *AuthService) generateToken(userID uint, email, role string) (string, error) {
	claims := &models.JWTClaim{
		UserID: userID,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
