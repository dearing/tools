package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

var NAME, VERSION string

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panicln(err)
	}

	json, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		log.Panicln(err)
	}

	print(string(json))
}
