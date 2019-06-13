#!/usr/bin/env bash

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! FAIL !!!!!!!!!!!!!!!!"
    exit 1
  fi
}
echo "测试chaincode"
./fabtest -r testcc -n mychannel
verifyResult $?

#echo "升级chaincode"
#./fabtest -r upgradecc -n mychannel
#verifyResult $?
