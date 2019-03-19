package main

import (
	gouuid "github.com/bradleypeabody/gouuidv6"
)

type Project struct {
	Name      string `json:"name"`
	Directory string `json:"directory"`
}

type store struct {
	Projects map[string]Project
}

func (s *store) newStore() *store {
	str := new(store)
	return str
}

//TODO - Add getRecentProjects, saveProject
func (s *store) init() {

	s.Projects = make(map[string]Project)

	var prj Project

	for i := 1; i < 5; i++ {

		ID := gouuid.NewB64().String()
		prj.Name = "Name for: " + ID
		prj.Directory = "Directory for: " + ID
		str.Projects[ID] = prj
	}

}

// Gets an open projects struct
func (s *store) getProject(ID string) (prj Project, ok bool) {
	prj, ok = s.Projects[ID]
	return prj, ok
}

// Gets a map of open projects
func (s *store) getProjects() (prj map[string]Project, ok bool) {
	return s.Projects, true
}

// Creates a new project in memory
func (s *store) postProject() (ID string, ok bool) {
	var prj Project
	ID = gouuid.NewB64().String()
	prj.Name = "Name for: " + ID
	prj.Directory = "Directory for: " + ID
	str.Projects[ID] = prj
	_, ok = str.Projects[ID]
	return ID, ok
}

// Deletes an open projects from memory
func (s *store) deleteProject(ID string) (ok bool) {
	delete(str.Projects, ID)
	_, ok = str.Projects[ID]
	ok = !ok

	return ok
}

//Saves a project to disk -- Dummy for now
func (s *store) putProject(ID string) (ok bool) {
	ok = true
	return ok
}
