package ddgsearch

// testEndpoints contains test URLs for DuckDuckGo API endpoints
var testEndpoints = map[string]string{}

// setTestEndpoints sets custom endpoints for testing
func setTestEndpoints(vqdURL, searchURL string) {
	testEndpoints["vqd"] = vqdURL
	testEndpoints["search"] = searchURL
}
