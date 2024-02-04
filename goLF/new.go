package goLF

import (
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/slogs"
)

func New() model.GoLF {
	var goLF model.GoLF

	goLF.Logger = slogs.NewLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	database.InitializeDB(&goLF, "")
	return goLF
}
