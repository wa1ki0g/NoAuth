package poc

func Twop(input []string) []string {
	var result []string
	for _, item := range input {
		result = append(result, item+"/..")
	}
	return result
}
