// Package commands provides constructors for all needed commands for the shelf.
package commands

import (
	"os"
	"os/exec"

	"github.com/clebs/gobatch"
	"github.com/clebs/shelf/project"
)

// StartIDE starts the given IDE by running its IDE.Exec and opens the workspace IDE.Workspace
func StartIDE(ide project.IDE) gobatch.Runner {
	return &gobatch.CommandRunner{Command: exec.Command("open", ide.Exec), Output: os.Stdout}

}

// UpdateRepositories updates all repositories in the project with git pull
func UpdateRepositories(repos ...string) []gobatch.Runner {
	var cmdRunners []gobatch.Runner
	for _, repo := range repos {
		cmd := exec.Command("git", "pull")
		cmd.Dir = repo
		cmdRunners = append(cmdRunners, &gobatch.CommandRunner{Command: cmd, Output: os.Stdout})
	}
	return cmdRunners
}

//StartServer Runs the command for the server on the path given
func StartServer(serv project.Server) gobatch.Runner {
	cmd := exec.Command(serv.Cmd[0])
	cmd.Dir = serv.Path
	cmd.Args = serv.Cmd
	return &gobatch.CommandRunner{Command: cmd, Output: os.Stdout}
}
