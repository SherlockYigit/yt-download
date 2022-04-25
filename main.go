package main

import (
	"fmt"
	"github.com/SherlockYigit/kommando"
	"github.com/SherlockYigit/youtube-go"
	"github.com/manifoldco/promptui"
	"strings"
)

func main() {
	handler := kommando.NewKommando(kommando.KommandoConfig{
		AppName:             "yt-download",
		Template:            "Welcome to {AppName}! That's a command list. Type 'help <command name>' to get help with any command.\n{CommandList}",
		CommandListTemplate: "{CommandName} | {CommandDescription}",
		CommandHelpTemplate: "{CommandName} | Info\n{CommandDescription}\n{FlagList}\n{CommandAliases}",
		FlagListTemplate:    "--{FlagName} {FlagDescription}",
	})

	handler.AddCommand(kommando.Command{
		Name:        "download",
		Description: "It searches and downloads the relevant file according to the selected result.",
		Flags: []kommando.Flag{
			kommando.Flag{
				Name:        "path",
				Description: "Location to save the downloaded file.",
			},
		},
		Execute: func(Res kommando.CommandResponse) {
			var squery string
			var flags []string

			for i := 0; len(Res.Args) > i; i++ {
				val := Res.Args[i]

				if !strings.Contains(val, "=") {
					squery = val
				} else {
					flags = append(flags, val)
				}
			}

			res := youtubego.Search(squery, youtubego.SearchOptions{
				Type:  "video",
				Limit: 15,
			})

			prompt := promptui.Select{
				Label: "Select video",
				Items: getTitles(res),
			}

			selectedIndex, _, err := prompt.Run()

			if err != nil {
				panic("Prompt failed " + err.Error())
			}

			fmt.Println(res[selectedIndex].Title)
		},
	})

	handler.Run()
}

func getTitles(val []youtubego.SearchResult) []string {
	var titles []string

	for i := 0; len(val) > i; i++ {
		titles = append(titles, val[i].Video.Title)
	}

	return titles
}
