package main

import "fmt"

const CLI_TAB_SPACE string = "    "

func help() {
	fmt.Println("WorkspaceGo is a tool for moving to and executing a programming project.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(CLI_TAB_SPACE, "wkgo.exe <command> <argument>")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println()
	fmt.Println(CLI_TAB_SPACE, "root                 ", CLI_TAB_SPACE, "Sets the projects root folder.")
	fmt.Println(CLI_TAB_SPACE, "Ex:                  ", CLI_TAB_SPACE, `wkgo.exe root C:\Users\PersonalProjects`)
	fmt.Println()
	fmt.Println(CLI_TAB_SPACE, `cmd <alias> "<command>"`, CLI_TAB_SPACE, "Adds an starting command alias.")
	fmt.Println(CLI_TAB_SPACE, "Ex:                  ", CLI_TAB_SPACE, `wkgo.exe cmd astro "pnpm run dev"`)
	fmt.Println()
	fmt.Println(CLI_TAB_SPACE, "<folder>    <alias>", CLI_TAB_SPACE, "Name of the project folder and command alias to start.")
	fmt.Println(CLI_TAB_SPACE, "Ex:                  ", CLI_TAB_SPACE, "wkgo.exe myblog astro")
	fmt.Println()
	fmt.Println(CLI_TAB_SPACE, "config               ", CLI_TAB_SPACE, "Shows the full configuration.")
}
