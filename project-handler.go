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

func (p *ProjectHandler) ProjectMainHandler(w http.ResponseWriter, r *http.Request) {
	if webutil.PathParse(r.URL.Path, "/api/project/") == nil {
		if r.Method == "POST" {
			//fmt.Println("in ProjectMainHandler")
			p.postProjectHandler(w, r, "")
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
		case "POST":
			p.postProjectHandler(w, r, id)
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

func (p *ProjectHandler) postProjectHandler(w http.ResponseWriter, r *http.Request, name string) {
	prj, ok := p.store.postProject(name)

	j, _ := json.Marshal(prj)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func (p *ProjectHandler) ListProjectsHandler(w http.ResponseWriter, r *http.Request) {

	prjList, ok := p.store.getProjects()
	j, _ := json.Marshal(prjList)
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
	prjList, ok := p.store.deleteProject(id)
	j, _ := json.Marshal(prjList)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}
