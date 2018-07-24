#!/usr/bin/env bash

set -x
#./fabtest -s nmon
#sleep 3

./fabtest -d jmeter
./fabtest -s jmeter
