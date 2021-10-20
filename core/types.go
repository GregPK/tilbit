package core

type Tilbit struct {
	Text string
	Data TilbitData
}

type TilbitData struct {
	AddedOn string
	Source  string
	Url			string
	Private string
}

type SourceMetadata struct {
	Source string
	Url    string
}
