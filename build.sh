#!/bin/bash
# courtesy of https://kylewbanks.com/blog/cross-compiling-go-applications-for-multiple-operating-systems-and-architectures

for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o bin/main-$GOOS-$GOARCH main.go 
    done
done
