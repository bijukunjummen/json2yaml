package main

import (
	"fmt"
	"io/ioutil"

	"github.com/voxelbrain/goptions"
)

func main() {
	options := struct {
		Jsonfile string        `goptions:"-j, --json, obligatory, description='File with json content'"`
		Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`
	}{ // Default values go here

	}
	goptions.ParseAndFail(&options)

	fileBytes, _ := ioutil.ReadFile(options.Jsonfile)

	jsonString := string(fileBytes[:])

	yamlString, err := ConvertToYaml(jsonString)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yamlString)


}
