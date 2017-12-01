OrdererOrgs:{{range $index,$value:= .orgs}}
  - Name: ord{{.org_id}}
    Domain: ord{{.org_id}}.{{$.domain_name}}
    Template:
      Count: 2{{end}}

PeerOrgs:{{range $index,$value:= .orgs}}
  - Name: org{{.org_id}}
    Domain: org{{.org_id}}.{{$.domain_name}}
    Template:
      Count: 2
    Users:
      Count: 1{{end}}

      