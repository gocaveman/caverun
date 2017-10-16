package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "bitbucket.org/cwrenard/scraper/jungle-rabbit/src/gopkg.in/yaml.v2"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	var state State
	// state.Projects = TestProjects()
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("Cave Runner")
	window.SetIconName("Cave Runner")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.NewVBox(false, 1)
	// vbox.Add(tab)
	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.NewMenuBar()

	//toolbar start

	toolbar := gtk.NewToolbar()
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

	btnnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	btnclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)
	separator := gtk.NewSeparatorToolItem()
	btnGlobalSettings := gtk.NewToolButton(nil, "Global Settings")
	btnmenu := gtk.NewMenuToolButtonFromStock("gtk.STOCK_CLOSE")
	btnmenu.SetArrowTooltipText("This is a tool tip")

	// btnclose.OnClicked()
	// btncustom.OnClicked()

	toolmenu := gtk.NewMenu()
	toolitem := gtk.NewMenuItemWithMnemonic("8")
	toolitem.Show()
	toolmenu.Append(toolitem)
	toolitem = gtk.NewMenuItemWithMnemonic("16")
	toolitem.Show()
	toolmenu.Append(toolitem)
	toolitem = gtk.NewMenuItemWithMnemonic("32")
	toolitem.Show()
	toolmenu.Append(toolitem)
	btnmenu.SetMenu(toolmenu)

	toolbar.Insert(btnnew, -1)
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)
	toolbar.Insert(btnGlobalSettings, -1)
	toolbar.Insert(btnmenu, -1)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	// vpaned := gtk.NewVPaned()

	// //--------------------------------------------------------
	// // GtkFrame
	// //--------------------------------------------------------
	// frame1 := gtk.NewFrame("")
	// framebox1 := gtk.NewHBox(false, 1)
	// frame1.Add(framebox1)

	// frame2 := gtk.NewFrame("")
	// framebox2 := gtk.NewVBox(false, 1)
	// frame2.Add(framebox2)

	// //--------------------------------------------------------
	// // GtkImage
	// //--------------------------------------------------------

	// label := gtk.NewLabel("Go Binding for GTK")
	// label.ModifyFontEasy("DejaVu Serif 15")

	button := gtk.NewButtonWithLabel("Choose a project")
	button.SetSizeRequest(5, 40)
	dir, _ := filepath.Split(os.Args[0])
	imagefile := filepath.Join(dir, "../../data/go-gtk-logo.png")
	image := gtk.NewImageFromFile(imagefile)
	button.SetImage(image)
	// framebox1.Add(button)

	// --------------------------------------------------------
	// GtkStatusbar
	// --------------------------------------------------------
	statusbar := gtk.NewStatusbar()
	context_id := statusbar.GetContextId("Cave Runner")
	statusbar.Push(context_id, "Cave Runner")

	// framebox2.PackStart(statusbar, false, false, 2)

	//--------------------------------------------------------
	// GtkTabs
	//--------------------------------------------------------

	notebook := gtk.NewNotebook()
	log.Printf("starting project ")

	for _, pjt := range state.Projects {
		MakeNotebookTab(&pjt, notebook)

	}
	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk.NewMenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	var menuitem, menuitem2 *gtk.MenuItem
	menuitem = gtk.NewMenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	})

	menuitem2 = gtk.NewMenuItemWithMnemonic("L_oad Project")
	menuitem2.Connect("activate", func() {

		OpenProject(&state, window, notebook)

	})
	submenu.Append(menuitem2)
	submenu.Append(menuitem)

	cascademenu = gtk.NewMenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk.NewMenu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.NewMenuItemWithMnemonic("_About")
	menuitem.Connect("activate", func() {
		dialog := gtk.NewAboutDialog()
		dialog.SetName("Cave Runner")
		dialog.SetWebsite("https://www.corriganrenard.com")
		dialog.SetProgramName("Cave Runner")

		dialog.SetLicense("Licence data...")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()

	})

	submenu.Append(menuitem)

	btnnew.OnClicked(func() {
		OpenProject(&state, window, notebook)
	})

	vbox.PackStart(menubar, false, false, 0)
	vbox.PackStart(toolbar, false, false, 0)

	vbox.PackStart(notebook, false, false, 0)
	// framebox1.PackStart(label, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)

	window.SetSizeRequest(600, 550)
	window.ShowAll()
	gtk.Main()
}

