#!/bin/bash -ex

rm -fr build

for os in "windows" "linux" "darwin"; do
    ext=""
    if [ ${os} = "windows" ]; then
        ext=".exe"
    fi

    mkdir -p build/${os}
    for arch in "386" "amd64"; do
        mkdir build/${os}/${arch}
        GOOS=${os} GOARCH=${arch} go build -a -tags netgo --ldflags '-s -w' -o build/${os}/${arch}/howdy${ext}
        nrsc build/${os}/${arch}/howdy${ext} app
    done
done
