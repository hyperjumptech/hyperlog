package hyperlog

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

type DummyHook struct {
	sound string
}

func (dummy *DummyHook) FireHook(entry *LogEntry) error {
	fmt.Println(dummy.sound + " - " + entry.Message)
	return nil
}

func TestAdd(t *testing.T) {
	bingHook := &DummyHook{sound: "Bing"}
	bongHook := &DummyHook{sound: "Bong"}
	bangHook := &DummyHook{sound: "Dang"}

	Add(bingHook, DebugLevel.Flag())
	Add(bongHook, TraceLevel.Flag())
	Add(bangHook, DebugLevel.Flag())

	Fire(&LogEntry{
		Level:      TraceLevel,
		Message:    "Bing",
		Attributes: nil,
		Time:       time.Now(),
	})
}

func TestBit(t *testing.T) {
	t.Logf("%d & %d = %d", 0b001011, 0b000100, 0b001011&0b000100)
	assert.Equal(t, 0, 11&4)
}
