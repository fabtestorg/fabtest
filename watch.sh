#!/bin/bash
#set -x
clear
while true
do
    for node in `cat /etc/hosts | grep  -e api |grep -v ^$|grep -v ".cn"|grep -v ".com"|awk '{print $2}'`
    do
        echo $node
        ssh root@$node -o ConnectTimeout=1 "docker stats --no-stream"
	for((i=0;i<8;i++));
	do	
	   for((j=1;j<=4;j++));
	   do
     		ssh root@$node -o ConnectTimeout=1 "if test -f ~/fabtest/event_server/peer${i}org1api${j}/eventserver.log;then cat ~/fabtest/event_server/peer${i}org1api${j}/eventserver.log|wc -l;fi"
	   done
	done
        echo 
    done
done
