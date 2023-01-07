package helper

import (
	"keycloak_api_go/app/model"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"issued_at": time.Now().Unix(),
		"expire_at": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

// func CurrentUser() (model.User, error) {
// 	err := ValidateJWT(context)
// 	if err != nil {
// 		return model.User{}, err
// 	}

// 	token, _ := getToken(context)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userID := uint(claims["id"].(float64))

// 	user, err := model.FindUserByID(userID)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return user, nil
// }
