#!/usr/bin/env bash


#5. 创建channel
./fabtest -c channel -n mychannel

#6. 更新锚节点
./fabtest -r updateanchor -n mychannel

#7. peer 加入channel
./fabtest -r joinchannel -n mychannel

#8. 安装chaincode
./fabtest -r installchaincode -ccoutpath $PWD/config/ccout/mycc.out

# #9. 实例化chaincode
./fabtest -r runchaincode -ccname mycc -n mychannel

#10. 启动api, event
# ./fabtest -s api
# ./fabtest -s event

#11. 获取eventlog  50线程 循环50次
# ./fabtest -g event -gn 50_1000

# ./fabtest -f haproxy
# ./fabtest -s haproxy
#12. 运行jmeter
#  ./fabtest -s jmeter

#13. 分析获取的eventlog  50线程 循环50次
#  ./fabtest -a event -gn 50_1000

#14. 删除节点
#   ./fabtest -d all/peer/order/api

#15. 停止或启动节点
#  ./fabtest -op start/stop
