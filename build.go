package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Project struct {
	Name string   `yaml:"name"`
	Path string   `yaml:"path"`
	Args []string `yaml:"args"`
}

//go tool functionality starts here

//update dependencies command (dep ensure)
func (project Project) Dep() {

}

func (project Project) BuildRun() {
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
		// buffer.Insert(&end, string(outputRun))

	} else {
		fmt.Println(string(outputRun))
		// buffer.Insert(&end, string(outputRun))

	}
}

func (project Project) Generate(comboText string) {

	// go generate
	//get GOPATH variable
	// gopath := os.Getenv("GOPATH")
	//go install - setting the directory project dir as stated in state
	cmdGen := exec.Command("go", "generate", comboText)
	cmdGen.Dir = project.Path
	outputGen, err := cmdGen.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(outputGen))
		// buffer.Insert(&end, string(outputGen))

	} else {
		fmt.Println(string(outputGen))
		// buffer.Insert(&end, string(outputGen))

	}
}
