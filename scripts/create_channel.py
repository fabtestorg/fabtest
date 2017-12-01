#!/usr/bin/env python
# -*- coding: utf-8 -*-

from fabric.api import local
import sys
import os
reload(sys)
sys.setdefaultencoding('utf8')

## create channel

def create_channel(bin_path, yaml_path, out_path, channel_name, domain_name):
    if not os.path.exists(yaml_path + "core.yaml"):
        local("cp %s/core.yaml %s"%(bin_path, yaml_path))
    ret = create_channeltx(bin_path, yaml_path, out_path, channel_name)
    print ret
    channeltx_name = channel_name + '.tx'
    msp_path = yaml_path + "crypto-config/peerOrganizations/org1.%s/users/Admin@org1.%s/msp"%(domain_name,domain_name)
    channel_dir = out_path + channel_name
    order_tls_path = yaml_path +  "crypto-config/ordererOrganizations/ord1.%s/orderers/orderer0.ord1.%s/msp/tlscacerts/tlsca.ord1.%s-cert.pem"%(domain_name,domain_name,domain_name)
    env = 'FABRIC_CFG_PATH=%s '%yaml_path
    env = env + 'CORE_PEER_LOCALMSPID=Org1MSP '
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s  '%msp_path
    bin = bin_path + "peer"
    order_address = "orderer0.ord1.%s:7050"%domain_name
    param = ' channel create -o %s -t 3000 -c %s -f %s/%s'%(order_address, channel_name, channel_dir, channeltx_name)

    tls = ' --tls --cafile %s'%order_tls_path

    command = env + bin + param + tls
    local(command)

def create_channeltx(bin_path, yaml_path, out_path, channel_name):
    bin = bin_path + "configtxgen"
    channel_dir = out_path + channel_name
    if not os.path.exists(channel_dir):
        local("mkdir -p %s"%channel_dir)
    channeltx_name = channel_name + '.tx'
    env = 'FABRIC_CFG_PATH=%s '%yaml_path
    param = ' -profile OrgsChannel -outputCreateChannelTx %s/%s -channelID %s'%(channel_dir, channeltx_name, channel_name)
    
    command = env + bin + param
    return local(command)

def update_anchor(bin_path, yaml_path, out_path, channel_name, msp_id, domain_name):

    create_anchor_tx(bin_path, yaml_path, out_path, channel_name, msp_id)

    channel_dir = out_path + channel_name

    order_tls_path = yaml_path +  "crypto-config/ordererOrganizations/ord1.%s/orderers/orderer0.ord1.%s/msp/tlscacerts/tlsca.ord1.%s-cert.pem"%(domain_name,domain_name,domain_name)
    order_address = "orderer0.ord1.%s:7050"%domain_name

    msp_path = yaml_path + "crypto-config/peerOrganizations/org1.%s/users/Admin@org1.%s/msp"%(domain_name,domain_name)
    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    env = env + ' CORE_PEER_LOCALMSPID=Org1MSP'
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s'%msp_path
    bin = bin_path + "peer"
    param = ' channel update -o %s -c %s -f %s/%sanchors.tx'%(order_address, channel_name, channel_dir, msp_id)
    tls = ' --tls --cafile %s'%order_tls_path

    command = env + bin + param + tls
    return local(command)

def create_anchor_tx(bin_path, yaml_path, out_path, channel_name, msp_id):
    channel_dir = out_path + channel_name
    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    param = ' -profile OrgsChannel -outputAnchorPeersUpdate %s/%sanchors.tx -channelID %s -asOrg %s'%(channel_dir, msp_id, channel_name, msp_id)

    bin = bin_path + "configtxgen"
    command = env + bin + param
    return local(command)

def join_channel(bin_path, yaml_path, out_path, channel_name, peer_address, peer_id, org_id, domain_name):
    channel_block = channel_name + '.block'
    tls_root_file = yaml_path + "crypto-config/peerOrganizations/org%s.%s/peers/peer%s.org%s.%s/tls/ca.crt"%(org_id,domain_name,peer_id,org_id,domain_name)
    msp_path = yaml_path + "crypto-config/peerOrganizations/org%s.%s/users/Admin@org%s.%s/msp"%(org_id,domain_name,org_id,domain_name)
    channel_dir = out_path + channel_name
    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    env = env + ' CORE_PEER_LOCALMSPID=Org%sMSP'%org_id
    env = env + ' CORE_PEER_TLS_ROOTCERT_FILE=%s'%tls_root_file
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s'%msp_path
    env = env + ' CORE_PEER_TLS_ENABLED=true'
    env = env + ' CORE_PEER_ADDRESS=%s'%peer_address
    bin = bin_path + "peer"
    param = ' channel join -b %s/%s'%(channel_dir, channel_block)

    command = env + bin + param
    local(command)
