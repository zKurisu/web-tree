package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"web-tree/conf"
)

var CONFIG_FILE = conf.GetConfigName()
var CONFIG_PATH = conf.GetConfigPath()

type Conf struct {
	Main struct {
		Sort string `yaml:"sort"`
		Help struct {
			Hidden bool `yaml:"hidden"`
		} `yaml:"help"`
	} `yaml:"main"`

	Search struct {
		Keybinding string      `yaml:"keybinding"`
		Style      interface{} `yaml:"style"`
	} `yaml:"search"`

	Tab struct {
		Keybinding string      `yaml:"keybinding"`
		Style      interface{} `yaml:"style"`
	} `yaml:"tab"`

	Tree struct {
		Keybinding struct {
			UP     string `yaml:"UP"`
			DOWN   string `yaml:"DOWN"`
			LEFT   string `yaml:"LEFT"`
			RIGHT  string `yaml:"RIGHT"`
			OPEN   string `yaml:"OPEN"`
			ADD    string `yaml:"ADD"`
			DELETE string `yaml:"DELETE"`
			EDIT   string `yaml:"EDIT"`
		} `yaml:keybinding`
		Style interface{} `yaml:"style"`
	} `yaml:"tree"`
}

func GetConf() Conf {
	yamlContent, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		fmt.Println("When read yaml file:")
		log.Fatal(err)
	}

	var c Conf
	if err = yaml.Unmarshal(yamlContent, &c); err != nil {
		fmt.Println("When running Unmarshal:")
		log.Fatal(err)
	}

	return c
}
