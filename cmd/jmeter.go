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
		value["jmeter"] = inputData["jmeter"]
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			err := tpl.Handler(param, TplJmeterConfig, dir+clientname+".jmx")
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
			obj := NewFabCmd("jmeter.py", value[APIIP].(string))
			err := obj.RunShow("start_jmeter", clientname, dir)
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
			obj := NewFabCmd("jmeter.py", value[APIIP].(string))
			err := obj.RunShow("get_jmeter_log", clientname, dir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetEventServerLog(logdir string) error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	obj := NewFabCmd("utils.py", "")
	err := obj.RunShow("rm_local", dir + logdir)
	if err != nil {
		return err
	}
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string)
			obj := NewFabCmd("jmeter.py", value[APIIP].(string))
			err := obj.RunShow("get_eventserver_log", clientname, dir, logdir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
