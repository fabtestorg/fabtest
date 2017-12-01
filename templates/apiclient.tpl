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
        - address: "peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051"
          eventHost: "peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}"
          eventPort: 7053
          primary: true
          localMspId: Org{{.org_id}}MSP
          tls:
              # Certificate location absolute path
              certificate: "./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/ca.crt"
              serverHostOverride: "peer{{.peer_id}}"

    orderer:
        - address: "orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}:7050"
          tls:
            # Certificate location absolute path
            certificate: "./crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.peer_id}}.{{.org_id}}.{{.peer_domain}}/msp/tlscacerts/tlsca.ord{{.org_id}}.{{.peer_domain}}-cert.pem"
            serverHostOverride: "orderer{{.peer_id}}"
###############################################################################
#
#    Peer section
#
###############################################################################
peer:
    listenAddress: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7053
    gomaxprocs: -1
    workers: 2
    tls:
        enabled: true
        rootcert:
            file: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/ca.crt
        serverhostoverride: peer{{.peer_id}}
    BCCSP:
        Default: SW
        SW:
            Hash: SHA2
            Security: 256
            FileKeyStore:
                KeyStore:
    mspConfigPath: ./crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId: Org{{.org_id}}MSP
###############################################################################
#
#    Chaincode section
#
###############################################################################
chaincode:
    id:
        name: ccname
        version: "1.0"
        chainID: channelname

user:
    alias: zhengfu998

apiserver:
    listenport: {{.apiport}}