Profiles:
    OrgsOrdererGenesis:
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            Organizations:{{range $index,$value:= .orderorgs}}
                - *OrdererOrg{{$value}}{{end}}
            Capabilities:
                <<: *OrdererCapabilities
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
            Capabilities:
                <<: *ApplicationCapabilities
Organizations:{{range $index,$value:= .orderorgs}}
    - &OrdererOrg{{$value}}
        Name: OrdererOrg{{$value}}
        ID: Orderer{{$value}}MSP
        MSPDir: crypto-config/ordererOrganizations/ord{{$value}}.finblockchain.cn/msp{{end}}
    {{range $index,$value:= .orgs}}
    - &Org{{$value}}
        Name: Org{{$value}}MSP
        ID: Org{{$value}}MSP
        MSPDir: crypto-config/peerOrganizations/org{{$value}}.finblockchain.cn/msp
        AnchorPeers:
            - Host: peer0.org{{$value}}.finblockchain.cn
              Port: 7051{{end}}
Orderer: &OrdererDefaults
    OrdererType: solo
    Addresses:{{range $index,$value:= .orderorgs}}
        - orderer0.ord{{$value}}.finblockchain.cn:7050{{end}}
    BatchTimeout: {{.batchTime}}
    BatchSize:
        MaxMessageCount: {{.batchSize}}
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: {{.batchPreferred}}
    Organizations:
Application: &ApplicationDefaults
    Organizations:

Capabilities:
    Global: &ChannelCapabilities
        V1_1: true
    Orderer: &OrdererCapabilities
        V1_1: true
    Application: &ApplicationCapabilities
        V1_1: true

