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

user:
    alias: zhengfu998
###############################################################################
#
#    other section
#
###############################################################################
other:
    mq_address:
      - "amqp://testpoc:123456@10.10.255.71:5672/"
      - "amqp://testpoc:123456@10.10.255.72:5672/"
    queue_name: "fftQueue"
    system_queue_name: "sys_fftQueue"
