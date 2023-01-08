package helper

import (
	"net/http"
	"strings"
)

func MakePostReqXWWWForm(url string, requestBody string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, string(url), strings.NewReader(requestBody))
	if err != nil {
		return &http.Response{}, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &http.Response{}, err
	}

	return response, err
}

func MakePostReqJsonAuth(url string, requestBody string, accessToken string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, string(url), strings.NewReader(requestBody))
	if err != nil {
		return &http.Response{}, err
	}

	request.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		request.Header.Add("Authorization", "Bearer "+accessToken)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &http.Response{}, err
	}

	return response, err
}
