package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/golang-jwt/jwt"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

func JwtToken(payload types.JwtPayload) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	jwt_expiration := time.Second * time.Duration(Jwt_Expiration_Int(os.Getenv("JWT_EXPIRATION")))
	userBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      string(userBytes),
		"expiredAt": time.Now().Add(jwt_expiration).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Print("here...")
		return "", err
	}
	return tokenString, nil
}
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
func Jwt_Expiration_Int(jwt_expiration_string string) int {

	duration, err := strconv.ParseInt(jwt_expiration_string, 10, 64)
	if err != nil {
		fmt.Println("error into int", err)
		return 0
	}
	return int(duration)

}
