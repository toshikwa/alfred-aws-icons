package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/toshikwa/alfred-aws-icons/icon"
	"github.com/toshikwa/alfred-aws-icons/workflow"
)

var (
	wf    *aw.Workflow
	abbrs icon.Abbreviations
	query string
	mode  string
)

func init() {
	wf = aw.New()
	abbrs = icon.LoadAbbreviations("abbreviations.yaml")
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
			"./assets/Architecture-Service-Icons/Arch_*",
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
			"./assets/Resource-Icons/Res_General-Icons/Res_Light",
			"Res_",
			"_64_Light.svg",
			"_48_Light.png",
			"Light",
			abbrs,
		)
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_General-Icons/Res_Dark",
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
	workflow.Run(wf, query)
}

func main() {
	wf.Run(run)
}
