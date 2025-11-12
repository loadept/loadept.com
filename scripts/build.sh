#!/usr/bin/env bash

export CGO_ENABLED=1

APP_NAME=app
DIR_NAME=loadept-core

if ! go build -o $APP_NAME cmd/loadept/main.go; then
    echo "Build binary failed"
    exit 1
fi

mkdir -p $DIR_NAME/logs

mv $APP_NAME $DIR_NAME

tar -czvf $DIR_NAME.tar.gz $DIR_NAME

rm -rf $DIR_NAME
