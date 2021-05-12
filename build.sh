#!/usr/bin/env bash

echo -n "select platform (linux/mac)? "

read platform

if [ ${platform} == "linux" ] ;then
    echo "build for Linux"
    export CGO_ENABLED=0 GOOS=linux GOARCH=amd64
else
    echo "build for Darwin Mac"
    platform="mac"
    export CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
fi

rm -rf bin/kcc-toolkit-${platform}

go build -o bin/kcc-toolkit-${platform} -ldflags "-s -w"

echo -n "need compress (yes/no)? "

read compress

if [ ${compress} == "yes" ] ;then
    upx -f -9 bin/kcc-toolkit-${platform}
fi
