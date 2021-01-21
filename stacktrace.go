package logger

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var (
	_stacktracePool = sync.Pool{
		New: func() interface{} {
			return newProgramCounters(64)
		},
	}
)

func TakeStacktrace() string {

	programCounters := _stacktracePool.Get().(*programCounters)
	defer _stacktracePool.Put(programCounters)

	var numFrames int
	for {
		// Skip the call to runtime.Counters and takeStacktrace so that the
		// program counters start at the caller of takeStacktrace.
		numFrames = runtime.Callers(2, programCounters.pcs)
		if numFrames < len(programCounters.pcs) {
			break
		}
		// Don't put the too-short counter slice back into the pool; this lets
		// the pool adjust if we consistently take deep stacktraces.
		programCounters = newProgramCounters(len(programCounters.pcs) * 2)
	}

	i := 0
	skipFrames := true // skip all consecutive zap frames at the beginning.
	frames := runtime.CallersFrames(programCounters.pcs[:numFrames])

	// Note: On the last iteration, frames.Next() returns false, with a valid
	// frame, but we ignore this frame. The last frame is a a runtime frame which
	// adds noise, since it's only either runtime.main or runtime.goexit.
	var str string
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		if skipFrames && strings.Contains(frame.File, "github.com/gzw13999/logger") {
			continue
		} else {
			skipFrames = false
		}

		if i != 0 {
			str += "\n"
		}
		i++
		str += fmt.Sprintf("%v\n\t%v:%v", frame.Function, frame.File, frame.Line)

	}

	return str
}

type programCounters struct {
	pcs []uintptr
}

func newProgramCounters(size int) *programCounters {
	return &programCounters{make([]uintptr, size)}
}
