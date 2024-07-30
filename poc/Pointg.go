package poc

import (
	"strings"
)

func InsertDots(path string) []string {
	parts := strings.Split(path, "/")
	var results []string
	var generate func(currentPath string, index int)
	generate = func(currentPath string, index int) {
		if index == len(parts) {
			if currentPath != "" {
				results = append(results, currentPath)
			}
			return
		}

		part := parts[index]
		if part != "" {
			newPathWithoutDot := currentPath + part
			if index < len(parts)-1 {
				newPathWithoutDot += "/"
			}
			generate(newPathWithoutDot, index+1)
			if currentPath != "" {
				newPathWithDot := currentPath + "./" + part
				if index < len(parts)-1 {
					newPathWithDot += "/"
				}
				generate(newPathWithDot, index+1)
			}
		} else if currentPath != "" {
			generate(currentPath, index+1)
		}
	}

	generate("/", 0)

	return results
}
