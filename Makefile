HOSTOS := $(shell uname -s)
export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)
export CGO_ENABLED ?= 0
export GOPATH := $(shell go env GOPATH)

all:
	go build
