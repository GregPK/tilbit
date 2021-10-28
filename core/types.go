package core

type Tilbit struct {
	Text string
	Data SourceMetadata
}

type SourceMetadata struct {
	Author  string
	Source  string
	Url     string
	AddedOn string
	Private string
}
