//Package project provides types and functions to represent project data in various ways
package project

import "fmt"

// Project holds all data relevant to boot or add a project and can be represented as JSON
type Project struct {
	Name   string
	IDE    IDE
	Repos  []string
	Server Server
}

// IDE represents the information necessary to open an IDE.
type IDE struct {
	Exec      string
	Workspace string
}

// Server contains the data necessary to boot a local server
type Server struct {
	Path string
	Cmd  []string // cmd + args
}

func (p Project) String() string {
	return fmt.Sprintf("Name: %s\nIDE: %s\nRepositories: %s\nServer: %s\n", p.Name, p.IDE, p.Repos, p.Server)
}
