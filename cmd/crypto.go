package cmd

import (
	"fmt"
	"github.com/peersafe/fabtest/tpl"
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
		inputData := GetJsonMap("configtx.json")
		return tpl.Handler(inputData, TplConfigtx, ConfigDir()+"configtx.yaml")
	} else if strType == "crypto-config" {
		inputData := GetJsonMap("crypto-config.json")
		return tpl.Handler(inputData, TplCryptoConfig, ConfigDir()+"crypto-config.yaml")
	} else if strType == "node" || strType == "client" {
		inputData := GetJsonMap("node.json")
		peerdomain := inputData[PeerDomain].(string)
		kfkdomain := inputData[KfkDomain].(string)
		list := inputData[List].([]interface{})
		for _, param := range list {
			value := param.(map[string]interface{})
			value[PeerDomain] = peerdomain
			value[KfkDomain] = kfkdomain
			nodeType := value[NodeType].(string)
			dir := ConfigDir()
			var outfile, tplfile, yamlname string
			if strType == "client" {
				if nodeType == TypePeer {
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
				yamlname = nodeType + value[ZkId].(string) + value[Zk2Id].(string)
				tplfile = TplZookeeper
			case TypeKafka:
				yamlname = nodeType + value[KfkId].(string)
				tplfile = TplKafka
			case TypeOrder:
				yamlname = nodeType + value[OrderId].(string) + "org" + value[OrgId].(string)
				tplfile = TplOrderer
			case TypePeer:
				yamlname = nodeType + value[PeerId].(string) + "org" + value[OrgId].(string)
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
			mspid := "Org" + value[OrgId].(string) + "MSP"
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
			initparam := `{"Args":["init"\,"a"\,"100"\,"b"\,"200"]}`
			policy := `'OR  ('Org1MSP.member'\,'Org2MSP.member'\,'Org3MSP.member'\,'Org5MSP.member')'`
			if orgid == "1" && peerid == "0" {
				err := obj.RunShow("instantiate_chaincode", BinPath(), ConfigDir(), peer_address, peerid, orgid, peerdomain, channelName, ccname, initparam, policy)
				if err != nil {
					return err
				}
			} else {
				txargs := `{"Args":["query"\,"a"]}`
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
		}
	}
	return nil
}
