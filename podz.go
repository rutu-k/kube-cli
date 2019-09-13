package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func listPodz(c *cli.Context) {
	fmt.Println("listing running Pods")
	clientset := getKubeHandle()

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fatal(fmt.Sprintf("error getting list of pods: %v", err))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pod Name", "Namespace"})

	fmt.Println("**- Currently Running Pods -**")
	for _, pod := range pods.Items {
		data := []string{pod.Name, pod.Namespace}
		table.Append(data)
	}
	table.Render()
}
