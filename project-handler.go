package main

import (
	"encoding/json"
	"fmt"
	webutil "github.com/gocaveman/caveman/webutil"
	"net/http"
)

type ProjectHandler struct {
	store *store
}

func (p *ProjectHandler) newStore() {
	p.store = new(store)
	p.store.init()
}

func (p *ProjectHandler) ProjectMainHandler(w http.ResponseWriter, r *http.Request) {
	if webutil.PathParse(r.URL.Path, "/api/project/") == nil {
		if r.Method == "POST" {
			//fmt.Println("in ProjectMainHandler")
			p.postProjectHandler(w, r)
		}
		return
	}

	if webutil.PathParse(r.URL.Path, "/api/project/list") == nil {
		if r.Method == "GET" {
			p.ListProjectsHandler(w, r)
		}
		return
	}

	var id string
	if webutil.PathParse(r.URL.Path, "/api/project/%s", &id) == nil {
		switch r.Method {
		case "GET":
			p.getProjectHandler(w, r, id)
		case "PUT":
			p.putProjectHandler(w, r, id)
		case "DELETE":
			p.deleteProjectHandler(w, r, id)
		}
		return
	}
}

func (p *ProjectHandler) postProjectHandler(w http.ResponseWriter, r *http.Request) {
	type projectID struct {
		ID string `json:"ID"`
	}
	var pID projectID
	//fmt.Println("in postProjectHandler")
	ID, ok := p.store.postProject()
	if ok {
		pID.ID = ID
	} else {
		pID.ID = ""
	}

	j, _ := json.Marshal(pID)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func (p *ProjectHandler) ListProjectsHandler(w http.ResponseWriter, r *http.Request) {
	prj, ok := p.store.getProjects()

	j, _ := json.Marshal(prj)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func (p *ProjectHandler) getProjectHandler(w http.ResponseWriter, r *http.Request, id string) {
	proj, ok := p.store.getProject(id)
	j, _ := json.Marshal(proj)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func (p *ProjectHandler) putProjectHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("in putProjectHandler", id)
}

func (p *ProjectHandler) deleteProjectHandler(w http.ResponseWriter, r *http.Request, id string) {

	type projectID struct {
		ID     string `json:"ID"`
		Status string `json:"status`
	}
	var pID projectID

	ok := p.store.deleteProject(id)
	pID.ID = id
	if ok {
		pID.Status = "Succeeded"
	} else {
		pID.Status = "Failed"
	}

	j, _ := json.Marshal(pID)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}
