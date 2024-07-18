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
var CONF = GetConf()

type Conf struct {
	Main struct {
		Browser []string `yaml:"browser"`
		Sort    string   `yaml:"sort"`
		Help    struct {
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

func IsConfExist() bool {
	// ...
	_, err := os.Stat(CONFIG_PATH)
	if err != nil {
		return false
	}
	return true
}

func InitConf() {
	// ...
	configDir := conf.GetConfigDir()
	configInitContent := conf.GetConfigTemplate()

	_, err := os.Stat(configDir)
	if err != nil {
		os.Mkdir(configDir, 0755)
		log.Println("Create dir: " + configDir)
	}

	f, err := os.OpenFile(CONFIG_PATH, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(configInitContent); err != nil {
		log.Fatal(err)
	}

	log.Println("Create config file: " + CONFIG_PATH + " ...ok")
	log.Println("Run [web-tree] again to open TUI!")
}

func GetConf() Conf {
	yamlContent, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		yamlContent = []byte(conf.GetConfigTemplate())
	}

	var c Conf
	if err = yaml.Unmarshal(yamlContent, &c); err != nil {
		fmt.Println("When running Unmarshal:")
		log.Fatal(err)
	}

	return c
}

func (c Conf) GetBrowser() []string {
	return c.Main.Browser
}
