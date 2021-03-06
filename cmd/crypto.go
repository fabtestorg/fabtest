package cmd

import (
	"fmt"
	"github.com/peersafe/fabtest/tpl"
	"time"
	"strconv"
)

func CreateCert() error {
	obj := NewFabCmd("apply_cert.py", "")
	err := obj.RunShow("generate_certs", BinPath(), ConfigDir(), ConfigDir())
	if err != nil {
		return err
	}
	return nil
}

func CreateYamlByJson(strType string) error {
	if strType == "configtx" {
		inputData := GetJsonMap("node.json")
		orgcounts := inputData[OrgCounts].(float64)
		var orgslist,kafkalist []string
		for i := 1; i <= int(orgcounts); i++ {
			orgslist = append(orgslist,strconv.Itoa(i))
		}
		kafkalist = append(kafkalist,findMapValue(TypeKafka,"1",""))
		kafkalist = append(kafkalist,findMapValue(TypeKafka,"2",""))
		kafkalist = append(kafkalist,findMapValue(TypeKafka,"3",""))
		inputData["orgs"] = orgslist
		inputData["kafkas"] = kafkalist
		return tpl.Handler(inputData, TplConfigtx, ConfigDir()+"configtx.yaml")
	} else if strType == "crypto-config" {
		inputData := GetJsonMap("node.json")
		orgcounts := inputData[OrgCounts].(float64)
		var orgslist []string
		for i := 1; i <= int(orgcounts); i++ {
			orgslist = append(orgslist,strconv.Itoa(i))
		}
		inputData["orgs"] = orgslist
		return tpl.Handler(inputData, TplCryptoConfig, ConfigDir()+"crypto-config.yaml")
	} else if strType == "node" || strType == "client" {
		inputData := GetJsonMap("node.json")
		peerdomain := inputData[PeerDomain].(string)
		kfkdomain := inputData[KfkDomain].(string)
		kfkversion := inputData[KfkVersion].(string)
		list := inputData[List].([]interface{})
		for _, param := range list {
			value := param.(map[string]interface{})
			value[PeerDomain] = peerdomain
			value[KfkDomain] = kfkdomain
			value[KfkVersion] = kfkversion
			nodeType := value[NodeType].(string)
			dir := ConfigDir()
			var outfile, tplfile, yamlname string
			if strType == "client" {
				if nodeType == TypePeer {
					curorgid := value[OrgId].(string)
					value[Order0_Address] = findMapValue(TypeOrder, "0", curorgid)
					value[Order1_Address] = findMapValue(TypeOrder, "1", curorgid)
					//生成api 和  event yaml文件
					clientname := nodeType + value[PeerId].(string) + "org" + value[OrgId].(string)
					err := tpl.Handler(value, TplApiClient, ConfigDir()+clientname+"apiclient.yaml")
					if err != nil {
						return err
					}
					err = tpl.Handler(value, TplEventClient, ConfigDir()+clientname+"eventclient.yaml")
					if err != nil {
						return err
					}
				}
				continue
			}
			switch nodeType {
			case TypeZookeeper:
				curzkid := value[ZkId].(string)
				if curzkid == "1"{
					value[IP3] = findMapValue(TypeZookeeper, "3", "")
					value[IP4] = value[IP3]
					value[IP5] = findMapValue(TypeZookeeper, "5", "")
				}else if curzkid == "3"{
					value[IP1] = findMapValue(TypeZookeeper, "1", "")
					value[IP2] = value[IP1]
					value[IP5] = findMapValue(TypeZookeeper, "5", "")
				}else if curzkid == "5" {
					value[IP1] = findMapValue(TypeZookeeper, "1", "")
					value[IP2] = value[IP1]
					value[IP3] = findMapValue(TypeZookeeper, "3", "")
					value[IP4] = value[IP3]
				}
				yamlname = nodeType + curzkid + value[Zk2Id].(string)
				tplfile = TplZookeeper
			case TypeKafka:
				value[Zk_IP1] = findMapValue(TypeZookeeper, "1", "")
				value[Zk_IP2] = value[Zk_IP1]
				value[Zk_IP3] = findMapValue(TypeZookeeper, "3", "")
				value[Zk_IP4] = value[Zk_IP3]
				value[Zk_IP5] = findMapValue(TypeZookeeper, "5", "")
				yamlname = nodeType + value[KfkId].(string)
				tplfile = TplKafka
			case TypeOrder:
				yamlname = nodeType + value[OrderId].(string) + "ord" + value[OrgId].(string)
				tplfile = TplOrderer
			case TypePeer:
				curid := value[PeerId].(string)
				curorgid := value[OrgId].(string)
				value[Order0_Address] = findMapValue(TypeOrder, "0", curorgid)
				value[Order1_Address] = findMapValue(TypeOrder, "1", curorgid)
				value[USECOUCHDB] = inputData[USECOUCHDB].(string)
				if curid == "0" {
					value[Other_PeerAddress] = findMapValue(TypePeer, "1", curorgid)
				} else if curid == "1" {
					value[Other_PeerAddress] = findMapValue(TypePeer, "0", curorgid)
				}
				yamlname = nodeType + curid + "org" + curorgid
				tplfile = TplPeer
			}
			//生成yaml文件
			outfile = dir + yamlname
			err := tpl.Handler(value, tplfile, outfile+".yaml")
			if err != nil {
				fmt.Errorf(err.Error())
			}
		}
	} else {
		return fmt.Errorf("%s not exist", strType)
	}
	return nil
}

