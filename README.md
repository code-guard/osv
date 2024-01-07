A super lightweight Go library to query the osv.dev database.

The library aims to be very easy to use, a query can be performed as shown below (for a complete example check out `examples/osv.go`):
```
pckInfo := osv.PackageInfo{Name: "org.apache.logging.log4j:log4j", Ecosystem: "Maven"}
vulnerabilityQuery := osv.VulnerabilityQuery{Version: "2.1", PckInfo: pckInfo}

osvClient := osvClient.OsvHttpClient{}
osvClient.OsvHttpClient(http.Client{})
myVulns, _ := osv.QueryVulnerabilities(&osvClient, &vulnerabilityQuery)

response, _ := osv.GetVulnerability(&osvClient, myVulns.Vulns[0].Id)
fmt.Println(response)

q1 := osv.VulnerabilityQuery{PckInfo: osv.PackageInfo{Purl: "pkg:pypi/antlr4-python3-runtime@4.7.2"}}
q2 := osv.VulnerabilityQuery{Commit: "6879efc2c1596d11a6a6ad296f80063b558d5e0f"}

queries := osv.VulnerabilityQueries{Queries: []osv.VulnerabilityQuery{q1, q2}}

response2, _ := osv.BulkQueryVulnerabilities(&osvClient, &queries)
fmt.Println(response2)
```
