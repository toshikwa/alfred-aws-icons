package workflow

import (
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/ku2482/alfred-aws-icons/awsutil"
)

func Run(wf *aw.Workflow, query string, yamlPath string) {
	if strings.Trim(query, " ") == "" {
		handleEmptyQuery(wf)
		return
	}
	defer finalize(wf)

	awsServices := awsutil.ParseAwsServices(yamlPath)
	awsutil.AddAwsServices(wf, awsServices)
	wf.Filter(strings.ToLower(query))
}

func handleEmptyQuery(wf *aw.Workflow) {
	wf.NewItem("Search for an AWS Icon...").
		Subtitle("e.g. ec2, s3, fargate ...")
}

func finalize(wf *aw.Workflow) {
	if r := recover(); r != nil {
		panic(r)
	}
	if wf.IsEmpty() {
		wf.NewItem("No matching AWS Icon found.").
			Subtitle("Try another query (e.g. `icon fargate`)").
			Icon(aw.IconNote)
	}
	wf.SendFeedback()
}
