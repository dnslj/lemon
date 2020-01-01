#!/bin/bash

BASE_DIR=$(echo $PWD)

buildServer() {
  echo 'server building...'
  export GOPROXY=https://goproxy.cn

  cd $BASE_DIR
  if [ ! -d logs ]; then mkdir logs; fi

  gofmt -w -s .

  if [ ! -d './cmd/bin' ]; then mkdir cmd/bin; fi
  go build -o cmd/bin/main cmd/main.go
  echo 'server built finished'
}

buildServer
echo 'all build done'
