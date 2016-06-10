#!/bin/bash -ex

make clean

go-bindata app/...

for os in "windows" "linux" "darwin"; do
    ext=""
    if [ ${os} = "windows" ]; then
        ext=".exe"
    fi

    for arch in "386" "amd64"; do
        GOOS=${os} GOARCH=${arch} go build -a -tags netgo --ldflags "-s -w" -o build/howdy_${os}_${arch}${ext}
    done
done
