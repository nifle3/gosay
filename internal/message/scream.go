package message

import (
	"strings"
	"unicode"
)

func GenerateScream(maxLenOfRow int, words []string) string {
	builder := &strings.Builder{}
	builder.WriteString("\033[0;31m")

	toUpper(words)
	generate(maxLenOfRow, words, builder)

	builder.WriteString("\033[0m")

	return builder.String()
}

func toUpper(words []string) {
	for _, value := range words {
		value = strings.ToUpperSpecial(unicode.SpecialCase{}, value)
	}
}
