package main

import (
	"github.com/LetsFocus/goLF/goLF"
	"github.com/LetsFocus/goLF/slogs"
	"net/http"
	"os"
	"time"
)

func main() {
	g := goLF.New()

	g.Logger.Logger.Info("hello")
	g.Logger.Logger.Error("error")
	g.Logger.Logger.Warn("Warn")
	g.Logger.Logger.Debug("debug")

	g.Logger.Logger.Info(g.Config.Get("LOG_LEVEL"))

	//go test(g.Logger)

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
