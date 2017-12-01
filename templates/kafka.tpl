version: '2'

services:
  kafka{{.kfk_id}}:
    image: hyperledger/fabric-kafka
    restart: always
    environment:
      - KAFKA_BROKER_ID={{.kfk_id}}
      - KAFKA_ZOOKEEPER_CONNECT=zk1.{{.kfk_domain}}:2181,zk2.{{.kfk_domain}}:12181,zk3.{{.kfk_domain}}:2181,zk4.{{.kfk_domain}}:12181,zk5.{{.kfk_domain}}:2181
      - KAFKA_DEFAULT_REPLICATION_FACTOR=2
      - KAFKA_UNCLEAN_LEADER_ELECTION_ENABLE=false
      - KAFKA_LOG_RETENTION_HOURS=876000
     #enable TLS
      - KAFKA_LISTENERS=PLAINTEXT://:8092,SSL://:9092
      - KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://{{.adver_address}},SSL://{{.blockchain_address}}
      - KAFKA_SSL_CLIENT_AUTH=required
      - KAFKA_SSL_KEYSTORE_LOCATION=/opt/kafka/ssl/server.keystore.jks
      - KAFKA_SSL_TRUSTSTORE_LOCATION=/opt/kafka/ssl/server.truststore.jks
      - KAFKA_SSL_KEY_PASSWORD=test1234
      - KAFKA_SSL_KEYSTORE_PASSWORD=test1234
      - KAFKA_SSL_TRUSTSTORE_PASSWORD=test1234
      - KAFKA_SSL_KEYSTORE_TYPE=JKS
      - KAFKA_SSL_TRUSTSTORE_TYPE=JKS
      - KAFKA_SSL_ENABLED_PROTOCOLS=TLSv1.2,TLSv1.1,TLSv1
      - KAFKA_SSL_INTER_BROKER_PROTOCOL=SSL
    ports:
      - 9092:9092
      - 8092:8092
    volumes:
      - ./kafkaTLSserver:/opt/kafka/ssl
      - /etc/localtime:/etc/localtime
      - ./kafka_log:/tmp/kafka-logs
    logging:
      driver: "json-file"
      options:
        max-size: "50m"
        max-file: "10"
    extra_hosts:
      zk1.{{.kfk_domain}}: {{.zk_ip1}}
      zk2.{{.kfk_domain}}: {{.zk_ip2}}
      zk3.{{.kfk_domain}}: {{.zk_ip3}}
      zk4.{{.kfk_domain}}: {{.zk_ip4}}
      zk5.{{.kfk_domain}}: {{.zk_ip5}}
      broker.finblockchain.cn: {{.broker_ip}}
      blockchain.hoperun.com: {{.broker_ip}}
