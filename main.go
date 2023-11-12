package main

import (
	"os"
)

func main() {
	// Check if the folder name argument is provided
	if len(os.Args) < 2 {
		help()
		return
	}

	// Load configurations to cache
	getConfig()

	// Get the folder name from the command-line argument
	command := os.Args[1]

	switch command {
	case "help":
		{
			help()
		}
	case "cmd":
		{
			saveConfig(os.Args[2], os.Args[3])
		}
	case "config":
		{
			showConfig()
		}
	case "root":
		{
			setProjectsRoot(os.Args[2])
		}
	default:
		{
			folderName := command
			findProjectAndExecute(folderName, os.Args[2])
		}
	}
}
