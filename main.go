package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peersafe/fabtest/cmd"
)

var (
	file        = flag.String("f", "", "configtx, crypto-config, node, client, jmeter, zabbix ' create yaml file '")
	start       = flag.String("s", "", "peer, order, zookeeper, kafka, all ,api, jmeter, zabbix 'start node or api'")
	image       = flag.String("i", "", "peer, order, zookeeper, kafka, all  'load image'")
	create      = flag.String("c", "", "crypto, genesisblock, channel, 'create source'")
	getlog      = flag.String("g", "", "get jmeter or event logs")
	logdir      = flag.String("gn","", "log dir name eg: 50_50  loop 50*50")
	channelname = flag.String("n", "", "channelname")
	ccname      = flag.String("ccname", "", "chaincode name")
	ccoutpath   = flag.String("ccoutpath", "", "chaincode .out path")
	run         = flag.String("r", "", "joinchannel,  updateanchor, installchaincode, runchaincode")
	put         = flag.String("p", "", "put all (include crypto-config and channel-artifacts to remote)")
	deleteobj   = flag.String("d", "", "delete peer or kafka or zookeeper or all or api")
	analyse     = flag.String("a", "", "event analyse")
)

func main() {
	flag.Parse()
	var err error
	if *file != "" {
		if *file == "jmeter" {
			err = cmd.CreateJmeterConfig()
		} else if *file == "zabbix" {
			err = cmd.CreateZabbixConfig()
		} else {
			err = cmd.CreateYamlByJson(*file)
		}
	} else if *start != "" {
		if *start == "jmeter" {
			err = cmd.StartJmeter()
		} else if *start == "zabbix" {
			err = cmd.StartZabbix()
		} else {
			err = cmd.StartNode(*start)
		}
	} else if *image != "" {
		err = cmd.LoadImage(*image)
	} else if *create == "genesisblock" {
		err = cmd.CreateGenesisBlock()
	} else if *create == "crypto" {
		err = cmd.CreateCert()
	} else if *create == "channel" {
		if *channelname == "" {
			flag.Usage()
			fmt.Println("channel name is nil")
			os.Exit(1)
		}
		err = cmd.CreateChannel(*channelname)
	} else if *run == "updateanchor" {
		if *channelname == "" {
			flag.Usage()
			fmt.Println("channel name is nil")
			os.Exit(1)
		}
		err = cmd.UpdateAnchor(*channelname)
	} else if *run == "joinchannel" {
		if *channelname == "" {
			flag.Usage()
			fmt.Println("channel name is nil")
			os.Exit(1)
		}
		err = cmd.JoinChannel(*channelname)
	} else if *run == "installchaincode" {
		if *ccoutpath == "" {
			flag.Usage()
			fmt.Println("ccname or ccoutpath is nil")
			os.Exit(1)
		}
		err = cmd.InstallChaincode(*ccoutpath)
	} else if *run == "runchaincode" {
		if *ccname == "" || *channelname == "" {
			flag.Usage()
			fmt.Println("ccname or channel name is nil")
			os.Exit(1)
		}
		err = cmd.RunChaincode(*ccname, *channelname)
	} else if *getlog == "jmeter" {
		err = cmd.GetJmeterLog()
	} else if *getlog == "event" {
		if *logdir == ""{
			flag.Usage()
			fmt.Println("logdir is nil")
			os.Exit(1)
		}
		err = cmd.GetEventServerLog(*logdir)
	} else if *put != "" {
		err = cmd.PutCryptoConfig()
	} else if *deleteobj != "" {
		err = cmd.DeleteObj(*deleteobj)
	} else if *analyse != "" {
		if *logdir == ""{
			flag.Usage()
			fmt.Println("logdir is nil")
			os.Exit(1)
		}
		err = cmd.EventAnalyse(*logdir)
	} else {
		fmt.Println("Both data and file are nil.")
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
	}
}
