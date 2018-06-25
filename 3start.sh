set -x
ssh ubuntu@kafka1 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@kafka2 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@kafka3 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@peer0 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@peer1 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@orderer0 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@orderer1 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@api0 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
ssh ubuntu@api1 "mkdir -p ~/nmon_log; cd ~/nmon_log;nmon -f -s 2 -c 600"
sleep 3

./fabtest -s jmeter
