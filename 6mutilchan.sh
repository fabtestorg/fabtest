#!/usr/bin/env bash
set -x
if [ "$1" == "" ]; then
   echo args 1  miss
   exit
fi
./createchannel.sh $1
./runcc.sh $1
