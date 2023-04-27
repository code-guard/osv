package osv

type VulnerabilityQuery struct {
	Commit  string      `json:"commit,omitempty"`
	Version string      `json:"version,omitempty"`
	PckInfo PackageInfo `json:"package,omitempty"`
}

type PackageInfo struct {
	Name      string `json:"name,omitempty"`
	Ecosystem string `json:"ecosystem,omitempty"`
	Purl      string `json:"purl,omitempty"`
}

type VulnerabilityQueries struct {
	Queries []VulnerabilityQuery `json:"queries,omitempty"`
}
