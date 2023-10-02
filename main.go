package main

import (
	"flag"
	"io"
	"os"
)

const BODY_TAG_POSITION = 158
const TITLE_TAG_POSITION = 137

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// defining the types of the
type TokenType int

const (
	Heading = iota
	UnorderedList
	OrderedList
	HorizontalLine
	Link
	Italic
	Bold
	BlockQuote
)

// making a token struct
type Token struct {
	Type  TokenType
	Value string
}

func NewToken(t TokenType, v string) Token {
	return Token{Type: t, Value: v}
}

// func WriteTitle(outputFile *os.File, title string) {

// }

func main() {

	templateSlice := []byte{60, 33, 68, 79, 67, 84, 89, 80, 69, 32, 104, 116, 109, 108, 62, 60, 104, 116, 109, 108, 32, 108, 97, 110, 103, 61, 34, 101, 110, 34, 62, 60, 104, 101, 97, 100, 62, 60, 109, 101, 116, 97, 32, 99, 104, 97, 114, 115, 101, 116, 61, 34, 85, 84, 70, 45, 56, 34, 62, 60, 109, 101, 116, 97, 32, 110, 97, 109, 101, 61, 34, 118, 105, 101, 119, 112, 111, 114, 116, 34, 32, 99, 111, 110, 116, 101, 110, 116, 61, 34, 119, 105, 100, 116, 104, 61, 100, 101, 118, 105, 99, 101, 45, 119, 105, 100, 116, 104, 44, 32, 105, 110, 105, 116, 105, 97, 108, 45, 115, 99, 97, 108, 101, 61, 49, 46, 48, 34, 62, 60, 116, 105, 116, 108, 101, 62, 60, 47, 116, 105, 116, 108, 101, 62, 60, 47, 104, 101, 97, 100, 62, 60, 98, 111, 100, 121, 62, 60, 47, 98, 111, 100, 121, 62, 60, 47, 104, 116, 109, 108, 62}

	inputPath := flag.String("i", "input.md", "path of the input markdown file")
	outputPath := flag.String("o", "output.html", "path of the output html file")
	//title := flag.String("t", *inputPath, "the title used in the html document")

	flag.Parse()

	inputFile, err := os.Open(*inputPath)
	check(err)

	outputFile, err := os.OpenFile(*outputPath, os.O_WRONLY|os.O_CREATE, 0644)
	check(err)

	inputFileByteSlice, err := io.ReadAll(inputFile) //all characters of the file
	check(err)

	inputString := string(inputFileByteSlice)

	outputFile.WriteAt(templateSlice, 0) //fill the output file with template

}
