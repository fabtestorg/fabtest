package cmd

import (
	"fmt"
	"sync"
)

func StartNode(stringType string) error {
	if err := WriteHost(); err != nil {
		return err
	}
	var wg sync.WaitGroup
	StartN := func(Ip, Sshuser, Sshpwd, NodeName string, w1 *sync.WaitGroup) {
		defer w1.Done()
		obj := NewFabCmd("add_node.py", Ip, Sshuser, Sshpwd)
		err := obj.RunShow("start_node", NodeName, ConfigDir())
		if err != nil {
			fmt.Println("start node err or")
		}
	}
	if stringType == "all" || stringType == TypeKafka {
		for _, kafka := range GlobalConfig.Kafkas {
			wg.Add(1)
			nodeName := fmt.Sprintf("kafka%s", kafka.Id)
			go StartN(kafka.Ip, kafka.SshUserName, kafka.SshPwd, nodeName, &wg)
		}
	}
	if stringType == "all" || stringType == TypeZookeeper {
		for _, zk := range GlobalConfig.Zookeepers {
			wg.Add(1)
			nodeName := fmt.Sprintf("zk%s", zk.Id)
			go StartN(zk.Ip, zk.SshUserName, zk.SshPwd, nodeName, &wg)
		}
	}
	if stringType == "all" || stringType == TypeOrder {
		for _, ord := range GlobalConfig.Orderers {
			wg.Add(1)
			nodeName := fmt.Sprintf("orderer%s.ord%s.%s", ord.Id, ord.OrgId, GlobalConfig.Domain)
			go StartN(ord.Ip, ord.SshUserName, ord.SshPwd, nodeName, &wg)
		}
	}
	if stringType == "all" || stringType == TypePeer {
		for _, peer := range GlobalConfig.Peers {
			wg.Add(1)
			nodeName := fmt.Sprintf("peer%s.org%s.%s", peer.Id, peer.OrgId, GlobalConfig.Domain)
			go StartN(peer.Ip, peer.SshUserName, peer.SshPwd, nodeName, &wg)
		}
	}
	wg.Wait()
	return nil
}

func WriteHost() error {
	for _, ord := range GlobalConfig.Orderers {
		if err := LocalHostsSet(ord.Ip, fmt.Sprintf("orderer%s.ord%s.%s", ord.Id, ord.OrgId, GlobalConfig.Domain)); err != nil {
			return err
		}
	}
	for _, peer := range GlobalConfig.Peers {
		if err := LocalHostsSet(peer.Ip, fmt.Sprintf("peer%s.org%s.%s", peer.Id, peer.OrgId, GlobalConfig.Domain)); err != nil {
			return err
		}
	}
	for _, kafka := range GlobalConfig.Kafkas {
		if err := LocalHostsSet(kafka.Ip, fmt.Sprintf("kafka%s", kafka.Id)); err != nil {
			return err
		}
	}
	for _, zk := range GlobalConfig.Zookeepers {
		if err := LocalHostsSet(zk.Ip, fmt.Sprintf("zk%s", zk.Id)); err != nil {
			return err
		}
	}

	return nil
}

func DeleteObj(stringType string) error {
	var wg sync.WaitGroup
	StopN := func(Ip, Sshuser, Sshpwd, Ty string, w1 *sync.WaitGroup) {
		defer w1.Done()
		obj := NewFabCmd("removenode.py", Ip, Sshuser, Sshpwd)
		err := obj.RunShow("remove_node", Ty)
		if err != nil {
			fmt.Println("stopnode err or")
		}
	}
	if stringType == "all" || stringType == TypeKafka {
		for _, kafka := range GlobalConfig.Kafkas {
			wg.Add(1)
			go StopN(kafka.Ip, kafka.SshUserName, kafka.SshPwd, TypeKafka, &wg)
		}
	}
	if stringType == "all" || stringType == TypeZookeeper {
		for _, zk := range GlobalConfig.Zookeepers {
			wg.Add(1)
			go StopN(zk.Ip, zk.SshUserName, zk.SshPwd, TypeZookeeper, &wg)
		}
	}
	if stringType == "all" || stringType == TypeOrder {
		for _, ord := range GlobalConfig.Orderers {
			wg.Add(1)
			go StopN(ord.Ip, ord.SshUserName, ord.SshPwd, TypeOrder, &wg)
		}
	}
	if stringType == "all" || stringType == TypePeer {
		for _, peer := range GlobalConfig.Peers {
			wg.Add(1)
			go StopN(peer.Ip, peer.SshUserName, peer.SshPwd, TypePeer, &wg)
		}
	}
	wg.Wait()
	return nil
}

func LocalHostsSet(ip, domain string) error {
	if ip == domain {
		return nil
	}
	if err := ModifyHosts("/etc/hosts", ip, domain); err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}

func CheckNode(stringType string) error {
	if stringType == "all" || stringType == TypeKafka {
		for _, kafka := range GlobalConfig.Kafkas {
			obj := NewFabCmd("add_node.py", kafka.Ip, kafka.SshUserName, kafka.SshPwd)
			if err := obj.RunShow("check_node"); err != nil {
				return err
			}
		}
	}
	if stringType == "all" || stringType == TypeZookeeper {
		for _, zk := range GlobalConfig.Zookeepers {
			obj := NewFabCmd("add_node.py", zk.Ip, zk.SshUserName, zk.SshPwd)
			if err := obj.RunShow("check_node"); err != nil {
				return err
			}
		}
	}
	if stringType == "all" || stringType == TypeOrder {
		for _, ord := range GlobalConfig.Orderers {
			obj := NewFabCmd("add_node.py", ord.Ip, ord.SshUserName, ord.SshPwd)
			if err := obj.RunShow("check_node"); err != nil {
				return err
			}
		}
	}
	if stringType == "all" || stringType == TypePeer {
		for _, peer := range GlobalConfig.Peers {
			obj := NewFabCmd("add_node.py", peer.Ip, peer.SshUserName, peer.SshPwd)
			if err := obj.RunShow("check_node"); err != nil {
				return err
			}
		}
	}
	return nil
}
