#!/usr/bin/env bash

set -x
cp /opt/gopath/src/github.com/hyperledger/fabric/build/docker/bin/peer .
cp /opt/gopath/src/github.com/hyperledger/fabric/build/docker/bin/cryptogen .
cp /opt/gopath/src/github.com/hyperledger/fabric/build/docker/bin/configtxgen .
cp /opt/gopath/src/github.com/peersafe/gm-crypto/usr/bin/LICENSE .
