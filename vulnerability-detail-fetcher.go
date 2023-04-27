package osv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	osvHttp "osv/http"
)

// Given a vulnerability OSV ID return the vulnerability details
func GetVulnerability(client *osvHttp.OsvHttpClient, id string) (*VulnerabilityDetails, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.osv.dev/v1/vulns/%s", id), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var vulnerabilityDetails VulnerabilityDetails
	json.Unmarshal(responseContent, &vulnerabilityDetails)

	return &vulnerabilityDetails, nil
}
