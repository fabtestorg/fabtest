Profiles:
    OrgsOrdererGenesis:
        Orderer:
            <<: *OrdererDefaults
            Organizations:{{range $index,$value:= .orgs}}
                - *OrdererOrg{{$value.org_id}}{{end}}
        Consortiums:
            SampleConsortium:
                Organizations:{{range $index,$value:= .orgs}}
                    - *Org{{$value.org_id}}{{end}}
    OrgsChannel:
        Consortium: SampleConsortium
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $index,$value:= .orgs}}
                - *Org{{$value.org_id}}{{end}}
Organizations:{{range $index,$value:= .orgs}}
    - &OrdererOrg{{$value.org_id}}
        Name: OrdererOrg{{$value.org_id}}
        ID: Orderer{{$value.org_id}}MSP
        MSPDir: crypto-config/ordererOrganizations/ord{{$value.org_id}}.finblockchain.cn/msp{{end}}
    {{range $index,$value:= .orgs}}
    - &Org{{$value.org_id}}
        Name: Org{{$value.org_id}}MSP
        ID: Org{{$value.org_id}}MSP
        MSPDir: crypto-config/peerOrganizations/org{{$value.org_id}}.finblockchain.cn/msp
        AnchorPeers:
            - Host: peer0.org{{$value.org_id}}.finblockchain.cn
              Port: 7051{{end}}
Orderer: &OrdererDefaults
    OrdererType: kafka
    Addresses:{{range $index,$value:= .orgs}}
        - orderer0.ord{{$value.org_id}}.finblockchain.cn:7050
        - orderer1.ord{{$value.org_id}}.finblockchain.cn:7050{{end}}
    BatchTimeout: 1s
    BatchSize:
        MaxMessageCount: 100
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: 512 KB
    Kafka:
        Brokers:{{range $index,$value:= .kafkas}}
            - {{$value}}{{end}}
    Organizations:
Application: &ApplicationDefaults
    Organizations:

