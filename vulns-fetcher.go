package osv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Given a vulnerability query returns the vulnerabilities associated with the query
func QueryVulns(queryVuln *VulnQuery) (*VulnsResponse, error) {

	jsonData, err := json.Marshal(&queryVuln)
	if err != nil {
		return nil, fmt.Errorf("cannot marshall the VulnQuery struct")
	}

	req, err := http.NewRequest("POST", "https://api.osv.dev/v1/query", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to generate the HTTP request for %s, version %s", queryVuln.PckInfo.Name, queryVuln.Version)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query the dependency vulnerabilities for %s, version %s", queryVuln.PckInfo.Name, queryVuln.Version)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read OSV response for %s, version %s", queryVuln.PckInfo.Name, queryVuln.Version)
	}

	var vulnResponse VulnsResponse

	json.Unmarshal(body, &vulnResponse)

	return &vulnResponse, nil

}

// Given vulnerability queries returns the vulnerabilities associated with the queries
func BulkQueryVulns(queriesVuln *VulnsQueries) (*BulkVulnsResponse, error) {

	jsonData, err := json.Marshal(&queriesVuln)

	if err != nil {
		return nil, fmt.Errorf("cannot marshall the VulnQuery struct")
	}

	req, err := http.NewRequest("POST", "https://api.osv.dev/v1/querybatch", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to generate the HTTP request for bulk query")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query the dependency vulnerabilities for bulk query")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read OSV response for bulk query")
	}

	var bullVulnResponse BulkVulnsResponse

	json.Unmarshal(body, &bullVulnResponse)

	return &bullVulnResponse, nil

}
