package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	yaml "bitbucket.org/cwrenard/scraper/jungle-rabbit/src/gopkg.in/yaml.v2"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	var state State
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Cave Runner")
	window.SetIconName("Cave Runner")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	})

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)

	//GTK toolbar
	toolbar := gtk.NewToolbar()
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

	btnnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	btnnew.SetArrowTooltipText("Load a new project")
	btnclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)
	btnclose.SetArrowTooltipText("Close active tab")
	separator := gtk.NewSeparatorToolItem()
	btnGlobalSettings := gtk.NewToolButton(nil, "Global Settings")
	btnGlobalSettings.SetArrowTooltipText("Modify caverunner global settings")
	btnmenu := gtk.NewMenuToolButtonFromStock("gtk.STOCK_CLOSE")
	btnmenu.SetArrowTooltipText("This is a tool tip")

	//GTK notebook
	notebook := gtk.NewNotebook()

	//handle click events on new and close buttons
	//opens a project from directory
	btnnew.OnClicked(func() {
		OpenProject(&state, window, notebook)
	})
	//closes a tab and removes project from state
	btnclose.OnClicked(func() {
		if len(state.Projects) != 0 {
			tab := notebook.GetCurrentPage()
			notebook.RemovePage(notebook, tab)
			//this deletes a project from state.Projects slice
			state.Projects = append(state.Projects[:tab], state.Projects[tab+1:]...)
		}
	})

	//closes a tab and removes project from state
	btnGlobalSettings.OnClicked(func() {
		//--------
		//project settings dialog working
		//--------
		settingsDialog := gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_OTHER, gtk.BUTTONS_NONE, "Global Settings")
		settingsDialog.Connect("destroy", func() { settingsDialog.Destroy() })

		settingsDialog.SetSizeRequest(500, 500)
		settingsDialog.SetPosition(gtk.WIN_POS_CENTER)

		//install git button
		installGit := gtk.NewButtonWithLabel("install git")
		installGit.Clicked(func() {
			ok := openBrowser("https://git-scm.com/downloads")
			if !ok {
				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading git download page. Please go to https://git-scm.com/downloads to install git")
			}
		})

		//install go button
		installGo := gtk.NewButtonWithLabel("install go")
		installGo.Clicked(func() {
			ok := openBrowser("https://golang.org/dl/")
			if !ok {
				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading go download page. Please go to https://golang.org/dl/ to install Golang")
			}
		})

		//install go button
		installDocker := gtk.NewButtonWithLabel("install docker")
		installDocker.Clicked(func() {
			ok := openBrowser("https://www.docker.com/community-edition#/download")
			if !ok {
				gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, "Error loading docker download page. Please go to https://www.docker.com/community-edition#/download to install docker")
			}
		})

		//setup workspace
		setupWorkspace := gtk.NewButtonWithLabel("setup go workspace")
		setupWorkspace.Clicked(func() {
			ok := setupGoWorkspace()
			if ok {
				log.Println("success")
			}
		})

		generateCheck := gtk.NewCheckButtonWithLabel("compress and minify")
		// frame := gtk.NewFrame("setup")
		fixed := gtk.NewFixed()
		// frame.Add(installDocker)
		// frame.Add(installGit)
		fixed.Put(installGit, 10, 10)
		fixed.Put(installGo, 10, 50)
		fixed.Put(setupWorkspace, 10, 90)
		fixed.Put(installDocker, 10, 130)
		fixed.Put(generateCheck, 10, 550)

		scrolledWindow := gtk.NewScrolledWindow(nil, nil)
		scrolledWindow.AddWithViewPort(fixed)
		vbox := settingsDialog.GetVBox()
		vbox.Add(scrolledWindow)

		settingsDialog.AddButton("Close", gtk.RESPONSE_CLOSE).Clicked(func() {
		})
		settingsDialog.AddButton("Save", gtk.RESPONSE_APPLY).Clicked(func() {
			log.Printf("this is a test of the emergency broadcast system")
			if generateCheck.GetActive() {
				log.Printf("generate button checked")
			}

		})
		vbox.ShowAll()
		settingsDialog.Run()
		settingsDialog.Destroy()

	})
	//TODO: Make a save button which saves the current project state into a yaml file in the projects directory
	//btnsave
	//onclick
	//get tab
	//get Project From State
	//marshal To New Yaml File
	//Save to disk

	//some menu stuff we're not using right now
	// toolmenu := gtk.NewMenu()
	// toolitem := gtk.NewMenuItemWithMnemonic("blue")
	// toolitem.Show()
	// toolmenu.Append(toolitem)
	// toolitem = gtk.NewMenuItemWithMnemonic("green")
	// toolitem.Show()
	// toolmenu.Append(toolitem)
	// toolitem = gtk.NewMenuItemWithMnemonic("red")
	// toolitem.Show()
	// toolmenu.Append(toolitem)
	// btnmenu.SetMenu(toolmenu)

	//adding an button with an image
	// imagefile := filepath.Join(dir, "../../data/go-gtk-logo.png")
	// image := gtk.NewImageFromFile(imagefile)
	// button.SetImage(image)
	// // framebox1.Add(button)

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btnGlobalSettings, -1)
	// toolbar.Insert(btnmenu, -1)

	vbox.PackStart(toolbar, false, false, 0)
	vbox.PackStart(notebook, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(900, 600)
	window.ShowAll()

	gtk.Main()

}

