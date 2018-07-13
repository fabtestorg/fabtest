version: '2'
services:
    zookeeper{{.zk_id}}:
        image: hyperledger/fabric-zookeeper
        restart: always
        environment:
          ZOO_MY_ID: {{.zk_id}}
          ZOO_SERVERS: server.1=zk0.{{.kfk_domain}}:2888:3888 server.2=zk1.{{.kfk_domain}}:2888:3888 server.3=zk2.{{.kfk_domain}}:2888:3888
        ports:
          - 2181:2181
          - 2888:2888
          - 3888:3888
        volumes:
          - /etc/localtime:/etc/localtime
          - /PATH1:/data
          - /PATH2:/datalog
        extra_hosts:
          zk0.{{.kfk_domain}}: {{.ip0}}
          zk1.{{.kfk_domain}}: {{.ip1}}
          zk2.{{.kfk_domain}}: {{.ip2}}
        logging:
          driver: "json-file"
          options:
            max-size: "50m"
            max-file: "10"