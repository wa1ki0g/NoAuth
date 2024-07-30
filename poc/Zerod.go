package poc

import (
	"strings"
)

func removeTrailingSlash(url string) string {
	if strings.HasSuffix(url, "/") {
		return url[:strings.LastIndex(url, "/")]
	}
	return url
}

func ConvertURL(baseURL string) []string {
	var results []string

	insertPosition := strings.LastIndex(baseURL, "/") + 1

	encodingChars := []string{"%0d", "%0a", ".", ";", "%3b", "%20", "%2e", "%2e%2e", "%25%32%66", "%2f"}

	for _, char := range encodingChars {
		if char == "%25%32%66" {
			convertedURL := removeTrailingSlash(baseURL[:insertPosition]) + char + baseURL[insertPosition:]
			results = append(results, convertedURL)
			results = append(results, "/"+convertedURL)
			convertedURL = convertedURL + ";"
			results = append(results, convertedURL)
			convertedURL = "/" + convertedURL
			results = append(results, convertedURL)
			continue
		}
		if char == "%2f" {
			convertedURL := removeTrailingSlash(baseURL[:insertPosition]) + char + baseURL[insertPosition:]
			results = append(results, convertedURL)
			results = append(results, "/"+convertedURL)
			convertedURL = convertedURL + ";"
			results = append(results, convertedURL)
			convertedURL = "/" + convertedURL
			results = append(results, convertedURL)
			continue
		}

		convertedURL := baseURL[:insertPosition] + char + baseURL[insertPosition:]
		results = append(results, convertedURL)
	}

	return results
}
