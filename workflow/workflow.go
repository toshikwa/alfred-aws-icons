package workflow

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

func Run(wf *aw.Workflow, query string) {
	if strings.Trim(query, " ") == "" {
		handleEmptyQuery(wf)
		return
	}
	defer finalize(wf)
	wf.Filter(strings.ToLower(query))
}

func handleEmptyQuery(wf *aw.Workflow) {
	wf.NewItem("Search for an AWS Icon...").
		Subtitle("e.g. `ic fargate`, `icr ecs task`")
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
