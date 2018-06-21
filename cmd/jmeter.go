package cmd

import (
	"fmt"
	"github.com/fabtestorg/fabtest/tpl"
)

const TplJmeterConfig = "./templates/jmeterconfig.tpl"
const TplHaproxyConfig = "./templates/haproxycfg.tpl"

func CreateJmeterConfig() error {
	inputData := GetJsonMap("node.json")
	dir := ConfigDir()
	err := tpl.Handler(inputData, TplJmeterConfig, dir+"jmeter.jmx")
	if err != nil {
		return err
	}
	return nil
}

func CreateHaproxyConfig() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	var apilist []string
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			apilist = append(apilist, value[APIIP].(string))
		}
	}
	inputData["apilist"] = apilist
	err := tpl.Handler(inputData, TplHaproxyConfig, dir+"haproxy_config/haproxy.cfg")
	if err != nil {
		return err
	}
	return nil
}

func StartJmeter() error {
	inputData := GetJsonMap("node.json")
	value := inputData[JMETER].(map[string]interface{})
	dir := ConfigDir()
	obj := NewFabCmd("jmeter.py", value[IP].(string))
	err := obj.RunShow("start_jmeter", dir)
	if err != nil {
		fmt.Println("******star_jmeter error******")
	}
	return nil
}

func StartHaproxy() error {
	inputData := GetJsonMap("node.json")
	value := inputData[JMETER].(map[string]interface{})
	dir := ConfigDir()
	obj := NewFabCmd("jmeter.py", value[IP].(string))
	err := obj.RunShow("start_haproxy", dir)
	if err != nil {
		fmt.Println("******start_haproxy error******")
	}
	return nil
}

func GetJmeterLog(logdir string) error {
	inputData := GetJsonMap("node.json")
	value := inputData[JMETER].(map[string]interface{})
	dir := ConfigDir()
	obj := NewFabCmd("jmeter.py", value[IP].(string))
	err := obj.RunShow("get_jmeter_log", dir, logdir)
	if err != nil {
		fmt.Println("******get_jmeter_log error******")
	}
	return nil
}

func GetEventServerLog(logdir string) error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	obj := NewFabCmd("utils.py", "")
	err := obj.RunShow("rm_local", dir+"event_logs/"+logdir)
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
