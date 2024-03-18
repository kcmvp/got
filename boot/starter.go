package boot

import (
	"fmt"
	"github.com/samber/do/v2"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var container *do.RootScope
var config *viper.Viper
var once sync.Once

func init() {
	// environment sensitive name, when in test environment the file name should be 'application-test'
	config.SetConfigName("application") // name of config file (without extension)
	config.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(".")           // optionally look for config in the working directory
	err := config.ReadInConfig()        // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	do.ProvideValue[*viper.Viper](Container(), config)
	//  if db then
	setupDb([]dataSource{})
}

func Container() *do.RootScope {
	if container == nil {
		once.Do(func() {
			container = do.NewWithOpts(&do.InjectorOpts{
				HookAfterRegistration: func(scope *do.Scope, serviceName string) {
					fmt.Printf("scope is %s, name is %s \n", scope.Name(), serviceName)
					//@todo, parse the mapping once
				},
				Logf: func(format string, args ...any) {
					log.Printf(format, args...)
				},
			})
		})
	}
	return container
}
