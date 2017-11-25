#!/bin/bash

set -e

VER="$1"
MSG="$2"

if [ -z "$VER" -o -z "$MSG" ]; then
	echo "You must provide a version number and a message, e.g. ./update-version.sh 2.0.0 'whatever update this is'"
	exit
fi

BRANCH=`git rev-parse --abbrev-ref HEAD`

if [ "master" != "$BRANCH" ]; then
	echo "You must be on the master branch"
	exit
fi

git update-index -q --refresh
if [ ! -z "$(git diff-index --name-only HEAD --)" ]; then
	echo "You you cannot have uncommitted files, please commit everything first"
	exit
fi

sed -i -E 's/"version": "[0-9.]+",/"version": "'"$VER"'",/g' caverun-app/package.json
sed -i -E 's/"version": "[0-9.]+",/"version": "'"$VER"'",/g' caverun-app/app/package.json
sed -i -E 's/VERSION = "[0-9.]+"/VERSION = "'"$VER"'"/g' caverun-srv/version.go

./server-build.sh

git commit -a -m "VERSION BUMP: $MSG"

git tag -a "v$VER" -m "$MSG"

git push --tags
