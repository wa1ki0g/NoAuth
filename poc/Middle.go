package poc

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func GenerateVariants(suffix string) []string {
	var results []string

	encodingChars := []string{
		"%25%32%66",
		"%2f",
		"%0a",
		"%0d",
	}

	mid := len(suffix) / 2

	for _, char := range encodingChars {
		convertedURL := suffix[:mid] + char + suffix[mid:]
		results = append(results, convertedURL)
		if char == "%2f" {
			results = append(results, convertedURL+";")
		}
	}

	suffixAsRunes := []rune(suffix)
	for i := 0; i < len(suffixAsRunes); i++ {
		rand.Seed(time.Now().UnixNano())
		if unicode.IsLetter(suffixAsRunes[i]) {
			if rand.Intn(2) == 0 {
				upperChar := unicode.ToUpper(suffixAsRunes[i])
				convertedURL := suffix[:i] + string(upperChar) + suffix[i+1:]
				results = append(results, convertedURL)
			} else {
				lowerChar := unicode.ToLower(suffixAsRunes[i])
				convertedURL := suffix[:i] + string(lowerChar) + suffix[i+1:]
				results = append(results, convertedURL)
			}
		}
	}

	results = append(results, suffix)

	return results
}

func ExtractAndModifyURL(baseURL string) []string {
	insertPosition := strings.LastIndex(baseURL, "/")
	if insertPosition == -1 {
		return []string{baseURL}
	}

	suffix := baseURL[insertPosition+1:]

	variants := GenerateVariants(suffix)

	var results []string
	for _, variant := range variants {
		results = append(results, baseURL[:insertPosition+1]+variant)
	}

	return results
}
