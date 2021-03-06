package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func StartNode(stringType string) error {
	inputData := GetJsonMap("node.json")
	peerdomain := inputData[PeerDomain].(string)
	kfkdomain := inputData[KfkDomain].(string)
	list := inputData[List].([]interface{})
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
				obj := NewFabCmd("add_node.py", value[APIIP].(string))
				err := obj.RunShow("start_api", peerid, orgid, ConfigDir())
				if err != nil {
					return err
				}
				continue
			} else if stringType == "event" && nodeType == TypePeer {
				//启动api
				peerid := value[PeerId].(string)
				orgid := value[OrgId].(string)
				obj := NewFabCmd("add_node.py", value[APIIP].(string))
				err := obj.RunShow("start_event", peerid, orgid, ConfigDir())
				if err != nil {
					return err
				}
				continue
			} else if stringType != "all" {
				continue
			}
		}
		var nodeId, yamlname string
		var ip = value[IP].(string)
		switch nodeType {
		case TypeZookeeper:
			nodeId = value[ZkId].(string) + value[Zk2Id].(string)
			yamlname = nodeType + value[ZkId].(string) + value[Zk2Id].(string)
		case TypeKafka:
			nodeId = value[KfkId].(string)
			yamlname = nodeType + value[KfkId].(string)
		case TypeOrder:
			nodeId = value[OrderId].(string)
			ordId := value[OrgId].(string)
			yamlname = nodeType + value[OrderId].(string) + "ord" + ordId
			LocalHostsSet(ip, fmt.Sprintf("orderer%s.ord%s.%s", nodeId, ordId, peerdomain))
		case TypePeer:
			nodeId = value[PeerId].(string)
			orgId := value[OrgId].(string)
			yamlname = nodeType + nodeId + "org" + orgId
			LocalHostsSet(ip, "peer"+nodeId+".org"+orgId+"."+peerdomain)
		}
		//启动节点
		obj := NewFabCmd("add_node.py", ip)
		err := obj.RunShow("start_node", nodeType, nodeId, yamlname, ConfigDir())
		if err != nil {
			return err
		}
	}

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
	for _, param := range list {
		value := param.(map[string]interface{})
		value[PeerDomain] = peerdomain
		value[KfkDomain] = kfkdomain
		nodeType := value[NodeType].(string)
		if nodeType == stringType {
			//删除节点
			obj := NewFabCmd("removenode.py", value[IP].(string))
			err := obj.RunShow("remove_node", stringType)
			if err != nil {
				return err
			}
		} else if stringType == TypeApi {
			if nodeType == TypePeer {
				obj := NewFabCmd("removenode.py", value[APIIP].(string))
				err := obj.RunShow("remove_client")
				if err != nil {
					return err
				}
			}
		} else if stringType == "all" && ( nodeType == TypeKafka || nodeType == TypeZookeeper ||
			nodeType == TypePeer || nodeType == TypeOrder ){
			//删除节点
			obj := NewFabCmd("removenode.py", value[IP].(string))
			err := obj.RunShow("remove_node", stringType)
			if err != nil {
				return err
			}
			if nodeType == TypePeer {
				obj := NewFabCmd("removenode.py", value[APIIP].(string))
				err := obj.RunShow("remove_client")
				if err != nil {
					return err
				}
			}
		}
	}

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
	file, err := os.Open("/etc/hosts")
	if err != nil {
		return err
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.HasSuffix(line, " "+domain) {
			if strings.HasPrefix(line, ip+" ") {
				return nil
			}
			return fmt.Errorf("The domain has been maped <%s>", line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	file, err = os.OpenFile("/etc/hosts", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	buf := []byte(fmt.Sprintf("%s       %s\n", ip, domain))
	_, err = file.Write(buf)
	if err != nil {
		return err
	}
	return file.Close()
}
