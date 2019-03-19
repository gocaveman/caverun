package main

import (
	"encoding/json"
	"net/http"
	// "strings"
	// "io"
	"fmt"
)

func postProjectHandler(w http.ResponseWriter, r *http.Request) {
	type projectID struct {
		ID string `json:"ID"`
	}
	var pID projectID
	ID, ok := str.postProject()
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

func ListProjectsHandler(w http.ResponseWriter, r *http.Request) {
	p, ok := str.getProjects()

	j, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func getProjectHandler(w http.ResponseWriter, r *http.Request, id string) {
	proj, ok := str.getProject(id)
	j, _ := json.Marshal(proj)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func putProjectHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("in putProjectHandler", id)
}

func deleteProjectHandler(w http.ResponseWriter, r *http.Request, id string) {

	type projectID struct {
		ID     string `json:"ID"`
		Status string `json:"status`
	}
	var pID projectID

	ok := str.deleteProject(id)
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
