#!/bin/bash

  protoc \
    --proto_path=protocol_buffers/definitions \
    --js_out=import_style=commonjs,binary:protocol_buffers/messages \
    protocol_buffers/definitions/*.proto