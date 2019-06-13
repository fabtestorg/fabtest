#!/usr/bin/env bash

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! FAIL !!!!!!!!!!!!!!!!"
    exit 1
  fi
}
echo "向节点传输证书、配置文件"
./fabtest -p all
verifyResult $?

echo "启动节点"
./fabtest -s orderer
verifyResult $?

./fabtest -s peer
verifyResult $?

