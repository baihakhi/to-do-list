package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/baihakhi/to-do-list/internal/models"
	"github.com/baihakhi/to-do-list/internal/models/enum"
	"github.com/golang-jwt/jwt"
)

// CreateToken use jwt to generate token when user login
func CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{}

	claims["account_id"] = user.Code
	claims["account_group"] = user.UserGroup
	claims["username"] = user.Name
	claims["email"] = user.Email

	if claims["account_group"] == enum.RoleSuperAdmin {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 36).Unix()
	} else if claims["account_group"] == enum.RoleAdmin {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 12).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
}

// TokenValid validate token from http request
func TokenValid(r *http.Request) (*models.User, error) {
	user := new(models.User)
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err = CreateUserFromMap(claims)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

// ExtractToken extract body token to get information
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// CreateFromMap function for convert map to user struct
func CreateUserFromMap(m map[string]interface{}) (*models.User, error) {
	data, _ := json.Marshal(m)
	var result = new(models.User)
	err := json.Unmarshal(data, &result)
	return result, err
}
