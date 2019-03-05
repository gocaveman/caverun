package main

type store struct {
	Projects map[string]Project
}

func (s *store) newStore() *store {
	str := new(store)
	return str
}
//Todo - Find out if there is a possibility of concurrent r/w from the store and implement locking as needed.
func (s *store) init() {
	
	s.Projects = make(map[string]Project)
	
	var prj Project
	prj.Name = "project1"
	prj.Directory = "Directory1"
	str.Projects["project1"]=prj

	prj.Name = "project2"
	prj.Directory = "Directory2"
	str.Projects["project2"]=prj
}

func (s *store) getProject(name string) (prj Project, ok bool) {
	prj, ok = s.Projects[name]
	return prj, ok
}

func (s *store) getProjects() (prj map[string]Project, ok bool) {
	return s.Projects, true
}


func (s *store) addProject(name string, prj Project) (ok bool) {
	s.Projects[name]=prj
	return ok
}