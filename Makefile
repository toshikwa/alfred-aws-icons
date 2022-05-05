SHELL := /bin/bash

PLIST=info.plist
EXEC_BIN=alfred-aws-icons
ICON=./icon.png
SERVICE=./services.yaml
IMAGES=./images
GO_SRCS=$(shell find -f . \( -name \*.go \))
DIST_FILE=aws-icons.alfredworkflow

all: $(DIST_FILE)

$(EXEC_BIN): $(GO_SRCS)
	go build -o $(EXEC_BIN) .

$(DIST_FILE): $(EXEC_BIN) $(PLIST) $(ICON) $(SERVICE) $(IMAGES)
	zip -r $(DIST_FILE) $(PLIST) $(EXEC_BIN) $(ICON) $(SERVICE) $(IMAGES)