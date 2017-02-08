package core

import (
	"os/exec"

	"github.com/clebs/shelf/project"
)

// StartIDECmd starts the given IDE by running its IDE.Exec and opens the workspace IDE.Workspace
func StartIDECmd(ide project.IDE) *exec.Cmd {
	return exec.Command("open", ide.Exec)

}

// UpdateRepositoriesCmd updates all repositories in the project with git pull
func UpdateRepositoriesCmd(repos ...string) []*exec.Cmd {
	var cmds []*exec.Cmd
	for _, repo := range repos {
		cmd := exec.Command("git", "pull")
		cmd.Dir = repo
		cmds = append(cmds, cmd)
	}
	return cmds
}

//StartServer Runs the command for the server on the path given
func StartServer(serv project.Server) *exec.Cmd {
	cmd := exec.Command(serv.Cmd[0])
	cmd.Dir = serv.Path
	cmd.Args = serv.Cmd
	return cmd
}
