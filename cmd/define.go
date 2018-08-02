package cmd

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

const (
	SshUserName     = "root"
	SshPwd          = "dev@peersafe"
	SudoPwd          = "dev@peersafe"
	TplZookeeper    = "./templates/zookeeper.tpl"
	TplKafka        = "./templates/kafka.tpl"
	TplOrderer      = "./templates/orderer.tpl"
	TplPeer         = "./templates/peer.tpl"
	TplCryptoConfig = "./templates/crypto-config.tpl"
	TplConfigtx     = "./templates/configtx.tpl"
	TplApiClient     = "./templates/apiclient.tpl"
	TplApiDocker     = "./templates/apidocker.tpl"
	TplEventClient     = "./templates/eventclient.tpl"

	TypePeer      = "peer"
	TypeOrder     = "order"
	TypeKafka     = "kafka"
	TypeZookeeper = "zookeeper"
	TypeApi = "api"
	PeerDomain = "peer_domain"
	OrgCounts = "org_counts"
	Nmon_Rate = "nmon_rate"
	Nmon_Times = "nmon_times"
	KfkDomain = "kfk_domain"
	KFK0_ADDRESS = "kfk0_address"
	KFK1_ADDRESS = "kfk1_address"
	KFK2_ADDRESS = "kfk2_address"
	KFK3_ADDRESS = "kfk3_address"
	KfkVersion = "kfk_version"
	ZabbixServerIp   = "zabbix_server_ip"
	ZabbixServerPort = "zabbix_server_port"
	ZabbixAgentIp = "zabbix_agent_ip"
	List = "list"
	NodeType = "node_type"
	ZkId = "zk_id"
	IP = "ip"
	APIIP = "apiip"
	JMETER= "jmeter"
	IP0 = "ip0"
	IP1 = "ip1"
	IP2 = "ip2"
	PeerId = "peer_id"
	OrgId = "org_id"
	OrderId = "order_id"
	KfkId = "kfk_id"
	OrderAddress = "order_address"
	BrokerAddress = "broker_address"
	Zk_IP1 = "zk_ip1"
	Zk_IP2 = "zk_ip2"
	Zk_IP0 = "zk_ip0"
	BrokerIp = "broker_ip"
	ChanCounts = "chan_counts"
	AdverAddress = "adver_address"
	BlockChainAddress = "blockchain_address"

	Order_Address = "order_address"
	Order1_Address = "order1_address"
	Other_PeerAddress = "other_peeraddress"
	USECOUCHDB = "usecouchdb"

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
