package message

import "strings"

func GenerateSay(maxLenOfRow int, words []string) string {
	builder := &strings.Builder{}

	generate(maxLenOfRow, words, builder)

	return builder.String()
}
