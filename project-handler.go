package main

import (
	"net/http"
	"encoding/json"
	"strings"
	//"fmt"
)

type Project struct{
	Name	string `json:"name"`
	Directory string `json:"directory"`
}
func (*Project) projectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s := strings.Split(r.URL.Path,"/")
		if len(s) > 3 {
			key := strings.Trim(s[3], " ")
			switch key {
			case "list":
				p, ok := str.getProjects()
				var proj []Project
				for _, v := range p {
					proj = append(proj, v)
				}
				j, _ := json.Marshal(proj)
				w.Header().Set("Content-Type","Application/Json")
				if ok {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
				w.Write(j)
			 default:
				proj,ok := str.getProject(key)
				j, _ := json.Marshal(proj)
				w.Header().Set("Content-Type","Application/Json")
				if ok {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
				w.Write(j)
			}
		}
	case "POST":
		s := strings.Split(r.URL.Path,"/")
		if len(s) > 3 {
			key := strings.Trim(s[3], " ")
			var prj Project
			prj.Name = key
			ok := str.addProject(key, prj)
			j, _ := json.Marshal(prj)
			w.Header().Set("Content-Type","Application/Json")
			if ok {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusFound)
			}
			w.Write(j)
			
		}		
	}
}
