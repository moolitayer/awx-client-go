#!/bin/bash
ret=0
bad_files=$(gofmt -s -l awx/ examples/)
if [[ -n "${bad_files}" ]]; then
    echo "gofmt needs to be run on the listed files"
    echo "${bad_files}"
    echo "Try running 'gofmt -w -d [path]'"
    ret=1
fi
bad_files=$(goimports -l awx/ examples/)
if [[ -n "${bad_files}" ]]; then
    echo "goimports needs to be run on the listed files"
    echo "${bad_files}"
    echo "Try running 'goimports -w -d [path]'"
    ret=1
fi

# Vetting the entirety of the awx package in one go
echo "Running go vet: "
go vet ./awx
if [[ "$?" -ne "0" ]]; then
    ret=1
fi

# Examples are self contained so they need to be vetted one-by-one
# otherwise you get errors for variables being re-declared
for f in examples/*.go; do
    go vet "$f"
    if [[ "$?" -ne "0" ]]; then
        ret=1
    fi
done

exit ${ret}
