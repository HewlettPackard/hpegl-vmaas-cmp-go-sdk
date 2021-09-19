#! /usr/bin/make
#(C) Copyright 2021 Hewlett Packard Enterprise Development LP
# Inspiration from https://github.com/rightscale/go-boilerplate/blob/master/Makefile

VERSION=0.0.1


default: all
.PHONY: default

vendor: go.mod go.sum
	go mod download
	go mod vendor

update up: really-clean vendor
.PHONY: update up

clean:
	rm -rf gathered_logs build .vendor/pkg $(testreport_dir) $(coverage_dir)
.PHONY: clean

really-clean clean-all cleanall: clean
	rm -rf vendor
.PHONY: really-clean clean-all cleanall

procs := $(shell grep -c ^processor /proc/cpuinfo 2>/dev/null || echo 1)
# TODO make --debug an option

lint: vendor golangci-lint-config.yaml
	@golangci-lint --version
	golangci-lint run --config golangci-lint-config.yaml
.PHONY: lint

testreport_dir := test-reports
test:
	go generate -v ./...
	go test -v ./...
.PHONY: test

coverage_dir := coverage/go
coverage: vendor
	@mkdir -p $(coverage_dir)/html
	go test -coverpkg=./... -coverprofile=$(coverage_dir)/coverage.out -v $$(go list ./... | grep -v /vendor/)
	@go tool cover -html=$(coverage_dir)/coverage.out -o $(coverage_dir)/html/main.html;
	@echo "Generated $(coverage_dir)/html/main.html";
.PHONY: coverage

build: vendor
	@go build ./...
.PHONY: build


all: lint test
.PHONY: all

tools:
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0
.PHONY: tools
