SHELL = /bin/sh
.PHONY: all build bump_version clean coverage dist format install lint maintainer-clean test test_all updatedeps version vet

# Set the src directory if it is not overwritten on the commandline.
# You can overwrite this by setting your build command to `make srcdir=path <target>`
ifndef SRCDIR
SRCDIR = ./src
endif

ifndef OSARCH
OSARCH = linux/amd64
endif

ifndef PKGBASE
PKGBASE = github.com
endif

ifndef PKG
PKG := $(shell pwd | awk -F/ '{ print $$NF }')
endif

ifndef OUTPUT
	ifeq ("$(OSARCH)","linux/amd64")
		OUT = ./bin/$(PKG)
	else
			OUT = ./bin/$(PKG)_{{.OS}}_{{.Arch}}
	endif
else
	ifeq ("$(OSARCH)","linux/amd64")
		OUT = ./bin/$(OUTPUT)
	else
			OUT = ./bin/$(OUTPUT)_{{.OS}}_{{.Arch}}
	endif
endif

ifndef TARGET_PATH
TARGET_PATH = target
endif

ifndef DESTDIR
DESTDIR = /usr/local/bin
endif

# This will format the code, run tests/linters, and then build/package the code
default: all

# We only care about go files at the moment so clear and explictly denote that
.SUFFIXES:
.SUFFIXES: .go

# build and then create a tarball in the target directory
# basically everything needed by drteeth to put it into artifactory
all:
	@echo "this needs to be implemented"

# Build a binary from the given package and drop it into the local bin
build:
		gox -osarch="$(OSARCH)" -output=$(OUT) $(PKGBASE)/$(PKG)/src

# this will bump the version of the project
bump_version:
	@echo "this needs to be implemented"

# delete all files used for building
clean:
	@echo "this needs to be implemented"

# run the go coverage tool
coverage:
	@echo "this needs to be implemented"

# pack everything up neatly
dist: build
	tar -czvpf $(TARGET_PATH)/$(PKG).tgz $$GOPATH/src/$(PKGBASE)/$(PKG)/bin/*

# run the go formatting tool on all files in the current src directory
format:
	@gofmt -w $(SRCDIR)/*.go

# install the binary locally for testing
install:
	cp $$GOPATH/src/$(PKGBASE)/$(PKG)/bin/* $(DESTDIR)

# run the go linting tool
lint:
	@golint $(SRCDIR)/*.go

# @echo 'This command is intended for maintainers to use; it'
# @echo 'deletes files that may need special tools to rebuild.'
maintainer-clean:
	@echo "this needs to be implemented"

# run unit tests and anything else testing wise needed
test:
	@echo "this needs to be implemented"

# run unit tests, vet, and liniting combined
test_all:
	@echo "this needs to be implemented"

# update all deps to the latest versions available
updatedeps:
	@go list ./... \
		| xargs go list -f '{{join .Deps "\n"}}' \
		| sort -u \
		| xargs go get -f -u -v

# print out the current version of the project
version:
	@echo "this needs to be implemented"

# run go vet
vet:
	@echo "this needs to be implemented"
