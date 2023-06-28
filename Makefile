SHELL := /bin/bash

PLIST=info.plist
BIN=alfred-aws-icons
ICON=./icon.png
YAML=./abbreviations.yaml
ASSETS=./assets
DIST_FILE=aws-icons.alfredworkflow

all: $(DIST_FILE)

$(BIN):
	go build -o $(BIN) ./main.go

$(DIST_FILE): $(BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)
	zip -r $(DIST_FILE) $(BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)