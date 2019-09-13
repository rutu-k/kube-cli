package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "kube-cli"
	app.Usage = "kube-statistics"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{Name: "podz", Usage: "Run pod-list", Action: listPodz},
		{Name: "deployments", Usage: "Run deployments-list", Action: listDeployments},
	}

	app.Run(os.Args)
}
