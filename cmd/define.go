package cmd

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

const (
	Domain_Name     = "finblockchain.cn"
	SshUserName     = "root"
	SshPwd          = "dev1@peersafe"
	TplZookeeper    = "./templates/zookeeper.tpl"
	TplKafka        = "./templates/kafka.tpl"
	TplOrderer      = "./templates/orderer.tpl"
	TplPeer         = "./templates/peer.tpl"
	TplCryptoConfig = "./templates/crypto-config.tpl"
	TplConfigtx     = "./templates/configtx.tpl"
	TplApiClient     = "./templates/apiclient.tpl"
	TplEventClient     = "./templates/eventclient.tpl"

	TypePeer      = "peer"
	TypeOrder     = "order"
	TypeKafka     = "kafka"
	TypeZookeeper = "zookeeper"

	PeerDomain = "peer_domain"
	KfkDomain = "kfk_domain"
	ZabbixServerIp   = "zabbix_server_ip"
	ZabbixServerPort = "zabbix_server_port"
	List = "list"
	NodeType = "node_type"
	ZkId = "zk_id"
	Zk2Id = "zk_2_id"
	IP = "ip"
	APIIP = "apiip"
	IP1 = "ip1"
	IP2 = "ip2"
	IP3 = "ip3"
	IP4 = "ip4"
	IP5 = "ip5"
	PeerId = "peer_id"
	OrgId = "org_id"
	OrderId = "order_id"
	KfkId = "kfk_id"
	OrderAddress = "order_address"
	BrokerAddress = "broker_address"
	Zk_IP1 = "zk_ip1"
	Zk_IP2 = "zk_ip2"
	Zk_IP3 = "zk_ip3"
	Zk_IP4 = "zk_ip4"
	Zk_IP5 = "zk_ip5"
	BrokerIp = "broker_ip"
	AdverAddress = "adver_address"
	BlockChainAddress = "blockchain_address"
)

func ConfigDir() string{
	return os.Getenv("PWD") +  "/config/"
}

func InputDir() string{
	return os.Getenv("PWD") +  "/data/"
}

func BinPath() string{
	return os.Getenv("PWD") +  "/bin/"
}

func ChannelPath() string{
	return os.Getenv("PWD") +  "/config/channel-artifacts/"
}

func ImagePath() string{
	return os.Getenv("PWD") +  "/images/"
}

func ScriptPath() string{
	return os.Getenv("PWD") +  "/scripts/"
}

func GetJsonMap(jsonfile string) map[string]interface{}{
	var inputData map[string]interface{}
	var jsonData []byte
	var err error

	inputfile := InputDir() + jsonfile
	jsonData, err = ioutil.ReadFile(inputfile)
	if err != nil {
		return inputData
	}

	err = json.Unmarshal(jsonData, &inputData)
	if err != nil {
		return inputData
	}
	return inputData

}