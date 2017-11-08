package main

import (
	"fmt"
	"go-gtk/gtk"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	yaml "gopkg.in/yaml.v2"
)

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

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func setupGoWorkspace() bool {

	//get home env variable
	home := UserHomeDir()

	//check if workspace already exists and create if not
	if _, err := os.Stat(home + "/go/src"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Printf("file doesn't exist in %v. creating files...", home)

			//create go directory
			mkGo := exec.Command("mkdir", "gotest")
			mkGo.Dir = home
			if mkGo.Start() == nil {
				log.Println("made go directory")
			}

			//create src directory
			mkSrc := exec.Command("mkdir", "go/src")
			mkSrc.Dir = home
			if mkSrc.Start() == nil {
				log.Println("made src directory")
			}

			//create bin directory
			mkBin := exec.Command("mkdir", "go/bin")
			mkBin.Dir = home
			if mkBin.Start() == nil {
				log.Println("made bin directory")
			}
			//create pkg directory
			mkPkg := exec.Command("mkdir", "go/pkg")
			mkPkg.Dir = home
			if mkPkg.Start() == nil {
				log.Println("made pkg directory")
			}
			//set the gopath

		} else {

			// other error
		}
	}
	log.Printf("file exists in %v", home)
	return true

}

func setGoPath() {

	home := UserHomeDir()
	os.Setenv("GOPATH", home+"go/src")

}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
