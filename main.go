package main

import (
	"github.com/meSATYA/WowGoAPI/app"
	"github.com/meSATYA/WowGoAPI/logger"
)

func main() {

	logger.Info("Starting WowGoAPI")
	app.Start()
}
