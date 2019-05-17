version: '2'

services:
  orderer{{.order_id}}.ord{{.org_id}}.{{.peer_domain}}:
    container_name: orderer{{.order_id}}.ord{{.org_id}}.{{.peer_domain}}
    image: hyperledger/fabric-orderer
    restart: always
    environment:
      - ORDERER_GENERAL_LOGLEVEL=INFO
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_GENESISMETHOD=file
      - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=Orderer{{.org_id}}MSP
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      # enabled TLS
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
      #- ORDERER_KAFKA_TOPIC_REPLICATIONFACTOR=1 TODO what is it?
      - ORDERER_KAFKA_VERBOSE=true
      - ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
      - ~/fabtest/channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ~/fabTestData/crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.order_id}}.ord{{.org_id}}.{{.peer_domain}}/msp:/var/hyperledger/orderer/msp
      - ~/fabTestData/crypto-config/ordererOrganizations/ord{{.org_id}}.{{.peer_domain}}/orderers/orderer{{.order_id}}.ord{{.org_id}}.{{.peer_domain}}/tls:/var/hyperledger/orderer/tls
      - /etc/localtime:/etc/localtime
    logging:
      driver: "json-file"
      options:
        max-size: "50m"
        max-file: "10"
    ports:
      - 7050:7050
    extra_hosts: 
       orderer0.ord1.example.com: {{.ip}}
       orderer1.ord1.example.com: {{.ip}}
       orderer2.ord1.example.com: {{.ip}}
       