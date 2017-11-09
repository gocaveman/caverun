package main

import (
	"container/list"
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var labelList = list.New()

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Cave Runner")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget())
	win.ShowAll()

	gtk.Main()
}

func windowWidget() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Just as a demonstration, we create and destroy a Label without ever
	// adding it to a container.  In native GTK, this would result in a
	// memory leak, since gtk_widget_destroy() will not deallocate any
	// memory when passed a GtkWidget with a floating reference.
	//
	// gotk3 handles this situation by always sinking floating references
	// of any struct type embedding a glib.InitiallyUnowned, and by setting
	// a finalizer to unreference the object when Go has lost scope of the
	// variable.  Due to this design, widgets may be allocated freely
	// without worrying about handling memory incorrectly.
	//
	// The following code is not entirely useful (except to demonstrate
	// this point), but it is also not "incorrect" as the C equivalent
	// would be.
	unused, err := gtk.LabelNew("This label is never used")
	if err != nil {
		// Calling Destroy() is also unnecessary in this case.  The
		// memory will still be freed with or without calling it.
		unused.Destroy()
	}

	notebook, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Unable to create notebook:", err)
	}

	toolbar, err := gtk.ToolbarNew()
	// toolbar.SetStyle(gtk.TOOLBAR_ICONS)
	btnnew, err := gtk.ToolButtonNew(nil, "new project")

	btnnew.SetTooltipText("Load a new project")
	btnclose, err := gtk.ToolButtonNew(nil, "close project")
	btnclose.SetTooltipText("Close active tab")
	separator, err := gtk.SeparatorToolItemNew()
	btnGlobalSettings, err := gtk.ToolButtonNew(nil, "Global Settings")
	btnGlobalSettings.SetTooltipText("Modify caverunner global settings")

	//GTK notebook
	// notebook, err := gtk.NotebookNew()

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btnGlobalSettings, -1)

	grid.Attach(toolbar, 0, 0, 2, 1)

	grid.Attach(notebook, 0, 1, 5, 4)

	// sw.SetHExpand(true)
	// sw.SetVExpand(true)

	labelsGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	labelsGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	buttonBuildRun, err := gtk.ButtonNewWithLabel("Build & Run")
	buttonGenerate, err := gtk.ButtonNewWithLabel("Generate")
	comboGenerate, err := gtk.ComboBoxTextNewWithEntry()
	buttonDeps, err := gtk.ButtonNewWithLabel("update deps")
	// space, err := gtk.LabelNew("                    ")

	tabGrid1, err := gtk.GridNew()

	tabGrid1.Attach(buttonBuildRun, -1, 1, 1, 1)
	tabGrid1.Attach(buttonGenerate, -1, 2, 1, 1)
	tabGrid1.Attach(comboGenerate, -2, 2, 1, 1)

	tabGrid1.Attach(buttonDeps, -1, 3, 1, 1)
	tabGrid1.SetRowSpacing(15)
	tabGrid1.SetMarginTop(20)
	tabGrid1.SetMarginBottom(20)

	tabGrid1.SetMarginStart(20)
	tabGrid1.SetMarginEnd(20)

	tabGrid1.SetHAlign(gtk.ALIGN_END)
	// tabGrid1.Attach(space, 1, 1, 3, 3)

	tabGrid2, err := gtk.GridNew()

	tabGrid3, err := gtk.GridNew()

	tabGrid1.SetOrientation(gtk.ORIENTATION_VERTICAL)
	tabLabel1, err := gtk.LabelNew("test tab1")
	tabLabel2, err := gtk.LabelNew("test tab2")
	tabLabel3, err := gtk.LabelNew("test tab3")

	notebook.AppendPage(tabGrid1, tabLabel1)
	notebook.AppendPage(tabGrid2, tabLabel2)

	notebook.AppendPage(tabGrid3, tabLabel3)

	labelsGrid.SetHExpand(true)

	insertBtn, err := gtk.ButtonNewWithLabel("Add a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	removeBtn, err := gtk.ButtonNewWithLabel("Remove a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	nLabels := 1
	insertBtn.Connect("clicked", func() {
		var s string
		if nLabels == 1 {
			s = fmt.Sprintf("Inserted %d label.", nLabels)
		} else {
			s = fmt.Sprintf("Inserted %d labels.", nLabels)
		}
		label, err := gtk.LabelNew(s)
		if err != nil {
			log.Print("Unable to create label:", err)
			return
		}

		labelList.PushBack(label)
		labelsGrid.Add(label)
		label.SetHExpand(true)
		labelsGrid.ShowAll()

		nLabels++
	})

	removeBtn.Connect("clicked", func() {
		e := labelList.Front()
		if e == nil {
			log.Print("Nothing to remove")
			return
		}
		lab, ok := labelList.Remove(e).(*gtk.Label)
		if !ok {
			log.Print("Element to remove is not a *gtk.Label")
			return
		}
		// (*Widget).Destroy() breaks this label's reference with all
		// other objects (in this case, the Grid container it was added
		// to).
		lab.Destroy()

		// At this point, only Go retains a reference to the GtkLabel.
		// When the lab variable goes out of scope when this function
		// returns, at the next garbage collector run, a finalizer will
		// be run to perform the final unreference and free the widget.
	})

	// grid.Attach(insertBtn, 0, 2, 1, 1)
	// grid.Attach(removeBtn, 1, 2, 1, 1)

	return &grid.Container.Widget
}

