package cmd

import (
	"github.com/peersafe/fabtest/tpl"
)

const TplJmeterConfig = "./templates/jmeterconfig.tpl"

func CreateJmeterConfig() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			err := tpl.Handler(value["jmeter"], TplJmeterConfig, dir+clientname+".jmx")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func StartJmeter() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			ip := value["jmeter"].(map[string]interface{})["ip"].(string)
			obj := NewFabCmd("jmeter.py", ip)
			err := obj.RunShow("cp_jmeter_config", clientname, dir)
			if err != nil {
				return err
			}
			err = obj.RunShow("start_jmeter", clientname)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetJmeterLog() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			ip := value["jmeter"].(map[string]interface{})["ip"].(string)
			obj := NewFabCmd("jmeter.py", ip)
			err := obj.RunShow("get_jmeter_log", clientname, dir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetEventServerLog() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			ip := value["jmeter"].(map[string]interface{})["ip"].(string)
			obj := NewFabCmd("jmeter.py", ip)
			err := obj.RunShow("get_event_log", clientname, dir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
