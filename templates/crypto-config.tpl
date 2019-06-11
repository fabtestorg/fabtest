OrdererOrgs:{{range $key,$value:= .ordList}}
  - Name: ord{{$key}}
    Domain: ord{{$key}}.{{$.domain}}
    Template:
      Count: {{$value}}{{end}}

PeerOrgs:{{range $key,$value:= .orgList}}
  - Name: org{{$key}}
    Domain: org{{$key}}.{{$.domain}}
    EnableNodeOUs: true
    Template:
      Count: {{$value}}
    Users:
      Count: 1{{end}}
