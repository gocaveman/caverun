package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"
	"runtime"
	"time"
	//"encoding/json"
	"strings"
	
	"github.com/gocaveman/caverun/files"
	
	"github.com/zserge/webview"
)
var str *store

func main() {
	str = new(store)
	str.init()

	headless := flag.Bool("headless", false, "Don't launch webview, just listen")
	portSpec := flag.String("port", "", "Port to listen on for HTTP (default is random)")
	debug := flag.Bool("debug", false, "Enable browser debugging")
	flag.Parse()

	pspec := *portSpec
	if pspec == "" {
		pspec = "0"
	}

	ln, err := net.Listen("tcp", "127.0.0.1:"+pspec)
	if err != nil {
		log.Fatal(err)
	}

	_, port, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Handler: &MainHandler{},
	}
	go func() {
		err := s.Serve(ln)
		if err != nil {
			log.Printf("Error from HTTP server: %v", err)
		}
	}()

	mainURL := fmt.Sprintf("http://127.0.0.1:%s/index.html", port)
	if *debug || *headless {
		log.Printf("Main URL: %s", mainURL)
	}
	if *debug {
		// https://github.com/zserge/webview#debugging-and-development-tips
		if runtime.GOOS == "windows" {
			mainURL += "#firebug" // TODO: can we use WebView.Bind() to expose settings instead?
		}
	}

	// endless loop until Ctrl+C if headless
	if *headless {
		for {
			time.Sleep(time.Second)
		}
	}

	settings := webview.Settings{
		URL:       mainURL,
		Title:     "caverun",
		Width:     1024,
		Height:    660,
		Resizable: true,
		Debug:     *debug,
	}

	w := webview.New(settings)
	w.Run() // blocks forever

}

type MainHandler struct{}

func (ws *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	// TODO: looks like all the browsers in question support websockets,
	// which is probably going to be the simplest way to message back
	// and forth between Go and the UI, rather than messing around
	// with fetch() polyfills and maybe SSE for pushing events, or
	// doing the webview.Bind() approach and hoping it handles complex
	// data without issues (and still doesn't provide a way to
	// push to the browser).
	
	// TODO: we'll also need to secure the endpoint somehow, probably would
	// work to generate a random key and Bind() that and require it
	// in each request as a security precaution.  Doesn't prevent sniffing
	// but that's probably fine, still prevents unwanted apps from having
	// direct access.
	
	if  strings.Contains(r.URL.Path, "api/project") {
		var p Project
		p.projectHandler(w,r)
		return
	}

	fpath := path.Join("/static", path.Clean("/"+r.URL.Path))

	f, err := files.EmbeddedAssets.Open(fpath)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	fst, err := f.Stat()
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}

	http.ServeContent(w, r, path.Base(fpath), fst.ModTime(), f)
}
