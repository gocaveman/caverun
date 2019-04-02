package main

import (
	gouuid "github.com/bradleypeabody/gouuidv6"
)

type Project struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	Directory string `json:"directory"`
}

type store struct {
	Projects map[string]Project
}

//TODO - Add getRecentProjects, saveProject

func (s *store) init() {

	s.Projects = make(map[string]Project)

	//var prj Project

	// for i := 1; i < 5; i++ {
	// 	ID := gouuid.NewB64().String()
	// 	prj.ID = ID
	// 	prj.Name = "Name for: " + ID
	// 	prj.Directory = "Directory for: " + ID
	// 	s.Projects[ID] = prj
	// }

}

// Gets an open projects struct
func (s *store) getProject(ID string) (prj Project, ok bool) {
	prj, ok = s.Projects[ID]
	return prj, ok
}

// Gets a map of open projects
func (s *store) getProjects() (prjList []Project, ok bool) {
	for _, v := range s.Projects {
		prjList = append(prjList, v)
	}
	return prjList, true
}

// Creates a new project in memory
func (s *store) postProject(name string) (prjList []Project, ok bool) {
	var prj Project
	prj.ID = gouuid.NewB64().String()
	if len(name) == 0 {
		prj.Name = "Name for: " + prj.ID
		prj.Directory = "Directory for: " + prj.ID
	} else {
		prj.Name = name
		prj.Directory = "Directory for: " + name

	}
	s.Projects[prj.ID] = prj
	_, ok = s.Projects[prj.ID]
	for _, v := range s.Projects {
		prjList = append(prjList, v)
	}
	return prjList, ok
}

// Deletes an open projects from memory
func (s *store) deleteProject(ID string) (prjList []Project, ok bool) {
	delete(s.Projects, ID)
	_, ok = s.Projects[ID]
	ok = !ok
	for _, v := range s.Projects {
		prjList = append(prjList, v)
	}
	return prjList, ok
}

//Saves a project to disk -- Dummy for now
func (s *store) putProject(ID string) (ok bool) {
	ok = true
	return ok
}
