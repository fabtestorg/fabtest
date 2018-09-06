crypto:
  family: ecdsa
  algorithm: P256-SHA256
  hash: SHA2-256
orderers:
  orderer:
    host: orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}:7050
    useTLS: true
    tlsPath: ./crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}/tls/server.crt
peers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgName: org{{.org_id}}
    useTLS: true
    tlsPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
eventPeers:
  peer:
    host: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
    orgName: org{{.org_id}}
    useTLS: true
    tlsPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
channel:
    mspConfigPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId:          Org{{.org_id}}MSP
    channelId:           testchannel
    chaincodeName:       testfabric
    chaincodeVersion:    1.0
    chaincodePolicy:
      orgs:
        - org{{.org_id}}
      rule: or
mq:
    mqAddress:
      - "amqp://guest:guest@localhost:5672/"
    queueName: "assetQueue"
    system_queue_name: "factoring_system"
log:
    logLevel: DEBUG
    logModelName: apiserver
user:
    alias: zhengfu1
apiserver:
    listenport: 5555
    probe_order: "orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}} 7050"

