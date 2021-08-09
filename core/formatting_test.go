package core

import (
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
)

func TestIsoDate(t *testing.T) {
	testza.AssertEqual(t, isoDate(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)), "2021-01-01") // -> Pass
}
