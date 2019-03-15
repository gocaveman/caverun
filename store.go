package main

import (
	gouuid "github.com/bradleypeabody/gouuidv6"
)

type Project struct {
	ID        gouuid.UUID `json:"id"`
	Name      string      `json:"name"`
	Directory string      `json:"directory"`
}

type store struct {
	Projects map[string]Project
}

func (s *store) newStore() *store {
	str := new(store)
	return str
}

//Todo - Find out if there is a possibility of concurrent r/w from the store and implement locking as needed.
//TODO - Add getRecentProjects, saveProject, deleteProject(to remove from list of open objects)
func (s *store) init() {

	s.Projects = make(map[string]Project)

	var prj Project
	prj.ID = gouuid.New()
	prj.Name = "project1"
	prj.Directory = "Directory1"
	str.Projects["project1"] = prj

	prj.ID = gouuid.New()
	prj.Name = "project2"
	prj.Directory = "Directory2"
	str.Projects["project2"] = prj

	prj.ID = gouuid.New()
	prj.Name = "project3"
	prj.Directory = "Directory3"
	str.Projects["project3"] = prj

	prj.ID = gouuid.New()
	prj.Name = "project4"
	prj.Directory = "Directory4"
	str.Projects["project4"] = prj

}

// Gets an open projects struct
func (s *store) getProject(name string) (prj Project, ok bool) {
	prj, ok = s.Projects[name]
	return prj, ok
}

// Gets a map of open projects
func (s *store) getProjects() (prj map[string]Project, ok bool) {
	return s.Projects, true
}

// Creates a new project in memory
func (s *store) addProject(name string, prj Project) (ok bool) {
	s.Projects[name] = prj
	return ok
}

// Removes a project from memory
// func(s *store) deleteProject......

//Saves a project to disk
//func(s *store) saveProject
