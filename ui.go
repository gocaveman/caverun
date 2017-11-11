package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	Projects     []Project
	Window       *gtk.Window
	Notebook     *gtk.Notebook
	GenerateArgs []string
}

func (ui UI) ProjectPos(value string) int {
	for p, v := range ui.Projects {
		if v.Name == value {
			return p
		}
	}
	return -1
}

func NewUI(window *gtk.Window, notebook *gtk.Notebook) *UI {
	return &UI{Window: window, Notebook: notebook}
}

//WindowWidget handles the main window
func (ui *UI) WindowWidget() *gtk.Widget {

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	ui.Notebook.SetSizeRequest(600, 200)

	toolbar, err := gtk.ToolbarNew()
	// toolbar.SetStyle(gtk.TOOLBAR_ICONS)
	btnnew, err := gtk.ToolButtonNew(nil, "new project")

	btnnew.SetTooltipText("Load a new project")
	btnclose, err := gtk.ToolButtonNew(nil, "close project")
	btnclose.SetTooltipText("Close active tab")
	separator, err := gtk.SeparatorToolItemNew()
	btnGlobalSettings, err := gtk.ToolButtonNew(nil, "Global Settings")
	btnGlobalSettings.SetTooltipText("Modify caverunner global settings")

	btnnew.Connect("clicked", func() {
		ui.OpenProject()
	})
	btnclose.Connect("clicked", func() {
		if len(ui.Projects) != 0 {
			tab := ui.Notebook.GetCurrentPage()
			ui.Notebook.RemovePage(tab)
			//this deletes a project from state.Projects slice
			ui.Projects = append(ui.Projects[:tab], ui.Projects[tab+1:]...)
		}
	})
	btnGlobalSettings.Connect("clicked", func() {

		project := &Project{
			Name: "general-settings",
			// Path: configPath,
		}

		//check if a tab is open with the same name
		tabExists := false
		for _, v := range ui.Projects {
			if project.Name == v.Name {
				tabExists = true
				break
			}

		}
		//if no matching tabs are open, add this to state and create tab
		if tabExists == false {
			ui.Projects = append(ui.Projects, *project)
			log.Printf("stuct: %v", ui.Projects)
			ui.MakeSettingsTab(project)
		} else {
			dialog := gtk.MessageDialogNew(
				ui.Window,
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
				project.Name+" is already open.")
			dialog.SetTitle("General Settings")
			// dialog.Response(func() {
			// 	dialog.Destroy()
			// })
			dialog.Run()
		}

	})

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btnGlobalSettings, -1)

	grid.Attach(toolbar, 0, 0, 2, 1)

	grid.Attach(ui.Notebook, 0, 1, 5, 4)

	grid.SetSizeRequest(600, 300)

	return &grid.Container.Widget
}

