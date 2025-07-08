package middlewares

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/pkg/errors"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))
var refreshTokenKey = []byte(os.Getenv("REFRESH_TOKEN"))

func GenerateJWT(userID string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"roles": roles,
		"typ":   "access",
		"exp":   time.Now().Add(5 * time.Minute).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
} //

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"typ": "refresh",
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshTokenKey)
} //

func ParseJWT(tokenStr string) (string, []string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		return "", nil, fmt.Errorf("invalid or expired token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", nil, errors.New("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return "", nil, errors.New("invalid 'sub' claim")
	}
	var roles []string
	if rolesClaim, ok := claims["roles"]; ok {
		if rolesSlice, ok := rolesClaim.([]interface{}); ok {
			for _, r := range rolesSlice {
				if roleStr, ok := r.(string); ok {
					roles = append(roles, roleStr)
				}
			}
		}
	}
	return sub, roles, nil
}
