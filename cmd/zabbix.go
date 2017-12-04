package cmd

import (
	"github.com/peersafe/fabtest/tpl"
)

const (
	zabbix_server_ip   = "192.168.0.31"
	zabbix_server_port = "10056"
	TplZabbixConfig    = "./templates/zabbix-agent.tpl"
)

func CreateZabbixConfig() error {
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
		err := tpl.Handler(map[string]string{
			"server_ip":   zabbix_server_ip,
			"server_port": zabbix_server_port,
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
