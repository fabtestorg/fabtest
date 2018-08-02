logging:
    peer:       warning
    node:       warning
    network:    warning
    version:    warning
    protoutils: warning
    error:      warning
    msp:        critical

    format: '%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}'
###############################################################################
#
#    client section
#
###############################################################################
client:
    peers:
        # peer0
        - address: "{{.ip}}:7051"
          eventHost: "{{.ip}}"
          eventPort: 7053
          primary: true
          localMspId: Org1MSP
          tls:
              # Certificate location absolute path
              certificate: "/root/fabTestData/crypto-config/peerOrganizations/org1.{{.peer_domain}}/peers/peer0.org1.{{.peer_domain}}/tls/ca.crt"
              serverHostOverride: "peer0"

    orderer:
        address: "{{.order_address}}:7050"
        tls:
             # Certificate location absolute path
             certificate: "/root/fabTestData/crypto-config/ordererOrganizations/ord1.{{.peer_domain}}/orderers/orderer0.ord1.{{.peer_domain}}/msp/tlscacerts/tlsca.ord1.{{.peer_domain}}-cert.pem"
             serverHostOverride: "orderer0"
###############################################################################
#
#    Peer section
#
###############################################################################
peer:
    listenAddress: {{.ip}}:7053
    gomaxprocs: -1
    workers: 2
    tls:
        enabled: true
        rootcert:
            file: /root/fabTestData/crypto-config/peerOrganizations/org1.{{.peer_domain}}/peers/peer0.org1.{{.peer_domain}}/tls/ca.crt
        serverhostoverride: peer0
    BCCSP:
        Default: SW
        SW:
            Hash: SHA2
            Security: 256
            FileKeyStore:
                KeyStore:
    mspConfigPath: /root/fabTestData/crypto-config/peerOrganizations/org1.{{.peer_domain}}/users/Admin@org1.{{.peer_domain}}/msp
    localMspId: Org1MSP

chaincode:
    id:
        name: factor{{.api_id}}
        version: "1.0"
        chainID: mychannel{{.api_id}}
user:
    alias: zhengfu998
###############################################################################
#
#    other section
#
###############################################################################
other:
    mq_address:
      - "amqp://guest:guest@{{.apiip}}:5672/"
    queue_name: "fftQueue"
    system_queue_name: "sys_fftQueue"
    


