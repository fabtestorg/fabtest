package cmd

import (
	"fmt"
	"github.com/fabtestorg/fabtest/tpl"
	"sync"
	"strconv"
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
			chancounts := inputData[ChanCounts].(float64)
			for i:= 1 ; i <= int(chancounts) ; i++ {
				apiid := strconv.Itoa(i)
				apilist = append(apilist, value[APIIP].(string)+":"+apiid+apiid+apiid+apiid)
			}
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
	var wg sync.WaitGroup
	for _, param := range list {
		value := param.(map[string]interface{})
		if value[NodeType].(string) == TypePeer {
			chancounts := inputData[ChanCounts].(float64)
			for i:= 1 ; i <= int(chancounts) ; i++ {
				apiid := strconv.Itoa(i)
				clientname := TypePeer + value[PeerId].(string) + "org" + value[OrgId].(string) + "api" + apiid
				wg.Add(1)
				go func(Ip, CliName, Dir, LogDir string) {
					obj := NewFabCmd("jmeter.py", Ip)
					err := obj.RunShow("get_eventserver_log", CliName, Dir, LogDir)
					if err != nil {
						fmt.Println(err)
					}
					wg.Done()
				}(value[APIIP].(string), clientname, dir, logdir)
			}
		}
	}
	wg.Wait()
	return nil
}

func StartNmon() error {
	inputData := GetJsonMap("node.json")
	nmonMap := make(map[string]string)
	rate := inputData[Nmon_Rate].(string)
	times := inputData[Nmon_Times].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		nodeType := value[NodeType].(string)
		ip := value[IP].(string)
		switch nodeType {
		case TypeZookeeper:
			nmonMap[ip] = fmt.Sprintf("zookeeper%s", value[ZkId].(string))
		case TypeKafka:
			nmonMap[ip] = fmt.Sprintf("kafka%s", value[KfkId].(string))
		case TypeOrder:
			nmonMap[ip] = fmt.Sprintf("orderer%sorg%s", value[OrderId].(string), value[OrgId].(string))
		case TypePeer:
			nmonMap[ip] = fmt.Sprintf("peer%sorg%s", value[PeerId].(string), value[OrgId].(string))
			nmonMap[value[APIIP].(string)] = fmt.Sprintf("api%sorg%s", value[PeerId].(string), value[OrgId].(string))
		}
	}
	for curIp, fileName := range nmonMap {
		obj := NewFabCmd("jmeter.py", curIp)
		err := obj.RunShow("start_nmon", rate, times, fileName)
		if err != nil {
			fmt.Println("******start_nmon error******")
		}
	}

	return nil
}

func GetNmonLog(logdir string) error {
	inputData := GetJsonMap("node.json")
	nmonMap := make(map[string]string)
	rate := inputData[Nmon_Rate].(string)
	times := inputData[Nmon_Times].(string)
	list := inputData[List].([]interface{})
	for _, param := range list {
		value := param.(map[string]interface{})
		nodeType := value[NodeType].(string)
		ip := value[IP].(string)
		switch nodeType {
		case TypeZookeeper:
			nmonMap[ip] = fmt.Sprintf("zookeeper%s", value[ZkId].(string))
		case TypeKafka:
			nmonMap[ip] = fmt.Sprintf("kafka%s", value[KfkId].(string))
		case TypeOrder:
			nmonMap[ip] = fmt.Sprintf("orderer%sorg%s", value[OrderId].(string), value[OrgId].(string))
		case TypePeer:
			nmonMap[ip] = fmt.Sprintf("peer%sorg%s", value[PeerId].(string), value[OrgId].(string))
			nmonMap[value[APIIP].(string)] = fmt.Sprintf("api%sorg%s", value[PeerId].(string), value[OrgId].(string))
		}
	}
	for curIp, fileName := range nmonMap {
		obj := NewFabCmd("jmeter.py", curIp)
		err := obj.RunShow("get_nmon_log", rate, times, fileName, ConfigDir(), logdir)
		if err != nil {
			fmt.Println("*****get_nmon_log error******")
		}
	}
	return nil
}
