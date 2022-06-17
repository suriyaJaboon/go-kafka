# Go-Kafka

Experimenting with **golang-1.18**

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xmlking/go-workspace)](https://github.com/xmlking/go-workspace/blob/main/go.mod)
[![Go](https://github.com/xmlking/go-workspace/actions/workflows/go.yml/badge.svg)](https://github.com/xmlking/go-workspace/actions/workflows/go.yml)

## Install

### Working with golang 1.18

```shell
brew install go 
$ go version 
go version go1.18 darwin/arm64
```

###
```shell
go test -v ./... 
```

### Workspace commands

```
$ go help work
Usage:

        go work <command> [arguments]

The commands are:

        edit        edit go.work from tools or scripts
        init        initialize workspace file
        sync        sync workspace build list to modules
        use         add modules to workspace file

```

Run `go work use -r ./` to recursively add directories in the argument directory with a go.mod file to your workspace. 
If a directory doesn’t have a go.mod file, or no longer exists, the use directive for that directory is removed from your go.work file.

```shell
# recursively add directories to go.work
go work use -r .
# pushes the dependencies in the go.work file back into the go.mod files of each workspace module.
go work sync
# provides a command-line interface for editing go.work, for use primarily by tools or scripts.
go work edit
# `go mod` examples
go mod download
go mod graph
go mod tidy
go mod verify
go work sync
go mod why -m github.com/ssoroka/slice
```