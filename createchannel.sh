#!/usr/bin/env bash
set -x

#5. 创建channel
for((i=1;i<=$1;i++));
do   
	./fabtest -c channel -n mychannel$i
done 
