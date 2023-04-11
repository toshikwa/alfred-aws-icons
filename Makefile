SHELL := /bin/bash

PLIST=info.plist
EXEC_BIN=alfred-aws-icons
ICON=./icon.png
YAML=./abbreviations.yaml
ASSETS=./assets
DIST_FILE=aws-icons.alfredworkflow

all: $(DIST_FILE)

$(EXEC_BIN):
	go build -o $(EXEC_BIN) .

$(DIST_FILE): $(EXEC_BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)
	zip -r $(DIST_FILE) $(EXEC_BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)