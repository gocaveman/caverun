#!/bin/bash

# NOTE: GOPATH must already be set properly in your environment or this will fail
# WINDOWS NOTE: For Windows users make sure your GOPATH is set to a proper Windows
# path not a Cygwin path (i.e. "C:\path\to\project" not "/c/path/to/project"), 
# even though you will probably be running this script in a Cygwin shell.

set -e

go install ./caverun-srv

cp $GOPATH/bin/caverun-srv* caverun-app/app/
