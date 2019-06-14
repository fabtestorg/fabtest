#!/usr/bin/env bash

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! FAIL !!!!!!!!!!!!!!!!"
    exit 1
  fi
}
echo "安装智能合约"
./fabtest -r installchaincode
verifyResult $?
echo "升级chaincode"
./fabtest -r upgradecc -n mychannel
verifyResult $?
