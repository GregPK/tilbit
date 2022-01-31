package core

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	frontmatterBreak = "---"
)

func ParseSource(source string) (tilbits []Tilbit, err error) {
	fileInfo, err := os.Stat(source)
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		tilbits, err = parseDirectory(source)
	} else {
		tilbits, err = parseFile(source)
	}

	return
}

func parseDirectory(source string) (tilbits []Tilbit, err error) {
	files, err := ioutil.ReadDir(source)
	for _, file := range files {
		name := filepath.Join(source, file.Name())

		if !file.IsDir() {
			bits, fileerr := parseFile(name)
			if fileerr != nil {
				log.Fatal(fmt.Sprintf("Error while parsing [%s]: %s", name, fileerr))
			}

			tilbits = append(tilbits, bits...)
		}
	}
	return
}

func parseFile(fileString string) (tilbits []Tilbit, err error) {
	fileParts := strings.Split(fileString, ".")
	extension := fileParts[len(fileParts)-1]

	if extension == "sqlite" {
		err, tilbits = ParseSqliteFile(fileString)
		return
	}

	data, err := ioutil.ReadFile(fileString)
	if err != nil {
		panic(err)
	}

	if extension == "md" {
		err, tilbits, _ = ParseMarkdownFile(string(data), fileString)
	} else if strings.Contains(fileString, "My Clippings.txt") {
		err, tilbits = ParseKindleClippingsFile(string(data), fileString)
	} else {
		err, tilbits = ParseTextFile(string(data), fileString)
	}
	return
}

func ParseIdsFromString(idString string) (ids []string) {
	re := regexp.MustCompile("\\W+")
	for _, id := range re.Split(idString, -1) {
		ids = append(ids, id)
	}
	return
}

func ParseSqliteFile(fileName string) (err error, tilbits []Tilbit) {
	// @todo implement parsing sqlite as input source
	return
}

func ParseMarkdownFile(fileContent string, fileName string) (err error, tilbits []Tilbit, metadata SourceMetadata) {
	frontmatter := ""
	body := ""

	parts := strings.Split(fileContent, frontmatterBreak)
	if len(parts) > 1 {
		frontmatter = parts[1]
		body = parts[2]
	} else {
		body = parts[0]
	}

	frontmatterLines := strings.Count(frontmatter, "\n") + strings.Count(fileContent, frontmatterBreak)

	err, metadata = parseMetadata(strings.TrimSpace(frontmatter))
	err, tilbits = parseMarkdownBody(strings.TrimSpace(body))

	for i, tilbit := range tilbits {
		tilbit.Data = metadata
		tilbit.Location.Uri = fileName
		tilbit.Location.LineNumber += frontmatterLines

		tilbits[i] = tilbit
	}

	return
}

func ParseTextFile(fileContent string, fileName string) (err error, tilbits []Tilbit) {
	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	lineNumber := 0
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		lineNumber += 1

		if line == "" {
			continue
		}

		parts := strings.Split(line, "{")
		if len(parts) != 2 {
			return errors.New(fmt.Sprintf("Unexpected parse of line with metadata: %s", line)), nil
		}

		text := strings.Trim(parts[0], " ")
		jsonstr := "{" + parts[1]

		var tilbit Tilbit
		var metadata SourceMetadata
		json.Unmarshal([]byte(jsonstr), &metadata)
		tilbit.Text = text
		tilbit.Data = metadata
		tilbit.Location.Uri = fileName
		tilbit.Location.LineNumber = lineNumber

		tilbits = append(tilbits, tilbit)
	}
	return
}

func ParseKindleClippingsFile(fileContent string, fileName string) (err error, tilbits []Tilbit) {
	items := strings.Split(fileContent, "==========")
	lineNumber := 1
	for _, item := range items {
		item = strings.Trim(item, " \n\r")
		if len(item) == 0 {
			lineNumber += 1
			continue
		}

		var tilbit Tilbit
		lines := linesFrom(item, false)

		tilbit.Data.Source = strings.Trim(lines[0], " \n\r")
		tilbit.Text = strings.Trim(strings.Join(lines[3:], "\n"), " \n\r")
		tilbit.Location.Uri = fileName
		tilbit.Location.LineNumber = lineNumber

		tilbits = append(tilbits, tilbit)
		lineNumber += len(lines) + 1
	}

	return
}

func parseMarkdownBody(input string) (err error, tilbits []Tilbit) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNumber := 0
	lastTilbitText := "" // needed because we're looking for paragraph (two line breaks)
	for scanner.Scan() {
		lineNumber += 1
		text := scanner.Text()

		if len(strings.Trim(text, " \n")) == 0 && len(lastTilbitText) > 0 {
			tilbitLineLen := strings.Count(lastTilbitText, "\n") + 1
			tilbits = append(tilbits, Tilbit{strings.Trim(lastTilbitText, "\n"), SourceMetadata{}, SourceLocation{LineNumber: lineNumber - tilbitLineLen}, -1, ""})
			lastTilbitText = ""
			continue
		}
		lastTilbitText += fmt.Sprintln(text)
	}
	tilbits = append(tilbits, Tilbit{strings.Trim(lastTilbitText, "\n"), SourceMetadata{}, SourceLocation{LineNumber: lineNumber}, -1, ""})

	return
}

func linesFrom(str string, numbered bool) (result []string) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	ln := 0
	for scanner.Scan() {
		line := scanner.Text()
		if numbered {
			line = strconv.Itoa(ln) + ":" + line
		}
		result = append(result, line)
		ln += 1
	}
	return result
}

func parseMetadata(input string) (err error, metadata SourceMetadata) {
	metadata = SourceMetadata{}

	input = strings.ReplaceAll(input, frontmatterBreak, "")
	input = strings.Trim(input, "\n ")
	err = yaml.Unmarshal([]byte(input), &metadata)

	return
}
