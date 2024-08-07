package hyperlog

import (
	"HyperIDP/pkg/masking"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type OutputFormat string

const (
	PlainFormat OutputFormat = "PLAIN"
	JSONFormat  OutputFormat = "JSON"
)

type LogLevel string

const (
	TraceLevel LogLevel = "trace"
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

type LogEngine interface {
	Trace(str ...interface{})
	Debug(str ...interface{})
	Info(str ...interface{})
	Warn(str ...interface{})
	Error(str ...interface{})
	Fatal(str ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(str string, args ...interface{})
	Infof(str string, args ...interface{})
	Warnf(str string, args ...interface{})
	Errorf(str string, args ...interface{})
	Fatalf(str string, args ...interface{})
	WithAttributes(attr map[string]string) LogEngine
	WithAttribute(key, value string) LogEngine
}

func (lvl LogLevel) Flag() int {
	switch lvl {
	case TraceLevel:
		return 1
	case DebugLevel:
		return 2
	case InfoLevel:
		return 4
	case WarnLevel:
		return 8
	case ErrorLevel:
		return 16
	case FatalLevel:
		return 32
	default:
		return -1
	}
}

func (lvl LogLevel) CanLogFor(thatLevel LogLevel) bool {
	return lvl.Flag() >= thatLevel.Flag()
}

func NewLogEntryf(lvl LogLevel, attr map[string]string, format string, args ...interface{}) *LogEntry {
	return &LogEntry{
		Level:      lvl,
		Time:       time.Now(),
		Message:    fmt.Sprintf(format, args...),
		Attributes: attr,
	}
}

func NewLogEntry(lvl LogLevel, attr map[string]string, message string) *LogEntry {
	return &LogEntry{
		Level:      lvl,
		Time:       time.Now(),
		Message:    message,
		Attributes: attr,
	}
}

type LogEntry struct {
	Level      LogLevel          `json:"lvl"`
	Time       time.Time         `json:"time"`
	Message    string            `json:"msg"`
	Attributes map[string]string `json:"attrs,omitempty"`
}

func (e *LogEntry) Mask() *LogEntry {
	le := &LogEntry{
		Level:      e.Level,
		Time:       e.Time,
		Message:    masking.Default.MaskSentence(e.Message),
		Attributes: nil,
	}
	if e.Attributes != nil && len(e.Attributes) > 0 {
		le.Attributes = make(map[string]string)
		for k, v := range e.Attributes {
			le.Attributes[k] = masking.Default.MaskSentence(v)
		}
	}
	return le
}

func (e *LogEntry) String() string {
	ret := make([]string, 0)
	ret = append(ret, fmt.Sprintf("[%s]", e.Level))
	ret = append(ret, e.Time.Format(time.RFC3339))
	ret = append(ret, e.Message)
	if e.Attributes != nil && len(e.Attributes) > 0 {
		for k, v := range e.Attributes {
			ret = append(ret, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return strings.Join(ret, " ")
}

func (e *LogEntry) JSONString() string {
	byts, err := json.Marshal(e)
	if err != nil {
		return e.String()
	}
	return string(byts)
}

type LogShutdownHandler interface {
	ShutdownLog()
}
