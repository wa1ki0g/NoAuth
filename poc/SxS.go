package poc

import "strings"

func SxS(noauth, auth string) []string {
	var resultList []string
	concatenated := auth + ";" + noauth
	resultList = append(resultList, concatenated)

	result := strings.ReplaceAll(noauth, "/", "%252f")
	concatenated = auth + ";" + result
	resultList = append(resultList, concatenated)
	return resultList
}
