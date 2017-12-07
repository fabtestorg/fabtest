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
	generateConfig := func(ip string) error {
		if _, ok := ipMap[ip]; ok {
			return nil
		}
		ipMap[ip] = struct{}{}
		return tpl.Handler(map[string]string{
			"server_ip":   server_ip,
			"server_port": server_port,
			"agent_ip":    ip,
			"agent_name":  ip,
		}, TplZabbixConfig, dir+ip+".conf")
	}
	for _, param := range list {
		value := param.(map[string]interface{})
		if err := generateConfig(value[IP].(string)); err != nil {
			return err
		}
		if ip, ok := value[APIIP].(string); ok {
			if err := generateConfig(ip); err != nil {
				return err
			}
		}
	}
	return nil
}

func StartZabbix() error {
	inputData := GetJsonMap("node.json")
	list := inputData[List].([]interface{})
	dir := ConfigDir()
	ipMap := make(map[string]struct{})
	start := func(ip string) error {
		if _, ok := ipMap[ip]; ok {
			return nil
		}
		ipMap[ip] = struct{}{}
		obj := NewFabCmd("zabbix.py", ip)
		return obj.RunShow("start_zabbix", ip, dir)
	}
	for _, param := range list {
		value := param.(map[string]interface{})
		if err := start(value[IP].(string)); err != nil {
			return err
		}
		if ip, ok := value[APIIP].(string); ok {
			if err := start(ip); err != nil {
				return err
			}
		}
	}
	return nil
}
