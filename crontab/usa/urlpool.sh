#!/bin/sh
export PATH=/usr/local/python3/bin:/data/app/go/bin:$PATH
export PYTHONPATH=/data/www/python/smartdo
export GOPATH=/data/www/web/go
export GOBIN=/data/app/go/bin
export GOROOT=/data/app/go
nohup go run /data/www/web/go/src/github.com/hunterhug/AmazonGo/spiders/usa/urlpool.go  > /data/www/web/go/src/github.com/hunterhug/AmazonGo/sh/usa/urlpool.log 2>&1