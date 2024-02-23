package goLF

import (
	"fmt"
	"github.com/LetsFocus/goLF/configs"
	"github.com/LetsFocus/goLF/cronjobs"
	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/metrics"
	"github.com/LetsFocus/goLF/slogs"
)

func New() model.GoLF {

	var goLF model.GoLF
	goLF.Logger = slogs.NewLogger()
	cron := cronjobs.NewCronManager()

	go func() {
		cron.AddJob("* * * * *", func() {
			goLF.Logger = slogs.NewLogger()
			goLF.Config = configs.NewConfig(goLF.Logger, "")
			fmt.Println("going")

			//database.InitializeDB(&goLF, "")
			//database.InitializeRedis(&goLF, "")
			//elasticstack.InitializeES(&goLF, "")

			goLF.Logger.Logger.Info("Log_level inside", "LEVEL", goLF.Config.Get("LOG_LEVEL"), "REFRESH", goLF.Config.Get("CONFIG_REFRESH"))
			if goLF.Config.Get("CONFIG_REFRESH") != "true" {
				cron.Stop()
			}

			goLF.Config.LoadConfigs(goLF.Logger, "")
		})
	}()

	cron.Start()

	defer cron.Stop()

	goLF.Metrics = metrics.NewMetricsServer()
	return goLF
}
