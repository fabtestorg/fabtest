#!/usr/bin/env bash

#chmod +x ./bin/*
#sudo chown ubuntu:ubuntu /etc/hosts
# 第一次要执行下 重启docker 服务
#./fabtest -s docker
# 重新生成证书
./fabtest -f crypto-config
./fabtest -c crypto
#1. 生成yaml文件  如果改变组织数(修改 config.json)
./fabtest -f configtx

#peer,order,zk,kfk yaml
./fabtest -f node

#sed -i  's/7050:7050/8050:7050/g' config/order1ord2.yaml

#sed -i  's/ord_address0/192.168.0.21/g' config/order0ord1.yaml
#sed -i  's/ord_address0/192.168.0.21/g' config/order1ord2.yaml

sed -i  's/ord_address1/192.168.0.21/g' config/order0ord1.yaml
#sed -i  's/ord_address1/192.168.0.21/g' config/order1ord2.yaml

#api , event client_sdk.yaml
# ./fabtest -f client

#2. 生成 创世区块
./fabtest -c genesisblock

#3. 向远程copy crypto channel-artifacts kafkaTLS配置文件
./fabtest -p all

# ./1star.sh
