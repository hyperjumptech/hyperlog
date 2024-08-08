# Hyperlog

Simple logging framework that works.

## Importing

```bash
$ go get github.com/hyperjumptech/hyperlog
```

## Features

It enable the following capabilities:

- Log Level
- JSON Log Output
- Log Hook and triggers on any level
- Shutdown Hook
- Log to any implementation io.Writer
- Writer implementation to write log for rolling based on time
- Writer implementation to write log for rolling based on size
- Log writer with "Attribute" (a.k.a "Field")
- Http middleware to log Request/Response

## Using hyperlog

### Preparation

Just import the library, as simple as that.

```go
import "github.com/hyperjumptech/hyperlog"
```

### Configure

### Logging

The following can be use for logging straight away...
```go
func Trace(str ...interface{})
func Debug(str ...interface{})
func Info(str ...interface{})
func Warn(str ...interface{})
func Error(str ...interface{})
func Fatal(str ...interface{})

func Tracef(format string,  ...interface{})
func Debugf(format string,  arg ...interface{})
func Infof(format string,  arg ...interface{})
func Warnf(format string,  arg ...interface{})
func Errorf(format string,  arg ...interface{})
func Fatalf(format string,  arg ...interface{})
```

So you can directly use them like ...

```go
import "github.com/hyperjumptech/hyperlog"

func DoSomething() {
    hyperlog.Warn("This is a warning")
}

func DoSomethingAgain() {
    hyperlog.Warnf("This is a %s", "warning")
}
```

Or you can create a logger using `WithAttribute` to add some specific attribute.

```go
import "github.com/hyperjumptech/hyperlog"

func LogIt() {
    log1 := hyperlog.WithAttribute("key", "Value One")
    log2 := hyperlog.WithAttribute("moreKey", "More Value")
    log3 := log2.WithAttribute("anotherKey", "More Value")

    log1.Info("Hello One")
    log2.Info("More Hello")
    log3.Info("Another log")
}
```

which will log something like the following (in plain format)

```text
[INFO] 2024-07-21T10:10:00 Hello One key="Value One"
[INFO] 2024-07-21T10:10:00 More Hello moreKey="More Value"
[INFO] 2024-07-21T10:10:00 Another log moreKey="More Value" anotherKey="More Value"
```

### Changing Output Format

You can generate log in either "JSON" or "Plain" format.

```go
import "github.com/hyperjumptech/hyperlog"

func init() {
    OutFormat = JSONFormat // this will produce log JSON format
	... or ...
    OutFormat = PlainFormat // this will produce log in Plain Text format
}
```

so, a call like this ...

```go
hyperlog.Error("an error has occurred")
```

will log ...

JSON

```text
{"lvl"="error","time":"2024-07-21T10:10:00","msg":an error has occurred"}
```

or Plain Text

```text
[error] 2024-07-21T10:10:00 an error has occurred
```

### Changing Writer (output target)

You can change the logging target like following.

```go
import "github.com/hyperjumptech/hyperlog"

func init() {
	hyperlog.SetWriter(os.StdOut)
}
```

or 

```go
import "github.com/hyperjumptech/hyperlog"

func init() {
	hyperlog.SetWriter(os.StdErr)
}
```

or 

```go
import "github.com/hyperjumptech/hyperlog"

type CustomWriter struct {}
func (c *CustomWriter) Write(b []byte) (written int, err error) {
	// some io.Writer() logic.
}

func init() {
	hyperlog.SetWriter(&CustomWriter{})
}
```

This you can set the log target as specified.

### Hooks

__More documentation will be added soon__

### File Based Writer

#### Rolling Log File based on Time

__More documentation will be added soon__

#### Rolling Log File based on file size

__More documentation will be added soon__
