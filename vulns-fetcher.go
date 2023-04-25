package osv

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	osvHttp "osv/http"
)

// Given a vulnerability query returns the vulnerabilities associated with the query
func QueryVulnerabilities(client *osvHttp.OsvHttpClient, query *VulnerabilityQuery) (*VulnerabilityResponses, error) {
	jsonData, err := json.Marshal(&query)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://api.osv.dev/v1/query", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var vulnerabilityResponses VulnerabilityResponses
	json.Unmarshal(body, &vulnerabilityResponses)

	return &vulnerabilityResponses, nil
}

// Given vulnerability queries returns the vulnerabilities associated with the queries
func BulkQueryVulnerabilities(client *osvHttp.OsvHttpClient, vulnerabilityQueries *VulnerabilityQueries) (*BulkVulnerabilityResponses, error) {
	jsonData, err := json.Marshal(&vulnerabilityQueries)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://api.osv.dev/v1/querybatch", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var bulkVulnerabilityResponses BulkVulnerabilityResponses
	json.Unmarshal(body, &bulkVulnerabilityResponses)

	return &bulkVulnerabilityResponses, nil

}
