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
	showCmd = ShowCmd(core.Tilbit{Text: "Tilbit text"})
	showCmd.SetOut(b)
	showCmd.SetArgs([]string{"random"})
	showCmd.Execute()

	out, err := ioutil.ReadAll(b)
	testza.AssertNoError(t, err)
	testza.AssertEqual(t, "Tilbit text\n   --  (id: 1ff954c5)", string(out))

	testza.AssertEqual(t, 1, 1) // -> Pass
}
