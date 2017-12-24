#!/bin/bash

set -x
# Find the directory we exist within
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd ${DIR}

VERSION=`git describe --always --long`

mkdir build
cp ../build/worldping-api build/
cp -a ../public build/
cp -a ../conf build/

docker build -t rednixon/worldping-api:$VERSION .
docker tag rednixon/worldping-api:$VERSION rednixon/worldping-api:latest
