package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

func GenerateJWTToken(id string, password string) (string, error) {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("error opening env file: ", envErr)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = jwt.MapClaims{
		"sub":  id,
		"name": password,
		"iat":  time.Now().Unix(),                     // Token issue date
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	signed, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ValidateJWTToken(token string) (bool, error) {
	parsed, err := jwt.Parse(token, func(parsed *jwt.Token) (interface{}, error) {
		if _, ok := parsed.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsed.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {

		// Validate token expiration date
		if !parsed.Valid {
			fmt.Printf("exp: %v", int64(claims["exp"].(float64)))
			return false, fmt.Errorf("token expiration error")
		}

		// Validate username and password
		if claims["sub"] == os.Getenv("ADMIN_USER_NAME") && claims["name"] == os.Getenv("ADMIN_USER_PASS") {
			return true, nil
		}
	}
	return false, fmt.Errorf("not a valid token")
}

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	},
})

// type Manager struct {
// cookieName string
// }
// func ValidateCookie() {}
