#!/bin/bash

cd tracker/
yarn build
cd ../

cd dashboard/
yarn build
cd ../

if [ $# -gt 0 ]; then
    if [ $1 = '--no-go' ]; then
        exit 0
    fi
fi

go build . -o build/
