version: '2'

services:
  orderer{{.id}}.ord{{.orgId}}.{{.domain}}:
    container_name: orderer{{.id}}.ord{{.orgId}}.{{.domain}}
    image: {{.imagePre}}/fabric-orderer:{{.imageTag}}
    restart: always
    environment:
      - GODEBUG=netdns=go
      - FABRIC_LOGGING_SPEC={{.log}}
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_GENESISMETHOD=file
      - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=Orderer{{.orgId}}MSP
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      - ORDERER_OPERATIONS_LISTENADDRESS=0.0.0.0:9443
      # enabled TLS
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
      - ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
      - ~/fabtest/channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
      - ~/fabtest/crypto-config/ordererOrganizations/ord{{.orgId}}.{{.domain}}/orderers/orderer{{.id}}.ord{{.orgId}}.{{.domain}}/msp:/var/hyperledger/orderer/msp
      - ~/fabtest/crypto-config/ordererOrganizations/ord{{.orgId}}.{{.domain}}/orderers/orderer{{.id}}.ord{{.orgId}}.{{.domain}}/tls:/var/hyperledger/orderer/tls
      - /data/orderer{{.id}}.ord{{.orgId}}.{{.domain}}:/var/hyperledger/production
    logging:
      driver: "json-file"
      options:
        max-size: "100m"
        max-file: "10"
    ports:{{range $index,$value:= .ports}}
      - {{$value}}{{end}}
    extra_hosts:{{range $index,$orderer := .orderers}}{{if or (ne $orderer.id $.id) (ne $orderer.orgId $.orgId)}}
       - "orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}:{{$orderer.ip}}"{{end}}{{end}}


       