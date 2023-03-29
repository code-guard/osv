package osv

import "time"

type VulnResp struct {
	SchemaVersion string    `json:"schemaVersion"`
	ID            string    `json:"id"`
	Published     time.Time `json:"published"`
	Modified      time.Time `json:"modified"`
	Withdrawn     time.Time `json:"withdrawn"`
	Aliases       []string  `json:"aliases"`
	Related       []string  `json:"related"`
	Summary       string    `json:"summary"`
	Details       string    `json:"details"`
	Affected      []struct {
		Package struct {
			Name      string `json:"name"`
			Ecosystem string `json:"ecosystem"`
			Purl      string `json:"purl"`
		} `json:"package"`
		Ranges []struct {
			Type   string `json:"type"`
			Repo   string `json:"repo"`
			Events []struct {
				Introduced string `json:"introduced"`
				Fixed      string `json:"fixed"`
				Limit      string `json:"limit"`
			} `json:"events"`
		} `json:"ranges"`
		Versions          []string               `json:"versions"`
		EcosystemSpecific map[string]interface{} `json:"ecosystemSpecific"`
		DatabaseSpecific  map[string]interface{} `json:"databaseSpecific"`
	} `json:"affected"`
	References []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"references"`
	Severity []struct {
		Type  string `json:"type"`
		Score string `json:"score"`
	} `json:"severity"`
	Credits []struct {
		Name    string   `json:"name"`
		Contact []string `json:"contact"`
	} `json:"credits"`
	DatabaseSpecific map[string]interface{} `json:"database_specific"`
}

type VulnsResponse struct {
	Vulns []VulnResp `json:"vulns"`
}

type BulkVulnsResponse struct {
	Results []VulnsResponse `json:"results"`
}
