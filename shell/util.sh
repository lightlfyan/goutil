#!/bin/bash

#Absolute
SCRIPT=$(readlink -f $0)
SCRIPTPATH=`dirname $SCRIPT`

SCRIPT=$(realpath -s $0)
SCRIPTPATH=`dirname $SCRIPT`

cd $(dirname $0)
SCRIPTPATH=$(pwd)