type Project struct {
	Name   string   `yaml:"name"`
	Path   string   `yaml:"path"`
	Args   []string `yaml:",flow"`
	Output []string
}

func TestProjects() []Project {

	var testProjects = []Project{
		Project{
			Name: "Testing123",
			Path: "Cory/Cory/",
		},
	}
	return testProjects
}

type State struct {
	Projects []Project
}

func OpenProject(state *State, window *gtk.Window, notebook *gtk.Notebook) {

	//--------------------------------------------------------
	// GtkFileChooserDialog
	//--------------------------------------------------------
	filechooserdialog := gtk.NewFileChooserDialog(
		"Choose Project...",
		window,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		gtk.STOCK_OK,
		gtk.RESPONSE_ACCEPT)
	filter := gtk.NewFileFilter()
	filter.SetName(".yaml")
	filter.AddPattern("*.yaml")
	filechooserdialog.AddFilter(filter)
	filechooserdialog.Response(func() {
		fmt.Println(filechooserdialog.GetFilename())
		configPath := filechooserdialog.GetFilename()
		var project *Project
		if configPath != "" {
			yamlFile, err := ioutil.ReadFile(configPath)
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
				state.Projects = append(state.Projects, *project)
				log.Printf("stuct: %v", state.Projects)

				MakeNotebookTab(project, notebook)
			} else {
				filechooserdialog.Destroy()
				dialog := gtk.NewMessageDialog(
					window,
					gtk.DIALOG_MODAL,
					gtk.MESSAGE_INFO,
					gtk.BUTTONS_OK,
					project.Name+" already exists. Please choose another project.")
				dialog.SetTitle("Project Exists!")
				dialog.Response(func() {
					dialog.Destroy()
				})
				dialog.Run()
			}
		}
		filechooserdialog.Destroy()

	})
	filechooserdialog.Run()

}

func MakeNotebookTab(project *Project, notebook *gtk.Notebook) {
	tabVbox := gtk.NewVBox(false, 10)
	fixed := gtk.NewFixed()
	tabLabel := gtk.NewLabel(project.Name)
	tabLabel.SetSizeRequest(125, 20)
	notebook.AppendPage(tabVbox, tabLabel)

	//notebook.SetTabLabel(child, tab_label)

	buttonUpdateDep := gtk.NewButtonWithLabel("Update Dependencies")
	buttonUpdateDep.SetSizeRequest(175, 50)
	buttonBuildRun := gtk.NewButtonWithLabel("Build & Run")
	buttonBuildRun.SetSizeRequest(175, 50)
	buttonGoGenerate := gtk.NewButtonWithLabel("Go Generate")
	buttonGoGenerate.SetSizeRequest(125, 40)
	buttonProjectSettings := gtk.NewButtonWithLabel("Project Settings")
	buttonProjectSettings.SetSizeRequest(150, 45)
	fixed.Put(buttonUpdateDep, 350, 0)
	fixed.Put(buttonBuildRun, 350, 70)
	fixed.Put(buttonGoGenerate, 350, 140)
	fixed.Put(buttonProjectSettings, 350, 210)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)
	textview := gtk.NewTextView()
	textview.SetSizeRequest(540, 100)
	var textWrap gtk.WrapMode
	textWrap = 3

	textview.SetWrapMode(textWrap)

	textview.SetEditable(false)
	var start, end gtk.TextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello there big World!")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "")
	tag := buffer.CreateTag("bold", map[string]string{
		"background": "#FF0000", "weight": "700"})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textview)

	buffer.Connect("changed", func() {
		fmt.Println("changed")
	})
	fixed.Put(swin, 30, 280)

	buttonUpdateDep.Connect("clicked", func() {
		go func() {
			log.Printf("clicked")

		}()
	})
	buttonBuildRun.Connect("clicked", func() {
		go func() {
			log.Printf("clicked")
			buffer.Insert(&end, project.Name)

		}()
	})
	buttonGoGenerate.Connect("clicked", func() {
		go func() {
			log.Printf("clicked")

		}()
	})
	buttonProjectSettings.Connect("clicked", func() {
		go func() {
			log.Printf("clicked")

		}()
	})
	tabVbox.PackStart(fixed, false, false, 10)
	notebook.ShowAll()

}
