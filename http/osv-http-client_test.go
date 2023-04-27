package http_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	osvHttp "osv/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SuccessRoundTripFunc func(req *http.Request) *http.Response

func (f SuccessRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
func SuccessNewTestClient(fn SuccessRoundTripFunc) *http.Client {
	return &http.Client{
		Transport: SuccessRoundTripFunc(fn),
	}
}

var errStandard = errors.New("Bla")

type FailureRoundTripFunc func(req *http.Request) *http.Response

func (f FailureRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errStandard
}
func FailureNewTestClient(fn FailureRoundTripFunc) *http.Client {
	return &http.Client{
		Transport: FailureRoundTripFunc(fn),
	}
}

func TestOsvClientHttpClientErrorResponse(t *testing.T) {
	client := FailureNewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://api.osv.dev")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	osvClient := osvHttp.OsvHttpClient{}
	osvClient.OsvHttpClient(*client)

	request, _ := http.NewRequest("GET", "https://api.osv.dev", nil)
	_, err := osvClient.Do(request)

	assert.ErrorIs(t, err, errStandard)
}

func TestOsvClientHttpClientSuccessResponse(t *testing.T) {
	client := SuccessNewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://api.osv.dev")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	osvClient := osvHttp.OsvHttpClient{}
	osvClient.OsvHttpClient(*client)

	request, _ := http.NewRequest("GET", "https://api.osv.dev", nil)
	response, _ := osvClient.Do(request)

	responseBody := make([]byte, 2)
	response.Body.Read(responseBody)

	assert.Equal(t, []byte("OK"), responseBody)
}

func TestOsvClientOsvErrorResponseWithInvalidFormat(t *testing.T) {
	client := SuccessNewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://api.osv.dev")
		return &http.Response{
			StatusCode: 500,
			Body:       io.NopCloser(bytes.NewBufferString(`{"message": "An Error occured", "details": []}`)),
			Header:     make(http.Header),
		}
	})

	osvClient := osvHttp.OsvHttpClient{}
	osvClient.OsvHttpClient(*client)

	request, _ := http.NewRequest("GET", "https://api.osv.dev", nil)
	_, error := osvClient.Do(request)

	assert.Equal(t, "an error occurred while reading OSV APIs error response", error.Error())
}

func TestOsvClientOsvErrorResponse(t *testing.T) {
	code := 321
	message := "An Error occured"
	client := SuccessNewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://api.osv.dev")
		return &http.Response{
			StatusCode: 500,
			Body:       io.NopCloser(bytes.NewBufferString(fmt.Sprintf(`{"code": %d, "message": "%s", "details": []}`, code, message))),
			Header:     make(http.Header),
		}
	})

	osvClient := osvHttp.OsvHttpClient{}
	osvClient.OsvHttpClient(*client)

	request, _ := http.NewRequest("GET", "https://api.osv.dev", nil)
	_, error := osvClient.Do(request)

	assert.Equal(t, fmt.Sprintf("the server replied with code %d: %s", code, message), error.Error())
}
