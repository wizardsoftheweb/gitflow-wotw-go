package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var (
	repo Repository
)

func BootstrapCli() *cli.App {
	app := cli.NewApp()
	app.Name = "git-flow"
	app.Compiled = time.Now()
	app.Commands = []cli.Command{
		CommandInit,
		CommandFeature,
		CommandHotfix,
		CommandRelease,
		CommandSupport,
		CommandVersion,
	}
	app.Flags = []cli.Flag{}
	return app
}

func main() {
	app := BootstrapCli()
	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}
