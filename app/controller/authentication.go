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

func Register(context *gin.Context) {
	var input model.AuthRegister

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get admin access token for register user
	resp := adminLogin()
	respJson := handleRespAdminLogin(resp)

	// register user into keycloak
	registerInKC(respJson.Access_token, input, context)

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
	apiUrl := getAdminLoginURL()
	apiBody := getAdminLoginBody()

	response, err := helper.MakePostReqXWWWForm(apiUrl, apiBody)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic(response)
	}

	responseBody, _ := ioutil.ReadAll(response.Body)

	return string(responseBody)
}

func getAdminLoginURL() string {
	KeycloakHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("get_token_path")

	u, _ := url.ParseRequestURI(KeycloakHost)
	u.Path = apiPath
	return u.String()
}

func getAdminLoginBody() string {
	clientId := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")
	grantType := os.Getenv("grant_type")

	requestBody := url.Values{}
	requestBody.Set("client_id", clientId)
	requestBody.Set("client_secret", clientSecret)
	requestBody.Set("grant_type", grantType)

	return requestBody.Encode()
}

func handleRespAdminLogin(resp string) helper.AdminLoginResponse {
	var respJson helper.AdminLoginResponse
	err := json.Unmarshal([]byte(resp), &respJson)
	if err != nil {
		panic(err)
	}

	return respJson
}

func registerInKC(accessToken string, input model.AuthRegister, context *gin.Context) error {
	apiUrl := getRegisterURL()
	apiBody := getRegisterBody(input)

	response, err := helper.MakePostReqJsonAuth(apiUrl, apiBody, accessToken)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusCreated {
		//success
		return nil
	}

	// error message response from keycloak
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	errRespJson := handleRespRegisterInKC(string(responseBody))
	context.JSON(response.StatusCode, gin.H{"message": errRespJson.Message})
	panic(errRespJson)
}

func getRegisterURL() string {
	KeycloakHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("register_user_path")

	u, _ := url.ParseRequestURI(KeycloakHost)
	u.Path = apiPath

	return u.String()
}

func getRegisterBody(input model.AuthRegister) string {
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

func handleRespRegisterInKC(resp string) helper.ErrorResponseKC {
	var errRespJson helper.ErrorResponseKC
	err := json.Unmarshal([]byte(resp), &errRespJson)
	if err != nil {
		panic(err)
	}
	return errRespJson
}
