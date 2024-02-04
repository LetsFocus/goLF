package goLF

import (
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

func New() model.GoLF {
	var goLF model.GoLF

	goLF.Logger = logger.NewCustomLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	database.InitializeDB(&goLF, "")
	return goLF
}
