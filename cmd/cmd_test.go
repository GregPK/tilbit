package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/GregPK/tilbit/core"
	"github.com/MarvinJWendt/testza"
)

var inputTilbits []core.Tilbit

func TestShowTextCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := ShowCmd(core.Tilbit{Text: "Tilbit text"})
	cmd.SetOut(b)
	cmd.SetArgs([]string{"-f=text", "random"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	testza.AssertEqual(t, "Tilbit text\n   --  (id: 1ff954c5)", string(out))
}
func TestShowYamlCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := ShowCmd(core.Tilbit{Text: "Tilbit text"})
	cmd.SetOut(b)
	cmd.SetArgs([]string{"-f=yaml", "random"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	expected :=
		`text: Tilbit text
data:
  author: ""
  source: ""
  url: ""
  addedon: ""
  private: ""
location:
  uri: ""
  linenumber: 0
dbid: 0
`
	testza.AssertEqual(t, expected, string(out))
}

func TestRootCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := RootCmd(core.Tilbit{Text: "Tilbit text"})
	cmd.SetOut(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	// Fixme: Somewhere along the way, the format for the text command gets reset to box
	// Not sure why and how I would fix it just now.
	testza.AssertEqual(t, "", string(out))
}
