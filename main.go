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
			"./assets/Architecture-Service-Icons/Arch_*/*64",
			"Arch_",
			"_64.svg",
			"_64@5x.png",
			abbrs,
		)
	} else if mode == "res" {
		// service resource icons
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_*",
			"Res_",
			"_48.svg",
			"_48.png",
			"",
			abbrs,
		)
		// General icons
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_General-Icons/Res_48_Light",
			"Res_",
			"_48_Light.svg",
			"_48_Light.png",
			"Light",
			abbrs,
		)
		icon.LoadResourceIcons(
			wf,
			"./assets/Resource-Icons/Res_General-Icons/Res_48_Dark",
			"Res_",
			"_48_Dark.svg",
			"_48_Dark.png",
			"Dark",
			abbrs,
		)
	}
	workflow.Run(wf, query)
}

func main() {
	wf.Run(run)
}