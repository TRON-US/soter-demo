package main

import (
	"fmt"

	"github.com/TRON-US/soter-demo/core"
	"github.com/TRON-US/soter-demo/log"
	"github.com/TRON-US/soter-demo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "soter controller server config file path")
)

func main() {
	pflag.Parse()

	// init config via viper
	if err := core.InitConfig(*cfg); err != nil {
		panic(err)
	}

	// create the gin engine
	g := gin.Default()

	// load routes
	router.Load(g)

	// start controller server
	log.Logger().Info(fmt.Sprintf("controller server is listening on port %s ...", viper.GetString("server_port")))
	if err := g.Run(viper.GetString("server_port")); err != nil {
		log.Logger().Fatal(fmt.Sprintf("controller server runs error: %s", err))
	}
	log.Logger().Info("controller server quit")
}
