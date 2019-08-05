#! /bin/bash

# kill any currently running version of the app
killall secret-share-li
killall secret-share-da

# build the static assets so they get bundled into the go binary
go-bindata -fs -o httpd/bindata.go views/ assets/ assets/images/

# cd into the httpd folder to build and run the new binary
cd httpd
env GOOS=linux GOARCH=amd64 go build -o ../builds/secret-share-linux
env GOOS=darwin GOARCH=amd64 go build -o ../builds/secret-share-darwin
env GOOS=windows GOARCH=amd64 go build -o ../builds/secret-share.exe

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    ../builds/secret-share-linux &
elif [[ "$OSTYPE" == "darwin"* ]]; then
    ../builds/secret-share-darwin &
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
