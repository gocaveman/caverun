#!/bin/bash

set -e

cd `dirname $0`
export BASEDIR=`pwd`

$BASEDIR/server-build.sh

cd caverun-app
npm run release
