package core

import (
	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Source string
	Url    string
}

func ParseMetadata(input string) (err error, metadata Metadata) {
	metadata = Metadata{}

	err = yaml.Unmarshal([]byte(input), &metadata)

	return err, metadata
}
