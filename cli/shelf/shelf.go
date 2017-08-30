//shelf is a nice servant that helps launching your favorite projects!!
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clebs/gobatch"
	"github.com/clebs/shelf/commands"
	"github.com/clebs/shelf/filesystem"
	"github.com/clebs/shelf/project"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		help()
	} else {
		switch args[0] {
		case "add":
			add()
		case "pick":
			if len(args) < 2 {
				help("pick")
				os.Exit(2)
			}
			pick(args[1])
		case "delete":
			if len(args) < 2 {
				help("delete")
				os.Exit(2)
			}
			delete(args[1])
		case "help":
			if len(args) < 2 {
				help()
			} else {
				help(args[1])
			}
		default:
			help(args[0])
		}
	}
}

//Adds a new project to the shelf
func add() {
	p := project.Project{
		Name: ask("Project name: (use it later to pick it up)"),
		IDE: project.IDE{
			Exec:      ask("IDE executable:"),
			Workspace: ask("IDE Workspace:")},
		Repos: strings.Split(ask("Repositories:"), ","),
		Server: project.Server{
			Path: ask("Server location:"),
			Cmd:  strings.Split(ask("Server command:"), " ")}}

	filesystem.Save(p, p.Name)

	println("Your project has been added:")
	fmt.Println(p)
}

//pick allows to pick up a project from the shelf by its name
func pick(name string) {
	p := filesystem.Load(name)

	println("Loading project:")
	fmt.Println(p)

	var serverBoot gobatch.SyncRunner
	serverBoot.Add(commands.UpdateRepositories(p.Server.Path)...)
	serverBoot.Add(commands.StartServer(p.Server))

	var projectBoot gobatch.AsyncRunner
	projectBoot.Add(commands.StartIDE(p.IDE))
	projectBoot.Add(commands.UpdateRepositories(p.Repos...)...)
	projectBoot.Add(serverBoot)
	projectBoot.Run()
	fmt.Printf("\nProject %s loaded.\n", name)
}

func delete(name string) {
	filesystem.Delete(name)
}

func help(args ...string) {
	if len(args) == 0 {
		println("Shelf is a project bootup storage. Store IDE/editor, repository and server launch configs to automatically open a working environment.\n")
		println("Usage:")
		println("    shelf command [arguments]\n")
		println("The commands are:")
		println("    add     add a new project")
		println("    pick    picks up a project from the shelf")
		println("    delete  removes a project from the shelf")
		return
	}

	switch args[0] {
	case "add":
		println("add     add a new project entering the following data:")
		println("    Project name:      name of the project, used later to find it on pick command")
		println("    IDE executable:    absolute path to the executable of your IDE, no args accepted")
		println("    IDE Workspace:     (optional) location of the workspace to open the IDE in")
		println("    Repositories:      list of comma separated repositories where the code is.")
		println("    Server location:   path to the folder where the local server is")
		println("    Server command:    command to run inside the server folder, accepts args (e.g. mvn clean install)")
	case "pick":
		println("pick     load a project previously added with the add command.")
		println("usage:   shelf pick [project name]")
	case "delete":
		println("delete   remove a project from the shelf.")
		println("usage:   shelf delete [project name]")
	default:
		println("Invalid arguments. Please run \"shelf help [command]\" to see a list of available arguments.")
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func ask(question string) (answer string) {
	println(question)
	scanner.Scan()
	return scanner.Text()
}
