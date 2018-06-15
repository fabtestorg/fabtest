version: '2'

services:
  apiserver:
    container_name: apiserver
    image: test_fabric/apiserver
    restart: always
    volumes:
    - ./schema:/opt/apiserver/schema
    - ./client_sdk.yaml:/opt/apiserver/client_sdk.yaml
    - ~/fabTestData/crypto-config/:/opt/apiserver/crypto-config
    - /etc/localtime:/etc/localtime
    working_dir: /opt/apiserver
    logging:
      driver: "json-file"
      options:
        max-size: "50m"
        max-file: "10"
    command: ./apiserver
    ports:
     - 5555:5555
    extra_hosts:
      {{if eq .peer_id "0"}}
       peer1.org{{.org_id}}.{{.peer_domain}}: {{.other_peeraddress}}
      {{else if eq .peer_id "1"}}
       peer0.org{{.org_id}}.{{.peer_domain}}: {{.other_peeraddress}}
      {{end}}
       orderer0.ord{{.org_id}}.{{.peer_domain}}: {{.order0_address}}
       orderer1.ord{{.org_id}}.{{.peer_domain}}: {{.order1_address}}