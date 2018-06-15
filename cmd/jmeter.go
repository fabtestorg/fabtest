package cmd

import (
	"fmt"
	"github.com/fabtestorg/fabtest/tpl"
	"sync"
)

const TplJmeterConfig = "./templates/jmeterconfig.tpl"
const TplHaproxyConfig = "./templates/haproxycfg.tpl"

func CreateJmeterConfig() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	tempMap := make(map[string]string)
	for _, param := range list {
		value := param.(map[string]interface{})
		value["jmeter"] = inputData["jmeter"]
		if value[NodeType].(string) == TypePeer {
			orgname := "org" + value[OrgId].(string)
			if _, ok := tempMap[orgname]; !ok {
				tempMap[orgname] = "already"
				//creat jmeter jmx request file
				err := tpl.Handler(param, TplJmeterConfig, dir+orgname+"jmeter.jmx")
				if err != nil {
					return err
				}
				//creat haproxy cfg
				value["apiip1"] = value[APIIP].(string)
				value["apiip2"] = findMapValue(TypePeer, "1", value[OrgId].(string), APIIP)
				err = tpl.Handler(param, TplHaproxyConfig, dir+orgname+"haproxy.cfg")
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func StartJmeter() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	tempMap := make(map[string]string)
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			orgname := "org" + value[OrgId].(string)
			if _, ok := tempMap[orgname]; !ok {
				tempMap[orgname] = "already"
				jmeterip := value[JMETERIP].(string)
				wg.Add(1)
				go func(filename, ip string) {
					obj := NewFabCmd("jmeter.py", ip)
					err := obj.RunShow("start_jmeter", filename, dir)
					if err != nil {
						fmt.Println("******star_jmeter error******", filename)
					}
					wg.Done()
				}(orgname, jmeterip)
			}
		}
	}
	wg.Wait()
	return nil
}

func StartHaproxy() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	tempMap := make(map[string]string)
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			orgname := "org" + value[OrgId].(string)
			if _, ok := tempMap[orgname]; !ok {
				tempMap[orgname] = "already"
				jmeterip := value[JMETERIP].(string)
				go func(filename, ip string) {
					obj := NewFabCmd("jmeter.py", ip)
					err := obj.RunShow("start_haproxy", filename, dir)
					if err != nil {
						fmt.Println("******start_haproxy error******", filename)
					}
				}(orgname, jmeterip)
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
