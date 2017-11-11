package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}
	notebook, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal(err)
	}
	ui := NewUI(window, notebook)

	ui.Window.SetTitle("Cave Runner")
	ui.Window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	ui.Window.Add(ui.WindowWidget())
	ui.Window.ShowAll()

	gtk.Main()
}
