package helper

import (
	"net/http"
	"strings"
)

func CallRequest(apiUrl string, requestBody string, contentType string, authorization string) (*http.Response, error) {
	if contentType == "" {
		contentType = "application/json"
	}

	request, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(requestBody))
	if err != nil {
		return &http.Response{}, err
	}
	request.Header.Add("Content-Type", contentType)
	if authorization != "" {
		request.Header.Add("Authorization", "Bearer "+authorization)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &http.Response{}, err
	}

	return response, err
}
