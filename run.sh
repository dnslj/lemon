#!/bin/bash

SERVICE_NAME="lemon-web"
SERVICE_HOME=$(echo $PWD)
SERVICE_LOG_PATH='/tmp/'
PROCESS_PATH="/cmd/bin/"
LEMON_WEB_PATH="${SERVICE_HOME}${PROCESS_PATH}${SERVICE_NAME}"
DATE_TIME=`date "+%Y-%m-%d"`

function usage()
{
    echo "sh $0 start|stop|restart"
    echo "    start: start $SERVICE_NAME only if $SERVICE_NAME not running"
    echo "    stop: stop $SERVICE_NAME only if $SERVICE_NAME is running"
    echo "    restart: restart $SERVICE_NAME only if $SERVICE_NAME is running, will not be restarted"
}

function stop_service()
{
    count=`ps -ef | grep $SERVICE_NAME | wc -l`
    if [ $count -gt 1 ]; then
        killall -15 $SERVICE_NAME
    fi
}

function op_start()
{
    count=`ps -ef | grep $SERVICE_NAME | wc -l`
    if [ $count -gt 2 ]; then
        echo "service: $SERVICE_NAME is running"
        exit 0
    fi
    $LEMON_WEB_PATH >> ${SERVICE_LOG_PATH}${SERVICE_NAME}-${DATE_TIME}.log 2>&1 &
    echo "service: $SERVICE_NAME is running, process details:"
    ps -ef | grep $SERVICE_NAME | grep -v grep
}

function op_stop()
{
    stop_service
    echo "service: $SERVICE_NAME is stopped"
    exit 0
}

function op_restart()
{
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
    start)
        op_start
        ;;
    stop)
        op_stop
        ;;
    restart)
        op_restart
        ;;
    *)
        usage
        ;;
esac