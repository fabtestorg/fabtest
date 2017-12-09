#!/usr/bin/env bash

#1. 生成yaml文件  如果改变组织数(修改 config.json)
./fabtest -f configtx

#启动 zabbix-agent
./fabtest -f zabbix
./fabtest -s zabbix

#peer,order,zk,kfk yaml
./fabtest -f node

#api , event client_sdk.yaml
./fabtest -f client

#2. 生成 创世区块
./fabtest -c genesisblock

#3. 向远程copy crypto channel-artifacts kafkaTLS配置文件
./fabtest -p all

#4. 启动节点zk, kfk, order, peer
./fabtest -s zookeeper
./fabtest -s kafka
./fabtest -s order
./fabtest -s peer

#44 机器添加 /etc/hosts    对 peer order

#5. 创建channel
./fabtest -c channel -n testchannel

#6. 更新锚节点
./fabtest -r updateanchor -n testchannel

#7. peer 加入channel
./fabtest -r joinchannel -n testchannel

#8. 安装chaincode
./fabtest -r installchaincode -ccoutpath $PWD/config/factor.out

#9. 实例化chaincode
./fabtest -r runchaincode -ccname factor -n testchannel

#10. 启动api, event
./fabtest -s api
./fabtest -s event
./fabtest -g event -gn templog