
Organizations:{{range $index,$value:= .orgs}}
    - &OrdererOrg{{$value}}
        Name: OrdererOrg{{$value}}
        ID: Orderer{{$value}}MSP
        MSPDir: crypto-config/ordererOrganizations/ord{{$value}}.example.com/msp{{end}}
        Policies:
            Readers:
                Type: Signature
                Rule: "OR('Orderer{{$value}}MSP.member')"
            Writers:
                Type: Signature
                Rule: "OR('Orderer{{$value}}MSP.member')"
            Admins:
                Type: Signature
                Rule: "OR('Orderer{{$value}}MSP.admin')"
    {{range $index,$value:= .orgs}}
    - &Org{{$value}}
        Name: Org{{$value}}MSP
        ID: Org{{$value}}MSP
        MSPDir: crypto-config/peerOrganizations/org{{$value}}.example.com/msp
        Policies:
             Readers:
                 Type: Signature
                 Rule: "OR('Org{{$value}}MSP.admin', 'Org{{$value}}MSP.peer', 'Org{{$value}}MSP.client')"
             Writers:
                 Type: Signature
                 Rule: "OR('Org{{$value}}MSP.admin', 'Org{{$value}}MSP.client')"
             Admins:
                 Type: Signature
                 Rule: "OR('Org{{$value}}MSP.admin')"
        AnchorPeers:
            - Host: peer0.org{{$value}}.example.com
              Port: 7051{{end}}
Capabilities:
    Channel: &ChannelCapabilities
        V1_3: true
    Orderer: &OrdererCapabilities
        V1_1: true
    Application: &ApplicationCapabilities
        V1_3: true

Application: &ApplicationDefaults
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
    Capabilities:
        <<: *ApplicationCapabilities
Orderer: &OrdererDefaults
    OrdererType: solo
    Addresses:{{range $index,$value:= .orgs}}
        - orderer0.ord{{$value}}.example.com:7050{{end}}
    BatchTimeout: {{.batchTime}}
    BatchSize:
        MaxMessageCount: {{.batchSize}}
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: {{.batchPreferred}}
    Kafka:
        Brokers:{{range $index,$value:= .kafkas}}
            - {{$value}}:9092{{end}}
    Organizations:
    Policies:
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"


Channel: &ChannelDefaults
    Policies:
        # Who may invoke the 'Deliver' API
        Readers:
            Type: ImplicitMeta
            Rule: "ANY Readers"
        # Who may invoke the 'Broadcast' API
        Writers:
            Type: ImplicitMeta
            Rule: "ANY Writers"
        # By default, who may modify elements at this config level
        Admins:
            Type: ImplicitMeta
            Rule: "MAJORITY Admins"
    Capabilities:
        <<: *ChannelCapabilities

Profiles:
     TwoOrgsOrdererGenesis:
        <<: *ChannelDefaults
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            OrdererType: etcdraft
            EtcdRaft:
                Consenters:{{range $index,$value:= .ords}}
                - Host: orderer{{$value}}.ord1.example.com
                  Port: 7050
                  ClientTLSCert: crypto-config/ordererOrganizations/ord1.example.com/orderers/orderer{{$value}}.ord1.example.com/tls/server.crt
                  ServerTLSCert: crypto-config/ordererOrganizations/ord1.example.com/orderers/orderer{{$value}}.ord1.example.com/tls/server.crt{{end}}
            Addresses:
            {{range $index,$value:= .ords}}
                - orderer{{$value}}.ord1.example.com:7050
            {{end}}
            Organizations:
                - *OrdererOrg1
            Capabilities:
                <<: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:
            - <<: *OrdererOrg1
        Consortiums:
            SampleConsortium:
                Organizations:{{range $index,$value:= .orgs}}
                    - *Org{{$value}}{{end}}
    TwoOrgsChannel:
        Consortium: SampleConsortium
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $index,$value:= .orgs}}
                - *Org{{$value}}{{end}}
            Capabilities:
                <<: *ApplicationCapabilities
