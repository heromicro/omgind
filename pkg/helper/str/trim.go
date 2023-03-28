package str

import (
	"regexp"
	"strings"
)

func RemoveDuplicateBlank(input string) string {

	re := regexp.MustCompile(`\s+`)

	out := strings.TrimSpace(input)
	out = re.ReplaceAllString(out, " ")

	return out
}