//OpenProject open a file chooser dialog and selects a folder to load as the current project tab
func (ui *UI) OpenProject() {

	//--------------------------------------------------------
	// GtkFileChooserDialog
	//--------------------------------------------------------
	filechooserdialog, err := gtk.FileChooserDialogNewWith2Buttons(
		"Open Project...",
		ui.Window,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		"Cancel",
		gtk.RESPONSE_CANCEL,
		"Open",
		gtk.RESPONSE_ACCEPT)
	if err != nil {
		log.Fatal("error creating filechooser dialog")
	}

	//event to run when dir is chosen

	ret := filechooserdialog.Run()
	log.Printf("ret is: %v\n", ret)
	if ret == -6 {
		filechooserdialog.Destroy()
	} else if ret == -3 {
		func() {
			configPath, err := filechooserdialog.GetCurrentFolder()
			if err != nil {
				log.Printf("couldn't get folder. error is: %v\n", err)
			}
			log.Printf("configpath: %v", configPath)
			var project *Project
			if configPath != "" {
				//check if config file already exists
				if _, err := os.Stat(configPath + "/caverun.yaml"); !os.IsNotExist(err) {
					if err != nil {
						panic(err)
					}
					log.Printf("file exists, reading caverun.yaml")
					yamlFile, err := ioutil.ReadFile(configPath + "/caverun.yaml")
					if err != nil {
						log.Fatal(err)
					}

					log.Printf("yaml file is: %v", yamlFile)
					err = yaml.Unmarshal(yamlFile, &project)
					if err != nil {
						log.Printf("error is: %v", err)
					}
					log.Printf("project is: %v", project)

					//loop through state.projects and see if project doesn't already exist
					projectExists := false
					for _, v := range ui.Projects {
						if project.Name == v.Name {
							projectExists = true
							break
						}
					}
					if projectExists == false {
						project.Path = configPath
						ui.Projects = append(ui.Projects, *project)
						log.Printf("stuct: %v", ui.Projects)

						ui.MakeNotebookTab(project)
					} else {
						filechooserdialog.Destroy()
						//if project already exists in a tab, tell them
						dialog := gtk.MessageDialogNew(
							ui.Window,
							gtk.DIALOG_MODAL,
							gtk.MESSAGE_INFO,
							gtk.BUTTONS_OK,
							project.Name+" is already open. Please choose another project.")
						dialog.SetTitle("Project open!")
						// dialog.Connect(func() {
						// 	dialog.Destroy()
						// })
						dialog.Run()
					}

				} else {
					log.Printf("file doesn't exist, creating new struct in memory")

					//file does not exist - create new struct in memory
					project = &Project{
						Name: filepath.Base(configPath),
						Path: configPath,
					}

					//check if a tab is open with the same name
					projectExists := false
					log.Printf("projects right before append: %v\n", ui.Projects)
					for _, v := range ui.Projects {
						if project.Name == v.Name {
							projectExists = true
							break
						}

					}
					//if no matching tabs are open, add this to state and create tab
					if projectExists == false {
						ui.Projects = append(ui.Projects, *project)
						log.Printf("stuct: %v", ui.Projects)

						ui.MakeNotebookTab(project)
					} else {
						filechooserdialog.Destroy()
						//if project already exists in a tab, tell them and don't add it
						dialog := gtk.MessageDialogNew(
							ui.Window,
							gtk.DIALOG_MODAL,
							gtk.MESSAGE_INFO,
							gtk.BUTTONS_OK,
							project.Name+" is already open. Please choose another project.")
						dialog.SetTitle("Project open!")
						// dialog.Response(func() {
						// 	dialog.Destroy()
						// })
						dialog.Run()
					}

				}

			}
			//if no folders are chosen, don't desroy window, just wait for a folder to be picked
			if configPath != "" {
				filechooserdialog.Destroy()
			}

		}()
	}

}

// //MakeNotebookTab makes a tab for a single project including the buttons, widgets, etc.
// //It also updates teh state to include the .yaml project data

