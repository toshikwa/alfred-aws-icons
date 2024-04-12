package main

import (
	"flag"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/toshikwa/alfred-aws-icons/icon"
)

var (
	query      string
	mode       string
	abbrs      icon.Abbreviations
	updateIcon *aw.Icon
	wf         *aw.Workflow
)

func init() {
	abbrs = icon.LoadAbbreviations("abbreviations.yaml")
	updateIcon = &aw.Icon{Value: "assets/update-available.png"}
	wf = aw.New(update.GitHub("toshikwa/alfred-aws-icons"))
}

func run() {
	// args
	flag.StringVar(&mode, "mode", "svc", "'svc' or 'res'")
	flag.StringVar(&query, "query", "", "query to use")
	flag.Parse()

	if mode == "svc" {
		// service icons
		icon.LoadArchitectureIcons(
			wf,
			"./assets/Architecture-Service-Icons/Arch_*/64",
			"Arch_",
			"_64.svg",
			"_64.png",
			abbrs,
		)
	} else if mode == "res" {
		// service resource icons
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_*",
			"Res_",
			"_64.svg",
			"_48.png",
			"",
			abbrs,
		)
		// general icons
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_General-Icons/Res_48_Light",
			"Res_",
			"_64_Light.svg",
			"_48_Light.png",
			"Light",
			abbrs,
		)
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_General-Icons/Res_48_Dark",
			"Res_",
			"_64_Dark.svg",
			"_48_Dark.png",
			"Dark",
			abbrs,
		)
		// group icons
		icon.LoadResourceIcons(
			wf,
			"./assets/Architecture-Group-Icons",
			"",
			"_32.svg",
			"_32.png",
			"",
			abbrs,
		)
	}

	defer finalize(wf)
	if strings.Trim(query, " ") == "" {
		// show example query
		wf.NewItem("Search for an AWS Icon...").Subtitle("e.g. `ic fargate`, `icr ecs task`")
		// check for update
		if wf.UpdateAvailable() {
			wf.Configure(aw.SuppressUIDs(true))
			wf.NewItem("An update is available!").
				Subtitle("Press Enter to install update").
				Valid(false).
				Autocomplete("workflow:update").
				Icon(updateIcon)
		}
	} else {
		wf.Filter(strings.ToLower(query))
	}
}

func finalize(wf *aw.Workflow) {
	if r := recover(); r != nil {
		panic(r)
	}
	if wf.IsEmpty() {
		wf.NewItem("No matching AWS Icon found.").
			Subtitle("Try another query (e.g. `ic fargate`, `icr ecs task`)").
			Icon(aw.IconNote)
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
