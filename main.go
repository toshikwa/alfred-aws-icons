package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/toshikwa/alfred-aws-icons/icon"
)

var (
	doCheck           bool
	query             string
	mode              string
	abbrs             = icon.LoadAbbreviations("assets/abbreviations.yaml")
	updateIcon        = &aw.Icon{Value: "assets/update-available.png"}
	repo              = "toshikwa/alfred-aws-icons"
	checkForUpdateJob = "checkForUpdate"
	wf                *aw.Workflow
)

func init() {
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	flag.StringVar(&mode, "mode", "svc", "'svc' or 'res'")
	wf = aw.New(update.GitHub(repo))
}

func run() {
	wf.Args()
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}
	defer finalize()

	// check for update
	if doCheck {
		wf.Configure(aw.TextErrors(true))
		log.Println("Checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	// execute command to check
	if wf.UpdateCheckDue() && !wf.IsRunning(checkForUpdateJob) {
		log.Println("Running update check in background...")
		cmd := exec.Command(os.Args[0], "-check")
		if err := wf.RunInBackground(checkForUpdateJob, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	// show update item
	if query == "" {
		if wf.UpdateAvailable() {
			wf.Configure(aw.SuppressUIDs(true))
			wf.NewItem("[alfred-aws-icons] An update is available!!").
				Subtitle("Press Enter to install update").
				Valid(false).
				Autocomplete("workflow:update").
				Icon(updateIcon)
		} else {
			wf.NewItem("Search for an AWS Icon...").Subtitle("e.g. `ic fargate`, `icr ecs task`")
		}
	}

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

	// filter results
	if query != "" {
		wf.Filter(strings.ToLower(query))
	}
}

func finalize() {
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
