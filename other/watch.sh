#!/bin/bash
clear
while true
do
    for node in `cat /etc/hosts | grep  -e api |grep -v ^$|grep -v ".cn"|grep -v ".com"|awk '{print $2}'`
    do
        echo $node
        ssh root@$node -o ConnectTimeout=1 "docker stats --no-stream"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer0org1api1/eventserver.log;then cat ~/fabtest/event_server/peer0org1api1/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer0org1api2/eventserver.log;then cat ~/fabtest/event_server/peer0org1api2/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer0org1api3/eventserver.log;then cat ~/fabtest/event_server/peer0org1api3/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer0org1api4/eventserver.log;then cat ~/fabtest/event_server/peer0org1api4/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer0org1api5/eventserver.log;then cat ~/fabtest/event_server/peer0org1api5/eventserver.log|wc -l;fi"

        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer1org1api1/eventserver.log;then cat ~/fabtest/event_server/peer1org1api1/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer1org1api2/eventserver.log;then cat ~/fabtest/event_server/peer1org1api2/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer1org1api3/eventserver.log;then cat ~/fabtest/event_server/peer1org1api3/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer1org1api4/eventserver.log;then cat ~/fabtest/event_server/peer1org1api4/eventserver.log|wc -l;fi"
        ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer1org1api5/eventserver.log;then cat ~/fabtest/event_server/peer1org1api5/eventserver.log|wc -l;fi"
        echo 
    done
done
