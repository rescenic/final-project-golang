// models/auth.go

package models

import "github.com/dgrijalva/jwt-go"

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token  string `json:"token"`
	Role   string `json:"role"`
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"nama_lengkap"`
	Active bool   `json:"active"`
	NoRM   string `json:"no_rm,omitempty"` // Only for patients
}

// JWTClaim represents the structure of our JWT claims
type JWTClaim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
