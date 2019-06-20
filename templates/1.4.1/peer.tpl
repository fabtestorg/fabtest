version: '2'

services:
  peer{{.id}}.org{{.orgId}}.{{.domain}}:
    image: peersafes/fabric-peer:{{.imageTag}}
    restart: always
    container_name: peer{{.id}}.org{{.orgId}}.{{.domain}}
    environment:
      # base env
      - GODEBUG=netdns=go
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE={{.defaultNetwork}}_default
      - FABRIC_LOGGING_SPEC={{.log}}
      - CORE_CHAINCODE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
      - CORE_CHAINCODE_BUILDER=peersafes/fabric-ccenv:{{.imageTag}}
      - CORE_CHAINCODE_GOLANG_RUNTIME=peersafes/fabric-baseos:{{.imageTag}}
      # improve env
      - CORE_PEER_ID=peer{{.id}}.org{{.orgId}}.{{.domain}}
      - CORE_PEER_ADDRESS=peer{{.id}}.org{{.orgId}}.{{.domain}}:7051
      - CORE_PEER_CHAINCODELISTENADDRESS=peer{{.id}}.org{{.orgId}}.{{.domain}}:7052
      - CORE_PEER_GOSSIP_BOOTSTRAP=peer{{.id}}.org{{.orgId}}.{{.domain}}:7051
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer{{.id}}.org{{.orgId}}.{{.domain}}:7051
      - CORE_OPERATIONS_LISTENADDRESS=0.0.0.0:9443
      - CORE_PEER_LOCALMSPID=Org{{.orgId}}MSP
      {{if eq .useCouchdb "true"}}
      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb:5984
      {{end}}
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
        - /var/run/:/host/var/run/
        - ~/fabtest/crypto-config/peerOrganizations/org{{.orgId}}.{{.domain}}/peers/peer{{.id}}.org{{.orgId}}.{{.domain}}/msp:/etc/hyperledger/fabric/msp
        - ~/fabtest/crypto-config/peerOrganizations/org{{.orgId}}.{{.domain}}/peers/peer{{.id}}.org{{.orgId}}.{{.domain}}/tls:/etc/hyperledger/fabric/tls
        - /data/peer{{.id}}.org{{.orgId}}.{{.domain}}:/var/hyperledger/production
    logging:
      driver: "json-file"
      options:
        max-size: "100m"
        max-file: "10"
    ports:{{range $index,$value:= .ports}}
      - {{$value}}{{end}}
    extra_hosts:{{range $index,$orderer:= .orderers}}
      - "orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}:{{$orderer.ip}}"{{end}}
      {{range $index,$peer:= .peers}}
      - "peer{{$peer.id}}.org{{$peer.orgId}}.{{$.domain}}:{{$peer.ip}}"{{end}}
  {{if eq .useCouchdb "true"}}
    depends_on:
      - couchdb
  couchdb:
    container_name: couchdb
    image: hyperledger/fabric-couchdb
    ports:
       - "5984:5984"
    volumes:
       - ./couchdb:/opt/couchdb/data
   {{end}}




