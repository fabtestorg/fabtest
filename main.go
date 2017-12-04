package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/peersafe/fabtest/cmd"
)

var (
	file        = flag.String("f", "", "configtx, crypto-config, node, client, jmeter ' create yaml file '")
	start       = flag.String("s", "", "peer, order, zookeeper, kafka, all ,api, jmeter 'start node or api'")
	image       = flag.String("i", "", "peer, order, zookeeper, kafka, all  'load image'")
	create      = flag.String("c", "", "crypto, genesisblock, channel, 'create source'")
	getlog      = flag.String("g", "", "get jmeter logs")
	channelname = flag.String("n", "", "channelname")
	ccname      = flag.String("ccname", "", "chaincode name")
	ccoutpath   = flag.String("ccoutpath", "", "chaincode path")
	run         = flag.String("r", "", "joinchannel,  updateanchor, installchaincode, runchaincode")
	put         = flag.String("p", "", "put all (include crypto-config and channel-artifacts to remote)")
)

func main() {
	flag.Parse()
	var err error
	if *file != "" {
		if *file == "jmeter" {
			err = cmd.CreateJmeterConfig()
		} else {
			err = cmd.CreateYamlByJson(*file)
		}
	} else if *start != "" {
		if *start == "jmeter" {
			err = cmd.StartJmeter()
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
		if *ccname == "" || *ccoutpath == "" {
			flag.Usage()
			fmt.Println("ccname or ccoutpath is nil")
			os.Exit(1)
		}
		err = cmd.InstallChaincode(*ccname, *ccoutpath)
	} else if *run == "runchaincode" {
		if *ccname == "" || *channelname == "" {
			flag.Usage()
			fmt.Println("ccname or channel name is nil")
			os.Exit(1)
		}
		err = cmd.RunChaincode(*ccname, *ccoutpath)
	} else if *getlog == "jmeter" {
		err = cmd.GetJmeterLog()
	} else if *put != "" {
		err = cmd.PutCryptoConfig()
	} else {
		fmt.Println("Both data and file are nil.")
		flag.Usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
	}
}
