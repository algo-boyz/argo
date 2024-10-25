SHELL := /bin/bash
.DEFAULT_GOAL := build

BIN         = $(CURDIR)/bin
BUILD_DIR   = $(CURDIR)/build
# Targets
NAME = $(shell basename $(CURDIR))
TARGET = $(BUILD_DIR)/$(NAME).o
OUT = $(BIN)/$(NAME)

VERSION     ?= $(shell git describe --tags --always --match=v*)
SHORT_COMMIT ?= $(shell git rev-parse --short HEAD)

# Printing
V ?= 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

$(BIN):
	@mkdir -p $@
$(BUILD_DIR):
	@mkdir -p $@


.PHONY: build-mac # Build mac executable
mac: $(BIN) $(BUILD_DIR)
	$(info $(M) building mac executable…)
	$Q as $(NAME).s -o $(TARGET)
	$Q ld $(TARGET) -o $(OUT) -l System -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _start -arch arm64
	@true

.PHONY: linux # Build linux executable
linux: $(BIN) $(BUILD_DIR)
	$(info $(M) building linux executable…)
	$Q set -o nounset
	$Q set -o errexit
	$Q gcc -o $(TARGET) $(NAME).s
	@true

.PHONY: fmt # Run gofmt on go source files
fmt:
	$(info $(M) running fmt…)
	$Q

.PHONY: clean # Cleanup project root
clean:
	$(info $(M) cleaning…)
	@rm -rf $(BIN)
	@rm -rf $(BUILD_DIR)

.PHONY: help # Display help
help:
	@grep  -E '^.PHONY' $(MAKEFILE_LIST) | awk 'BEGIN {FS = "#|: "}; {printf "\033[36m%-20s\033[0m %s\n",$$2,$$3}'