type Project struct {
	Name   string   `yaml:"name"`
	Path   string   `yaml:"path"`
	Args   []string `yaml:",flow"`
	Output []string
}

type State struct {
	Projects []Project
}

var generateArgs []string

//OpenProject open a file chooser dialog and selects a folder to load as the current project tab
func OpenProject(state *State, window *gtk.Window, notebook *gtk.Notebook) {

	//--------------------------------------------------------
	// GtkFileChooserDialog
	//--------------------------------------------------------
	filechooserdialog := gtk.NewFileChooserDialog(
		"Choose Project...",
		window,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		gtk.STOCK_OK,
		gtk.RESPONSE_ACCEPT)

	//not using this filter stuff
	// filter := gtk.NewFileFilter()
	// filter.SetName(".yaml")
	// filter.AddPattern("*.yaml")
	// filechooserdialog.AddFilter(filter)
	// filechooserdialog.AddButton("generate yaml", gtk.RESPONSE_NONE).Clicked(func() {
	// 	generatePath := filechooserdialog.GetFilename()

	//event to run when dir is chosen
	filechooserdialog.Response(func() {
		configPath := filechooserdialog.GetCurrentFolder()
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
				for _, v := range state.Projects {
					if project.Name == v.Name {
						projectExists = true
						break
					}
				}
				if projectExists == false {
					project.Path = configPath
					state.Projects = append(state.Projects, *project)
					log.Printf("stuct: %v", state.Projects)

					MakeNotebookTab(project, notebook, window)
				} else {
					filechooserdialog.Destroy()
					//if project already exists in a tab, tell them
					dialog := gtk.NewMessageDialog(
						window,
						gtk.DIALOG_MODAL,
						gtk.MESSAGE_INFO,
						gtk.BUTTONS_OK,
						project.Name+" is already open. Please choose another project.")
					dialog.SetTitle("Project open!")
					dialog.Response(func() {
						dialog.Destroy()
					})
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
				for _, v := range state.Projects {
					if project.Name == v.Name {
						projectExists = true
						break
					}

				}
				//if no matching tabs are open, add this to state and create tab
				if projectExists == false {
					state.Projects = append(state.Projects, *project)
					log.Printf("stuct: %v", state.Projects)

					MakeNotebookTab(project, notebook, window)
				} else {
					filechooserdialog.Destroy()
					//if project already exists in a tab, tell them and don't add it
					dialog := gtk.NewMessageDialog(
						window,
						gtk.DIALOG_MODAL,
						gtk.MESSAGE_INFO,
						gtk.BUTTONS_OK,
						project.Name+" is already open. Please choose another project.")
					dialog.SetTitle("Project open!")
					dialog.Response(func() {
						dialog.Destroy()
					})
					dialog.Run()
				}

			}

		}
		//if no folders are chosen, don't desroy window, just wait for a folder to be picked
		if configPath != "" {
			filechooserdialog.Destroy()
		}

	})
	filechooserdialog.Run()

}

//MakeNotebookTab makes a tab for a single project including the buttons, widgets, etc.
//It also updates teh state to include the .yaml project data

