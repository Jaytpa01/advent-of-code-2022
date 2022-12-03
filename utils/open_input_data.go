package utils

import (
	"bufio"
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

type FileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileScanner(filePath string) *FileScanner {
	file := loadFile(filePath)
	scanner := newLineScannerFromFile(file)

	return &FileScanner{file, scanner}
}

func (fs *FileScanner) Scan() bool {
	return fs.scanner.Scan()
}

func (fs *FileScanner) Text() string {
	return fs.scanner.Text()
}

func (fs *FileScanner) Close() error {
	return fs.file.Close()
}

func loadFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to load input file: %v", err)
	}

	return f
}

func newLineScannerFromFile(f *os.File) *bufio.Scanner {
	return bufio.NewScanner(f)
}
