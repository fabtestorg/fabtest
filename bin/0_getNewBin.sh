#!/usr/bin/env bash

set -x

cd /opt/gopath/src/github.com/hyperledger/fabric/common/tools/cryptogen
go build --ldflags "-extldflags -static"
mv cryptogen /opt/gopath/src/github.com/peersafe/bcap/fabfile/bin

cd /opt/gopath/src/github.com/hyperledger/fabric/common/configtx/tool/configtxgen
go build --ldflags "-extldflags -static"
mv configtxgen /opt/gopath/src/github.com/peersafe/bcap/fabfile/bin

cd /opt/gopath/src/github.com/hyperledger/fabric/common/tools/configtxlator
go build --ldflags "-extldflags -static"
mv configtxlator /opt/gopath/src/github.com/peersafe/bcap/fabfile/bin

cd /opt/gopath/src/github.com/hyperledger/fabric/peer
go build --ldflags "-extldflags -static"
mv peer /opt/gopath/src/github.com/peersafe/bcap/fabfile/bin

cd /opt/gopath/src/github.com/hyperledger/fabric-ca/cmd/fabric-ca-client/
go build --ldflags "-extldflags -static"
mv fabric-ca-client /opt/gopath/src/github.com/peersafe/bcap/fabfile/bin
