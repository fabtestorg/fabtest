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
			err := tpl.Handler(value["jemter"], TplJmeterConfig, dir+clientname+".jmx")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
