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
    tls:
      enabled: true
    peers:
        # peer0
        - address: "{{.ip}}:7051"
          eventHost: "{{.ip}}"
          eventPort: 7053
          primary: true
          localMspId: Org1MSP
          tls:
              # Certificate location absolute path
              certificate: "./crypto-config/peerOrganizations/org1.{{.peer_domain}}/peers/peer0.org1.{{.peer_domain}}/tls/ca.crt"
              serverHostOverride: "peer0"

    orderer:
        - address: "{{.order_address}}:7050"
          tls:
            # Certificate location absolute path
            certificate: "./crypto-config/ordererOrganizations/ord1.{{.peer_domain}}/orderers/orderer0.ord1.{{.peer_domain}}/msp/tlscacerts/tlsca.ord1.{{.peer_domain}}-cert.pem"
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
            file: ./crypto-config/peerOrganizations/org1.{{.peer_domain}}/peers/peer0.org1.{{.peer_domain}}/tls/ca.crt
        serverhostoverride: peer0
    BCCSP:
        Default: SW
        SW:
            Hash: SHA2
            Security: 256
            FileKeyStore:
                KeyStore:
    mspConfigPath: ./crypto-config/peerOrganizations/org1.{{.peer_domain}}/users/Admin@org1.{{.peer_domain}}/msp
    localMspId: Org1MSP
###############################################################################
#
#    Chaincode section
#
###############################################################################
chaincode:
    id:
        name: factor{{.peer_id}}
        version: "1.0"
        chainID: mychannel{{.peer_id}}

user:
    alias: zhengfu998

apiserver:
    listenport: 5555
    probe_order: "{{.order_address}} 7050"



