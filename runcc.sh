#!/usr/bin/env bash
set -x
#7. peer 加入channel
for((i=1;i<=$1;i++));
do   
	./fabtest -r joinchannel -n mychannel$i
	./fabtest -r installchaincode -ccoutpath $PWD/config/ccout/factor$i.out
	./fabtest -r runchaincode -ccname factor$i -n mychannel$i
done  
