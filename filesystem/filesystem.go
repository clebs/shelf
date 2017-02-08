package filesystem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/clebs/shelf/errorout"
	"github.com/clebs/shelf/project"
)

const shelfDir = ".shelf"

//Save persists the provided object in JSON format into the filesystem under "{user}/.shelf/{pID}.json"
func Save(p project.Project, pID string) {
	data, err := json.MarshalIndent(p, "", "    ")
	errorout.ErrQuit(err, "Could not serialize project.")

	err = ioutil.WriteFile(getProjectFilePath(pID), data, 0644)
	errorout.ErrQuit(err, "Could not add Project to the shelf")
}

//Load locates a project on the filesystem by its name and loads its data
func Load(name string) project.Project {
	filePath := getProjectFilePath(name)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		errorout.ErrQuit(err, fmt.Sprintf("The project %s could not be found.", name))
	}

	projectData, err := ioutil.ReadFile(filePath)
	errorout.ErrQuit(err, fmt.Sprintf("Could not load project %s", name))

	var p project.Project
	err = json.Unmarshal(projectData, &p)
	errorout.ErrQuit(err, "Error parsing project data.")
	return p
}

func getProjectFilePath(name string) string {
	user, err := user.Current()
	errorout.ErrQuit(err, "Could not obtain user home directory.")

	projectsFolder := user.HomeDir + "/" + shelfDir
	if _, err := os.Stat(projectsFolder); os.IsNotExist(err) {
		err := os.Mkdir(projectsFolder, 0700)
		errorout.ErrQuit(err, "Could not create Project bot directory")
	}
	return projectsFolder + "/" + name + ".json"
}
