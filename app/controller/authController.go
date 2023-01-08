package controller

import (
	"encoding/json"
	"io/ioutil"
	"keycloak_api_go/app/helper"
	"keycloak_api_go/app/model"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input model.AuthRegister

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get admin access token for register user
	resp := userLogin(context, true, nil)
	respJson := handleRespUserLogin(resp) // error handling itself

	// register user into keycloak
	registerInKC(respJson.Access_token, input, context) // error handling itself

	context.JSON(http.StatusCreated, gin.H{"message": "user created!"})
}

func Login(context *gin.Context) {
	var input model.AuthLogin

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// login user
	resp := userLogin(context, false, &input)
	respJson := handleRespUserLogin(resp) // error handling itself
	respMap := structs.Map(&respJson)

	context.JSON(http.StatusOK, respMap)
}

func userLogin(context *gin.Context, isAdmin bool, loginInfo *model.AuthLogin) string {
	apiUrl := getUserLoginURL()
	var apiBody string
	if isAdmin {
		apiBody = getAdminLoginBody()
	} else {
		apiBody = getUserLoginBody(loginInfo)
	}

	response, err := helper.MakePostReqXWWWForm(apiUrl, apiBody)
	if err != nil {
		panic(err)
	}

	responseBody, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		errRespJson := handleRespRegisterInKC(string(responseBody))
		context.JSON(response.StatusCode, gin.H{"message": errRespJson.Message})
		context.Abort()
		panic(response)
	}

	return string(responseBody)
}

func getUserLoginURL() string {
	KCKHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("user_login_path")

	u, _ := url.ParseRequestURI(KCKHost)
	u.Path = apiPath
	return u.String()
}

func getAdminLoginBody() string {
	clientId := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")
	grantType := "client_credentials"

	requestBody := url.Values{}
	requestBody.Set("client_id", clientId)
	requestBody.Set("client_secret", clientSecret)
	requestBody.Set("grant_type", grantType)

	return requestBody.Encode()
}

func getUserLoginBody(loginInfo *model.AuthLogin) string {
	clientId := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")
	grantType := "password"

	requestBody := url.Values{}
	requestBody.Set("client_id", clientId)
	requestBody.Set("client_secret", clientSecret)
	requestBody.Set("grant_type", grantType)
	requestBody.Set("username", loginInfo.Username)
	requestBody.Set("password", loginInfo.Password)

	return requestBody.Encode()
}

func handleRespUserLogin(resp string) helper.AdminLoginResponse {
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
	context.Abort()
	panic(errRespJson)
}

func getRegisterURL() string {
	KCKHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("user_register_path")

	u, _ := url.ParseRequestURI(KCKHost)
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
