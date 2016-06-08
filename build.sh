#!/bin/bash -ex

rm -fr build

for os in "windows" "linux" "darwin"; do
    ext=""
    if [ ${os} = "windows" ]; then
        ext=".exe"
    fi

    for arch in "386" "amd64"; do
        GOOS=${os} GOARCH=${arch} go build -a -tags netgo --ldflags '-s -w' -o build/howdy_${os}_${arch}${ext}
        nrsc build/howdy_${os}_${arch}${ext} app
    done
done
