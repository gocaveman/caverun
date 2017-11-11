package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

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
