package cmd

import (
	"fmt"
	"sync"
)

func InstallChaincode(ccname, ccversion, ccpath string) error {
	if ccpath == "" {
		ccpath = GlobalConfig.CCPath
	}
	if ccversion == "" {
		ccversion = GlobalConfig.CCVersion
	}
	if ccname == "" {
		ccname = GlobalConfig.CCName
	}
	var wg sync.WaitGroup
	for _, peer := range GlobalConfig.Peers {
		wg.Add(1)
		peerAddress := fmt.Sprintf("peer%s.org%s.%s:%s", peer.Id, peer.OrgId, GlobalConfig.Domain, peer.ConfigTxPort)
		go func(binPath, configDir, peerAds, PeerId, OrgId, Pdn string) {
			defer wg.Done()
			obj := NewFabCmd("chaincode.py", peer.Ip, peer.SshUserName, peer.SshPwd)
			err := obj.RunShow("install_chaincode", binPath, configDir, peerAds, PeerId, OrgId, Pdn, ccname, ccversion, ccpath)
			if err != nil {
				fmt.Printf(err.Error())
			}
		}(BinPath(), ConfigDir(), peerAddress, peer.Id, peer.OrgId, GlobalConfig.Domain)
	}
	wg.Wait()
	return nil
}

func RunChaincode(ccname, ccversion, channelName, opration string) error {
	if channelName == "" {
		return fmt.Errorf("channel is nil")
	}
	if ccversion == "" {
		ccversion = GlobalConfig.CCVersion
	}
	if ccname == "" {
		ccname = GlobalConfig.CCName
	}
	ordererAddress := ""
	for _, ord := range GlobalConfig.Orderers {
		ordererAddress = fmt.Sprintf("orderer%s.ord%s.%s:%s", ord.Id, ord.OrgId, GlobalConfig.Domain, ord.ConfigTxPort)
		break
	}
	var wg sync.WaitGroup
	for _, peer := range GlobalConfig.Peers {
		wg.Add(1)
		peerAddress := fmt.Sprintf("peer%s.org%s.%s:%s", peer.Id, peer.OrgId, GlobalConfig.Domain, peer.ConfigTxPort)
		go func(binPath, configDir, peerAds, PeerId, OrgId, Pdn string) {
			defer wg.Done()
			obj := NewFabCmd("chaincode.py", peer.Ip, peer.SshUserName, peer.SshPwd)
			initparam := fmt.Sprintf(`%s`, GlobalConfig.CCInit)
			policy := fmt.Sprintf("%s", GlobalConfig.CCPolicy)
			err := obj.RunShow("instantiate_chaincode", BinPath(), opration, ConfigDir(), peerAds, ordererAddress, PeerId, OrgId, GlobalConfig.Domain, channelName, ccname, ccversion, initparam, policy)
			if err != nil {
				fmt.Println(err)
			}
		}(BinPath(), ConfigDir(), peerAddress, peer.Id, peer.OrgId, GlobalConfig.Domain)
	}
	wg.Wait()
	return nil
}

func TestChaincode(ccname, channelName, testArgs string) error {
	if channelName == "" {
		return fmt.Errorf("channel is nil")
	}
	if ccname == "" {
		ccname = GlobalConfig.CCName
	}
	if testArgs == "" {
		testArgs = GlobalConfig.TestArgs
	}
	ordererAddress := ""
	for _, ord := range GlobalConfig.Orderers {
		ordererAddress = fmt.Sprintf("orderer%s.ord%s.%s:%s", ord.Id, ord.OrgId, GlobalConfig.Domain, ord.ConfigTxPort)
		break
	}
	var wg sync.WaitGroup
	for _, peer := range GlobalConfig.Peers {
		wg.Add(1)
		peerAddress := fmt.Sprintf("peer%s.org%s.%s:%s", peer.Id, peer.OrgId, GlobalConfig.Domain, peer.ConfigTxPort)
		go func(binPath, configDir, peerAds, PeerId, OrgId, Pdn string) {
			defer wg.Done()
			obj := NewFabCmd("chaincode.py", peer.Ip, peer.SshUserName, peer.SshPwd)
			err := obj.RunShow("test_chaincode", BinPath(), ConfigDir(), peerAds, ordererAddress, PeerId, OrgId, GlobalConfig.Domain, channelName, ccname, testArgs)
			if err != nil {
				fmt.Println(err)
			}
		}(BinPath(), ConfigDir(), peerAddress, peer.Id, peer.OrgId, GlobalConfig.Domain)
	}
	wg.Wait()
	return nil
}
