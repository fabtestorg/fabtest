package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

func StartNode(stringType string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	kfkdomain := inputData[KfkDomain].(string)
	list := inputData[List].([]interface{})
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		value[KfkDomain] = kfkdomain
		nodeType := value[NodeType].(string)
		if nodeType != stringType {
			if stringType == "api" && nodeType == TypePeer {
				//启动api
				peerid := value[PeerId].(string)
				orgid := value[OrgId].(string)
				wg.Add(1)
				go func(IP, PeerId, OrgId string) {
					obj := NewFabCmd("add_node.py", IP)
					err := obj.RunShow("start_api", PeerId, OrgId, ConfigDir())
					if err != nil {
						fmt.Printf(err.Error())
					}
					wg.Done()
				}(value[APIIP].(string), peerid, orgid)
				err := LocalHostsSet(value[APIIP].(string), fmt.Sprintf("api%s%s", orgid, peerid))
				if err != nil {
					return err
				}
				continue
			} else if stringType == "event" && nodeType == TypePeer {
				//启动api
				peerid := value[PeerId].(string)
				orgid := value[OrgId].(string)
				//wg.Add(1)
				//go func(IP, PeerId, OrgId string) {
				obj := NewFabCmd("add_node.py", value[APIIP].(string))
				err := obj.RunShow("start_event", peerid, orgid, ConfigDir(), "event")
				if err != nil {
					fmt.Println(err)
				}
				//wg.Done()
				//}(value[APIIP].(string), peerid, orgid)
				continue
			} else if stringType != "all" {
				continue
			}
		}
		var nodeId, yamlname string
		var ip = value[IP].(string)
		switch nodeType {
		case TypeZookeeper:
			nodeId = value[ZkId].(string)
			yamlname = nodeType + value[ZkId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("zk%s", nodeId))
			if err != nil {
				return err
			}
		case TypeKafka:
			nodeId = value[KfkId].(string)
			yamlname = nodeType + value[KfkId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("kafka%s", nodeId))
			if err != nil {
				return err
			}
		case TypeOrder:
			nodeId = value[OrderId].(string)
			ordId := value[OrgId].(string)
			yamlname = nodeType + value[OrderId].(string) + "ord" + ordId
			err := LocalHostsSet(ip, fmt.Sprintf("orderer%s.ord%s.%s", nodeId, ordId, peerdomain))
			if err != nil {
				return err
			}
			err = LocalHostsSet(ip, fmt.Sprintf("orderer%s%s", ordId, nodeId))
			if err != nil {
				return err
			}
		case TypePeer:
			nodeId = value[PeerId].(string)
			orgId := value[OrgId].(string)
			yamlname = nodeType + nodeId + "org" + orgId
			err := LocalHostsSet(ip, "peer"+nodeId+".org"+orgId+"."+peerdomain)
			if err != nil {
				return err
			}
			err = LocalHostsSet(ip, fmt.Sprintf("peer%s%s", orgId, nodeId))
			if err != nil {
				return err
			}
		}
		//启动节点
		wg.Add(1)
		go func(Ip, NodeType, NodeId, YamleName string) {
			obj := NewFabCmd("add_node.py", Ip)
			err := obj.RunShow("start_node", NodeType, NodeId, YamleName, ConfigDir())
			if err != nil {
				fmt.Errorf("start node err or")
				return
			}
			wg.Done()
		}(ip, nodeType, nodeId, yamlname)
	}
	wg.Wait()
	return nil
}

