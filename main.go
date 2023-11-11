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

	// TODO: Set executable file in system env path
	// file, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file = fmt.Sprintf(`%s\wkgo.exe`, file)
	// fmt.Printf(file)

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
			getConfig()
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
