PROJECT_DIR := $(CURDIR)

setup:
	cd $(PROJECT_DIR)/main_module && go mod tidy

build:
	cd $(PROJECT_DIR)/main_module && go build

run:
	cd $(PROJECT_DIR)/main_module && go run main.go

all: setup build run

.PHONY: setup build run all