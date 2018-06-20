# fabtest

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

#peer,order,zk,kfk yaml
./fabtest -f node

#api , event client_sdk.yaml
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

12. 运行jmeter

./fabtest -f haproxy
./fabtest -s haproxy

./fabtest -s jmeter

13. 分析获取的eventlog  50线程 循环50次
./fabtest -a event -gn 50_1000

14. 删除节点
./fabtest -d all/peer/order/api

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

