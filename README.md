# fabtest

备注: 正确使用步骤
1. 根据镜像创建所需个数和配置的测试机器(用户名: root 密码: dev@peersafe)
   控制机需要外网ip，其他机器不需要
2. 选一台为控制机器，（同时为jmeter+haproxy)
3. 进入控制机cd $GOPATH/src/github.com/fabtestorg/
4. vi data/node.json   //编辑不同测试所需参数
5. 执行脚本 ./1star.sh   //生成配置文件，证书目录等， 并启动zk，kfk，order，peer 节点
6. 执行脚本 ./2start.sh  //加入channl， 更新锚节点，实例化cc， 部署haproxy， api，event
7. 执行脚本 ./3start.sh   //启动nmon 服务，  运行jmeter
8. 看eventlog 是否全部落账
   1) 控制机执行脚本 ./copyisa.sh  //将控制机器的ssh公钥部署到所有机器(执行一次即可）
   2) 控制机执行脚本 ./loginpc.sh api10 //进入组织 1， event 0 的机器
   3) api主机执行命令 watch -n 1 wc -l fabtest/event_server/eventserver.log   //看是否全部落账
9. 控制机执行脚本 ./analysis.sh $dirname   //dirname: 自己要保存的统计结果的目录
10. 测试完成后删除所有节点： ./4stop.sh

以上为测试工具使用方式， 具体细节可以根据需求自行调整脚本里的命令


命令介绍:
#基于 python fabric 1.13.1 版本
# pip install fabric==1.13.1    //root 执行

#chmod +x ./bin/*
#sudo chown ubuntu:ubuntu /etc/hosts
# 生产 yaml
#./fabtest -f crypto-config
# 生产证书文件
#./fabtest -c crypto

1. 生成yaml文件  如果改变组织数(修改 config.json)
#configtx.yaml
./fabtest -f configtx

启动 zabbix-agent
#./fabtest -f zabbix
#./fabtest -s zabbix

#生成配置文件 peer,order,zk,kfk yaml
./fabtest -f node

#生产client 用配置文件api , event client_sdk.yaml
./fabtest -f client

2. 生成 创世区块
./fabtest -c genesisblock

3. 向远程copy crypto channel-artifacts kafkaTLS配置文件
./fabtest -p all
 
4. 启动节点zk, kfk, order, peer
./fabtest -s zookeeper
./fabtest -s kafka

./fabtest -s order
./fabtest -s peer

5. 创建channel
./fabtest -c channel -n testchannel

6. 更新锚节点
./fabtest -r updateanchor -n testchannel

7. peer 加入channel
./fabtest -r joinchannel -n testchannel

8. 安装chaincode
./fabtest -r installchaincode -ccoutpath $PWD/config/testfabric.out

9. 实例化chaincode
./fabtest -r runchaincode -ccname testfabric -n testchannel

10. 启动api, event
./fabtest -s api
./fabtest -s event    //替换event重启也用这个命令

11. 获取eventlog 和jmeter.jtl 50线程 循环50次
./fabtest -g event -gn 50_1000

12. 运行haproxy
./fabtest -f haproxy
./fabtest -s haproxy

12. 运行jmeter
./fabtest -s jmeter
13. 停止jmeter
./fabtest -d jmeter

13. 分析获取的eventlog  50线程 循环50次
./fabtest -a event -gn 50_1000

14. 运行nmon
./fabtest -s nmon

15.  获取nmon日志文件
./fabtest -a nmon -gn 50_1000

14. 删除节点
./fabtest -d all/peer/order/api/jmeter

15. 停止或启动节点
./fabtest -op start/stop

16. 生成 zabbix agent配置文件 Ip.conf
./fabtest -f zabbix

17. 启动zabbix agent
./fabtest -s zabbix

18. 手动启动 zabbix   44主机
systemctl status zabbix-agent
systemctl start zabbix-agent

systemctl status zabbix-server
systemctl start zabbix-server

19. 删除节点
./fabtest -d all/peer/order/api

20. 替换镜像
./fabtest -t order/peer

21. 重启docker服务
./fabtest -s docker
