package main

import (
	"web-tree/cmd"
	"web-tree/utils"
)

func main() {
	if utils.IsConfExist() {
		cmd.Execute()
	} else {
		utils.InitConf()
	}
}
