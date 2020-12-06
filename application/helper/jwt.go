package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//MyCustomClaims jwt claims struct
type MyCustomClaims struct {
	AccountID uint64
	Email     string
	Role      string
	jwt.StandardClaims
}

// GetExpiryTime for jwt
func GetExpiryTime(exp time.Duration) int64 {
	return time.Now().Add(time.Hour * exp).Unix()
}

// CreateClaims .
func CreateClaims(accountID uint64, email string, role string, exp time.Duration) MyCustomClaims {
	return MyCustomClaims{
		accountID,
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: GetExpiryTime(exp),
			IssuedAt:  time.Now().Unix(),
		},
	}
}
