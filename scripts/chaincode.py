#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
from fabric.api import local

reload(sys)
sys.setdefaultencoding('utf8')


def install_chaincode(bin_path, yaml_path, peer_address, peer_id, org_id, domain_name, ccoutpath):
    tls_root_file = yaml_path + "crypto-config/peerOrganizations/org%s.%s/peers/peer%s.org%s.%s/tls/ca.crt"%(org_id,domain_name,peer_id,org_id,domain_name)
    msp_path = yaml_path + "crypto-config/peerOrganizations/org%s.%s/users/Admin@org%s.%s/msp"%(org_id,domain_name,org_id,domain_name)
    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    env = env + ' CORE_PEER_LOCALMSPID=Org%sMSP'%org_id
    env = env + ' CORE_PEER_TLS_ROOTCERT_FILE=%s'%tls_root_file
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s'%msp_path
    env = env + ' CORE_PEER_TLS_ENABLED=true'
    env = env + ' CORE_PEER_ADDRESS=%s '%peer_address
    bin = bin_path + "peer"

    param = ' chaincode install %s'%ccoutpath

    command = env + bin + param
    local(command)


def instantiate_chaincode(bin_path, yaml_path, peer_address, peer_id, org_id, domain_name, channel_name ,ccname , init_param, policy):
    tls_root_file = yaml_path + "crypto-config/peerOrganizations/org%s.%s/peers/peer%s.org%s.%s/tls/ca.crt"%(org_id,domain_name,peer_id,org_id,domain_name)
    msp_path = yaml_path + "crypto-config/peerOrganizations/org%s.%s/users/Admin@org%s.%s/msp"%(org_id,domain_name,org_id,domain_name)
    order_tls_path = yaml_path +  "crypto-config/ordererOrganizations/ord1.%s/orderers/orderer0.ord1.%s/msp/tlscacerts/tlsca.ord1.%s-cert.pem"%(domain_name,domain_name,domain_name)
    order_address = "orderer0.ord1.%s:7050"%domain_name

    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    env = env + ' CORE_PEER_LOCALMSPID=Org%sMSP'%org_id
    env = env + ' CORE_PEER_TLS_ROOTCERT_FILE=%s'%tls_root_file
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s'%msp_path
    env = env + ' CORE_PEER_TLS_ENABLED=true'
    env = env + ' CORE_PEER_ADDRESS=%s '%peer_address
    bin = bin_path + "peer"
    init_param = '\'{"Args":["init","LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFTFVlRWJrYWtxOVVNc3JpeldJcTdDVGM0NGpYTAo1TFpYVVA1aFNxMWFnUENaNGw4a1pSekVoRjFPRE04aTdyU1MrMEFvMnBQc1hDKzh3MkR6bVBXYUpnPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="]}\''
    # --collections-config ${COLLECTION_DEFINE}
    col_cfg = yaml_path + "ccout/collection_define.json"
    param = ' chaincode instantiate -o %s -C %s -n %s -v %s -c %s -P %s --collections-config %s'%(order_address, channel_name, ccname, "1.0", init_param, policy,col_cfg)
    tls = ' --tls --cafile %s'%order_tls_path

    command = env + bin + param + tls
    local(command)

def test_query_tx(bin_path, yaml_path, peer_address, peer_id, org_id, domain_name, channel_name, ccname, tx_args):
    tls_root_file = yaml_path + "crypto-config/peerOrganizations/org%s.%s/peers/peer%s.org%s.%s/tls/ca.crt"%(org_id,domain_name,peer_id,org_id,domain_name)
    msp_path = yaml_path + "crypto-config/peerOrganizations/org%s.%s/users/Admin@org%s.%s/msp"%(org_id,domain_name,org_id,domain_name)
    env = ' FABRIC_CFG_PATH=%s '%yaml_path
    env = env + ' CORE_PEER_LOCALMSPID=Org%sMSP'%org_id
    env = env + ' CORE_PEER_TLS_ROOTCERT_FILE=%s'%tls_root_file
    env = env + ' CORE_PEER_MSPCONFIGPATH=%s'%msp_path
    env = env + ' CORE_PEER_TLS_ENABLED=true'
    env = env + ' CORE_PEER_ADDRESS=%s '%peer_address
    bin = bin_path + "peer"
    param = '  chaincode query -C %s -n %s -c %s '%(channel_name, ccname,tx_args)
    command = env + bin + param
    local(command)
