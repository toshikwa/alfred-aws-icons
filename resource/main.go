package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/toshikwa/alfred-aws-icons/icon"
	"github.com/toshikwa/alfred-aws-icons/workflow"
)

var (
	reswf *aw.Workflow
	abbrs icon.Abbreviations
	query string
)

func init() {
	reswf = aw.New()
	abbrs = icon.LoadAbbreviations("abbreviations.yaml")
	icon.LoadResourceIcons(
		reswf,
		"./assets/Resource-Icons/Res_*/Res_48_Light",
		"Res_",
		"_48_Light.svg",
		"_48_Light.png",
		"Light",
		abbrs,
	)
	icon.LoadResourceIcons(
		reswf,
		"./assets/Resource-Icons/Res_*/Res_48_Dark",
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
	workflow.Run(reswf, query)
}

func main() {
	reswf.Run(run)
}
