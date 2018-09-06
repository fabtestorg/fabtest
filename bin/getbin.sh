#!/usr/bin/env bash

set -x

cd /opt/gopath/src/github.com/hyperledger/fabric/common/tools/cryptogen
go build
cd -

cd /opt/gopath/src/github.com/hyperledger/fabric/common/tools/configtxgen
go build
cd -

#cd /opt/gopath/src/github.com/hyperledger/fabric/common/tools/configtxlator
#go build
#cd -

cd /opt/gopath/src/github.com/hyperledger/fabric/peer
go build
cd -

mv /opt/gopath/src/github.com/hyperledger/fabric/peer/peer .
mv /opt/gopath/src/github.com/hyperledger/fabric/common/tools/cryptogen/cryptogen . 
mv /opt/gopath/src/github.com/hyperledger/fabric/common/tools/configtxgen/configtxgen . 
#mv /opt/gopath/src/github.com/hyperledger/fabric/common/tools/configtxlator/configtxlator . 
cp /opt/gopath/src/github.com/hyperledger/fabric/sampleconfig/core.yaml .
