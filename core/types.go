package core

type Tilbit struct {
	Text     string
	Data     SourceMetadata
	Location SourceLocation
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
