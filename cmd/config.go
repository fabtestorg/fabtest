package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	TplZookeeper    = "zookeeper.tpl"
	TplKafka        = "kafka.tpl"
	TplOrderer      = "orderer.tpl"
	TplPeer         = "peer.tpl"
	TplCryptoConfig = "crypto-config.tpl"
	TplConfigtx     = "configtx.tpl"
	TplApiClient    = "apiclient.tpl"
	TplApiDocker    = "apidocker.tpl"
	TplEventClient  = "eventclient.tpl"

	TypePeer         = "peer"
	TypeOrder        = "orderer"
	TypeKafka        = "kafka"
	TypeZookeeper    = "zookeeper"
	TypeApi          = "api"
	ZabbixServerIp   = "zabbix_server_ip"
	ZabbixServerPort = "zabbix_server_port"
	ZabbixAgentIp    = "zabbix_agent_ip"
	List             = "list"
	NodeType         = "node_type"
	ZkId             = "zk_id"
	IP               = "ip"
	APIIP            = "apiip"
	JMETER           = "jmeter"
	PeerId           = "peer_id"
	OrgId            = "org_id"
	OrderId          = "order_id"
	KfkId            = "kfk_id"
	ChanCounts       = "chan_counts"
)

var GlobalConfig *ConfigObj

type ConfigObj struct {
	FabricVersion  string         `json:"fabricVersion"`
	TestArgs       string         `json:"testArgs"`
	CCInit         string         `json:"ccInit"`
	CCPolicy       string         `json:"ccPolicy"`
	CCPath         string         `json:"ccPath"`
	CCName         string         `json:"ccName"`
	CCVersion      string         `json:"ccVersion"`
	ConsensusType  string         `json:"consensusType"`
	BatchTime      string         `json:"batchTime"`
	BatchPreferred string         `json:"batchPreferred"`
	BatchSize      int            `json:"batchSize"`
	Zookeepers     []NodeObj      `json:"zookeepers"`
	Kafkas         []NodeObj      `json:"kafkas"`
	OrdList        map[string]int `json:"ordList"`
	OrgList        map[string]int `json:"orgList"`
	Expand
}

type Expand struct {
	SshUserName    string    `json:"sshUserName"`
	SshPwd         string    `json:"sshPwd"`
	Log            string    `json:"log"`
	UseCouchdb     string    `json:"useCouchdb"`
	Domain         string    `json:"domain"`
	Orderers       []NodeObj `json:"orderers"`
	Peers          []NodeObj `json:"peers"`
	ImageTag       string    `json:"imageTag"`
	ImagePre       string    `json:"imagePre"`
	DefaultNetwork string    `json:"defaultNetwork"`
}

type NodeObj struct {
	Ip           string   `json:"ip"`
	ApiIp        string   `json:"apiIp"`
	Id           string   `json:"id"`
	OrgId        string   `json:"orgId"`
	Ports        []string `json:"ports"`
	ConfigTxPort string   `json:"configTxPort"`
	Expand
}

func ConfigDir() string {
	return os.Getenv("PWD") + "/config/"
}

func InputDir() string {
	return os.Getenv("PWD") + "/data/"
}

func TplPath(name string) string {
	return fmt.Sprintf("%s/templates/%s/%s", os.Getenv("PWD"), GlobalConfig.FabricVersion, name)
}

func BinPath() string {
	return fmt.Sprintf("%s/bin/%s/", os.Getenv("PWD"), GlobalConfig.FabricVersion)
}

func ChannelPath() string {
	return os.Getenv("PWD") + "/config/channel-artifacts/"
}

func ImagePath() string {
	return os.Getenv("PWD") + "/images/"
}

func ScriptPath() string {
	return os.Getenv("PWD") + "/scripts/"
}

func ParseJson(jsonfile string) (*ConfigObj, error) {
	var obj ConfigObj
	file := InputDir() + jsonfile
	fmt.Printf("json file %s\n", file)
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		return &obj, err
	}
	obj.OrdList = make(map[string]int)
	obj.OrgList = make(map[string]int)
	err = json.Unmarshal(jsonData, &obj)
	if err != nil {
		return &obj, err
	}
	for i, v := range obj.Peers {
		obj.OrgList[v.OrgId] = obj.OrgList[v.OrgId] + 1
		configTxPort, err := findConfigTxPort(v.Ports, "7051")
		if err != nil {
			return &obj, err
		}
		obj.Peers[i].ConfigTxPort = configTxPort
	}
	for i, v := range obj.Orderers {
		obj.OrdList[v.OrgId] = obj.OrgList[v.OrgId] + 1
		configTxPort, err := findConfigTxPort(v.Ports, "7050")
		if err != nil {
			return &obj, err
		}
		obj.Orderers[i].ConfigTxPort = configTxPort
	}
	for i, v := range obj.Kafkas {
		configTxPort, err := findConfigTxPort(v.Ports, "9092")
		if err != nil {
			return &obj, err
		}
		obj.Kafkas[i].ConfigTxPort = configTxPort
	}
	if obj.ImagePre == "" {
		obj.ImagePre = "peersafes"
	}

	//fmt.Printf("config obj is %#v\n", obj)
	return &obj, nil
}

func GetJsonMap(jsonfile string) map[string]interface{} {
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

func findConfigTxPort(list []string, destPort string) (string, error) {
	for _, v := range list {
		curLine := strings.Split(v, ":")
		if len(curLine) != 2 {
			return "", fmt.Errorf("findConfigTxPort err %s", v)
		}
		if curLine[1] == destPort {
			return curLine[0], nil
		}
	}
	return "", fmt.Errorf("findConfigTxPort err destPort %s not exist ", destPort)
}
