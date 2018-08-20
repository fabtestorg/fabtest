#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
import os
from fabric.api import local, lcd

reload(sys)
sys.setdefaultencoding('utf8')

def generate_genesis_block(bin_path, cfg_path ,out_path):
    if not os.path.exists(out_path + "crypto-config"):
        with lcd(out_path):
            local("tar -zxvf crypto-config.tar.gz")
    if not os.path.exists(cfg_path + "core.yaml"):
        local("cp %s/core.yaml %s"%(bin_path, cfg_path))
    tool = bin_path + "configtxgen"
    channel_path = out_path + "channel-artifacts"
    local("rm -rf %s"%channel_path)
    local("mkdir -p %s"%channel_path)
    env = "FABRIC_CFG_PATH=%s"%cfg_path
    local("%s %s -profile OrgsOrdererGenesis -outputBlock %s/genesis.block"%(env,tool,channel_path))
    with lcd(out_path):
        local("tar -zcvf channel-artifacts.tar.gz channel-artifacts")

## Generates orderer Org certs using cryptogen tool
def generate_certs(bin_path, cfg_path ,out_path):
    cryptotool = bin_path + "cryptogen"
    yamlfile =  cfg_path + "crypto-config.yaml"
    mm_path = out_path + "crypto-config"

    with lcd(out_path):
        local("rm -rf crypto-config.tar.gz crypto-config")
    local("%s generate --config=%s --output='%s'"%(cryptotool,yamlfile,mm_path))
    with lcd(out_path):
        local("tar -zxvf cryptobench.tar.gz")
        local("rm -rf crypto-config/peerOrganizations")
        local("cp -r cryptobench/peerOrganizations crypto-config/")
        local("tar -zcvf crypto-config.tar.gz crypto-config")