func CreateGenesisBlock() error {
	obj := NewFabCmd("apply_cert.py", "")
	err := obj.RunShow("generate_genesis_block", BinPath(), ConfigDir(), ConfigDir())
	if err != nil {
		return err
	}
	return nil
}

func CreateChannel(channelName string) error {
	obj := NewFabCmd("create_channel.py", "")
	err := obj.RunShow("create_channel", BinPath(), ConfigDir(), ChannelPath(), channelName, Domain_Name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAnchor(channelName string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer && value[PeerId].(string) == "0" {
			obj := NewFabCmd("create_channel.py", "")
			mspid := value[OrgId].(string)
			err := obj.RunShow("update_anchor", BinPath(), ConfigDir(), ChannelPath(), channelName, mspid, peerdomain)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func JoinChannel(channelName string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer {
			obj := NewFabCmd("create_channel.py", "")
			orgid := value[OrgId].(string)
			peerid := value[PeerId].(string)
			peer_address := "peer" + peerid + ".org" + orgid + "." + peerdomain + ":7051"
			err := obj.RunShow("join_channel", BinPath(), ConfigDir(), ChannelPath(), channelName, peer_address, peerid, orgid, peerdomain)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InstallChaincode(ccoutpath string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer {
			obj := NewFabCmd("chaincode.py", "")
			orgid := value[OrgId].(string)
			peerid := value[PeerId].(string)
			peer_address := "peer" + peerid + ".org" + orgid + "." + peerdomain + ":7051"
			err := obj.RunShow("install_chaincode", BinPath(), ConfigDir(), peer_address, peerid, orgid, peerdomain, ccoutpath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RunChaincode(ccname, channelName string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer {
			obj := NewFabCmd("chaincode.py", "")
			orgid := value[OrgId].(string)
			peerid := value[PeerId].(string)
			peer_address := "peer" + peerid + ".org" + orgid + "." + peerdomain + ":7051"
			initparam := `'{"Args":["init"\,"a"\,"100"\,"b"\,"200"]}'`
			policy := "\"OR  ('Org1MSP.member'\\,'Org2MSP.member'\\,'Org3MSP.member'\\,'Org4MSP.member'\\,'Org5MSP.member'" +
				"\\,'Org6MSP.member'\\,'Org7MSP.member'\\,'Org8MSP.member'\\,'Org9MSP.member'\\,'Org10MSP.member'" +
				"\\,'Org11MSP.member'\\,'Org12MSP.member'\\,'Org13MSP.member'\\,'Org14MSP.member'\\,'Org15MSP.member'" +
				"\\,'Org16MSP.member'\\,'Org17MSP.member'\\,'Org18MSP.member'\\,'Org19MSP.member'\\,'Org20MSP.member'" +
				"\\,'Org21MSP.member'\\,'Org22MSP.member'\\,'Org23MSP.member'\\,'Org24MSP.member'\\,'Org25MSP.member'" +
				"\\,'Org26MSP.member'\\,'Org27MSP.member'\\,'Org28MSP.member'\\,'Org29MSP.member'\\,'Org30MSP.member'" +
				"\\,'Org31MSP.member'\\,'Org32MSP.member'\\,'Org33MSP.member'\\,'Org34MSP.member'\\,'Org35MSP.member'" +
				"\\,'Org36MSP.member'\\,'Org37MSP.member'\\,'Org38MSP.member'\\,'Org39MSP.member'\\,'Org40MSP.member'" +
				"\\,'Org41MSP.member'\\,'Org42MSP.member'\\,'Org43MSP.member'\\,'Org44MSP.member'\\,'Org45MSP.member'" +
				"\\,'Org46MSP.member'\\,'Org47MSP.member'\\,'Org48MSP.member'\\,'Org49MSP.member'\\,'Org50MSP.member'" +
				"\\,'Org51MSP.member'\\,'Org52MSP.member'\\,'Org53MSP.member'\\,'Org54MSP.member'\\,'Org55MSP.member'" +
				"\\,'Org56MSP.member'\\,'Org57MSP.member'\\,'Org58MSP.member'\\,'Org59MSP.member'\\,'Org60MSP.member'" +
				"\\,'Org61MSP.member'\\,'Org62MSP.member'\\,'Org63MSP.member'\\,'Org64MSP.member'\\,'Org65MSP.member'" +
					")\""
			if orgid == "1" && peerid == "0" {
				err := obj.RunShow("instantiate_chaincode", BinPath(), ConfigDir(), peer_address, peerid, orgid, peerdomain, channelName, ccname, initparam, policy)
				if err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			} else {
				//txargs := `'{"Args":["query"\,"a"]}'`
				txargs := `'{"Args":["DslQuery"\,"trackid"\,"{\"dslSyntax\":\"{\\\"selector\\\":{\\\"sender\\\":\\\"zhengfu0\\\"}}\"}"]}'`
				err := obj.RunShow("test_query_tx", BinPath(), ConfigDir(), peer_address, peerid, orgid, peerdomain, channelName, ccname, txargs)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func PutCryptoConfig() error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer || nodeType == TypeOrder || nodeType == TypeKafka {
			obj := NewFabCmd("create_channel.py", value[IP].(string))
			err := obj.RunShow("put_cryptoconfig", ConfigDir(), nodeType)
			if err != nil {
				return err
			}
			if nodeType == TypePeer {
				obj := NewFabCmd("create_channel.py", value[APIIP].(string))
				err := obj.RunShow("put_cryptoconfig", ConfigDir(), TypeApi)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func findMapValue(findType, findid, findorgid string) string {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		nodeType := value[NodeType].(string)

		if nodeType == findType {
			switch findType {
			case TypeZookeeper:
				zkid := value[ZkId].(string)
				if zkid == findid{
					return value[IP].(string)
				}
			case TypeKafka:
				kfkid := value[KfkId].(string)
				if kfkid == findid{
					return value[IP].(string)
				}
			case TypeOrder:
				orderid := value[OrderId].(string)
				orgid := value[OrgId].(string)
				if orderid == findid && orgid == findorgid {
					return value[IP].(string)
				}
			case TypePeer:
				peerid := value[PeerId].(string)
				orgid := value[OrgId].(string)
				if peerid == findid && orgid == findorgid {
					return value[IP].(string)
				}
			}
		}
	}
	return "139.0.0.1"
}
