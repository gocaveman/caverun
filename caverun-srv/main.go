package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

var q = flag.Bool("q", false, "Set quiet mode, minimal console output.")

func main() {
	startTime := time.Now()

	publicdir := flag.String("publicdir", "./public/", "Where to server content from (supports .asar)")
	listen := flag.String("listen", ":6753", "IP:Port to listen on for HTTP server")

	flag.Parse()

	// automatically looks for .asar file and opens it, otherwise supports normal dir
	fileSystem, err := NewAsarFileSystem(*publicdir)
	if err != nil {
		log.Fatal(err)
	}

	var mainHandler http.HandlerFunc

	mainHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Cache-Control", "no-store")

		p := r.URL.Path
		p = path.Clean("/" + p)

		if p == "/api/status.action" {
			w.Header().Set("content-type", "text/plain")
			fmt.Fprintf(w, "publicdir: %q\n", *publicdir)
			fmt.Fprintf(w, "listen: %q\n", *listen)
			fmt.Fprintf(w, "uptime: %v\n", time.Since(startTime))
			return
		}

		hf, err := fileSystem.Open(p)
		if err != nil {
			log.Printf("error opening file (%q): %v", p, err)
			http.NotFound(w, r)
			return
		}
		defer hf.Close()

		// do template rendering for .html files
		if strings.HasSuffix(p, ".html") {

			b, err := ioutil.ReadAll(hf)
			if err != nil {
				log.Printf("error reading file (%q): %v", p, err)
				http.Error(w, "error reading file", 500)
				return
			}

			t, err := template.New(p).Parse(string(b))
			if err != nil {
				log.Printf("error parsing file (%q): %v", p, err)
				http.Error(w, "error parsing file", 500)
				return
			}

			err = TmplIncludeAll(fileSystem, t)
			if err != nil {
				log.Printf("error handling includes on file (%q): %v", p, err)
				http.Error(w, "error handling includes on file", 500)
				return
			}

			err = t.ExecuteTemplate(w, p, map[string]interface{}{
				"Request": r,
				"VERSION": VERSION,
			})
			if err != nil {
				log.Printf("error processing template on file (%q): %v", p, err)
				http.Error(w, "error processing template on file", 500)
				return
			}

			return

		}

		st, err := hf.Stat()
		if err != nil {
			log.Printf("error stat'ing file (%q): %v", p, err)
			http.NotFound(w, r)
			return
		}

		http.ServeContent(w, r, p, st.ModTime(), hf)

		return

	})

	http.HandleFunc("/", mainHandler)

	if !*q {
		log.Printf("Listening at: %q", *listen)
	}
	log.Fatal(http.ListenAndServe(*listen, nil))

}
