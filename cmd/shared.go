package cmd

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/GregPK/tilbit/core"
)

func dbDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tilbit/data/"
}

func privateDbFilename() string {
	return dbDir() + "private.txt"
}

func parseFile(fileString string) (tilbits []core.Tilbit) {
	fileParts := strings.Split(fileString, ".")
	extension := fileParts[len(fileParts)-1]

	data, err := ioutil.ReadFile(fileString)
	if err != nil {
		panic(err)
	}

	if extension == "md" {
		err, tilbits = core.ParseMarkdownBody(string(data))
	} else {
		err, tilbits = core.ParseTextFile(string(data))
	}
	return
}
