
package core

type Tilbit struct {
	Text string
	Data TilbitData
}

type TilbitData struct {
	AddedOn string
	Source  string
	Private string
}
