---
crypto:
  family: {{.crypto_family}}
  algorithm: {{.crypto_algorithm}}
  hash: {{.crypto_hash}}
orderers:
  orderer0:
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
    channelId:           mychannel{{.api_id}}
    chaincodeName:       factor{{.api_id}}
    chaincodeVersion:    1.0
    chaincodePolicy:
      orgs:
        - org{{.org_id}}
      rule: or
mq:
    mqEnable: false
    mqAddress:
      - "amqp://guest:guest@localhost:5672/"
    queueName: "TradeQueue"
log:
    logLevel: ERROR
    logModelName: eventserver
user:
    id: bankA
apiserver:
    authorization:
      user: "123456"
      root: "root"
    listenport: 5555
    probe_order: "orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}} 7050"
other:
    check_time: 2m
