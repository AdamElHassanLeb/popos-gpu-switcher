APP_NAME := popos-gpu-switcher
BUILD_DIR := build
CMD := ./cmd

GO := go
GOFLAGS := -trimpath

.PHONY: all build clean run deps

all: build

deps:
	$(GO) mod tidy

build:
	mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(CMD)

run:
	$(GO) run $(CMD)

clean:
	rm -rf $(BUILD_DIR)
