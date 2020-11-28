package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Template represents the JSON file structure to consume
type Template struct {
	Name    string       `json:"name"`
	Folders []FolderInfo `json:"folders"`
}

// FolderInfo provides the name and Gen command alias for generating files
type FolderInfo struct {
	Name string `json:"name"`
	Gen  string `json:"gen"`
}

// CreateFromTemplates creates the folder inside the JSON structure
func CreateFromTemplate(template *Template) {

	for _, folder := range template.Folders {
		if err := os.Mkdir(folder.Name, os.ModePerm); err != nil {
			panic(err)
		}

	}
}

func main() {
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initTemplate := initCommand.String("template", "structure-template.json", "Path to the json file you want to struct from")

	saveCommand := flag.NewFlagSet("save", flag.ExitOnError)
	templateFile := saveCommand.String("file", "structure-template.json", "The file you want to submit to server")

	if len(os.Args) < 2 {
		fmt.Println("Please provide the command you want to execute, for help enter -help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initCommand.Parse(os.Args[2:])
		jsonTemplate, err := os.Open(path.Join(".", *initTemplate))

		defer jsonTemplate.Close()
		if err != nil {
			fmt.Printf("Unable to locate the json template file. Please provide it using -template.")
			os.Exit(1)
		}

		byteValue, _ := ioutil.ReadAll(jsonTemplate)

		var template Template

		json.Unmarshal(byteValue, &template)

		CreateFromTemplate(&template)

	case "save":
		saveCommand.Parse(os.Args[2:])
		jsonTemplate, err := os.Open(path.Join(".", *templateFile))

		defer jsonTemplate.Close()
		if err != nil {
			fmt.Printf("Unable to locate json template file.")
			os.Exit(1)
		}
		// TODO: SEND TO API
	}

}
