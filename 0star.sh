#!/usr/bin/env bash

#chmod +x ./bin/*
#sudo chown ubuntu:ubuntu /etc/hosts
# 重新生成证书
./fabtest -f crypto-config
./fabtest -c crypto-config
#1. 生成yaml文件  如果改变组织数(修改 config.json)
./fabtest -f configtx

#peer,order,zk,kfk yaml
./fabtest -f node

#2. 生成 创世区块
./fabtest -c genesisblock

#3. 向远程copy crypto channel-artifacts kafkaTLS配置文件
./fabtest -p all
