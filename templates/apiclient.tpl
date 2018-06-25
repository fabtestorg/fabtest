crypto:
  family: ecdsa
  algorithm: P256-SHA256
  hash: SHA2-256
orderers:
  orderer:
    host: orderer0.ord1.{{.peer_domain}}:7050
    useTLS: true
    tlsPath: ./crypto-config/ordererOrganizations/ord1.{{.peer_domain}}/orderers/orderer0.ord1.{{.peer_domain}}/tls/server.crt
peers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
eventPeers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
other:
    mspConfigPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId:          Org{{.org_id}}MSP
    channelId:           testchannel
    chaincodeName:       testfabric
    chaincodeVersion:    1.0
policy:
    orgs: org{{.org_id}}
    rule: or

apiserver:
    #The alias should not be changed manually, unless you know what it means.
    #And the file can not have another alias in other fileds.
    alias: zhengfu1
    listenport: 5555
    probe_order: "orderer0.ord1.{{.peer_domain}} 7050"

