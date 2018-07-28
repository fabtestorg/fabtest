#!/usr/bin/env bash
#set -x
tmp="aa"
if [ "$1" == "" -o -d config/event_logs/$1 ]; then
    echo error
    exit
fi
tmp=$1

./fabtest -g event -gn $tmp
./fabtest -g jmeter -gn $tmp
./fabtest -a event -gn $tmp > config/event_logs/$tmp/analysis.txt

# kill remote nmon and rename logfile
#./fabtest -g nmon -gn $tmp

cat config/event_logs/$tmp/analysis.txt
cat config/event_logs/$tmp/jmeter_send1.txt | grep Summariser
echo "######"
ls config/event_logs/$tmp/

