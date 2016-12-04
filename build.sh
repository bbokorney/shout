#!/bin/bash

set -ex

go test -v ./...

hash=$(git rev-parse HEAD)
version=$(git show-ref --tags -d | grep $hash | sed -e 's,.* refs/tags/,,' -e 's/\^{}//')
if [[ -z $version ]]; then
  echo "No version for current commit, defaulting to 'pre-release'"
  version='pre-release'
fi
hash=$(echo $hash | cut -c 1-10)
echo "Building version $version (commit $hash)"

arch='amd64'
for os in $(echo linux windows darwin); do
  GOOS=$os GOARCH=$arch go build -o build/shout-$os-$arch -ldflags "-X main.version=$version -X main.hash=$hash"
done
