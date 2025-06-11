package main

import (
	"flag"
	"fmt"
	"os"

	"wild-cli/internal/cli"
	"wild-cli/internal/daemon"
)

func main() {
	var uiMode bool
	flag.BoolVar(&uiMode, "ui", false, "start in daemon/web UI mode")
	flag.Parse()

	if uiMode {
		fmt.Println("Starting wild in daemon/web UI mode...")
		if err := daemon.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting daemon: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Starting wild in CLI mode...")
		if err := cli.Run(flag.Args()); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}