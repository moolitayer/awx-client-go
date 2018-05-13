#!/bin/bash
if [ ! -d "_out" ]; then
    mkdir _out
fi
pushd _out || exit 1
for i in ../examples/*.go; do go build "${i}"; done
popd