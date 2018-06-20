#!/usr/bin/env bash
set -x
./fabtest -g event -gn $1
./fabtest -a event -gn $1 > config/event_logs/$1/analysis.txt
cat config/event_logs/$1/analysis.txt
