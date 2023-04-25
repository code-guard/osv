package osv

import "osv/http"

func main() {

	pckInfo := PackageInfo{Name: "org.apache.logging.log4j:log4j", Ecosystem: "Maven"}
	vulnQuery := VulnQuery{Version: "2.1", PckInfo: pckInfo}

	myVulns, _ := QueryVulns(&vulnQuery)

	GetVulnerability(&http.OsvHttpClient{}, myVulns.Vulns[0].ID)

	q1 := VulnQuery{PckInfo: PackageInfo{Purl: "pkg:pypi/antlr4-python3-runtime@4.7.2"}}
	q2 := VulnQuery{Commit: "6879efc2c1596d11a6a6ad296f80063b558d5e0f"}

	queries := VulnsQueries{VulnsQuery: []VulnQuery{q1, q2}}

	BulkQueryVulns(&queries)

}
