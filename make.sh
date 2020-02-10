#!/bin/sh
gofmt -w *.go && gofmt -w */*.go

gSource=$(cd `dirname $0`; pwd)

export GO111MODULE=on
export GOPROXY=https://goproxy.io
export CGO_ENABLED=0
export GOOS=linux

###返回源码路径，编译项目
cd ${gSource}
COMMIT_NUMBER=`git rev-list HEAD | wc -l | awk '{print $1}'`
HASH_NUMBER=`git rev-parse --short HEAD`
LD_FLAGS="-X newplatform/version.HashNumber=${HASH_NUMBER} -X newplatform/version.CommitNumber=${COMMIT_NUMBER}"
go build -ldflags "$LD_FLAGS" -o platform
