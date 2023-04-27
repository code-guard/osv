package osv

type Event struct {
	Introduced   string `json:"introduced"`
	Fixed        string `json:"fixed"`
	LastAffected string `json:"last_affected"`
	Limit        string `json:"limit"`
}

type Range struct {
	Type             string   `json:"type"`
	Repo             string   `json:"repo"`
	Events           []Event  `json:"events"`
	DatabaseSpecific struct{} `json:"database_specific"`
}

type Package struct {
	Ecosystem string `json:"ecosystem"`
	Name      string `json:"name"`
	Purl      string `json:"purl"`
}

type Affected struct {
	Package  Package `json:"package"`
	Severity []struct {
		Type  string `json:"type"`
		Score string `json:"score"`
	} `json:"severity"`
	Ranges            []Range  `json:"ranges"`
	Versions          []string `json:"versions"`
	EcosystemSpecific struct{} `json:"ecosystem_specific"`
	DatabaseSpecific  struct{} `json:"database_specific"`
}

type Reference struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Credit struct {
	Name    string   `json:"name"`
	Contact []string `json:"contact"`
	Type    []string `json:"type"`
}

type Severity struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

type DatabaseSpecific struct{}

type VulnerabilityDetails struct {
	SchemaVersion    string           `json:"schema_version"`
	ID               string           `json:"id"`
	Modified         string           `json:"modified"`
	Published        string           `json:"published"`
	Withdrawn        string           `json:"withdrawn"`
	Aliases          []string         `json:"aliases"`
	Related          []string         `json:"related"`
	Summary          string           `json:"summary"`
	Details          string           `json:"details"`
	Severity         []Severity       `json:"severity"`
	Affected         []Affected       `json:"affected"`
	References       []Reference      `json:"references"`
	Credits          []Credit         `json:"credits"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific"`
}
