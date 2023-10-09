package main

import "github.com/sdual/mlserving/apps/serving/server"

func main() {
	s := server.GRPCServer{}
	s.Start()
}
