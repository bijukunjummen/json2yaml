package main

import (
	"strings"
	"testing"
)

func TestConvertToYaml(t *testing.T) {
	json := `{"hm_emailer_options": {
	"enabled": false,
	"tls": false,
	"recipients": {
	"value": null
	}}}`

	yaml, err := ConvertToYaml(json)

	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(yaml, "hm_emailer_options") {
		t.Errorf("Bad prefix. Expected %v, Got %v!", "hm_emailer_options", yaml)
	}
}
