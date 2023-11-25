#!/bin/bash

# Simple script to cross-compile Krampus for multiple platforms quicker.

readonly krampus_version='v12' # Change per version release ( v1.2 / v12 ).

pre_GOOS=$GOOS
pre_GOARCH=$GOARCH

# Linux
export GOOS='linux'; export GOARCH='amd64'; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS='linux'; export GOARCH='arm64'; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"

# Windows
export GOOS="windows"; export GOARCH="amd64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS="windows"; export GOARCH="386"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS="windows"; export GOARCH="arm64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"

# Darwin
export GOOS="darwin"; export GOARCH="amd64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS="darwin"; export GOARCH="arm64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"

# FreeBSD
export GOOS="freebsd"; export GOARCH="amd64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS="freebsd"; export GOARCH="arm64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"

# OpenBSD
export GOOS="openbsd"; export GOARCH="amd64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"
export GOOS="openbsd"; export GOARCH="arm64"; go build -o krampus_${krampus_version}_${GOOS}_${GOARCH} -ldflags="-s -w"

# Set current shell back to pre-compilation ENV variables.
export GOOS=$pre_GOOS
export GOARCH=$pre_GOARCH