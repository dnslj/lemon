#!/bin/sh

SERVICE_NAME="frontend"
SERVICE_HOME=$(echo $PWD)
SERVICE_LOG_PATH='/tmp/'
PROCESS_PATH="/cmd/bin/"
SERVICE_PATH="${SERVICE_HOME}${PROCESS_PATH}${SERVICE_NAME}"
DATE_TIME=$(date "+%Y-%m-%d")

# set msg
info_msg="\033[;34m[INFO]\033[0m"
warn_msg="\033[;33m[WARN]\033[0m"
error_msg="\033[;31m[ERROR]\033[0m"
success_msg="\033[;42m[SUCCESS]\033[0m"

buildServer() {
  echo "${info_msg} server building..."
  export GOPROXY=https://goproxy.cn

  if [ ! -d logs ]; then mkdir logs; fi

  gofmt -w -s .

  BIN_FOLDER=${SERVICE_HOME}${PROCESS_PATH}
  if [ ! -d $BIN_FOLDER ]; then mkdir -p $BIN_FOLDER; fi
  go build -o $SERVICE_PATH ${SERVICE_HOME}/cmd/main.go
  chmod o+x $SERVICE_PATH
  echo "${success_msg} server built finished"
}

function usage() {
  echo "${error_msg} sh $0 start|stop|restart"
  echo "    start: start $SERVICE_NAME only if $SERVICE_NAME not running"
  echo "    stop: stop $SERVICE_NAME only if $SERVICE_NAME is running"
  echo "    restart: restart $SERVICE_NAME only if $SERVICE_NAME is running, will not be restarted"
}

function stop_service() {
  count=$(ps -ef | grep $SERVICE_NAME | wc -l)
  if [ $count -gt 1 ]; then
    killall -15 $SERVICE_NAME
  fi
}

function op_start() {
  # 开启服务
  count=$(ps -ef | grep $SERVICE_NAME | wc -l)
  if [ $count -gt 2 ]; then
    echo "${info_msg} service: $SERVICE_NAME is running"
    exit 0
  fi
  $SERVICE_PATH >>${SERVICE_LOG_PATH}${SERVICE_NAME}-${DATE_TIME}.log 2>&1 &
  echo "\n${info_msg}service: $SERVICE_NAME is running, process details:"
  ps -ef | grep $SERVICE_NAME | grep -v grep
}

function op_stop() {
  stop_service
  echo "${info_msg} service: $SERVICE_NAME is stopped"
  exit 0
}

function op_restart() {
  stop_service
  op_start
  exit 0
}

op=$1
if [ "$op" == "" ]; then
  usage
  exit 0
fi

case $op in
build)
  buildServer
  ;;
start)
  op_start
  ;;
stop)
  op_stop
  ;;
reload)
  op_restart
  ;;
*)
  usage
  ;;
esac
