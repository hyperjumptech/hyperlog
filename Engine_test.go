package hyperlog

import (
	"bytes"
	"testing"
)

func TestLogOutput(t *testing.T) {
	buff := bytes.NewBuffer(make([]byte, 0))
	Level = DebugLevel
	SetWriter(buff)
	OutFormat = JSONFormat
	Mask = false

	Debug("Hallo Hello Yoooo Houuu")
	// Tracef("%d %d %d", 1, 2, 3)
	t.Log(buff.String())
}
