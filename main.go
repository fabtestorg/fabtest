package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fabtestorg/fabtest/cmd"
)

var (
	file        = flag.String("f", "", "configtx, crypto-config, node, client, jmeter, zabbix ' create yaml file '")
	start       = flag.String("s", "", "peer, order, zookeeper, kafka, all ,api, jmeter,nmon, zabbix 'start node or api'")
	create      = flag.String("c", "", "crypto, genesisblock, channel, 'create source'")
	getlog      = flag.String("g", "", "get jmeter or event or nmon logs")
	logdir      = flag.String("gn", "", "log dir name eg: 50_50  loop 50*50")
	channelname = flag.String("n", "", "channelname")
	ccname      = flag.String("ccname", "", "chaincode name")
	ccversion   = flag.String("version", "", "chaincode version")
	ccpath      = flag.String("ccpath", "", "chaincode go path")
	testArgs    = flag.String("args", "", "test chaincode args")
	function    = flag.String("func", "invoke", "invoke or query")
	run         = flag.String("r", "", "joinchannel,  updateanchor, installchaincode, runchaincode, checknode, upgradecc,testcc")
	put         = flag.String("p", "", "put all (include crypto-config and channel-artifacts to remote)")
	deleteobj   = flag.String("d", "", "delete peer or kafka or zookeeper or all or api")
	analyse     = flag.String("a", "", "event analyse")
)

func main() {
	flag.Parse()
	var err error
	cmd.GlobalConfig, err = cmd.ParseJson("node.json")
	if err != nil {
		panic(err)
	}
	if *file != "" {
		if *file == "jmeter" {
			err = cmd.CreateJmeterConfig()
		} else if *file == "haproxy" {
			err = cmd.CreateHaproxyConfig()
		} else {
			err = cmd.CreateYamlByJson(*file)
		}
	} else if *start != "" {
		if *start == "jmeter" {
			err = cmd.CreateJmeterConfig()
			if err == nil {
				err = cmd.StartJmeter()
			}
		} else if *start == "haproxy" {
			err = cmd.StartHaproxy()
		} else {
			err = cmd.StartNode(*start)
		}
	} else if *create == "genesisblock" {
		err = cmd.CreateGenesisBlock()
	} else if *create == "crypto-config" {
		err = cmd.CreateCert()
	} else if *create == "channel" {
		err = cmd.CreateChannel(*channelname)
	} else if *run == "updateanchor" {
		err = cmd.UpdateAnchor(*channelname)
	} else if *run == "joinchannel" {
		err = cmd.JoinChannel(*channelname)
	} else if *run == "installchaincode" {
		err = cmd.InstallChaincode(*ccname, *ccversion, *ccpath)
	} else if *run == "runchaincode" {
		err = cmd.RunChaincode(*ccname, *ccversion, *channelname, "instantiate")
	} else if *run == "upgradecc" {
		err = cmd.RunChaincode(*ccname, *ccversion, *channelname, "upgrade")
	} else if *run == "testcc" {
		err = cmd.TestChaincode(*ccname, *channelname, *function, *testArgs)
	} else if *run == "checknode" {
		err = cmd.CheckNode("all")
	} else if *getlog == "jmeter" {
		err = cmd.GetJmeterLog(*logdir)
	} else if *getlog == "event" {
		err = cmd.GetEventServerLog(*logdir)
	} else if *put != "" {
		err = cmd.PutCryptoConfig()
	} else if *deleteobj != "" {
		err = cmd.DeleteObj(*deleteobj)
	} else if *analyse != "" {
		err = cmd.EventAnalyse(*logdir)
	} else {
		fmt.Println("Both data and file are nil.")
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
}
