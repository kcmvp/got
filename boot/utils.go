package boot

import (
	"runtime"
	"strings"

	"github.com/samber/lo"
)

// CleanStr Function to remove non-printable characters
func CleanStr(str string) string {
	cleanStr := func(r rune) rune {
		if r >= 32 && r != 127 {
			return r
		}
		return -1
	}
	return strings.Map(cleanStr, str)
}

func Caller() (bool, string) {
	var test bool
	var file string
	callers := make([]uintptr, 10)
	n := runtime.Callers(0, callers)
	frames := runtime.CallersFrames(callers[:n])
	for !test {
		frame, more := frames.Next()
		if !more {
			break
		}
		// fmt.Printf("%s->%s:%d\n", frame.File, frame.Function, frame.Line)
		if strings.HasPrefix(RootDir(), frame.File) {
			test = strings.HasSuffix(frame.File, "_test.go")
			items := strings.Split(frame.File, "/")
			items = lo.Map(items[len(items)-2:], func(item string, _ int) string {
				return strings.ReplaceAll(item, ".go", "")
			})
			uniqueNames := strings.Split(frame.Function, ".")
			items = append(items, uniqueNames[len(uniqueNames)-1])
			file = strings.Join(items, "_")
		}
	}
	return test, file
}
