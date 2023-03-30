package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// =========================================================================
// App configuration

type config struct {
	listenAddress string
}

type application struct {
	config config
	logger *logrus.Logger
}

func main() {

	// =========================================================================
	// Init app configuration

	var cfg config

	var log = logrus.New()
	log.Out = os.Stdout

	app := &application{
		config: cfg,
		logger: log,
	}

	// =========================================================================
	// Setup CLI flags and config

	cliApp := &cli.App{
		Name:  "pod-metrics-exporter",
		Usage: "monitor all the things!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "listen-addr",
				Value:       ":8080",
				Usage:       "listen address of the server",
				Destination: &app.config.listenAddress,
			},
		},
		Action: app.serve,
	}

	// =========================================================================
	// Run the app

	err := cliApp.Run(os.Args)
	if err != nil {
		app.logger.Fatal(err)
	}
}
