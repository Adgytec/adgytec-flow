package normalize

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Unaccent(src string) (string, error) {
	var b strings.Builder
	b.Grow(len(src))

	for _, r := range src {
		if repl, ok := replacements[r]; ok {
			b.WriteString(repl)
		} else {
			b.WriteRune(r)
		}
	}
	src = b.String()

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, src)
	if err != nil {
		return "", err
	}

	return result, nil
}
