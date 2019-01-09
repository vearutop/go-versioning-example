package main

import "go.uber.org/zap"

var version = "dev"

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Sugar().Infof("App Version: %s\n", version)
}
