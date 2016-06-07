package main
import (
	"gopkg.in/yaml.v2"
)

func ConvertToYaml(json string) (string, error) {
	jsonDatastructure := &yaml.MapSlice{}
	if err := yaml.Unmarshal([]byte(json), &jsonDatastructure); err != nil {
		return "", err
	}

	yamlBytes, errOnMarshal := yaml.Marshal(jsonDatastructure)

	if errOnMarshal != nil {
		return "", errOnMarshal
	}

	return string(yamlBytes[:]), nil
}