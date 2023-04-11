SHELL := /bin/bash

PLIST=info.plist
SVC_BIN=svc.alfred-aws-icons
RES_BIN=res.alfred-aws-icons
ICON=./icon.png
YAML=./abbreviations.yaml
ASSETS=./assets
DIST_FILE=aws-icons.alfredworkflow

all: $(DIST_FILE)

$(SVC_BIN):
	go build -o $(SVC_BIN) ./service/main.go

$(RES_BIN):
	go build -o $(RES_BIN) ./resource/main.go

$(DIST_FILE): $(SVC_BIN) $(RES_BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)
	zip -r $(DIST_FILE) $(SVC_BIN) $(RES_BIN) $(PLIST) $(ICON) $(YAML) $(ASSETS)