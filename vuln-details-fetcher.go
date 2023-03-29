package osv

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Given a vulnerability OSV ID return the vulnerability details
func GetVulnDetails(vulnId string) (*VulnDetails, error) {

	urlRequest := fmt.Sprintf("https://api.osv.dev/v1/vulns/%s", vulnId)

	resp, err := http.Get(urlRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to query the dependency vulnerabilities for vulnerability %s", vulnId)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read OSV response for vulnerability %s", vulnId)
	}

	var vulnDetailsResponse VulnDetails

	json.Unmarshal(body, &vulnDetailsResponse)

	return &vulnDetailsResponse, nil

}
