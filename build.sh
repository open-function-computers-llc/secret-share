#! /bin/bash

# kill any currently running version of the app

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    killall secret-share-li
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    killall secret-share-da
fi

# build and run the new binary
env GOOS=linux GOARCH=amd64 go build -o builds/secret-share-linux
env GOOS=darwin GOARCH=amd64 go build -o builds/secret-share-darwin
env GOOS=windows GOARCH=amd64 go build -o builds/secret-share.exe

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    builds/secret-share-linux &
elif [[ "$OSTYPE" == "darwin"* ]]; then
    builds/secret-share-darwin &
elif [[ "$OSTYPE" == "cygwin" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "msys" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "win32" ]]; then
    echo "windows?"
elif [[ "$OSTYPE" == "freebsd"* ]]; then
    echo "free-bsd?"
else
    echo "i have no idea what's going on"
fi
