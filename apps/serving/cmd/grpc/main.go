package main

import (
	"fmt"

	"github.com/sdual/mlserving/apps/serving/internal/config"
	"github.com/sdual/mlserving/apps/serving/internal/server"
	"github.com/sdual/mlserving/pkg/appconf"
	"github.com/sdual/mlserving/pkg/logger"
)

const appName = "serving"

func main() {
	appConfig := appconf.Load(appName, &config.AppConfig{}).(*config.AppConfig)
	logger.SetUpLogger(appConfig.Logger.LogLevel)
	fmt.Println(appConfig)

	s := server.GRPCServer{}
	s.Start(appConfig.GRPC)
}
