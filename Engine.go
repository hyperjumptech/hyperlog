package hyperlog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	mutex sync.Mutex

	writer    io.Writer    = os.Stdout
	Level     LogLevel     = TraceLevel
	OutFormat OutputFormat = JSONFormat
	Mask      bool         = false
)

func GetWriter() io.Writer {
	return writer
}

func ShutdownWriter() {
	writer = nil
}

func SetWriter(w io.Writer) {
	if w != nil && writer != nil {
		writer = w
	}
}

func Trace(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	msg := fmt.Sprint(str...)
	if TraceLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(TraceLevel, nil, msg)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Debug(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if DebugLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(DebugLevel, nil, fmt.Sprint(str...))
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Info(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if InfoLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(InfoLevel, nil, fmt.Sprint(str...))
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Warn(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if WarnLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(WarnLevel, nil, fmt.Sprint(str...))
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Error(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if ErrorLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(ErrorLevel, nil, fmt.Sprint(str...))
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Fatal(str ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if FatalLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntry(FatalLevel, nil, fmt.Sprint(str...))
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}

func Tracef(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if TraceLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(TraceLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Debugf(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if DebugLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(DebugLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Infof(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if InfoLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(InfoLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Warnf(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if WarnLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(WarnLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Errorf(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if ErrorLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(ErrorLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}
func Fatalf(format string, args ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	if FatalLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		entry := NewLogEntryf(FatalLevel, nil, format, args...)
		if OutFormat == JSONFormat {
			_, _ = GetWriter().Write([]byte(entry.JSONString() + "\n"))
		} else {
			_, _ = GetWriter().Write([]byte(entry.String() + "\n"))
		}
		Fire(entry)
	}
}

func WithAttributes(attr map[string]string) LogEngine {
	return &LogEngineImpl{
		attributes: attr,
	}
}

func WithAttribute(key, value string) LogEngine {
	return WithAttributes(map[string]string{
		key: value,
	})
}

type LogEngineImpl struct {
	entryMutex sync.Mutex
	attributes map[string]string
}

func (engine *LogEngineImpl) toString(lvl LogLevel, mask bool, str string) (string, *LogEntry) {
	entry := NewLogEntry(lvl, engine.attributes, str)
	if OutFormat == JSONFormat {
		return entry.JSONString() + "\n", entry
	}
	if OutFormat == PlainFormat {
		return entry.String() + "\n", entry
	}
	return entry.String() + "\n", entry
}

func (engine *LogEngineImpl) toStringf(lvl LogLevel, mask bool, format string, args ...interface{}) (string, *LogEntry) {
	entry := NewLogEntryf(lvl, engine.attributes, format, args...)
	if OutFormat == JSONFormat {
		return entry.JSONString() + "\n", entry
	}
	if OutFormat == PlainFormat {
		return entry.String() + "\n", entry
	}
	return entry.String() + "\n", entry
}

func (engine *LogEngineImpl) Trace(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if TraceLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(TraceLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Debug(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if DebugLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(DebugLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Info(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if InfoLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(InfoLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Warn(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if WarnLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(WarnLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Error(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if ErrorLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(ErrorLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Fatal(itfs ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if FatalLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toString(FatalLevel, Mask, fmt.Sprint(itfs...))
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Tracef(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if TraceLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(TraceLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Debugf(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if DebugLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(DebugLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Infof(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if InfoLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(InfoLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Warnf(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if WarnLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(WarnLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Errorf(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if ErrorLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(ErrorLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}
func (engine *LogEngineImpl) Fatalf(format string, args ...interface{}) {
	engine.entryMutex.Lock()
	defer engine.entryMutex.Unlock()
	if FatalLevel.CanLogFor(Level) {
		if GetWriter() == nil {
			return
		}
		msg, entry := engine.toStringf(FatalLevel, Mask, format, args...)
		_, _ = GetWriter().Write([]byte(msg))
		Fire(entry)
	}
}

func (engine *LogEngineImpl) WithAttributes(attr map[string]string) LogEngine {
	m := make(map[string]string)
	for k, v := range engine.attributes {
		m[k] = v
	}
	for k, v := range attr {
		m[k] = v
	}
	newLogStruct := &LogEngineImpl{
		attributes: attr,
	}

	return newLogStruct
}

func (engine *LogEngineImpl) WithAttribute(key, value string) LogEngine {
	return engine.WithAttributes(map[string]string{
		key: value,
	})
}
