package poc

import (
	"strings"
)

func GFG(path string) []string {
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
			newPathWithoutSemicolon := currentPath + part
			if index < len(parts)-1 {
				newPathWithoutSemicolon += "/"
			}
			generate(newPathWithoutSemicolon, index+1)

			if currentPath != "" {
				newPathWithSemicolon := currentPath + "/;//" + part
				if index < len(parts)-1 {
					newPathWithSemicolon += "/"
				}
				generate(newPathWithSemicolon, index+1)
			}
		} else if currentPath != "" {
			generate(currentPath, index+1)
		}
	}

	generate("/", 0)

	return results
}
