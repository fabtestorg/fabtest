#!/usr/bin/env bash
set -x
tmp="aa"
if [ $1 != "" ]; then
    tmp="$1"
fi
echo $tmp

./fabtest -g event -gn $tmp
./fabtest -g jmeter -gn $tmp
./fabtest -a event -gn $tmp > config/event_logs/$tmp/analysis.txt

# kill remote nmon and rename logfile
ssh ubuntu@kafka1 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/kafka1.nmon"
ssh ubuntu@kafka2 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/kafka2.nmon"
ssh ubuntu@kafka3 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/kafka3.nmon"
ssh ubuntu@peer0 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/peer0.nmon"
ssh ubuntu@peer1 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/peer1.nmon"
ssh ubuntu@orderer0 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/orderer0.nmon"
ssh ubuntu@orderer1 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/orderer1.nmon"
ssh ubuntu@api0 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/api0.nmon"
ssh ubuntu@api1 "pgrep nmon| xargs kill -9;mv ~/nmon_log/*.nmon ~/nmon_log/api1.nmon"

# scp nmon log files to local
scp ubuntu@kafka1:~/nmon_log/kafka1.nmon config/event_logs/$tmp
scp ubuntu@kafka2:~/nmon_log/kafka2.nmon config/event_logs/$tmp
scp ubuntu@kafka3:~/nmon_log/kafka3.nmon config/event_logs/$tmp
scp ubuntu@peer0:~/nmon_log/peer0.nmon config/event_logs/$tmp
scp ubuntu@peer1:~/nmon_log/peer1.nmon config/event_logs/$tmp
scp ubuntu@orderer0:~/nmon_log/orderer0.nmon config/event_logs/$tmp
scp ubuntu@orderer1:~/nmon_log/orderer1.nmon config/event_logs/$tmp
scp ubuntu@api0:~/nmon_log/api0.nmon config/event_logs/$tmp
scp ubuntu@api1:~/nmon_log/api1.nmon config/event_logs/$tmp

# rm remote nmon log files
ssh ubuntu@kafka1 "rm ~/nmon_log/*.nmon"
ssh ubuntu@kafka2 "rm ~/nmon_log/*.nmon"
ssh ubuntu@kafka3 "rm ~/nmon_log/*.nmon"
ssh ubuntu@peer0 "rm ~/nmon_log/*.nmon"
ssh ubuntu@peer1 "rm ~/nmon_log/*.nmon"
ssh ubuntu@orderer0 "rm ~/nmon_log/*.nmon"
ssh ubuntu@orderer1 "rm ~/nmon_log/*.nmon"
ssh ubuntu@api0 "rm ~/nmon_log/*.nmon"
ssh ubuntu@api1 "rm ~/nmon_log/*.nmon"

cat config/event_logs/$tmp/analysis.txt
echo "######"
ls config/event_logs/$tmp/

