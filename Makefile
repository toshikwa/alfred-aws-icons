SHELL := /bin/bash

PLIST=info.plist
EXEC_BIN=alfred-aws-icons
ICON=./icon.png
YAML=./icons.yaml
IMAGES=./images
DIST_FILE=aws-icons.alfredworkflow

all: $(DIST_FILE)

$(EXEC_BIN):
	go build -o $(EXEC_BIN) .

$(DIST_FILE): $(EXEC_BIN) $(PLIST) $(ICON) $(YAML) $(IMAGES)
	zip -r $(DIST_FILE) $(EXEC_BIN) $(PLIST) $(ICON) $(YAML) $(IMAGES)