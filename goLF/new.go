package goLF

import (
	"github.com/LetsFocus/goLF/cmd"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/metrics"
)

func New() model.GoLF {
	var goLF model.GoLF

	goLF.Logger = logger.NewCustomLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	goLF.Metrics = metrics.NewMetricsServer()

	initializeDatabases(&goLF)

	return goLF
}

func NewCMD() model.GoLF {
	var goLF model.GoLF

	goLF.Logger = logger.NewCustomLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	goLF.Metrics = metrics.NewMetricsServer()

	initializeDatabases(&goLF)

	goLF.CLI = cmd.NewCMD()
	goLF.CLI.ToolName = goLF.Config.Get("CMD_TOOL_NAME")
	if goLF.CLI.ToolName == "" {
		goLF.CLI.ToolName = "myTool"
	}
	goLF.CLI.Version = goLF.Config.Get("CMD_TOOL_VERSION")
	if goLF.CLI.Version == "" {
		goLF.CLI.Version = "0.0.0"
	}

	return goLF
}
