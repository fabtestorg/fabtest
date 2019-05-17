OrdererOrgs:{{range $index,$value:= .orgs}}
  - Name: ord{{$value}}
    Domain: ord{{$value}}.{{$.peer_domain}}
    EnableNodeOUs: true
    CA:
        Country: US
        Province: California
        Locality: San Francisco
    Template:
      Count: 3{{end}}

PeerOrgs:{{range $index,$value:= .orgs}}
  - Name: org{{$value}}
    Domain: org{{$value}}.{{$.peer_domain}}
    EnableNodeOUs: true
    CA:
        Country: US
        Province: California
        Locality: San Francisco
    Template:
      Count: 2
    Users:
      Count: 1{{end}}

      