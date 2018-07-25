Profiles:
    OrgsOrdererGenesis:
        Orderer:
            <<: *OrdererDefaults
            Organizations:{{range $index,$value:= .orgs}}
                - *OrdererOrg{{$value}}{{end}}
        Consortiums:
            SampleConsortium:
                Organizations:{{range $index,$value:= .orgs}}
                    - *Org{{$value}}{{end}}
    OrgsChannel:
        Consortium: SampleConsortium
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $index,$value:= .orgs}}
                - *Org{{$value}}{{end}}
Organizations:{{range $index,$value:= .orgs}}
    - &OrdererOrg{{$value}}
        Name: OrdererOrg{{$value}}
        ID: Orderer{{$value}}MSP
        MSPDir: crypto-config/ordererOrganizations/ord{{$value}}.example.com/msp{{end}}
    {{range $index,$value:= .orgs}}
    - &Org{{$value}}
        Name: Org{{$value}}MSP
        ID: Org{{$value}}MSP
        MSPDir: crypto-config/peerOrganizations/org{{$value}}.example.com/msp
        AnchorPeers:
            - Host: peer0.org{{$value}}.example.com
              Port: 7051{{end}}
Orderer: &OrdererDefaults
    OrdererType: kafka
    Addresses:{{range $index,$value:= .orgs}}
        - orderer0.ord{{$value}}.example.com:7050
        - orderer1.ord{{$value}}.example.com:7050{{end}}
    BatchTimeout: {{.batchTime}}
    BatchSize:
        MaxMessageCount: {{.batchSize}}
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: {{.batchPreferred}}
    Kafka:
        Brokers:{{range $index,$value:= .kafkas}}
            - {{$value}}:9092{{end}}
    Organizations:
Application: &ApplicationDefaults
    Organizations:

