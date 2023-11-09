package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the folder name argument is provided
	if len(os.Args) < 2 {
		log.Fatal("Folder name argument is missing.")
	}

	// Get the folder name from the command-line argument
	folderName := os.Args[1]

	// Traverse the file system to find the folder
	var targetFolderPath string
	err := filepath.Walk(`E:\Trabajos`, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current entry is a directory and matches the folder name
		if info.IsDir() && strings.EqualFold(info.Name(), folderName) {
			targetFolderPath = path
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		log.Fatalf("Failed to traverse the file system: %v", err)
	}

	// Check if the target folder was found
	if targetFolderPath == "" {
		log.Fatalf("Folder '%s' not found.", folderName)
	}

	// Change directory to the target folder
	err = os.Chdir(targetFolderPath)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	// Open VSCode at the target folder location
	cmd := exec.Command("code", targetFolderPath)
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Failed to open VSCode: %v", err)
	}

	// Run "pnpm run dev" command
	cmd = exec.Command("pnpm", "run", "dev")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run 'pnpm run dev': %v", err)
	}
}
