package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) == 2 { // 1: bearer 2: token
		return splitToken[1]
	}
	return ""
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
