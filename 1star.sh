#!/usr/bin/env bash

#4. 启动节点zk, kfk, order, peer
./fabtest -s zookeeper
./fabtest -s kafka
sleep 20
./fabtest -s order
sleep 15
./fabtest -c channel -n mychannel1
./fabtest -c channel -n mychannel2
./fabtest -c channel -n mychannel3
./fabtest -c channel -n mychannel4
