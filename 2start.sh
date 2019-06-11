#!/usr/bin/env bash

#5. 创建channel
./fabtest -c channel -n mychannel

#6. 更新锚节点
./fabtest -r updateanchor -n mychannel

#7. peer 加入channel
./fabtest -r joinchannel -n mychannel

#8. 安装chaincode
./fabtest -r installchaincode

# #9. 实例化chaincode
./fabtest -r runchaincode -n mychannel
