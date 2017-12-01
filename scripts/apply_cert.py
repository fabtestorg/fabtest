#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
import os
from fabric.api import local

reload(sys)
sys.setdefaultencoding('utf8')

def generate_genesis_block(bin_path, cfg_path ,out_path):
    if not os.path.exists(cfg_path + "core.yaml"):
        local("cp %s/core.yaml %s"%(bin_path, cfg_path))
    tool = bin_path + "configtxgen"
    out_path = out_path + "channel-artifacts"
    local("mkdir -p %s"%out_path)
    env = "FABRIC_CFG_PATH=%s"%cfg_path
    local("%s %s -profile OrgsOrdererGenesis -outputBlock %s/genesis.block"%(env,tool,out_path))

## Generates orderer Org certs using cryptogen tool
def generate_certs(bin_path, cfg_path ,out_path):
    cryptotool = bin_path + "cryptogen"
    yamlfile =  cfg_path + "crypto-config.yaml"
    outpath = out_path + "crypto-config"

    local("%s generate --config=%s --output='%s'"%(cryptotool,yamlfile,outpath))

