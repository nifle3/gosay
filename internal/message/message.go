package message

import "unicode/utf8"

func Generate(maxWorldInRow int, message string) string {
	messages := splitMessageInRow(maxWorldInRow, message)

	if len(messages) > 1 {

	}

	return message
}

func splitMessageInRow(maxWorldInRow int, message string) []string {
	if utf8.RuneCountInString(message) < maxWorldInRow {
		return []string{message}
	}

	return nil
}

func generateMultiRow(maxWorldInRow int, messsage []string) {

}

func generateOneRow(message string) {

}