func MakeNotebookTab(project *Project, notebook *gtk.Notebook, window *gtk.Window) {
	tabVbox := gtk.NewVBox(false, 10)
	fixed := gtk.NewFixed()

	tabLabel := gtk.NewLabel(project.Name)
	tabLabel.SetSizeRequest(125, 20)

	notebook.AppendPage(tabVbox, tabLabel)

	buttonUpdateDep := gtk.NewButtonWithLabel("Update Dependencies")
	buttonUpdateDep.SetSizeRequest(175, 50)
	buttonBuildRun := gtk.NewButtonWithLabel("Build & Run")
	buttonBuildRun.SetSizeRequest(175, 50)
	buttonGoGenerate := gtk.NewButtonWithLabel("Go Generate")
	buttonGoGenerate.SetSizeRequest(100, 40)
	buttonProjectSettings := gtk.NewButtonWithLabel("Project Settings")
	buttonProjectSettings.SetSizeRequest(150, 45)
	// comboBoxEntry := gtk.NewComboBoxEntryNewText()
	// comboBoxWithEntry := gtk.NewComboBoxWithEntry()
	// comboBoxWithEntry.
	// comboBoxWithEntry.Add(comboBoxEntry)
	// comboBoxWithEntry.SetSizeRequest(150, 40)

	comboboxentry := gtk.NewComboBoxEntryNewText()

	// comboboxentry.Connect("insert-text", func(ctx *glib.CallbackContext) {
	// 	a := (*[2000]uint8)(unsafe.Pointer(ctx.Args(0)))
	// 	p := (*int)(unsafe.Pointer(ctx.Args(2)))
	// 	i := 0
	// 	for a[i] != 0 {
	// 		i++
	// 	}
	// 	s := string(a[0:i])
	// 	if s == "." {
	// 		if *p == 0 {
	// 			comboboxentry.StopEmission("insert-text")
	// 		}
	// 	} else {
	// 		_, err := strconv.Atoi(s)
	// 		if err != nil {
	// 			comboboxentry.StopEmission("insert-text")
	// 		}
	// 	}
	// 	log.Printf("input is %v\n", comboboxentry.GetActiveText())
	// })

	fixed.Put(buttonUpdateDep, 550, 0)
	fixed.Put(buttonBuildRun, 550, 70)
	fixed.Put(buttonGoGenerate, 550, 140)
	fixed.Put(buttonProjectSettings, 550, 240)
	fixed.Put(comboboxentry, 670, 145)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	textview := gtk.NewTextView()
	textview.SetSizeRequest(840, 100)
	var textWrap gtk.WrapMode
	textWrap = 3

	textview.SetWrapMode(textWrap)

	textview.SetEditable(false)
	var start, end gtk.TextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "Path is: "+project.Path)
	tag := buffer.CreateTag("bold", map[string]string{
		"background": "#fff", "weight": "700"})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textview)

	buffer.Connect("changed", func() {
		fmt.Println("changed")
	})
	fixed.Put(swin, 30, 330)

	//go tool functionality starts here

	//update dependencies command (dep ensure)
	buttonUpdateDep.Connect("clicked", func() {
		go func() {
			//dep ensure

		}()
	})

	//Build and Run command
	buttonBuildRun.Connect("clicked", func() {
		go func() {

			//TODO: inserting error data into output box
			buffer.Insert(&end, project.Name)

			//get GOPATH variable
			gopath := os.Getenv("GOPATH")
			//go install - setting the directory project dir as stated in state
			cmdBuild := exec.Command("go", "install")
			cmdBuild.Dir = project.Path
			//run the resulting executable
			cmdRun := exec.Command(gopath + "/bin/" + project.Name)

			outputBuild, err := cmdBuild.CombinedOutput()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(outputBuild))

			} else {
				fmt.Println(string(outputBuild))
			}
			log.Printf("running exe file")

			outputRun, err := cmdRun.CombinedOutput()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(outputRun))
				buffer.Insert(&end, string(outputRun))

			} else {
				fmt.Println(string(outputRun))
				buffer.Insert(&end, string(outputRun))

			}

		}()
	})
	//Go Generate command
	buttonGoGenerate.Connect("clicked", func() {
		go func() {
			log.Printf("clicked %v", project.Name)

			generateArgExists := false
			for _, v := range generateArgs {
				if v == comboboxentry.GetActiveText() {
					generateArgExists = true
				}
			}
			if generateArgExists == false {
				generateArgs = append(generateArgs, comboboxentry.GetActiveText())
				comboboxentry.AppendText(comboboxentry.GetActiveText())
			}

			// go generate
			//get GOPATH variable
			// gopath := os.Getenv("GOPATH")
			//go install - setting the directory project dir as stated in state
			cmdGen := exec.Command("go", "generate", comboboxentry.GetActiveText())
			cmdGen.Dir = project.Path
			outputGen, err := cmdGen.CombinedOutput()
			if err != nil {
				fmt.Println(fmt.Sprint(err) + ": " + string(outputGen))
				buffer.Insert(&end, string(outputGen))

			} else {
				fmt.Println(string(outputGen))
				buffer.Insert(&end, string(outputGen))

			}

		}()
	})
	//View project settings page
	buttonProjectSettings.Clicked(func() {

		//--------
		//project settings dialog working
		//--------
		settingsDialog := gtk.NewMessageDialog(window, gtk.DIALOG_MODAL, gtk.MESSAGE_OTHER, gtk.BUTTONS_NONE, "Project Settings")
		settingsDialog.Connect("destroy", func() { settingsDialog.Destroy() })

		settingsDialog.SetSizeRequest(500, 500)
		settingsDialog.SetPosition(gtk.WIN_POS_CENTER)
		//testing checkbuttons
		generateCheck := gtk.NewCheckButtonWithLabel("compress and minify")
		fixed := gtk.NewFixed()
		fixed.Put(generateCheck, 10, 550)
		scrolledWindow := gtk.NewScrolledWindow(nil, nil)
		scrolledWindow.AddWithViewPort(fixed)
		vbox := settingsDialog.GetVBox()
		vbox.Add(scrolledWindow)
		settingsDialog.AddButton("Close", gtk.RESPONSE_CLOSE).Clicked(func() {
			log.Printf("project name is: %v", project.Name)
		})
		settingsDialog.AddButton("Save", gtk.RESPONSE_APPLY).Clicked(func() {
			log.Printf("this is a test of the emergency broadcast system")
			if generateCheck.GetActive() {
				log.Printf("generate button checked")
			}

		})
		vbox.ShowAll()
		settingsDialog.Run()
		settingsDialog.Destroy()

	})
	tabVbox.PackStart(fixed, false, false, 10)
	notebook.ShowAll()

}
