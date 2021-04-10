package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	Hits           int64       `yaml:"hits"`
	Time           int64       `yaml:"time"`
	MissingInteger int64       `yaml:"missing_integer"`
	MissingFloat   float64     `yaml:"missing_float"`
	MissingString  string      `yaml:"missing_field"`
	MissingMap     interface{} `yaml:"missing_map"`
	FreeMap        interface{} `yaml:"free_map"`
	Name           string      `yaml:"name"`
	StructOnly1    string
	StructOnly2    string
	Integers       []int64   `yaml:"integers"`
	Floats         []float64 `yaml:"floats"`
}

func (c *config) loadConfig(filename string) *config {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c config

	// Fields that are not in the struct are ignored!
	c.loadConfig("./manifest.yaml")

	c.StructOnly1 = "StructOnly1"

	// When printed a values are ordered based on the field declaration order in struct
	fmt.Println(c)
	if c.MissingString == "" {
		fmt.Println(c.Name)
		fmt.Println(c.MissingInteger)
		fmt.Println(c.MissingFloat)
	}
	if c.MissingMap == nil {
		fmt.Println("Missing section set to <nil>")
	} else {
		fmt.Println("Found some content")
	}
	if c.FreeMap != nil {
		fmt.Println("free_map is set")
	}
}
