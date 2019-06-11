Organizations:{{range $key,$value:= .ordList}}
    - &OrdererOrg{{$key}}
        Name: OrdererOrg{{$key}}
        ID: Orderer{{$key}}MSP
        MSPDir: crypto-config/ordererOrganizations/ord{{$key}}.{{$.domain}}/msp
        Policies:
            Readers:
                Type: Signature
                Rule: "OR('Orderer{{$key}}MSP.member')"
            Writers:
                Type: Signature
                Rule: "OR('Orderer{{$key}}MSP.member')"
            Admins:
                Type: Signature
                Rule: "OR('Orderer{{$key}}MSP.admin')"{{end}}
    {{range $key,$value:= .orgList}}
    - &Org{{$key}}
        Name: Org{{$key}}MSP
        ID: Org{{$key}}MSP
        MSPDir: crypto-config/peerOrganizations/org{{$key}}.{{$.domain}}/msp
        Policies:
             Readers:
                 Type: Signature
                 Rule: "OR('Org{{$key}}MSP.admin', 'Org{{$key}}MSP.peer', 'Org{{$key}}MSP.client')"
             Writers:
                 Type: Signature
                 Rule: "OR('Org{{$key}}MSP.admin', 'Org{{$key}}MSP.client')"
             Admins:
                 Type: Signature
                 Rule: "OR('Org{{$key}}MSP.admin')"
        AnchorPeers:{{range $index,$peer:= $.peers}} {{if eq $peer.orgId $key}} {{if eq $peer.id "0"}}
            - Host: peer0.org{{$peer.orgId}}.{{$.domain}}
              Port: {{$peer.configTxPort}}{{end}}{{end}}{{end}}{{end}}

Capabilities:
    Channel: &ChannelCapabilities
        V1_3: true
    Orderer: &OrdererCapabilities
        V1_1: true
    Application: &ApplicationCapabilities
        V1_3: true
        V1_2: false
        V1_1: false

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
    Addresses:
        - orderer0.ord1.example.com:7050
    BatchTimeout: {{.batchTime}}
    BatchSize:
        MaxMessageCount: {{.batchSize}}
        AbsoluteMaxBytes: 98 MB
        PreferredMaxBytes: {{.batchPreferred}}
    Kafka:
        Brokers:
            - 127.0.0.1:9092
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
    OrgsOrdererGenesis:
        <<: *ChannelDefaults
        Orderer:
            <<: *OrdererDefaults
            Organizations:{{range $key,$value:= .ordList}}
                - *OrdererOrg{{$key}}{{end}}
            Capabilities:
                <<: *OrdererCapabilities
        Consortiums:
            SampleConsortium:
                Organizations:{{range $key,$value:= .orgList}}
                    - *Org{{$key}}{{end}}
    OrgsChannel:
        Consortium: SampleConsortium
        <<: *ChannelDefaults
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $key,$value:= .orgList}}
                - *Org{{$key}}{{end}}
            Capabilities:
                <<: *ApplicationCapabilities


    SampleDevModeKafka:
        <<: *ChannelDefaults
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            OrdererType: kafka
            Kafka:
                Brokers:{{range $index,$kafka:= .kafkas}}
                - kafka{{$kafka.id}}:{{$kafka.configTxPort}}{{end}}
            Organizations:{{range $key,$value:= .ordList}}
            - *OrdererOrg{{$key}}{{end}}
            Capabilities:
                <<: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $key,$value:= .ordList}}
            - <<: *OrdererOrg{{$key}}{{end}}
        Consortiums:
            SampleConsortium:
                Organizations:{{range $key,$value:= .orgList}}
                - *Org{{$key}}{{end}}


    SampleMultiNodeEtcdRaft:
        <<: *ChannelDefaults
        Capabilities:
            <<: *ChannelCapabilities
        Orderer:
            <<: *OrdererDefaults
            OrdererType: etcdraft
            EtcdRaft:
                Consenters:{{range $index,$orderer:= .orderers}}
                - Host: orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}
                  Port: {{$orderer.configTxPort}}
                  ClientTLSCert: crypto-config/ordererOrganizations/ord{{$orderer.orgId}}.{{$.domain}}/orderers/orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}/tls/server.crt
                  ServerTLSCert: crypto-config/ordererOrganizations/ord{{$orderer.orgId}}.{{$.domain}}/orderers/orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}/tls/server.crt{{end}}
            Addresses:{{range $index,$orderer:= .orderers}}
                - orderer{{$orderer.id}}.ord{{$orderer.orgId}}.{{$.domain}}:{{$orderer.configTxPort}}{{end}}
            Organizations:{{range $key,$value:= .ordList}}
            - *OrdererOrg{{$key}}{{end}}
            Capabilities:
                <<: *OrdererCapabilities
        Application:
            <<: *ApplicationDefaults
            Organizations:{{range $key,$value:= .ordList}}
            - <<: *OrdererOrg{{$key}}{{end}}
        Consortiums:
            SampleConsortium:
                Organizations:{{range $key,$value:= .orgList}}
                - *Org{{$key}}{{end}}

