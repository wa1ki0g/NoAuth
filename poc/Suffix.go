package poc

func GenerateURLs(baseURL string) []string {
	suffixes := []string{
		";.rand",
		".rand",
		";.js",
		";.css",
		";.html",
		";.tmpl",
		".json",
		".js",
		".css",
		".html",
		".tmpl",
		"/",
		"/%20/",
		"/%3b",
		"/;",
		"..;/",
		"/12123123123123.jsp",
		";/12123123123123.jsp",
	}

	urls := make([]string, len(suffixes))

	for i, suffix := range suffixes {
		urls[i] = baseURL + suffix
	}

	return urls
}
