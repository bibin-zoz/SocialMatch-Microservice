package helper

import (
	"fmt"
	"regexp"

	"github.com/golang-jwt/jwt"
)

func GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}

	return header
}
func ExtractUserIDFromToken(tokenString string) (int, string, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	return claims.Id, claims.Email, nil

}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@gmail.com$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
func ValidatePassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	numberRegex := regexp.MustCompile(`[0-9]`)
	if !lowercaseRegex.MatchString(password) || !uppercaseRegex.MatchString(password) || !numberRegex.MatchString(password) {
		return false
	}
	return true
}
