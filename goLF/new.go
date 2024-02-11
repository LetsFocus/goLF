package goLF

import (
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/cronjobs"
	"github.com/LetsFocus/goLF/database"
	"github.com/LetsFocus/goLF/elasticstack"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/metrics"
	"github.com/LetsFocus/goLF/slogs"
)

func New() model.GoLF {
	defer func() {
		if r := recover(); r != nil {
			//add some log here
		}
	}()

	var goLF model.GoLF

	goLF.Logger = slogs.NewLogger()
	goLF.Config = configs.NewConfig(goLF.Logger)

	database.InitializeDB(&goLF, "")
	database.InitializeRedis(&goLF, "")
	elasticstack.InitializeES(&goLF, "")

	cron := cronjobs.NewCronManager()
	cron.AddJob("* * * * *", func() {
		if goLF.Config.Get("CONFIG_REFRESH") == "true" {
			configs.
		}
	})

	cron.Start()

	defer cron.Stop()

	goLF.Metrics = metrics.NewMetricsServer()
	return goLF
}
