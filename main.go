package main

import (
	"fmt"

	"github.com/TRON-US/soter-demo/conf"
	"github.com/TRON-US/soter-demo/log"
	"github.com/TRON-US/soter-demo/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "BTFS soter demo server config file path")
)

func main() {
	pflag.Parse()

	// init config via viper
	if err := conf.InitConfig(*cfg); err != nil {
		panic(err)
	}

	// create the gin engine
	g := gin.Default()

	// load routes
	router.Load(g)

	// start controller server
	log.Logger().Info(fmt.Sprintf("soter demo server is listening on port %s ...", viper.GetString("server_port")))
	if err := g.Run(viper.GetString("server_port")); err != nil {
		log.Logger().Fatal(fmt.Sprintf("soter demo server runs error: %s", err))
	}
	log.Logger().Info("soter demo server quit")
}
