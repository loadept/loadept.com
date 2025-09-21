#!/usr/bin/env bash

export CGO_ENABLED=1

APP_NAME=app
DIR_NAME=loadept-web

pushd ./web > /dev/null
if ! pnpm i --frozen-lockfile; then
    echo "Install dependencies failed"
    exit 1
fi

export API_URL=/
if ! pnpm run build; then
    echo "Build static failed"
    exit 1
fi
popd > /dev/null

if ! go build -o $APP_NAME cmd/loadept/main.go; then
    echo "Build binary failed"
    exit 1
fi

mkdir -p $DIR_NAME/logs
mkdir $DIR_NAME/web

mv $APP_NAME $DIR_NAME
mv web/dist $DIR_NAME/web
cp ecosystem.config.js $DIR_NAME

tar -czvf $DIR_NAME.tar.gz $DIR_NAME

rm -rf $DIR_NAME
