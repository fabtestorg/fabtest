crypto:
  family: ecdsa
  algorithm: P256-SHA256
  hash: SHA2-256
orderers:
  orderer:
    {{if eq .peer_id "0"}}
    host: {{.order0_address}}:7050
    {{else if eq .peer_id "1"}}
    host: {{.order1_address}}:7050
    {{end}}
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}/tls/server.crt
peers:
  peer:
    host: {{.ip}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
eventPeers:
  peer:
    host: {{.ip}}:7051
    orgname: org{{.org_id}}
    useTLS: true
    tlsPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/server.crt
other:
    mspConfigPath: /home/ubuntu/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId:          Org{{.org_id}}MSP
    channelId:           mychannel
    chaincodeName:       testfabric
    chaincodeVersion:    1.0
policy:
    orgs: org{{.org_id}}
    rule: or
mq:
    mq_address:
      - "amqp://guest:guest@localhost:5672/"
    queue_name: "assetQueue"
apiserver:
    #The alias should not be changed manually, unless you know what it means.
    #And the file can not have another alias in other fileds.
    alias: zhengfu1
    listenport: 5555
    {{if eq .peer_id "0"}}
    probe_order: "{{.order0_address}} 7050"
    {{else if eq .peer_id "1"}}
    probe_order: "{{.order1_address}} 7050"
    {{end}}

