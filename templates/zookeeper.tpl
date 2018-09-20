version: '2'
services:
    zookeeper{{.zk_id}}:
        image: peersafes/fabric-zookeeper:1.1.1-gm
        restart: always
        environment:
          ZOO_MY_ID: {{.zk_id}}
          {{if eq .zk_id "0"}}
          ZOO_SERVERS: server.0=zookeeper0:2888:3888 server.1=zk1.{{.kfk_domain}}:2888:3888 server.2=zk2.{{.kfk_domain}}:2888:3888
          {{else if eq .zk_id "1"}}
          ZOO_SERVERS: server.0=zk0.{{.kfk_domain}}:2888:3888 server.1=zookeeper1:2888:3888 server.2=zk2.{{.kfk_domain}}:2888:3888
          {{else if eq .zk_id "2"}}
          ZOO_SERVERS: server.0=zk0.{{.kfk_domain}}:2888:3888 server.1=zk1.{{.kfk_domain}}:2888:3888  server.2=zookeeper:2888:3888
          {{end}}
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