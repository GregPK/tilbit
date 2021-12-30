package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/GregPK/tilbit/core"
	"github.com/MarvinJWendt/testza"
)

var inputTilbits []core.Tilbit

func TestShowCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	Config.outputFormat = "text"
	cmd := ShowCmd(core.Tilbit{Text: "Tilbit text"})
	cmd.SetOut(b)
	cmd.SetArgs([]string{"random"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	testza.AssertEqual(t, "Tilbit text\n   --  (id: 1ff954c5)", string(out))
}

func TestRootCommand(t *testing.T) {
	b := bytes.NewBufferString("")
	Config.outputFormat = "text"
	cmd := RootCmd(core.Tilbit{Text: "Tilbit text"})
	cmd.SetOut(b)
	cmd.SetArgs([]string{"random"})
	cmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	testza.AssertEqual(t, "Tilbit text\n   --  (id: 1ff954c5)", string(out))
}
