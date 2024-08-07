# Hyperlog

Simple logging framework that works.
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