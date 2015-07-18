package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

//go:generate goversion -major=1 -minor=0 -patch=0
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