// //--------------------------------------------------------
// 	// GtkVBox
// 	//--------------------------------------------------------
// 	vbox := gtk.NewVBox(false, 1)

// 	//GTK toolbar
// 	toolbar := gtk.NewToolbar()
// 	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

// 	btnnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
// 	btnnew.SetArrowTooltipText("Load a new project")
// 	btnclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)
// 	btnclose.SetArrowTooltipText("Close active tab")
// 	separator := gtk.NewSeparatorToolItem()
// 	btnGlobalSettings := gtk.NewToolButton(nil, "Global Settings")
// 	btnGlobalSettings.SetArrowTooltipText("Modify caverunner global settings")
// 	btnmenu := gtk.NewMenuToolButtonFromStock("gtk.STOCK_CLOSE")
// 	btnmenu.SetArrowTooltipText("This is a tool tip")

// 	//GTK notebook
// 	notebook := gtk.NewNotebook()

// 	//handle click events on new and close buttons
// 	//opens a project from directory
// 	btnnew.OnClicked(func() {
// 		OpenProject(&state, window, notebook)
// 	})
// 	//closes a tab and removes project from state
// 	btnclose.OnClicked(func() {
// 		if len(state.Projects) != 0 {
// 			tab := notebook.GetCurrentPage()
// 			notebook.RemovePage(notebook, tab)
// 			//this deletes a project from state.Projects slice
// 			state.Projects = append(state.Projects[:tab], state.Projects[tab+1:]...)
// 		}
// 	})

// 	//closes a tab and removes project from state
// 	btnGlobalSettings.OnClicked(func() {
// 		//--------
// 		//project settings dialog working
// 		//--------
// 		settingsDialog := gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_OTHER, gtk.BUTTONS_NONE, "Global Settings")
// 		settingsDialog.Connect("destroy", func() { settingsDialog.Destroy() })

// 		settingsDialog.SetSizeRequest(500, 500)
// 		settingsDialog.SetPosition(gtk.WIN_POS_CENTER)

// 		//install git button
// 		installGit := gtk.NewButtonWithLabel("install git")
// 		installGit.Clicked(func() {
// 			ok := openBrowser("https://git-scm.com/downloads")
// 			if !ok {
// 				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading git download page. Please go to https://git-scm.com/downloads to install git")
// 			}
// 		})

