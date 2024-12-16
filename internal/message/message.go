package message

import (
	"strings"
	"unicode/utf8"
)

func generate(maxLenOfRow int, words []string, builder *strings.Builder) {
	maxLenOfRow = getMaxLenOfRow(words, maxLenOfRow)
	generateMessage(maxLenOfRow, words, builder)
	generateTail(3, maxLenOfRow, builder)
}

func getMaxLenOfRow(words []string, maxlenOfRow int) int {
	for _, val := range words {
		valLen := utf8.RuneCountInString(val)
		if valLen > maxlenOfRow {
			maxlenOfRow = valLen
		}
	}

	return maxlenOfRow
}

func generateMessage(maxLenOfRow int, message []string, builder *strings.Builder) {
	generateTopAndBottom(maxLenOfRow, builder)
	genrateBodyOfMessage(maxLenOfRow, message, builder)
	generateTopAndBottom(maxLenOfRow, builder)
}

func genrateBodyOfMessage(maxLenOfRow int, words []string, builder *strings.Builder) {
	rowBuilder := strings.Builder{}
	currentLenOfRow := 0

	for _, val := range words {
		valLen := utf8.RuneCountInString(val)

		if currentLenOfRow+valLen+1 <= maxLenOfRow {
			rowBuilder.WriteString(val + " ")
			currentLenOfRow += valLen + 1
			continue
		}

		if rowBuilder.Len() > 0 {
			row := rowBuilder.String()
			generateRowOfMessage(row, builder)
		}

		rowBuilder.Reset()
		rowBuilder.WriteString(val + " ")
		currentLenOfRow = valLen + 1
	}

	if rowBuilder.Len() > 0 {
		row := rowBuilder.String()
		generateRowOfMessage(row, builder)
	}
}

func generateTopAndBottom(length int, builder *strings.Builder) {
	builder.WriteString(" ")
	builder.WriteString(strings.Repeat("-", length))
	builder.WriteString("\n")
}

func generateRowOfMessage(message string, builder *strings.Builder) {
	builder.WriteString("|")
	builder.WriteString(message)
	builder.WriteString("|\n")
}

func generateTail(length, width int, builder *strings.Builder) {
	for i := 0; i < length; i++ {
		builder.WriteString(strings.Repeat(" ", width+i))
		builder.WriteString("\\\n")
	}
}
