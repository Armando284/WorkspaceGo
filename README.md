# WORKSPACE GO

A cli tool, made with Go to open and run your workspace in just one terminal command

Never happened to you that every time you are going to start working on your project you need to open your terminal cd into your folder, open your editor and run your program? 🤯 This monotonous task can be easily automated, saving you a few moments and making your start faster. 👍

## What it does

1. Finds the project workspace folder you pass to it inside of your projects folder
2. Navigates automatically your terminal into that folder 💯
3. Opens Visual Studio Code in that folder
4. Runs an starting command set by you
5. Everything in just one line

## How to use

1. Download executable from `./output/wkgo.exe` or build the program yourself [build instructions](#how-to-build-the-project)
2. Recommended to add executable to the ***Systems Environment Variables: Path***
3. Open a terminal and test the program with the command `wkgo`
4. If everything is well a help menu should be displayed
5. Add your projects root folder. Ex: `wkgo root C:\User\MyProjects`
   1. This makes the program run faster since it doesn't have to search your whole computer, it also helps with permission issues
6. Once root is set you can config starting commands and aliases. Ex: `wkgo cmd astro "pnpm run dev"`
   1. You can add as many starting commands as you need 👍
7. Once you have everything set you can run your workspace from anywhere in your terminal.
8. Open the terminal and run `wkgo <project folder name> <start alias>`
9. That easy 👌

## How to build the project

1. Go into the project folder and execute in a terminal `go build -o wkgo.exe`, you must have golang installed for this
2. Open your terminal at the root of the project and execute the program. Ex: `.\wkgo.exe`
3. If everything went well you should see a help message
4. For submitting changes with a pull request use the utility `gp.sh`
5. The utility will test the module, build it and push the changes to the origin main
6. To use it don't forget to make it executable by opening a Power Shell terminal and running this command `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`
7. Then just use the utility after committing all your changes

### TODO

1. Add the program automatically to the system or user path
2. Add support for multiple editors
3. Add support for other Operating Systems
