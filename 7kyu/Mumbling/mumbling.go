package mumbling

import(
	"strings"
)

func Accum(s string) string {
	var words []string

	for i, char := range s {
		part := strings.ToUpper(string(char)) + strings.Repeat(strings.ToLower(string(char)), i)
		words = append(words, part)
	}

	return strings.Join(words, "-")
}