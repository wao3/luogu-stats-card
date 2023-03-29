#!/bin/bash

# 编译前端代码
# shellcheck disable=SC2164
cd ./web
npm install
npm run build

# 返回上级目录，编译golang代码
# shellcheck disable=SC2103
cd ..
GOOS=windows GOARCH=amd64 go build -o ./dist/luogu_stats_card_windows_amd64.exe
GOOS=linux GOARCH=amd64 go build -o ./dist/luogu_stats_card_linux_amd64
GOOS=darwin GOARCH=amd64 go build -o ./dist/luogu_stats_card_darwin_amd64

docker build -t luogu .