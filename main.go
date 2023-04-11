package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/ku2482/alfred-aws-icons/icon"
	"github.com/ku2482/alfred-aws-icons/workflow"
)

var (
	wf    *aw.Workflow
	abbrs icon.Abbreviations
	query string
)

func init() {
	wf = aw.New()
	abbrs = icon.LoadAbbreviations("abbreviations.yaml")
	icon.LoadArchitectureIcons(
		wf,
		"assets/Architecture-Service-Icons/Arch_*/*64",
		"Arch_",
		"_64.svg",
		"_64@5x.png",
		abbrs,
	)
	icon.LoadResourceIcons(
		wf,
		"assets/Resource-Icons/Res_*/Res_48_Light",
		"Res_",
		"_48_Light.svg",
		"_48_Light.png",
		"Light",
		abbrs,
	)
	icon.LoadResourceIcons(
		wf,
		"assets/Resource-Icons/Res_*/Res_48_Dark",
		"Res_",
		"_48_Dark.svg",
		"_48_Dark.png",
		"Dark",
		abbrs,
	)
}

func run() {
	flag.StringVar(&query, "query", "", "query to use")
	flag.Parse()
	workflow.Run(wf, query)
}

func main() {
	wf.Run(run)
}
