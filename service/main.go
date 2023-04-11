package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/ku2482/alfred-aws-icons/icon"
	"github.com/ku2482/alfred-aws-icons/workflow"
)

var (
	svcwf *aw.Workflow
	abbrs icon.Abbreviations
	query string
)

func init() {
	svcwf = aw.New()
	abbrs = icon.LoadAbbreviations("abbreviations.yaml")
	icon.LoadArchitectureIcons(
		svcwf,
		"./assets/Architecture-Service-Icons/Arch_*/*64",
		"Arch_",
		"_64.svg",
		"_64@5x.png",
		abbrs,
	)
}

func run() {
	flag.StringVar(&query, "query", "", "query to use")
	flag.Parse()
	workflow.Run(svcwf, query)
}

func main() {
	svcwf.Run(run)
}
