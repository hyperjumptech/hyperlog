package hyperlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogLevel_IsUnder(t *testing.T) {
	assert.True(t, TraceLevel.CanLogFor(TraceLevel))
	assert.False(t, TraceLevel.CanLogFor(DebugLevel))
	assert.False(t, TraceLevel.CanLogFor(InfoLevel))
	assert.False(t, TraceLevel.CanLogFor(WarnLevel))
	assert.False(t, TraceLevel.CanLogFor(ErrorLevel))
	assert.False(t, TraceLevel.CanLogFor(FatalLevel))

	assert.True(t, DebugLevel.CanLogFor(TraceLevel))
	assert.True(t, DebugLevel.CanLogFor(DebugLevel))
	assert.False(t, DebugLevel.CanLogFor(InfoLevel))
	assert.False(t, DebugLevel.CanLogFor(WarnLevel))
	assert.False(t, DebugLevel.CanLogFor(ErrorLevel))
	assert.False(t, DebugLevel.CanLogFor(FatalLevel))
}
