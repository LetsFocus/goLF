package main

import (
	"github.com/LetsFocus/goLF/goLF"
	"github.com/LetsFocus/goLF/slogs"
	"net/http"
	"os"
	"time"
)

func main() {
	goLF.New()
	logger := slogs.NewLogger()

	logger.Logger.Info("hello")
	logger.Logger.Error("error")
	logger.Logger.Warn("Warn")
	logger.Logger.Debug("debug")

	logger.Logger.Info(os.Getenv("LOG_LEVEL"))

	go test(logger)

	http.ListenAndServe(":8000", nil)
}

func test(log slogs.Log) {
	log.Logger.Info("hello")
	log.Logger.Error("error")
	log.Logger.Warn("Warn")
	log.Logger.Debug("debug")
	log.Logger.Info(os.Getenv("LOG_LEVEL"))

	time.Sleep(time.Second * 5)
}
