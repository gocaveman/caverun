package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gotk3/gotk3/gtk"
)

type UI struct {
	Projects []Project
	Window   *gtk.Window
	Notebook *gtk.Notebook
}

func (ui UI) GetTabByName(value string) int {
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

type ProjectDupError struct {
	DupErr string
}

func (e *ProjectDupError) Error() string {
	return fmt.Sprintf("Duplicate project error: %s", e.DupErr)
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
	separator, err := gtk.SeparatorToolItemNew()
	btnGlobalSettings, err := gtk.ToolButtonNew(nil, "Global Settings")
	btnGlobalSettings.SetTooltipText("Modify caverunner global settings")

	btnnew.Connect("clicked", func() {
		ui.RunFileChooser()
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
	toolbar.Insert(separator, -1)
	toolbar.Insert(btnGlobalSettings, -1)

	grid.Attach(toolbar, 0, 0, 2, 1)

	grid.Attach(ui.Notebook, 0, 1, 5, 4)

	grid.SetSizeRequest(600, 300)

	return &grid.Container.Widget
}

func (ui *UI) RunFileChooser() {

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
	//cancel button pushed
	if ret == -6 {
		filechooserdialog.Destroy()
		//open button pushed
	} else if ret == -3 {
		configPath, err := filechooserdialog.GetCurrentFolder()
		if err != nil {
			log.Printf("couldn't get folder. error is: %v\n", err)
		}
		projectName := filepath.Base(configPath)
		//verify project tab is not open
		proceed, err := ui.VerifyProject(projectName)
		if err != nil {
			//error check for dup project
			if de, ok := err.(*ProjectDupError); ok {
				log.Println(de.DupErr)
				//already open
				filechooserdialog.Destroy()
				//if project already exists in a tab, tell them
				dialog := gtk.MessageDialogNew(
					ui.Window,
					gtk.DIALOG_MODAL,
					gtk.MESSAGE_INFO,
					gtk.BUTTONS_OK,
					projectName+" is already open. Please choose another project.")
				dialog.SetTitle("Project open!")
				dialog.Run()
			}
			log.Printf("error is: %v\n", err)
		}
		if proceed {

			//if we get to here then the project is not open yet
			//create project in memory and either populate it or start from scratch
			var project *Project
			//check for yaml file
			if _, err := os.Stat(configPath + "/caverun.yaml"); !os.IsNotExist(err) {
				if err != nil {
					panic(err)
				}
				//if yaml file is there then open it
				file, err := os.Open(configPath + "/caverun.yaml")
				if err != nil {
					log.Printf("error is %v\n", err)
				}
				//make project from yaml
				project, err = NewProjectFromYaml(file)
				if err != nil {
					log.Printf("error is %v\n", err)
				}
			} else {
				//yaml file does not exist - create new empty project
				project = ui.NewEmptyProject(configPath)
			}
			ui.Projects = append(ui.Projects, *project)
			ui.MakeNotebookTab(project)
		}

		filechooserdialog.Destroy()
	}
}

func (ui *UI) VerifyProject(projectName string) (bool, error) {
	//check if the project is already open in a tab/memory
	for _, v := range ui.Projects {
		if projectName == v.Name {
			return false, &ProjectDupError{"Project already open."}
		}
	}
	//not open yet
	return true, nil
}

func (ui *UI) NewEmptyProject(configPath string) *Project {

	project := &Project{
		Name: filepath.Base(configPath),
		Path: configPath,
	}

	return project
}

// //MakeNotebookTab makes a tab for a single project including the buttons, widgets, etc.
// //It also updates teh state to include the .yaml project data

func (ui *UI) MakeNotebookTab(project *Project) {

	buttonBuildRun, err := gtk.ButtonNewWithLabel("Build & Run")
	if err != nil {
		log.Fatalf("error making buttons: %v\n", err)
	}
	buttonGoGenerate, err := gtk.ButtonNewWithLabel("Generate")
	if err != nil {
		log.Fatalf("error making buttons: %v\n", err)
	}
	comboGenerate, err := gtk.ComboBoxTextNewWithEntry()
	if err != nil {
		log.Fatalf("error making buttons: %v\n", err)
	}

	//TODO: for loop to add combo items (project.Args) to combo box

	for _, v := range project.Args {
		comboGenerate.AppendText(v)
	}

	buttonUpdateDep, err := gtk.ButtonNewWithLabel("update deps")
	if err != nil {
		log.Fatalf("error making buttons: %v\n", err)
	}
	buttonProjectSettings, err := gtk.ButtonNewWithLabel("Project Settings")
	if err != nil {
		log.Fatalf("error making buttons: %v\n", err)
	}
	// space, err := gtk.LabelNew("                    ")

	tabGrid1, err := gtk.GridNew()
	if err != nil {
		log.Fatalf("error making tab grid: %v\n", err)
	}

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
	if err != nil {
		log.Fatalf("error making tab label: %v\n", err)
	}
	tabButton, err := gtk.ButtonNewFromIconName("window-close", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatalf("error making tab button: %v\n", err)
	}

	tabButton.Connect("clicked", func() {
		tab := ui.GetTabByName(project.Name)

		ui.Notebook.RemovePage(tab)

		//this deletes a project from state.Projects slice
		ui.Projects = append(ui.Projects[:tab], ui.Projects[tab+1:]...)

	})
	tabCloseGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatalf("error making tab close grid: %v\n", err)
	}
	tabCloseGrid.Attach(tabLabel1, 1, -1, 1, 1)
	tabCloseGrid.Attach(tabButton, 2, -1, 1, 1)
	tabCloseGrid.ShowAll()

	//FIXME: returns an int
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
		log.Printf("page is %v\n", page)
		comboText := comboGenerate.GetActiveText()

		exists, err := ui.Projects[page].CheckArgs(comboText)
		if err != nil {
			log.Printf("generate args error: %v\n", err)
		}

		//add arg to dropdown
		if !exists {
			log.Println("Doesn't Exist!")
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
	if err != nil {
		log.Fatalf("error making install go button: %v\n", err)
	}

	installGo.Connect("clicked", func() {
		ok := openBrowser("https://golang.org/dl/")
		if !ok {
			gtk.MessageDialogNew(ui.Window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading go download page. Please go to https://golang.org/dl/ to install Golang")
		}
	})

	//install go button
	installDocker, err := gtk.ButtonNewWithLabel("install docker")
	if err != nil {
		log.Fatalf("error making docker button: %v\n", err)
	}
	installDocker.Connect("clicked", func() {
		ok := openBrowser("https://www.docker.com/community-edition#/download")
		if !ok {
			gtk.MessageDialogNew(ui.Window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading docker download page. Please go to https://www.docker.com/community-edition#/download to install docker")
		}
	})

	//setup workspace
	setupWorkspace, err := gtk.ButtonNewWithLabel("setup go workspace")
	if err != nil {
		log.Fatalf("error making workspace button: %v\n", err)
	}
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
	if err != nil {
		log.Fatalf("error making general settings label: %v\n", err)
	}
	tabButton, err := gtk.ButtonNewFromIconName("window-close", gtk.ICON_SIZE_BUTTON)
	if err != nil {
		log.Fatalf("error making tab close button: %v\n", err)
	}
	tabButton.Connect("clicked", func() {

		tab := ui.GetTabByName(project.Name)

		ui.Notebook.RemovePage(tab)

		//this deletes a project from state.Projects slice
		ui.Projects = append(ui.Projects[:tab], ui.Projects[tab+1:]...)

	})
	tabCloseGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatalf("error making tab close grid: %v\n", err)
	}
	tabCloseGrid.Attach(tabLabel1, 1, -1, 1, 1)
	tabCloseGrid.Attach(tabButton, 2, -1, 1, 1)
	tabCloseGrid.ShowAll()

	ui.Notebook.AppendPage(tabGrid1, tabCloseGrid)
	ui.Notebook.ShowAll()
}
