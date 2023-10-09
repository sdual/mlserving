package main

import (
	"github.com/sdual/mlserving/apps/serving/internal/server"
)

func main() {
	s := server.GRPCServer{}
	s.Start()
}
