package main

import (
	"github.com/sdual/mlserving/apps/serving/internal/config"
	"github.com/sdual/mlserving/apps/serving/internal/server"
	"github.com/sdual/mlserving/pkg/appconf"
)

const appName = "serving"

func main() {
	appConfig := appconf.Load(appName, &config.AppConfig{}).(*config.AppConfig)
	s := server.GRPCServer{}
	s.Start(appConfig.GRPC)
}
