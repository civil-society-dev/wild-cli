package cli

import (
	"fmt"
)

func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	command := args[0]
	switch command {
	case "help":
		showHelp()
	case "version":
		showVersion()
	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
}

func showHelp() {
	fmt.Println("wild - A powerful CLI tool")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  wild [command]")
	fmt.Println("  wild --ui    # Start in daemon/web UI mode")
	fmt.Println("")
	fmt.Println("Available Commands:")
	fmt.Println("  help         Show this help message")
	fmt.Println("  version      Show version information")
}

func showVersion() {
	fmt.Println("wild version 0.1.0")
}