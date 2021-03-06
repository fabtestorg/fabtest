version: '2'
services:
  peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:
    image: hyperledger/fabric-peer
    restart: always
    container_name: peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}
    environment:
      # base env
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=peer{{.peer_id}}_default
      - CORE_LOGGING_LEVEL=DEBUG
      - CORE_PEER_TLS_ENABLED=true
      - CORE_PEER_ENDORSER_ENABLED=true
      - CORE_PEER_GOSSIP_USELEADERELECTION=false
      - CORE_PEER_GOSSIP_ORGLEADER=true
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
      - CORE_PEER_GOSSIP_RECONNECTMAXPERIOD=300
      - CORE_PEER_GOSSIP_RECONNECTMINPERIOD=5
      - CORE_PEER_GOSSIP_RECONNECTMINPERIODATTEMPTTIME=10
      - CORE_PEER_GOSSIP_DEFMAXBLOCKDISTANCD=100
      {{if eq .peer_id "0"}}
      - CORE_PEER_GOSSIP_DEFAULTORDERERADDRESS=orderer0.ord{{.org_id}}.{{.peer_domain}}:7050
      {{else if eq .peer_id "1"}}
      - CORE_PEER_GOSSIP_DEFAULTORDERERADDRESS=orderer1.ord{{.org_id}}.{{.peer_domain}}:7050
      - CORE_PEER_GOSSIP_BOOTSTRAP=peer0.org{{.org_id}}.{{.peer_domain}}:7051
      {{end}}
      # improve env
      - CORE_PEER_ID=peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}
      - CORE_PEER_ADDRESS=peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
      - CORE_PEER_CHAINCODELISTENADDRESS=peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7052
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}:7051
      - CORE_PEER_LOCALMSPID=Org{{.org_id}}MSP
      {{if eq .usecouchdb "true"}}
      - CORE_LEDGER_STATE_STATEDATABASE=CouchDB
      - CORE_LEDGER_STATE_COUCHDBCONFIG_COUCHDBADDRESS=couchdb:5984
      {{end}}
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
    volumes:
        - /var/run/:/host/var/run/
        - ~/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/msp:/etc/hyperledger/fabric/msp
        - ~/fabTestData/crypto-config/peerOrganizations/org{{.org_id}}.{{.peer_domain}}/peers/peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}/tls:/etc/hyperledger/fabric/tls
        - /etc/localtime:/etc/localtime
        - ./peer_data:/var/hyperledger/production
    logging:
      driver: "json-file"
      options:
        max-size: "50m"
        max-file: "10"
    ports:
      - 7051:7051
      - 7052:7052
      - 7053:7053
    extra_hosts:
      {{if eq .peer_id "0"}}
       peer1.org{{.org_id}}.{{.peer_domain}}: {{.other_peeraddress}}
      {{else if eq .peer_id "1"}}
       peer0.org{{.org_id}}.{{.peer_domain}}: {{.other_peeraddress}}
      {{end}}
       orderer0.ord{{.org_id}}.{{.peer_domain}}: {{.order0_address}}
       orderer1.ord{{.org_id}}.{{.peer_domain}}: {{.order1_address}}
  {{if eq .usecouchdb "true"}}
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




