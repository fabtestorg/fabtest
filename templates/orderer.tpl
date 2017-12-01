version: '2'

services:
  orderer{{.order_id}}.org{{.org_id}}.{{.peer_domain}}:
    container_name: orderer{{.order_id}}.org{{.org_id}}.{{.peer_domain}}
    image: hyperledger/fabric-orderer
    restart: always
    environment:
      - ORDERER_GENERAL_LOGLEVEL=debug
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
      # enabled kafka client TLS
      - ORDERER_KAFKA_SERVER=kafkaserver
      - ORDERER_KAFKA_VERBOSE=true
      - ORDERER_KAFKA_TLS_ENABLED=true
      - ORDERER_KAFKA_TLS_PRIVATEKEY=/var/hyperledger/orderer/kafka/tls/client.key
      - ORDERER_KAFKA_TLS_CERTIFICATE=/var/hyperledger/orderer/kafka/tls/client.crt
      - ORDERER_KAFKA_TLS_ROOTCAS=[/var/hyperledger/orderer/kafka/tls/ca.crt]
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
      - ./channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ./orderer{{.order_id}}.org{{.org_id}}.{{.peer_domain}}/msp:/var/hyperledger/orderer/msp
      - ./orderer{{.order_id}}.org{{.org_id}}.{{.peer_domain}}/tls:/var/hyperledger/orderer/tls
      - ./kafkaTLSclient:/var/hyperledger/orderer/kafka/tls
      - /etc/localtime:/etc/localtime
      - ./orderer_data:/var/hyperledger/production
    logging:
      driver: "json-file"
      options:
        max-size: "50m"
        max-file: "10"
    ports:
      - 7001:7050
      - 7002:7050
      - 7050:7050
    extra_hosts:
      broker.finblockchain.cn: {{.broker_address}}
