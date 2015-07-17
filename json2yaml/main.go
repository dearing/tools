package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	json, err := yaml.JSONToYAML(bytes)
	if err != nil {
		log.Panicln(err)
	}

	print(string(json))
}

func convert() {

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	json, err := yaml.JSONToYAML(bytes)
	if err != nil {
		log.Panicln(err)
	}
	
}