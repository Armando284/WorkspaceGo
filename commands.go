package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const ROOT_KEY string = "ProjectsRoot"

var projectsRootPath string = ""
var config map[string]string = make(map[string]string)

func setProjectsRoot(rootDir string) {
	if rootDir == "" {
		log.Fatal("Invalid empty path for projects root directory!")
	}
	saveConfig(ROOT_KEY, rootDir)
}

func showConfig() {
	configs := getConfig()
	fmt.Println("{")
	for key, value := range configs {
		fmt.Printf(`    %s: "%s",%s`, key, value, "\n")
	}
	fmt.Println("}")
}

func getProjectsRoot() {
	projectsRootPath = getConfigValue(ROOT_KEY)
}

func findProjectAndExecute(folderName string, alias string) {
	if projectsRootPath == "" {
		getProjectsRoot()
	}
	// Traverse the file system to find the folder
	var targetFolderPath string
	err := filepath.Walk(projectsRootPath, func(path string, info os.FileInfo, err error) error {
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
	fmt.Println("path:", targetFolderPath)
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

	if alias == "" {
		return
	}

	commands := strings.Split(getConfigValue(alias), " ")
	fmt.Println(commands[0], commands[1:])
	// Run "pnpm run dev" command
	cmd = exec.Command(commands[0], commands[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run %s: %v", commands, err)
	}
}
