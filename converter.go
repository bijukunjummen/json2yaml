package main

import (
	"gopkg.in/yaml.v2"
)

func ConvertToYaml(json string) (string, error) {
	jsonDatastructure := &yaml.MapSlice{}
	if err := yaml.Unmarshal([]byte(json), &jsonDatastructure); err != nil {
		return "", err
	}

	yamlBytes, err := yaml.Marshal(jsonDatastructure)

	if err != nil {
		return "", err
	}

	return string(yamlBytes[:]), nil
}
