package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/xoap-io/packer-plugin-dsc/provisioner/dsc"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(dsc.Provisioner))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		fmt.Fprintln(os.Stderr, "Attempting to run with internal arguments")
		errCommand := pps.RunCommand("start", "provisioner", plugin.DEFAULT_NAME)
		if errCommand != nil {
			fmt.Fprintln(os.Stderr, errCommand.Error())
			os.Exit(1)
		}
	}
}
