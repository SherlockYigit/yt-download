package main

import (
	"fmt"
	"github.com/SherlockYigit/kommando"
	"github.com/SherlockYigit/youtube-go"
	"github.com/manifoldco/promptui"
)

func main() {
	handler := kommando.NewKommando(kommando.KommandoConfig{
		AppName:             "yt-download",
		Template:            "Welcome to {AppName}! That's a command list. Type 'help <command name>' to get help with any command.\n{CommandList}",
		CommandListTemplate: "{CommandName} | {CommandDescription}",
		CommandHelpTemplate: "{CommandName} | Info\n{CommandDescription}\n{FlagList}\n{CommandAliases}",
		FlagListTemplate:    "--{FlagName} {FlagDescription}",
	})

	handler.Run()
}
