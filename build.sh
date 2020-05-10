#!/bin/bash

SERVICE_NAME="lemon-web"
SERVICE_HOME=$(echo $PWD)
PROCESS_PATH="/cmd/bin/"
LEMON_WEB_PATH="${SERVICE_HOME}${PROCESS_PATH}${SERVICE_NAME}"

buildServer() {
  echo 'server building...'
  export GOPROXY=https://goproxy.cn

  if [ ! -d logs ]; then mkdir logs; fi

  gofmt -w -s .

  BIN_FOLDER=${SERVICE_HOME}${PROCESS_PATH}
  if [ ! -d $BIN_FOLDER ]; then mkdir -p $BIN_FOLDER; fi
  go build -o $LEMON_WEB_PATH ${SERVICE_HOME}/cmd/main.go
  echo 'server built finished'
}

buildServer
echo 'all build done'
