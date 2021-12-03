#!/bin/bash
# protoc --proto_path=`pwd`/window.proto \
#     --go_out=`pwd`/ \
#     `pwd`/window.proto
protoc -I ./ *.proto --go_out=.