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
          localMspId: Org{{.org_id}}MSP
          tls:
              # Certificate location absolute path
              certificate: "/root/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/ca.crt"
              serverHostOverride: "peer{{.peer_id}}"

    orderer:
        {{if eq .peer_id "0"}}
        address: "{{.order0_address}}:7050"
        {{else if eq .peer_id "1"}}
        address: "{{.order1_address}}:7050"
        {{end}}
        tls:
             # Certificate location absolute path
             certificate: "/root/fabTestData/crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}/msp/tlscacerts/tlsca.ord{{.org_id}}.{{.peer_domain}}-cert.pem"
             serverHostOverride: "orderer{{.peer_id}}"
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
            file: /root/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls/ca.crt
        serverhostoverride: peer{{.peer_id}}
    BCCSP:
        Default: SW
        SW:
            Hash: SHA2
            Security: 256
            FileKeyStore:
                KeyStore:
    mspConfigPath: /root/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/users/Admin@org{{.org_id}}.{{.peer_domain}}/msp
    localMspId: Org{{.org_id}}MSP

chaincode:
    id:
        name: factor
        version: "1.0"
        chainID: testchannel
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


