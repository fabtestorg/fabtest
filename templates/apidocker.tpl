version: '2'

services:
  apiserver:
    container_name: apiserver{{.api_id}}
    image: factoring/apiserver
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
     - {{.api_id}}{{.api_id}}{{.api_id}}{{.api_id}}:5555
    extra_hosts:
       peer{{.peer_id}}.org{{.org_id}}.{{.peer_domain}}: {{.ip}}
       orderer{{.peer_id}}.ord{{.org_id}}.{{.peer_domain}}: {{.order_address}}


