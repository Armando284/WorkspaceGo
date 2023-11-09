# WORKSPACE GO

A golang program to open and run your workspace in just one terminal command

Never happened to you that every time you are going to start working on your project you need to open your terminal cd into your folder, open your editor and run your program? This monotonous task can be easily automated, saving you a few moments and making your start faster.

## What it does

1. Right now finds projects on a `E:\Trabajos` folder
2. Opens the folder in Visual Studio Code
3. Runs the project `pnpm run dev`

## How to use

1. Build the program
   1. Go into the project folder and execute in a terminal `go build`, you must have golang installed for this
2. Go to your terminal and execute the program passing the name of your project folder as argument. Ex: `workspace-go.exe my-project-name`
   1. Recommended to use an alias for easier use
   2. To make a powershell alias you can run this command `New-Alias -Name MyAlias -Value "C:\Path\To\Program.exe"`
   3. Then to use it would just be `MyAlias my-project-name`

### TODO

1. Add support for multiple starting scripts
2. Add support for multiple root folders
