package cmd

import (
	"github.com/peersafe/fabtest/tpl"
)

const (
	TplZabbixConfig = "./templates/zabbix-agent.tpl"
)

func CreateZabbixConfig() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	server_ip := inputData[ZabbixServerIp].(string)
	server_port := inputData[ZabbixServerPort].(string)
	dir := ConfigDir()
	ipMap := make(map[string]struct{})
	for _, param := range list {
		value := param.(map[string]interface{})
		ip := value[IP].(string)
		if _, ok := ipMap[ip]; ok {
			continue
		}
		err := tpl.Handler(map[string]string{
			"server_ip":   server_ip,
			"server_port": server_port,
			"agent_ip":    ip,
			"agent_name":  ip,
		}, TplZabbixConfig, dir+ip+".conf")
		if err != nil {
			return err
		}
	}
	return nil
}

func StartZabbix() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	ipMap := make(map[string]struct{})
	for _, param := range list {
		value := param.(map[string]interface{})
		ip := value[IP].(string)
		if _, ok := ipMap[ip]; ok {
			continue
		}
		obj := NewFabCmd("zabbix.py", ip)
		err := obj.RunShow("cp_zabbix_config", ip, dir)
		if err != nil {
			return err
		}
		err = obj.RunShow("start_zabbix")
		if err != nil {
			return err
		}
	}
	return nil
}
