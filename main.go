package main

import (
	"flag"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/ku2482/alfred-aws-icons/workflow"
)

var (
	wf       *aw.Workflow
	query    string
	yamlPath string
)

func init() {
	flag.StringVar(&query, "query", "", "query to use")
	flag.StringVar(&yamlPath, "yaml_path", "icons.yaml", "config file")
	flag.Parse()
	wf = aw.New()
}

func run() {
	workflow.Run(wf, strings.TrimLeft(query, " "), yamlPath)
}

func main() {
	wf.Run(run)
}
