crypto:
  family: ecdsa
  algorithm: P256-SHA256
  hash: SHA2-256
orderers:
  orderer:
    host: orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}:7050
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}/tls/server.crt
peers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
eventPeers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
other:
    mspConfigPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId:          Org{{.org_id}}MSP
    channelId:           testchannel
    chaincodeName:       testfabric
    chaincodeVersion:    1.0
    mq_address:
      - "amqp://guest:guest@localhost:5672/"
    queue_name: "assetQueue"
policy:
    orgs: org{{.org_id}}
    rule: or
user:
    alias: zhengfu1

