package model

import (
	"github.com/LetsFocus/goLF/cmd"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/metrics"
)

type GoLF struct {
	database.Database
	Config  configs.Config
	Logger  *logger.CustomLogger
	Metrics *metrics.Metrics
	*cmd.CLI
}
