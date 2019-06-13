#!/usr/bin/env bash

#chmod +x ./bin/*
#sudo chown ubuntu:ubuntu /etc/hosts
verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! FAIL !!!!!!!!!!!!!!!!"
    exit 1
  fi
}
echo "生成证书配置文件"
./fabtest -f crypto-config
verifyResult $?
echo "生成证书目录"
./fabtest -c crypto-config
verifyResult $?

echo "生成configtx配置文件"
./fabtest -f configtx
verifyResult $?

echo "生成节点docker启动文件"
./fabtest -f node
verifyResult $?

echo "生成创世区块"
./fabtest -c genesisblock
verifyResult $?