func WriteHost(stringType string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	kfkdomain := inputData[KfkDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		value[KfkDomain] = kfkdomain
		nodeType := value[NodeType].(string)
		if stringType != "all" {
			if nodeType != stringType {
				continue
			}
		}
		var nodeId string
		var ip = value[IP].(string)
		switch nodeType {
		case TypeZookeeper:
			nodeId = value[ZkId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("zk%s", nodeId))
			if err != nil {
				return err
			}
		case TypeKafka:
			nodeId = value[KfkId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("kafka%s", nodeId))
			if err != nil {
				return err
			}
		case TypeOrder:
			nodeId = value[OrderId].(string)
			ordId := value[OrgId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("orderer%s%s", ordId, nodeId))
			if err != nil {
				return err
			}
		case TypePeer:
			nodeId = value[PeerId].(string)
			orgId := value[OrgId].(string)
			err := LocalHostsSet(ip, fmt.Sprintf("peer%s%s", orgId, nodeId))
			if err != nil {
				return err
			}
			err = LocalHostsSet(value[APIIP].(string), fmt.Sprintf("api%s%s", orgId, nodeId))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func ReplaceImage(imagesType, id string) error {
	var inputData map[string]interface{}
	var jsonData []byte
	var err error

	inputfile := InputDir() + "node.json"
	jsonData, err = ioutil.ReadFile(inputfile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &inputData)
	if err != nil {
		return err
	}
	list := inputData[List].([]interface{})
	var nodeId string
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		nodeType := value[NodeType].(string)
		if id != "" {
			switch nodeType {
			case TypeZookeeper:
				nodeId = value[ZkId].(string)
			case TypeKafka:
				nodeId = value[KfkId].(string)
			case TypeOrder:
				nodeId = value[OrderId].(string)
			case TypePeer:
				nodeId = value[PeerId].(string)
			}
			if nodeId != id {
				continue
			}
		}
		if nodeType == imagesType {
			//copy images
			wg.Add(1)
			go func(Ip, Ty string) {
				obj := NewFabCmd("add_node.py", Ip)
				err := obj.RunShow("replace_images", Ty, ConfigDir())
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}(value[IP].(string), nodeType)
		}
	}
	wg.Wait()
	return nil
}
func LoadImage(stringType string) error {
	var inputData map[string]interface{}
	var jsonData []byte
	var err error

	inputfile := InputDir() + "node.json"
	jsonData, err = ioutil.ReadFile(inputfile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, &inputData)
	if err != nil {
		return err
	}

	if stringType == "all" {
		list := inputData[List].([]interface{})
		for _, param := range list {
			value := param.(map[string]interface{})
			nodeType := value[NodeType].(string)
			//copy images
			obj := NewFabCmd("add_node.py", value[IP].(string))
			err = obj.RunShow("load_images", nodeType, ImagePath())
			if err != nil {
				return err
			}
			if nodeType == TypePeer {
				err = obj.RunShow("load_images", "baseos", ImagePath())
				if err != nil {
					return err
				}
				err = obj.RunShow("load_images", "ccenv", ImagePath())
				if err != nil {
					return err
				}
			}
		}
	} else {
		return fmt.Errorf("%s not exist", stringType)
	}
	return nil
}

func DeleteObj(stringType string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	kfkdomain := inputData[KfkDomain].(string)
	list := inputData[List].([]interface{})
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		value[KfkDomain] = kfkdomain
		nodeType := value[NodeType].(string)
		if nodeType == stringType {
			//删除节点
			wg.Add(1)
			go func(Ip, Str string) {
				obj := NewFabCmd("removenode.py", Ip)
				err := obj.RunShow("remove_node", Str)
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}(value[IP].(string), stringType)
		} else if stringType == TypeApi {
			if nodeType == TypePeer {
				wg.Add(1)
				go func(Ip string) {
					obj := NewFabCmd("removenode.py", Ip)
					err := obj.RunShow("remove_client")
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(value[APIIP].(string))
			}
		} else if stringType == "all" && (nodeType == TypeKafka || nodeType == TypeZookeeper ||
			nodeType == TypePeer || nodeType == TypeOrder) {
			//删除节点
			wg.Add(1)
			go func(Ip, str string) {
				obj := NewFabCmd("removenode.py", Ip)
				err := obj.RunShow("remove_node", str)
				if err != nil {
					fmt.Println(err.Error())
				}
				wg.Done()
			}(value[IP].(string), stringType)
			if nodeType == TypePeer {
				wg.Add(1)
				go func(Ip string) {
					obj := NewFabCmd("removenode.py", Ip)
					err := obj.RunShow("remove_client")
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(value[APIIP].(string))
			}
		} else if stringType == "jmeter" && nodeType == TypeHaproxy {
			wg.Add(1)
			go func(Ip string) {
				defer wg.Done()
				obj := NewFabCmd("removenode.py", Ip)
				if err := obj.RunShow("remove_jmeter"); err != nil {
					fmt.Println(err)
				}
			}(value[IP].(string))
		}
	}
	wg.Wait()
	return nil
}

func OperationNode(cmdstr string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	kfkdomain := inputData[KfkDomain].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		value[KfkDomain] = kfkdomain
		nodeType := value[NodeType].(string)
		if nodeType == TypePeer || nodeType == TypeOrder {
			var nodeId, yamlname string
			if nodeType == TypeOrder {
				nodeId = value[OrderId].(string)
				yamlname = nodeType + value[OrderId].(string) + "ord" + value[OrgId].(string)
			} else if nodeType == TypePeer {
				nodeId = value[PeerId].(string)
				yamlname = nodeType + value[PeerId].(string) + "org" + value[OrgId].(string)
			}
			//删除节点
			obj := NewFabCmd("add_node.py", value[IP].(string))
			var err error
			if cmdstr == "stop" {
				err = obj.RunShow("stop_node", nodeType, nodeId, yamlname)
			} else if cmdstr == "start" {
				err = obj.RunShow("restart_node", nodeType, nodeId, yamlname)
			}
			if err != nil {
				return err
			}
		}
	}

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

func StartDocker() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		nodeType := value[NodeType].(string)
		if nodeType != TypePeer && nodeType != TypeKafka && nodeType != TypeZookeeper && nodeType != TypeOrder {
			continue
		}
		//启动docker server
		wg.Add(1)
		go func(Ip string) {
			obj := NewFabCmd("add_node.py", Ip)
			err := obj.RunShow("start_docker")
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(value[IP].(string))

		if nodeType == TypePeer {
			wg.Add(1)
			go func(Ip string) {
				defer wg.Done()
				obj := NewFabCmd("add_node.py", Ip)
				err := obj.RunShow("start_docker")
				if err != nil {
					fmt.Println(err)
				}
			}(value[APIIP].(string))
		}

	}
	wg.Wait()
	return nil
}
