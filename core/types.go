package core

import (
	"crypto/md5"
	"fmt"
	"io"
)

type Tilbit struct {
	Text     string
	Data     SourceMetadata
	Location SourceLocation
	DbID     int64 // Temporary field for database transition
}

func (t Tilbit) Id() string {
	return t.Hash()[:8]
}

func (t Tilbit) Hash() string {
	h := md5.New()
	io.WriteString(h, t.Text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type SourceMetadata struct {
	Author  string
	Source  string
	Url     string
	AddedOn string
	Private string
}

type SourceLocation struct {
	Uri        string
	LineNumber int
}
