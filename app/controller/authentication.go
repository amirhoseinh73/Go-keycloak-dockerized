package controller

import (
	"encoding/json"
	"io/ioutil"
	"keycloak_api_go/app/helper"
	"keycloak_api_go/app/model"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type adminLoginResponse struct {
	Access_token       string `json:"access_token"`
	Expires_in         int    `json:"expires_in"`
	Refresh_expires_in int    `json:"refresh_expires_in"`
	Token_type         string `json:"token_type"`
	NotBeforePolicy    int    `json:"not-before-policy"`
	Scope              string `json:"scope"`
}

func Register(context *gin.Context) {
	var input model.AuthRegister

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get admin access token for register user
	adminLoginResp := adminLogin()

	var adminLoginJson adminLoginResponse
	err := json.Unmarshal([]byte(adminLoginResp), &adminLoginJson)
	if err != nil {
		panic(err)
	}

	accessToken := adminLoginJson.Access_token

	// register user into keycloak
	_ = registerIntoKeyclock(accessToken, input)

	context.JSON(http.StatusCreated, gin.H{"message": "user created!"})
}

func Login(context *gin.Context) {
	// var input model.Authentication

	// if err := context.ShouldBindJSON(&input); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{
	// 		"error json": err.Error(),
	// 		"data":       &input,
	// 	})
	// 	return
	// }

	// user, err := model.FindUserByUsername(input.Username)

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error username": err.Error()})
	// 	return
	// }

	// err = user.ValidatePassword(input.Password)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error password": err.Error()})
	// 	return
	// }

	// jwt, err := helper.GenerateJWT(user)
	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error jwt": err.Error()})
	// 	return
	// }

	context.JSON(http.StatusOK, gin.H{"jwt": "jwt"})
}

func adminLogin() string {
	KeycloakHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("get_token_path")
	contentType := "application/x-www-form-urlencoded"

	u, _ := url.ParseRequestURI(KeycloakHost)
	u.Path = apiPath
	apiUrl := u.String()

	response, err := helper.CallRequest(apiUrl, getAdminLoginBody(), contentType, "")
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic(response)
	}

	responseBody, _ := ioutil.ReadAll(response.Body)

	return string(responseBody)
}

func getAdminLoginBody() string {
	client_id := os.Getenv("client_id")
	client_secret := os.Getenv("client_secret")
	grant_type := os.Getenv("grant_type")

	requestBody := url.Values{}
	requestBody.Set("client_id", client_id)
	requestBody.Set("client_secret", client_secret)
	requestBody.Set("grant_type", grant_type)

	return requestBody.Encode()
}

func registerIntoKeyclock(accessToken string, input model.AuthRegister) error {
	KeycloakHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("register_user_path")

	u, _ := url.ParseRequestURI(KeycloakHost)
	u.Path = apiPath
	apiUrl := u.String()

	response, err := helper.CallRequest(apiUrl, getRegisterIntoKeycloakBody(input), "", accessToken)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusCreated {
		panic(response)
	}

	return nil
}

func getRegisterIntoKeycloakBody(input model.AuthRegister) string {
	credential := map[string]any{
		"type":      "password",
		"value":     input.Password,
		"temporary": false,
	}
	credentials := make([]map[string]any, 1)
	credentials[0] = credential

	requestBody := map[string]any{
		"username":  input.Username,
		"firstName": input.Firstname,
		"lastName":  input.Lastname,
		"email":     input.Email,
		"enabled":   true,

		"credentials": credentials,
	}

	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	return string(bodyJson)
}
