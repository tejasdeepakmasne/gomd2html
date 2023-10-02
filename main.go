package main

import (
	"flag"
	"io"
	"os"
	"regexp"
	"strings"
)

const BODY_TAG_POSITION = 158
const TITLE_TAG_POSITION = 137

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_heading_1(text string) string {
	re := regexp.MustCompile(`#\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h1>" + text[2:] + "</h1>"
		return processed_string
	}
}

func check_heading_2(text string) string {
	re := regexp.MustCompile(`##\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h2>" + text[3:] + "</h2>"
		return processed_string
	}
}

func check_heading_3(text string) string {
	re := regexp.MustCompile(`###\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h3>" + text[4:] + "</h3>"
		return processed_string
	}
}

func check_heading_4(text string) string {
	re := regexp.MustCompile(`####\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h4>" + text[5:] + "</h4>"
		return processed_string
	}
}

func check_heading_5(text string) string {
	re := regexp.MustCompile(`#####\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h5>" + text[6:] + "</h5>"
		return processed_string
	}
}

func check_heading_6(text string) string {
	re := regexp.MustCompile(`######\s`)
	isPresent := re.MatchString(text)

	if !isPresent {
		return text
	} else {
		processed_string := "<h6>" + text[7:] + "</h6>"
		return processed_string
	}
}

func check_headings(text string) string {

}
func Parser(text string) []string {
	lines := strings.Split(text, "\n")
	tags_to_write := make([]string, 0)

	for _, line := range lines {
		if line[1] == '#' {
			line = check_headings(line)
			tags_to_write = append(tags_to_write, line)
		}
	}

	return tags_to_write
}

func Writer(tags_to_write []string, outputfile *os.File) {
	//write_pos := BODY_TAG_POSITION
	for i := 0; i < len(tags_to_write); i++ {
		outputfile.WriteString(tags_to_write[i])
	}
}

// func WriteTitle(outputFile *os.File, title string) {

// }

func main() {

	//templateSlice := []byte{60, 33, 68, 79, 67, 84, 89, 80, 69, 32, 104, 116, 109, 108, 62, 60, 104, 116, 109, 108, 32, 108, 97, 110, 103, 61, 34, 101, 110, 34, 62, 60, 104, 101, 97, 100, 62, 60, 109, 101, 116, 97, 32, 99, 104, 97, 114, 115, 101, 116, 61, 34, 85, 84, 70, 45, 56, 34, 62, 60, 109, 101, 116, 97, 32, 110, 97, 109, 101, 61, 34, 118, 105, 101, 119, 112, 111, 114, 116, 34, 32, 99, 111, 110, 116, 101, 110, 116, 61, 34, 119, 105, 100, 116, 104, 61, 100, 101, 118, 105, 99, 101, 45, 119, 105, 100, 116, 104, 44, 32, 105, 110, 105, 116, 105, 97, 108, 45, 115, 99, 97, 108, 101, 61, 49, 46, 48, 34, 62, 60, 116, 105, 116, 108, 101, 62, 60, 47, 116, 105, 116, 108, 101, 62, 60, 47, 104, 101, 97, 100, 62, 60, 98, 111, 100, 121, 62, 60, 47, 98, 111, 100, 121, 62, 60, 47, 104, 116, 109, 108, 62}

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
	beginning_template := "<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Document</title></head><body>"
	outputFile.WriteString(beginning_template) //outputFile.WriteAt(templateSlice, 0) //fill the output file with template
	parsed_text := Parser(inputString)
	Writer(parsed_text, outputFile)
	ending_template := "</body></html>"
	outputFile.WriteString(ending_template)
}