func (ui *UI) MakeNotebookTab(project *Project) {

	buttonBuildRun, err := gtk.ButtonNewWithLabel("Build & Run")
	buttonGoGenerate, err := gtk.ButtonNewWithLabel("Generate")
	comboGenerate, err := gtk.ComboBoxTextNewWithEntry()
	buttonUpdateDep, err := gtk.ButtonNewWithLabel("update deps")
	buttonProjectSettings, err := gtk.ButtonNewWithLabel("Project Settings")

	if err != nil {
		log.Printf("error making buttons: %v\n", err)
	}
	// space, err := gtk.LabelNew("                    ")

	tabGrid1, err := gtk.GridNew()

	tabGrid1.Attach(buttonBuildRun, -1, 1, 1, 1)
	tabGrid1.Attach(buttonGoGenerate, -1, 2, 1, 1)
	tabGrid1.Attach(comboGenerate, -2, 2, 1, 1)
	tabGrid1.Attach(buttonUpdateDep, -1, 3, 1, 1)
	tabGrid1.Attach(buttonProjectSettings, -1, 4, 1, 1)

	tabGrid1.SetRowSpacing(15)
	tabGrid1.SetMarginTop(20)
	tabGrid1.SetMarginBottom(20)

	tabGrid1.SetMarginStart(20)
	tabGrid1.SetMarginEnd(20)

	tabGrid1.SetHAlign(gtk.ALIGN_END)
	// tabGrid1.Attach(space, 1, 1, 3, 3)

	tabLabel1, err := gtk.LabelNew(project.Name)
	tabButton, err := gtk.ButtonNewFromIconName("window-close", gtk.ICON_SIZE_BUTTON)

	tabButton.Connect("clicked", func() {
		tab := ui.ProjectPos(project.Name)

		ui.Notebook.RemovePage(tab)

		//this deletes a project from state.Projects slice
		ui.Projects = append(ui.Projects[:tab], ui.Projects[tab+1:]...)

	})
	tabCloseGrid, err := gtk.GridNew()
	tabCloseGrid.Attach(tabLabel1, 1, -1, 1, 1)
	tabCloseGrid.Attach(tabButton, 2, -1, 1, 1)
	tabCloseGrid.ShowAll()

	ui.Notebook.AppendPage(tabGrid1, tabCloseGrid)

	// //--------------------------------------------------------
	// // GtkTextView
	// //--------------------------------------------------------
	// swin := gtk.NewScrolledWindow(nil, nil)
	// swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	// swin.SetShadowType(gtk.SHADOW_IN)
	// textview := gtk.NewTextView()
	// textview.SetSizeRequest(840, 100)
	// var textWrap gtk.WrapMode
	// textWrap = 3

	// textview.SetWrapMode(textWrap)

	// textview.SetEditable(false)
	// var start, end gtk.TextIter
	// buffer := textview.GetBuffer()
	// buffer.GetStartIter(&start)
	// buffer.Insert(&start, "")
	// buffer.GetEndIter(&end)
	// buffer.Insert(&end, "Path is: "+project.Path)
	// tag := buffer.CreateTag("bold", map[string]string{
	// 	"background": "#fff", "weight": "700"})
	// buffer.GetStartIter(&start)
	// buffer.GetEndIter(&end)
	// buffer.ApplyTag(tag, &start, &end)
	// swin.Add(textview)

	// buffer.Connect("changed", func() {
	// 	fmt.Println("changed")
	// })
	// fixed.Put(swin, 30, 330)

	//go tool functionality starts here

	//update dependencies command (dep ensure)
	buttonUpdateDep.Connect("clicked", func() {
		page := ui.Notebook.GetCurrentPage()
		ui.Projects[page].Dep()
	})

	//Build and Run command
	buttonBuildRun.Connect("clicked", func() {
		page := ui.Notebook.GetCurrentPage()
		ui.Projects[page].BuildRun()
	})
	//Go Generate command
	buttonGoGenerate.Connect("clicked", func() {
		page := ui.Notebook.GetCurrentPage()
		comboText := comboGenerate.GetActiveText()
		generateArgExists := false
		for _, v := range ui.GenerateArgs {
			if v == comboText {
				generateArgExists = true
			}
		}
		if generateArgExists == false {
			ui.GenerateArgs = append(ui.GenerateArgs, comboText)
			comboGenerate.AppendText(comboText)
		}
		ui.Projects[page].Generate(comboText)
	})
	// //View project settings page
	// buttonProjectSettings.Clicked(func() {

	// 	//--------
	// 	//project settings dialog working
	// 	//--------
	// 	settingsDialog := gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_OTHER, gtk.BUTTONS_NONE, "Project Settings")
	// 	settingsDialog.Connect("destroy", func() { settingsDialog.Destroy() })

	// 	settingsDialog.SetSizeRequest(500, 500)
	// 	settingsDialog.SetPosition(gtk.WIN_POS_CENTER)
	// 	//testing checkbuttons
	// 	generateCheck := gtk.NewCheckButtonWithLabel("compress and minify")
	// 	fixed := gtk.NewFixed()
	// 	fixed.Put(generateCheck, 10, 550)
	// 	scrolledWindow := gtk.NewScrolledWindow(nil, nil)
	// 	scrolledWindow.AddWithViewPort(fixed)
	// 	vbox := settingsDialog.GetVBox()
	// 	vbox.Add(scrolledWindow)
	// 	settingsDialog.AddButton("Close", gtk.RESPONSE_CLOSE).Clicked(func() {
	// 		log.Printf("project name is: %v", project.Name)
	// 	})
	// 	settingsDialog.AddButton("Save", gtk.RESPONSE_APPLY).Clicked(func() {
	// 		log.Printf("this is a test of the emergency broadcast system")
	// 		if generateCheck.GetActive() {
	// 			log.Printf("generate button checked")
	// 		}

	// 	})
	// 	vbox.ShowAll()
	// 	settingsDialog.Run()
	// 	settingsDialog.Destroy()

	// })
	// tabVbox.PackStart(fixed, false, false, 10)
	ui.Notebook.ShowAll()

}

