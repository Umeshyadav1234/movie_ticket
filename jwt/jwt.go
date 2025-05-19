package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
var RolePrivileges = map[string][]string{
    "admin": {
        "create_movie",
		"get_movies",
        "add_schedule",
    },
    "customer": {
        "view_movies",
        "view_schedules",
		"view_seatsbymovie",
        "book_ticket",
		"get_report",
    },
}


var SecretKey = []byte("MTBS") // Replace with env variable in production

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Privileges []string `json:"privileges"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT token with email and role
func GenerateToken(email string, role string) (string, error) {
	privileges := RolePrivileges[role]
	claims := Claims{
		Email: email,
		Role:  role,
		Privileges: privileges,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// ValidateToken parses the token and returns claims
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
