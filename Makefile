SHELL=/bin/bash -o pipefail

export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on

DOCKER_DOC_TAG ?= latest

.PHONY: all
all: build

.PHONY: mod
mod:
	@go mod download

.PHONY: build
build: build-clients build-servers ## build all modules

.PHONY: build-clients
build-clients: build-v1-client  ## build all clients

.PHONY: build-v1-client
build-v1-client:  ## build only client of v1
	./v1/indirect/update-client.sh
	@go mod tidy
	@go build ./v1/client

.PHONY: build-servers
build-servers: build-v1-server ## build all servers

.PHONY: build-v1-server
build-v1-server: ## build only server of v1
	./v1/indirect/update-server.sh
	@go build ./v1/server

.PHONY: open
open: ## open all index.html files
	 open ./**/html/index.html

.PHONY: help
help: ## display help screen
	@grep -E '^[a-zA-Z_GO111MODULE=on go mod tidy-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

HOOK_FILE_SOURCE=".indirect/pre-commit.sh"
HOOK_FILE_TARGET=".git/hooks/pre-commit"

SELECT_LIST= yes no
.PHONY: hook
hook: ## add pre-commit git hook file
	@if test -f "${HOOK_FILE_TARGET}"; then \
	    echo -e "\nPre-commit file already exist: ${HOOK_FILE_TARGET}\nReplace file?"; \
	    select SELECTED in ${SELECT_LIST}; \
	    do test -n "$$SELECTED" && break; \
	        echo -e "\x1b[31;01mInvalid Selection\x1b[0m"; \
	    done; \
        if [ "$${SELECTED}" != "yes" ]; then \
            echo -e "\033[0;32mSkipped\033[0m"; \
            exit 0; \
        fi; \
	fi; \
	echo -n > .git/hooks/pre-commit; \
    echo -e "#!/bin/bash" >> .git/hooks/pre-commit; \
    echo -e ".indirect/pre-commit.sh" >> .git/hooks/pre-commit; \
    chmod 0755 ${HOOK_FILE_TARGET}; \
    echo -e "\033[0;32mHook added successfully: ${HOOK_FILE_TARGET}\033[0m";
