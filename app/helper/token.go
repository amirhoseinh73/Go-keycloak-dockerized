package helper

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
