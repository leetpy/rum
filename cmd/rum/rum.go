package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"

	rum "github.com/leetpy/rum/internal"
	"github.com/leetpy/rum/manager"
	"gopkg.in/op/go-logging.v1"
	"github.com/gin-gonic/gin"
)

var logger = logging.MustGetLogger("rum")

func startManager(c *cli.Context) error {
	rum.InitLogger(c.Bool("verbose"), c.Bool("debug"), c.Bool("with-systemd"))
	cfg, err := manager.LoadConfig(c.String("config"), c)
	if err != nil {
		logger.Errorf("Error loading config: %s", err.Error())
		os.Exit(1)
	}
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "rum"
	app.Usage = "rum tool"
	app.EnableBashCompletion = true
	app.Version = rum.Version
	app.Commands = []cli.Command{
		{
			Name: "manager",
			Aliases: []string{"m"},
			Action: startManager,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "config, c",
					Usage: "Load manager configurations from `FILE`",
				},
			},
		},
	}
	app.Run(os.Args)
}
