package appconf

import (
	"fmt"

	"github.com/sdual/mlserving/pkg/env"
	"github.com/spf13/viper"
)

func Load(appName string, config any) any {
	currentEnv := env.CurrentEnv()
	viper.SetConfigName(tomlFileType.fileName(currentEnv))
	viper.SetConfigType(tomlFileType.extension)

	viper.AddConfigPath(confFilePath(currentEnv, appName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("app config file is not found: %w", err))
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config file: %w", err))
	}
	return config
}

func confFilePath(currentEnv env.SystemEnv, appName string) string {
	if currentEnv == env.Test {
		return "apps/" + appName + "/config/"
	}
	return "/"
}
