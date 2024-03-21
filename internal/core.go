package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kcmvp/got/util"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
)

var (
	Container *do.RootScope
	RootDir   string
)

func init() {
	Container = do.NewWithOpts(&do.InjectorOpts{
		HookAfterRegistration: func(scope *do.Scope, serviceName string) {
			fmt.Printf("scope is %s, name is %s \n", scope.Name(), serviceName)
			//@todo, parse the mapping once
		},
		Logf: func(format string, args ...any) {
			log.Printf(format, args...)
		},
	})

	if output, err := exec.Command("go", "list", "-f", "{{.Dir}}", "./...").CombinedOutput(); err == nil {
		RootDir = util.CleanStr(string(output))
		_, err = os.Stat(filepath.Join(RootDir, "go.mod"))
		for err != nil {
			RootDir = filepath.Dir(RootDir)
			_, err = os.Stat(filepath.Join(RootDir, "go.mod"))
		}
	} else {
		RootDir, _ = os.Executable()
	}
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
		if strings.HasPrefix(frame.File, RootDir) {
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