// 		//install go button
// 		installGo := gtk.NewButtonWithLabel("install go")
// 		installGo.Clicked(func() {
// 			ok := openBrowser("https://golang.org/dl/")
// 			if !ok {
// 				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading go download page. Please go to https://golang.org/dl/ to install Golang")
// 			}
// 		})

// 		//install go button
// 		installDocker := gtk.NewButtonWithLabel("install docker")
// 		installDocker.Clicked(func() {
// 			ok := openBrowser("https://www.docker.com/community-edition#/download")
// 			if !ok {
// 				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading docker download page. Please go to https://www.docker.com/community-edition#/download to install docker")
// 			}
// 		})

// 		//setup workspace
// 		setupWorkspace := gtk.NewButtonWithLabel("setup go workspace")
// 		setupWorkspace.Clicked(func() {
// 			ok := setupGoWorkspace()
// 			if ok {
// 				log.Println("success")
// 			}
// 		})

// 		generateCheck := gtk.NewCheckButtonWithLabel("compress and minify")
// 		// frame := gtk.NewFrame("setup")
// 		fixed := gtk.NewFixed()
// 		// frame.Add(installDocker)
// 		// frame.Add(installGit)
// 		fixed.Put(installGit, 10, 10)
// 		fixed.Put(installGo, 10, 50)
// 		fixed.Put(setupWorkspace, 10, 90)
// 		fixed.Put(installDocker, 10, 130)
// 		fixed.Put(generateCheck, 10, 550)

// 		scrolledWindow := gtk.NewScrolledWindow(nil, nil)
// 		scrolledWindow.AddWithViewPort(fixed)
// 		vbox := settingsDialog.GetVBox()
// 		vbox.Add(scrolledWindow)

// 		settingsDialog.AddButton("Close", gtk.RESPONSE_CLOSE).Clicked(func() {
// 		})
// 		settingsDialog.AddButton("Save", gtk.RESPONSE_APPLY).Clicked(func() {
// 			log.Printf("this is a test of the emergency broadcast system")
// 			if generateCheck.GetActive() {
// 				log.Printf("generate button checked")
// 			}

// 		})
// 		vbox.ShowAll()
// 		settingsDialog.Run()
// 		settingsDialog.Destroy()

// 	})
// 	//TODO: Make a save button which saves the current project state into a yaml file in the projects directory
// 	//btnsave
// 	//onclick
// 	//get tab
// 	//get Project From State
// 	//marshal To New Yaml File
// 	//Save to disk

// 	//some menu stuff we're not using right now
// 	// toolmenu := gtk.NewMenu()
// 	// toolitem := gtk.NewMenuItemWithMnemonic("blue")
// 	// toolitem.Show()
// 	// toolmenu.Append(toolitem)
// 	// toolitem = gtk.NewMenuItemWithMnemonic("green")
// 	// toolitem.Show()
// 	// toolmenu.Append(toolitem)
// 	// toolitem = gtk.NewMenuItemWithMnemonic("red")
// 	// toolitem.Show()
// 	// toolmenu.Append(toolitem)
// 	// btnmenu.SetMenu(toolmenu)

// 	//adding an button with an image
// 	// imagefile := filepath.Join(dir, "../../data/go-gtk-logo.png")
// 	// image := gtk.NewImageFromFile(imagefile)
// 	// button.SetImage(image)
// 	// // framebox1.Add(button)

// 	toolbar.Insert(btnnew, -1)
// 	toolbar.Insert(btnclose, -1)
// 	toolbar.Insert(separator, -1)
// 	toolbar.Insert(btnGlobalSettings, -1)
// 	// toolbar.Insert(btnmenu, -1)

// 	vbox.PackStart(toolbar, false, false, 0)
// 	vbox.PackStart(notebook, false, false, 0)

// 	//--------------------------------------------------------
// 	// Event
// 	//--------------------------------------------------------
// 	window.Add(vbox)
// 	window.SetSizeRequest(900, 600)
// 	window.ShowAll()

// 	gtk.Main()
