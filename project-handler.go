package main

import (
	"encoding/json"
	"net/http"

	// "strings"
	// "io"
	"fmt"
)

func createProjectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in createProjectHandler")

}

func ListProjectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in ListProjectsHandler")
	p, ok := str.getProjects()
	var proj []Project
	for _, v := range p {
		proj = append(proj, v)
	}
	j, _ := json.Marshal(proj)
	w.Header().Set("Content-Type", "Application/Json")
	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(j)
}

func getProjectHandler(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("in getProjectHandler", id)
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
	fmt.Println("in deleteProjectHandler", id)
}

// func projectHandler(w http.ResponseWriter, r *http.Request) {
// 	var id string
// 	PathParse(r.URL.Path, "/api/project/%s", &id)
// 	fmt.Println("%+v\n", id)
// 	switch r.Method {
// 	case "GET":
// 		s := strings.Split(r.URL.Path,"/")
// 		if len(s) > 3 {
// 			key := strings.Trim(s[3], " ")
// 			switch key {
// 			case "list":
// 				p, ok := str.getProjects()
// 				var proj []Project
// 				for _, v := range p {
// 					proj = append(proj, v)
// 				}
// 				j, _ := json.Marshal(proj)
// 				w.Header().Set("Content-Type","Application/Json")
// 				if ok {
// 					w.WriteHeader(http.StatusOK)
// 				} else {
// 					w.WriteHeader(http.StatusNotFound)
// 				}
// 				w.Write(j)
// 			 default:
// 				proj,ok := str.getProject(key)
// 				j, _ := json.Marshal(proj)
// 				w.Header().Set("Content-Type","Application/Json")
// 				if ok {
// 					w.WriteHeader(http.StatusOK)
// 				} else {
// 					w.WriteHeader(http.StatusNotFound)
// 				}
// 				w.Write(j)
// 			}
// 		}
// 	case "POST":
// 		s := strings.Split(r.URL.Path,"/")
// 		if len(s) > 3 {
// 			key := strings.Trim(s[3], " ")
// 			var prj Project
// 			prj.Name = key
// 			ok := str.addProject(key, prj)
// 			j, _ := json.Marshal(prj)
// 			w.Header().Set("Content-Type","Application/Json")
// 			if ok {
// 				w.WriteHeader(http.StatusOK)
// 			} else {
// 				w.WriteHeader(http.StatusFound)
// 			}
// 			w.Write(j)

// 		}
// 	}
// }

// func (*Project) projectHandler(w http.ResponseWriter, r *http.Request) {
// 	var id string
// 	PathParse(r.URL.Path, "/api/project/%s", &id)
// 	fmt.Println("%+v\n", id)
// 	switch r.Method {
// 	case "GET":
// 		s := strings.Split(r.URL.Path,"/")
// 		if len(s) > 3 {
// 			key := strings.Trim(s[3], " ")
// 			switch key {
// 			case "list":
// 				p, ok := str.getProjects()
// 				var proj []Project
// 				for _, v := range p {
// 					proj = append(proj, v)
// 				}
// 				j, _ := json.Marshal(proj)
// 				w.Header().Set("Content-Type","Application/Json")
// 				if ok {
// 					w.WriteHeader(http.StatusOK)
// 				} else {
// 					w.WriteHeader(http.StatusNotFound)
// 				}
// 				w.Write(j)
// 			 default:
// 				proj,ok := str.getProject(key)
// 				j, _ := json.Marshal(proj)
// 				w.Header().Set("Content-Type","Application/Json")
// 				if ok {
// 					w.WriteHeader(http.StatusOK)
// 				} else {
// 					w.WriteHeader(http.StatusNotFound)
// 				}
// 				w.Write(j)
// 			}
// 		}
// 	case "POST":
// 		s := strings.Split(r.URL.Path,"/")
// 		if len(s) > 3 {
// 			key := strings.Trim(s[3], " ")
// 			var prj Project
// 			prj.Name = key
// 			ok := str.addProject(key, prj)
// 			j, _ := json.Marshal(prj)
// 			w.Header().Set("Content-Type","Application/Json")
// 			if ok {
// 				w.WriteHeader(http.StatusOK)
// 			} else {
// 				w.WriteHeader(http.StatusFound)
// 			}
// 			w.Write(j)

// 		}
// 	}
// }
