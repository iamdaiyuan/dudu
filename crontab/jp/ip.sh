#!/bin/sh
export PATH=/usr/local/python3/bin:/data/app/go/bin:$PATH
export PYTHONPATH=/data/www/python/smartdo
export GOPATH=/data/www/web/go
export GOBIN=/data/app/go/bin
export GOROOT=/data/app/go
nohup go run /data/www/web/go/src/github.com/beautytop/dudu/spiders/jp/ippool.go  > /data/www/web/go/src/github.com/beautytop/dudu/sh/jp/ippool.log 2>&1 &