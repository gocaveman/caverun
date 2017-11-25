#!/bin/bash

set -e

cd `dirname $0`
export BASEDIR=`pwd`

$BASEDIR/server-build.sh

cd caverun-app
if [ "-debug" = "$1" ]; then
	npm run debug
else
	npm start
fi
