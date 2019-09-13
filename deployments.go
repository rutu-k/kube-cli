package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listDeployments(c *cli.Context) {
	fmt.Println("listing running Deployments")
	clientset := getKubeHandle()

	deps, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		fatal(fmt.Sprintf("error getting list of deployments: %v", err))
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Deployments", "Namespace"})

	fmt.Println("**- Currently Running Pods -**")
	for _, dep := range deps.Items {
		data := []string{dep.Name, dep.Namespace}
		table.Append(data)
	}
	table.Render()
}
