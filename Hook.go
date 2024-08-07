package hyperlog

import (
	"sync"
)

func MakeHookFlag(trace, debug, info, warn, error, fatal bool) int {
	i := 0
	if trace {
		i = i | 1
	}
	if debug {
		i = i | 2
	}
	if info {
		i = i | 4
	}
	if warn {
		i = i | 8
	}
	if error {
		i = i | 16
	}
	if fatal {
		i = i | 32
	}
	return i
}

var (
	hookMutext          sync.Mutex
	hookMap             = make(map[int][]LogHook)
	logShutdownHandlers = make([]LogShutdownHandler, 0)
)

type LogHook interface {
	FireHook(entry *LogEntry) error
}

func Fire(entry *LogEntry) {
	hookMutext.Lock()
	defer hookMutext.Unlock()

	for fl, hooks := range hookMap {
		if fl&entry.Level.Flag() == entry.Level.Flag() {
			for _, hook := range hooks {
				hook.FireHook(entry)
			}
		}
	}
}

func Add(hook LogHook, flags int) {
	hookMutext.Lock()
	defer hookMutext.Unlock()

	if flags == 0 || hook == nil {
		return
	}

	if hooks, ok := hookMap[flags]; ok {
		if hooks == nil {
			hookMap[flags] = []LogHook{
				hook,
			}
		} else {
			hookMap[flags] = append(hookMap[flags], hook)
		}
	} else {
		hookMap[flags] = []LogHook{
			hook,
		}
	}
}

func Shutdown() {
	for _, handler := range logShutdownHandlers {
		handler.ShutdownLog()
	}
}

func AddLogShutdownHandler(handler LogShutdownHandler) {
	logShutdownHandlers = append(logShutdownHandlers, handler)
}
