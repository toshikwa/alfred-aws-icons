package awsutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
	"gopkg.in/yaml.v2"
)

type AwsService struct {
	Id               string   `yaml:"id"`
	Name             string   `yaml:"name"`
	ShortName        string   `yaml:"short_name"`
	ExtraSearchTerms []string `yaml:"extra_search_terms"`
}

func ParseAwsServices(yamlPath string) []AwsService {
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Fatal(err)
	}

	awsServices := []AwsService{}
	if err = yaml.Unmarshal(yamlFile, &awsServices); err != nil {
		log.Fatal(err)
	}
	return awsServices
}

func AddAwsService(wf *aw.Workflow, awsService AwsService) {
	title := awsService.Id
	subtitle := awsService.Name
	alt := awsService.Name
	match := awsService.Name

	if awsService.ShortName != "" {
		alt = awsService.ShortName
		match = awsService.ShortName + " " + match
		subtitle += " (" + awsService.ShortName + ")"
	}

	if len(awsService.ExtraSearchTerms) > 0 {
		match += " " + strings.Join(awsService.ExtraSearchTerms, " ")
	}

	image_path, err := filepath.Abs("images/svg/" + awsService.Id + ".svg")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	match = strings.ToLower(match)
	match = strings.Replace(match, "amazon ", "", -1)
	match = strings.Replace(match, "aws ", "", -1)
	match = strings.Replace(match, "(", "", -1)
	match = strings.Replace(match, ")", "", -1)

	item := wf.NewItem(title).
		Valid(true).
		Subtitle(subtitle).
		Var("action", "run-script").
		Match(match).
		UID(awsService.Id).
		Arg(image_path + "," + alt)
	item.Cmd().
		Subtitle(subtitle).
		Var("action", "run-script-cmd")

	item.Icon(&aw.Icon{Value: "images/png/" + awsService.Id + ".png"})
}

func AddAwsServices(wf *aw.Workflow, awsServices []AwsService) {
	for _, awsService := range awsServices {
		AddAwsService(wf, awsService)
	}
}
