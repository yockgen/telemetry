package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/docker/machine/libmachine/drivers/plugin"
)

var Version string = "latest"

func main() {
	version := flag.Bool("v", false, "prints current edge-iaas-node-driver version")
	flag.Parse()
	if *version {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	plugin.RegisterDriver(NewDriver("", ""))
}