func (ui *UI) MakeSettingsTab(project *Project) {
	tabGrid1, err := gtk.GridNew()

	//install git button
	installGit, err := gtk.ButtonNewWithLabel("install git")
	if err != nil {
		log.Printf("error making buttons: %v\n", err)
	}

	installGit.Connect("clicked", func() {
		ok := openBrowser("https://git-scm.com/downloads")
		if !ok {
			gtk.MessageDialogNew(ui.Window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading git download page. Please go to https://git-scm.com/downloads to install git")
		}
	})

	//install go button
	installGo, err := gtk.ButtonNewWithLabel("install go")
	installGo.Connect("clicked", func() {
		ok := openBrowser("https://golang.org/dl/")
		if !ok {
			gtk.MessageDialogNew(ui.Window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading go download page. Please go to https://golang.org/dl/ to install Golang")
		}
	})

	//install go button
	installDocker, err := gtk.ButtonNewWithLabel("install docker")
	installDocker.Connect("clicked", func() {
		ok := openBrowser("https://www.docker.com/community-edition#/download")
		if !ok {
			gtk.MessageDialogNew(ui.Window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading docker download page. Please go to https://www.docker.com/community-edition#/download to install docker")
		}
	})

	//setup workspace
	setupWorkspace, err := gtk.ButtonNewWithLabel("setup go workspace")
	setupWorkspace.Connect("clicked", func() {
		ok := setupGoWorkspace()
		if ok {
			log.Println("success")
		}
	})

	tabGrid1.Attach(installGit, -1, 1, 1, 1)
	tabGrid1.Attach(installDocker, -1, 2, 1, 1)
	tabGrid1.Attach(installGo, -1, 3, 1, 1)
	tabGrid1.Attach(setupWorkspace, -1, 4, 1, 1)
	// tabGrid1.Attach(buttonProjectSettings, -1, 4, 1, 1)

	tabGrid1.SetRowSpacing(15)
	tabGrid1.SetMarginTop(20)
	tabGrid1.SetMarginBottom(20)

	tabGrid1.SetMarginStart(20)
	tabGrid1.SetMarginEnd(20)

	tabGrid1.SetHAlign(gtk.ALIGN_START)
	// tabGrid1.Attach(space, 1, 1, 3, 3)

	tabLabel1, err := gtk.LabelNew("General Settings")
	tabButton, err := gtk.ButtonNewFromIconName("window-close", gtk.ICON_SIZE_BUTTON)
	tabButton.Connect("clicked", func() {

		tab := ui.Notebook.GetCurrentPage()
		ui.Notebook.RemovePage(tab)
		//this deletes a project from state.Projects slice
		ui.Projects = append(ui.Projects[:tab], ui.Projects[tab+1:]...)

	})
	tabCloseGrid, err := gtk.GridNew()
	tabCloseGrid.Attach(tabLabel1, 1, -1, 1, 1)
	tabCloseGrid.Attach(tabButton, 2, -1, 1, 1)
	tabCloseGrid.ShowAll()

	ui.Notebook.AppendPage(tabGrid1, tabCloseGrid)
	ui.Notebook.ShowAll()
}
