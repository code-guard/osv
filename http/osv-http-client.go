package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OsvErrorReponseDetail struct {
	TypeUrl string `json:"typeUrl"`
	Value   string `json:"value"`
}

type OsvErrorReponse struct {
	Code    *int                    `json:"code"`
	Message string                  `json:"message"`
	Details []OsvErrorReponseDetail `json:"details"`
}

func (err OsvErrorReponse) Error() string {
	return fmt.Sprintf("the server replied with code %d: %s", *err.Code, err.Message)
}

type OsvHttpClient struct {
	client http.Client
}

func (client *OsvHttpClient) OsvHttpClient(httpClient http.Client) {
	client.client = httpClient
}

func (client *OsvHttpClient) Do(request *http.Request) (*http.Response, error) {
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.client.Do(request)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return resp, nil
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var errorResponse OsvErrorReponse
	json.Unmarshal(body, &errorResponse)

	if errorResponse.Code == nil {
		return nil, fmt.Errorf("an error occurred while reading OSV APIs error response")
	}

	return nil, errorResponse
}
