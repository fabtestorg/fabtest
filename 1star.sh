#!/usr/bin/env bash

#4. 启动节点zk, kfk, order, peer
./fabtest -s zookeeper
./fabtest -s kafka
sleep 10
./fabtest -s order
./fabtest -s peer
