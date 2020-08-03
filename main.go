package main

import (
	"chango/cmd"
	"chango/internal/config"
)

func main() {
	config.Config.UpdateConfig()
	cmd.Execute()
}
