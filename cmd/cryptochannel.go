package cmd

import (
	"fmt"
	"github.com/fabtestorg/fabtest/tpl"
	"strings"
	"sync"
)

func CreateCert() error {
	obj := NewLocalFabCmd("apply_cert.py")
	err := obj.RunShow("generate_certs", BinPath(), ConfigDir(), ConfigDir())
	if err != nil {
		return err
	}
	return nil
}

func CreateYamlByJson(strType string) error {
	if strType == "configtx" {
		return tpl.Handler(GlobalConfig, TplPath(TplConfigtx), ConfigDir()+"configtx.yaml")
	} else if strType == "crypto-config" {
		return tpl.Handler(GlobalConfig, TplPath(TplCryptoConfig), ConfigDir()+"crypto-config.yaml")
	} else if strType == "node" || strType == "client" {
		for _, ord := range GlobalConfig.Orderers {
			CopyConfig(&ord)
			outfile := ConfigDir() + fmt.Sprintf("orderer%s.ord%s.%s", ord.Id, ord.OrgId, GlobalConfig.Domain)
			if err := tpl.Handler(ord, TplPath(TplOrderer), outfile+".yaml"); err != nil {
				return err
			}
		}
		for _, peer := range GlobalConfig.Peers {
			CopyConfig(&peer)
			outfile := ConfigDir() + fmt.Sprintf("peer%s.org%s.%s", peer.Id, peer.OrgId, GlobalConfig.Domain)
			peer.DefaultNetwork = strings.Replace(fmt.Sprintf("peer%s.org%s.%s", peer.Id, peer.OrgId, GlobalConfig.Domain), ".", "", -1)
			if err := tpl.Handler(peer, TplPath(TplPeer), outfile+".yaml"); err != nil {
				return err
			}
		}
		for _, kafka := range GlobalConfig.Kafkas {
			outfile := ConfigDir() + fmt.Sprintf("kafka%s", kafka.Id)
			if err := tpl.Handler(kafka, TplPath(TplPeer), outfile+".yaml"); err != nil {
				return err
			}
		}
		for _, zk := range GlobalConfig.Zookeepers {
			outfile := ConfigDir() + fmt.Sprintf("zk%s", zk.Id)
			if err := tpl.Handler(zk, TplPath(TplZookeeper), outfile+".yaml"); err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("%s not exist", strType)
	}
	return nil
}

func CreateGenesisBlock() error {
	model := ""
	if GlobalConfig.ConsensusType == "solo" {
		model = "OrgsOrdererGenesis"
	} else if GlobalConfig.ConsensusType == "kafka" {
		model = "SampleDevModeKafka"
	} else if GlobalConfig.ConsensusType == "raft" {
		model = "SampleMultiNodeEtcdRaft"
	} else {
		return fmt.Errorf("ConsensusType %s unknow", GlobalConfig.ConsensusType)
	}
	obj := NewLocalFabCmd("apply_cert.py")
	err := obj.RunShow("generate_genesis_block", model, BinPath(), ConfigDir(), ConfigDir())
	if err != nil {
		return err
	}
	return nil
}

func CreateChannel(channelName string) error {
	if channelName == "" {
		return fmt.Errorf("channel name is nil")
	}
	obj := NewLocalFabCmd("create_channel.py")
	ordererAddress := ""
	for _, ord := range GlobalConfig.Orderers {
		ordererAddress = fmt.Sprintf("orderer%s.ord%s.%s:%s", ord.Id, ord.OrgId, GlobalConfig.Domain, ord.ConfigTxPort)
		break
	}
	err := obj.RunShow("create_channel", BinPath(), ConfigDir(), ChannelPath(), channelName, ordererAddress, GlobalConfig.Domain)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAnchor(channelName string) error {
	if channelName == "" {
		return fmt.Errorf("channel name is nil")
	}
	ordererAddress := ""
	for _, ord := range GlobalConfig.Orderers {
		ordererAddress = fmt.Sprintf("orderer%s.ord%s.%s:%s", ord.Id, ord.OrgId, GlobalConfig.Domain, ord.ConfigTxPort)
		break
	}
	for _, peer := range GlobalConfig.Peers {
		if peer.Id == "0" {
			obj := NewFabCmd("create_channel.py", peer.Ip, peer.SshUserName, peer.SshPwd)
			mspid := peer.OrgId
			err := obj.RunShow("update_anchor", BinPath(), ConfigDir(), ChannelPath(), channelName, mspid, ordererAddress, GlobalConfig.Domain)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CopyConfig(obj *NodeObj) {
	obj.Domain = GlobalConfig.Domain
	obj.Log = GlobalConfig.Log
	obj.Orderers = GlobalConfig.Orderers
	obj.Peers = GlobalConfig.Peers
	obj.UseCouchdb = GlobalConfig.UseCouchdb
	obj.ImageTag = GlobalConfig.ImageTag
	obj.ImagePre = GlobalConfig.ImagePre
}

func JoinChannel(channelName string) error {
	if channelName == "" {
		return fmt.Errorf("channel name is nil")
	}
	for _, peer := range GlobalConfig.Peers {
		peerAddress := fmt.Sprintf("peer%s.org%s.%s:%s", peer.Id, peer.OrgId, GlobalConfig.Domain, peer.ConfigTxPort)
		obj := NewLocalFabCmd("create_channel.py")
		err := obj.RunShow("join_channel", BinPath(), ConfigDir(), ChannelPath(), channelName, peerAddress, peer.Id, peer.OrgId, GlobalConfig.Domain)
		if err != nil {
			return err
		}
	}
	return nil
}

func PutCryptoConfig() error {
	var wg sync.WaitGroup
	putCrypto := func(ip, sshuser, sshpwd, cfg, nodeTy string, w1 *sync.WaitGroup) {
		obj := NewFabCmd("create_channel.py", ip, sshuser, sshpwd)
		err := obj.RunShow("put_cryptoconfig", cfg, nodeTy)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer w1.Done()
	}
	for _, kafka := range GlobalConfig.Kafkas {
		wg.Add(1)
		go putCrypto(kafka.Ip, kafka.SshUserName, kafka.SshPwd, ConfigDir(), "kafka", &wg)
	}
	for _, zk := range GlobalConfig.Zookeepers {
		wg.Add(1)
		go putCrypto(zk.Ip, zk.SshUserName, zk.SshPwd, ConfigDir(), "zk", &wg)
	}
	for _, ord := range GlobalConfig.Orderers {
		wg.Add(1)
		go putCrypto(ord.Ip, ord.SshUserName, ord.SshPwd, ConfigDir(), "orderer", &wg)
	}
	for _, peer := range GlobalConfig.Peers {
		wg.Add(1)
		go putCrypto(peer.Ip, peer.SshUserName, peer.SshPwd, ConfigDir(), "peer", &wg)
	}
	wg.Wait()
	return nil
}
