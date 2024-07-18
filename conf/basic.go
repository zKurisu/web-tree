package conf

import (
	"os"
	"path/filepath"
)

const CONFIG_FILE = "example.yaml"

var CONFIG_DIR = filepath.Join(os.Getenv("HOME"), "/.config/web-tree/")
var STORE_DIR = filepath.Join(CONFIG_DIR, "data")
var BACK_DIR = filepath.Join(STORE_DIR, ".backup")
var TREE_TEMPELATE = `name:  %
tree: []
nodes: []
`
var CONFIG_TEMPLATE = `main:
  browser: 
    - chromium
    - edge
    - firefox
  sort: "label"
  help:
    hidden: false

search:
  keybinding: "<C-s>"
  style: "None"

tab:
  keybinding: "<Tab>"
  style: "None"

tree:
  keybinding:
    UP:     "k"
    DOWN:   "j"
    LEFT:   "h"
    RIGHT:  "l"
    OPEN:   "o"
    ADD:    "a"
    DELETE: "d"
    EDIT:   "e"
  style: "None"
`

func GetConfigName() string {
	return CONFIG_FILE
}

func GetStoreDir() string {
	return STORE_DIR
}

func GetBackDir() string {
	return BACK_DIR
}

func GetConfigDir() string {
	return CONFIG_DIR
}

func GetConfigPath() string {
	return filepath.Join(CONFIG_DIR, CONFIG_FILE)
}

func GetTreeTemp() string {
	return TREE_TEMPELATE
}

func GetConfigTemplate() string {
	return CONFIG_TEMPLATE
}
