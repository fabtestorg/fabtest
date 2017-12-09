OrdererOrgs:{{range $index,$value:= .orgs}}
  - Name: ord{{$value}}
    Domain: ord{{$value}}.{{$.peer_domain}}
    Template:
      Count: 2{{end}}

PeerOrgs:{{range $index,$value:= .orgs}}
  - Name: org{{$value}}
    Domain: org{{$value}}.{{$.peer_domain}}
    Template:
      Count: 2
    Users:
      Count: 1{{end}}

      