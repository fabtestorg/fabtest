# 北京众享比特科技有限公司

## 区块链层部署文档

## 检查部署所需资源文件

chaincode目录： 包含智能合约源码

node.json 文件：节点服务器配置文件

## 根据需求修改配置文件node.json

下面为模板文件： 采用fabric1.4.1版本， etcdraft 共识类型，3个orderer，2个peer 节点

*具体参数需根据实际数据进行调整*，主要关注“xxx” 这些参数修改，ip写内网ip

```json
{
  "fabricVersion":"1.4.1","domain":"example.com",
  "ccInit":"'{\"Args\":[\"init\"\\,\"xxx\"]}'",
  "ccPolicy":"\"OR  ('Org1MSP.member'\\,'Org2MSP.member')\"",
  "ccName":"name","ccVersion":"1.0",
  "ccPath":"github.com/chaincode",
  "chan_counts":1,
  "consensusType":"raft", "imageTag":"1.4.1","log":"INFO",
  "batchTime":"2s", "batchSize":100, "batchPreferred":"512 KB", "useCouchdb":"false",
  "orderers":[
   {"sshUserName":"xxx","sshPwd":"xxx","ip":"xxx","id":"0","orgId":"1","ports":["7050:7050"]},
   {"sshUserName":"xxx","sshPwd":"xxx","ip":"xxx","id":"1","orgId":"1","ports":["7050:7050"]},
  {"sshUserName":"xxx","sshPwd":"xxx","ip":"xxx","id":"2","orgId":"1","ports":["7050:7050"]}
  ],
  "peers": [
   {"sshUserName":"xxx","sshPwd":"xxx","ip":"xxx","id":"0","orgId":"1","ports":["7051:7051"]},
   {"sshUserName":"xxx","sshPwd":"xxx","ip":"xxx","id":"1","orgId":"1","ports":["7051:7051"]}
  ]
}

```

## 部署

### 创建部署目录

```bash
mkdir -p ~/deploy && cd  ~/deploy
```

### 复制资源文件 

复制（chaincode目录，修改正确数据的 node.json 配置文件) 到deploy目录下

### 启动部署工具

*这个过程需要拉取部署工具镜像文件,第一次执行需要等待几分钟，以下命令为一行*

```bash
docker run -it -d --name manager -v $PWD/config:/opt/fabtest/config -v $PWD/node.json:/opt/fabtest/data/node.json -v $PWD/chaincode:/opt/gopath/src/github.com/peersafe/aiwan/fabric/chaincode peersafes/deploy-cli:latest
```

如果报错，需要根据具体错误调整

### -下面一定要按顺序执行-

### 检测配置文件登录参数是否配置正确

```bash
docker exec manager bash -c ./0-checknode.sh
```

如果连接失败或卡住，就要检测node.json里面的ip和ssh配置，或者用ssh命令先测试一下

### 生成fabric需要配置文件

```bash
docker exec manager bash -c ./1-makeConfig.sh
```

### 启动fabric节点

```bash
docker exec manager bash -c ./2-startNode.sh
```

### 启动fabric智能合约

```bash
docker exec manager bash -c ./3-runChaincode.sh
```

### 检查所有节点是否启动成功

```bash
docker exec manager bash -c ./0-checknode.sh
```



如果上面部署命令执行有错误，先根据错误日志判定是否node.json文件参数错误，然后要清理环境再重新部署。



## 后台服务客户端所需证书目录

```bash
~/depoly/config/crypto-config
```

## 重新部署-清理环境

```bash
docker exec manager bash -c ./deletenode.sh
```

然后再重头开始执行

## 业务变动升级fabric智能合约

修改node.json文件里面的ccVersion为新版本号：  "ccVersion":"1.1"

```bash
docker exec manager bash -c ./upgradecc.sh
```

### 

## 特别说明：

以上部署需要依赖外网环境

如果想要部署在内网环境机器，可以先用外网机器拉取所需镜像，在将镜像导入到内网服务器

```bash
docker pull peersafes/deploy-cli:latest   		#部署工具镜像
docker pull peersafes/fabric-zookeeper:1.4.1	#zookeeper镜像
docker pull peersafes/fabric-kafka:1.4.1		#kafka镜像
docker pull peersafes/fabric-orderer:1.4.1		#orderer镜像
docker pull peersafes/fabric-peer:1.4.1			#peer镜像
docker pull peersafes/fabric-ccenv:1.4.1		#编译智能合约依赖镜像
docker pull peersafes/fabric-baseos:1.4.1		#智能合约运行镜像
docker pull peersafes/fabric-ca:1.4.1			#ca镜像
```



