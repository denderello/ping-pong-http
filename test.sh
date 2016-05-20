#!/bin/sh

set -e

packages=$(go list ./... | grep -v /vendor/)
echo $packages | xargs go vet
echo $packages | xargs go test -v
