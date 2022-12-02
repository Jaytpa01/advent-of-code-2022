package utils

import (
	"log"
	"os"
	"strings"
)

func OpenInputFile(path string) string {
	inputBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", path, err)
	}

	input := string(inputBytes)

	// replace all CRLF's with LF's:
	// https://developer.mozilla.org/en-US/docs/Glossary/CRLF
	input = strings.ReplaceAll(input, "\r\n", "\n")

	return input
}
