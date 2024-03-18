package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hculpan/go-web-base-project-1/internal/db"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET not defined")
	}

	jwtKey = []byte(secret) // Replace this with your secret key from .env
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token for a given username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func IsAuthorized(r *http.Request) bool {
	c, err := r.Cookie("token")
	if err != nil {
		return false
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}

	if !tkn.Valid {
		return false
	}

	return true
}

// This is a placeholder for the actual login logic.
// You need to replace it with your own login validation logic.
func ValidateUser(username, password string) bool {
	if username == "" || password == "" {
		return false
	}

	hashedPassword, err := db.RetrieveUserHash(username)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
