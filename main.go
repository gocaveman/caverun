package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	ui := NewUI()

	ui.Window.Add(ui.WindowWidget())
	ui.Window.ShowAll()
	gtk.Main()
}
