package controller

import (
	"encoding/json"
	"io/ioutil"
	"keycloak_api_go/app/helper"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func Info(context *gin.Context) {
	apiUrl := getInfoURL()
	accessToken := helper.GetTokenFromRequest(context)

	response, err := helper.MakePostReqJsonAuth(apiUrl, "", accessToken)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		// context.Abort()
		panic(err)
	}

	responseBody, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {

		var respJson helper.ErrorResponseKC
		err := json.Unmarshal(responseBody, &respJson)
		if err != nil {
			panic(err)
		}
		respMap := structs.Map(&respJson)
		context.JSON(response.StatusCode, respMap)
		// context.Abort()
		panic(response)
	}

	var respJson helper.UserInfoResponse
	err = json.Unmarshal(responseBody, &respJson)
	respMap := structs.Map(&respJson)
	if err != nil {
		context.JSON(http.StatusBadRequest, respMap)
		// fmt.Println("err1", err)
		// context.Abort()
		panic(err)
	}

	context.JSON(response.StatusCode, respMap)
}

func getInfoURL() string {
	KCKHost := os.Getenv("keycloak_host")
	apiPath := os.Getenv("user_info_path")

	u, _ := url.ParseRequestURI(KCKHost)
	u.Path = apiPath

	return u.String()
}
