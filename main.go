package main

import (
	"fmt"
	"io/ioutil"

	"github.com/voxelbrain/goptions"
	"gopkg.in/yaml.v2"
)

func main() {
	options := struct {
		Jsonfile string        `goptions:"-j, --json, obligatory, description='File with json content'"`
		Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`
	}{ // Default values goes here

	}
	goptions.ParseAndFail(&options)

	fileBytes, _ := ioutil.ReadFile(options.Jsonfile)

	f1 := &yaml.MapSlice{}
	if err := yaml.Unmarshal(fileBytes, &f1); err != nil {
		fmt.Println(err)
	}
	yamlBytes, err2 := yaml.Marshal(f1)
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(yamlBytes[:]))
}
