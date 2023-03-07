SHELL := /bin/bash
current_dir = $(shell pwd)

# detect GOPATH if not set
ifndef $(GOPATH)
    $(info GOPATH is not set, autodetecting..)
    TESTPATH := $(dir $(abspath ../../..))
    DIRS := bin pkg src
    # create a ; separated line of tests and pass it to shell
    MISSING_DIRS := $(shell $(foreach entry,$(DIRS),test -d "$(TESTPATH)$(entry)" || echo "$(entry)";))
    ifeq ($(MISSING_DIRS),)
        $(info Found GOPATH: $(TESTPATH))
        export GOPATH := $(TESTPATH)
    else
        $(info ..missing dirs "$(MISSING_DIRS)" in "$(TESTDIR)")
        $(info GOPATH autodetection failed)
    endif
endif

ifeq (, $(shell which golangci-lint))
	LINTER := docker run -it --rm -v $(shell pwd):/app -v "$(GOPATH)/pkg:/go/pkg" -w /app  golangci/golangci-lint:v1.50.1 golangci-lint run
else
	LINTER := golangci-lint run
endif

all: help

help :
	@echo "Help information, please run specific target:"
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for help_line in $${help_lines[@]}; do \
		IFS=$$'#' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf " - %-20s %s\n" $$help_command $$help_info ; \
	done


install: ## Setting up API
	@docker-compose down -v
	@docker-compose build
	@docker-compose up -d db && printf "init DB ..." && sleep 10 && echo " continue" && sleep 1
	@docker-compose up -d

local: ## Setting up API in dev mode
	@cp deployments/docker-compose.dev.override.yml docker-compose.override.yml
	@touch .env.local
	@make install

destroy: ## Uninstall project environment
	@read -p $$'\e  \033[0;36mDestroy old environment?\033[0m \033[0;33mWARNING: It will remove all images, local configurations and database!\033[0m [\033[0;32mno\033[0m]: ' -r destroy; \
	destroy=$${destroy:-"no"}; \
	if [[ $$destroy =~ ^[yY][eE][sS]|[yY]$$ ]]; then \
		printf "\033[0;36mStop and remove docker containers, images, volumes ...\033[0m\n"; \
		docker-compose down -v --remove-orphans --rmi local; \
		printf "\033[0;32mApplication environment destroyed!\033[0m\n"; \
	fi;

ifeq (lint, $(firstword $(MAKECMDGOALS)))
  lintargs := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
  $(eval $(lintargs):;@true)
endif

lint: ## Run golang linter
	@go mod tidy
	@$(LINTER) $(lintargs)