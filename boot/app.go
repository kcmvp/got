package boot

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/samber/do/v2"
	"github.com/spf13/viper"
)

const (
	CfgName = "application"
)

var (
	container *do.RootScope
	cfg       *viper.Viper
	rootDir   string
	once      sync.Once
)

func init() {
	if output, err := exec.Command("go", "list", "-f", "{{.Dir}}", "./...").CombinedOutput(); err == nil {
		rootDir = CleanStr(string(output))
		_, err = os.Stat(filepath.Join(rootDir, "go.mod"))
		for err != nil {
			rootDir = filepath.Base(rootDir)
			_, err = os.Stat(filepath.Join(rootDir, "go.mod"))
		}
	} else {
		rootDir, _ = os.Executable()
	}
	container = do.NewWithOpts(&do.InjectorOpts{
		HookAfterRegistration: func(scope *do.Scope, serviceName string) {
			fmt.Printf("scope is %s, name is %s \n", scope.Name(), serviceName)
			//@todo, parse the mapping once
		},
		Logf: func(format string, args ...any) {
			log.Printf(format, args...)
		},
	})
}

func RootDir() string {
	return rootDir
}

func Container() *do.RootScope {
	return container
}

func InitApp() {
	if cfg == nil {
		once.Do(func() {
			cfg = viper.New()
			cfg.SetConfigName(CfgName)   // name of cfg file (without extension)
			cfg.SetConfigType("yaml")    // REQUIRED if the cfg file does not have the extension in the name
			cfg.AddConfigPath(RootDir()) // optionally look for cfg in the working directory
			err := cfg.ReadInConfig()    // Find and read the cfg file
			if err != nil {              // Handle errors reading the cfg file
				panic(fmt.Errorf("fatal error cfg file: %w", err))
			}
			if test, _ := Caller(); test {
				if testCfg, err := os.Open(filepath.Join(RootDir(), fmt.Sprintf("%s_test.yaml", CfgName))); err == nil {
					if err = cfg.MergeConfig(testCfg); err != nil {
						panic(fmt.Errorf("failed to merge test configuration file: %w", err))
					}
				}
			}
			setupDb()
		})
	}
}
