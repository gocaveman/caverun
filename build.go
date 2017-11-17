package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

type Project struct {
	Name string   `yaml:"name"`
	Path string   `yaml:"path"`
	Args []string `yaml:"args"`
}

//go tool functionality starts here

//update dependencies command (dep ensure)
func (project *Project) Dep() {

}

//CheckArgs takes an input string which is the arg to check for and checks if it's already in the Args slice.
//If it's not there it adds it
func (project *Project) CheckArgs(input string) (bool, error) {

	log.Printf("input: %v\n", input)

	//check if the arguments already exist
	//loop through the generate args
	generateArgExists := false
	log.Printf("project is: %v\n", project)

	for _, v := range project.Args {
		log.Printf("arg: %v\n", v)
		log.Printf("input: %v\n", input)
		if v == input {
			generateArgExists = true
			break
		}
	}
	//if they don't add them to the project.generateArgs
	if !generateArgExists {
		// ui.GenerateArgs = append(ui.GenerateArgs, comboText)
		log.Printf("appending: %v\n", input)
		project.Args = append(project.Args, input)

		return false, nil

	}

	return true, nil

}

func NewProjectFromYaml(r io.Reader) (*Project, error) {

	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	var project *Project

	err = yaml.Unmarshal(content, &project)
	if err != nil {
		log.Printf("error is: %v", err)
		return nil, err
	}

	return project, nil

}

func (project *Project) Build() ([]byte, error) {
	//get GOPATH variable
	// gopath := os.Getenv("GOPATH")
	//go install - setting the directory project dir as stated in state
	cmdBuild := exec.Command("go", "install")
	cmdBuild.Dir = project.Path

	outputBuild, err := cmdBuild.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(outputBuild))
		return outputBuild, err

	} else {
		fmt.Println(string(outputBuild))
		fmt.Println("build successful")

		return outputBuild, nil
	}

}
func (project *Project) Run() ([]byte, error) {
	//get GOPATH variable
	gopath := os.Getenv("GOPATH")

	log.Printf("running %v\n", project.Name)

	//run the resulting executable

	cmdRun := exec.Command(gopath + "/bin/" + project.Name)

	outputRun, err := cmdRun.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(outputRun))
		return outputRun, err
		// buffer.Insert(&end, string(outputRun))

	} else {
		fmt.Println(string(outputRun))
		return outputRun, nil
		// buffer.Insert(&end, string(outputRun))

	}
}

func (project *Project) Generate(comboText string) {

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
