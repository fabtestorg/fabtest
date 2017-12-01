version: '2'
services:
    zookeeper{{.zk_id}}:
        image: hyperledger/fabric-zookeeper
        restart: always
        environment:
          ZOO_MY_ID: {{.zk_id}}
          {{if eq .zk_id "1"}}
          ZOO_SERVERS: server.1=zookeeper1:2888:3888 server.2=zookeeper2:2888:3888 server.3=zk3.{{.kfk_domain}}:2888:3888 server.4=zk4.{{.kfk_domain}}:12888:13888  server.5=zk5.{{.kfk_domain}}:2888:3888
          {{else if eq .zk_id "3"}}
          ZOO_SERVERS: server.1=zk1.{{.kfk_domain}}:2888:3888 server.2=zk2.{{.kfk_domain}}:12888:13888  server.3=zookeeper3:2888:3888 server.4=zookeeper4:2888:3888 server.5=zk5.{{.kfk_domain}}:2888:3888
          {{else if eq .zk_id "5"}}
          ZOO_SERVERS: server.1=zk1.{{.kfk_domain}}:2888:3888 server.2=zk2.{{.kfk_domain}}:12888:13888  server.3=zk3.{{.kfk_domain}}:2888:3888 server.4=zk4.{{.kfk_domain}}:12888:13888 server.5=zookeeper5:2888:3888
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
          {{if eq .zk_id "1"}}
          zk3.{{.kfk_domain}}: {{.ip3}}
          zk4.{{.kfk_domain}}: {{.ip4}}
          zk5.{{.kfk_domain}}: {{.ip5}}
          {{else if eq .zk_id "3"}}
          zk1.{{.kfk_domain}}: {{.ip1}}
          zk2.{{.kfk_domain}}: {{.ip2}}
          zk5.{{.kfk_domain}}: {{.ip5}}
          {{else if eq .zk_id "5"}}
          zk1.{{.kfk_domain}}: {{.ip1}}
          zk2.{{.kfk_domain}}: {{.ip2}}
          zk3.{{.kfk_domain}}: {{.ip3}}
          zk4.{{.kfk_domain}}: {{.ip4}}
          {{end}}
        logging:
          driver: "json-file"
          options:
            max-size: "50m"
            max-file: "10"

    {{if ne .zk_id "5"}}
    zookeeper{{.zk_2_id}}:
        image: hyperledger/fabric-zookeeper
        restart: always
        environment:
          ZOO_MY_ID: {{.zk_2_id}}
          {{if eq .zk_id "1"}}
          ZOO_SERVERS: server.1=zookeeper1:2888:3888 server.2=zookeeper2:2888:3888 server.3=zk3.{{.kfk_domain}}:2888:3888 server.4=zk4.{{.kfk_domain}}:12888:13888  server.5=zk5.{{.kfk_domain}}:2888:3888
          {{else if eq .zk_id "3"}}
          ZOO_SERVERS: server.1=zk1.{{.kfk_domain}}:2888:3888 server.2=zk2.{{.kfk_domain}}:12888:13888  server.3=zookeeper3:2888:3888 server.4=zookeeper4:2888:3888 server.5=zk5.{{.kfk_domain}}:2888:3888
          {{end}}
        ports:
          - 12181:2181
          - 12888:2888
          - 13888:3888
        volumes:
          - /etc/localtime:/etc/localtime
          - /PATH1:/data
          - /PATH2:/datalog
        extra_hosts:
          {{if eq .zk_id "1"}}
          zk3.{{.kfk_domain}}: {{.ip3}}
          zk4.{{.kfk_domain}}: {{.ip4}}
          zk5.{{.kfk_domain}}: {{.ip5}}
          {{else if eq .zk_id "3"}}
          zk1.{{.kfk_domain}}: {{.ip1}}
          zk2.{{.kfk_domain}}: {{.ip2}}
          zk5.{{.kfk_domain}}: {{.ip5}}
          {{end}}
        logging:
          driver: "json-file"
          options:
            max-size: "50m"
            max-file: "10"
    {{end}}