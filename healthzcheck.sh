#!/usr/bin/env bash

#set -x
while [ true ]
    do
        echo "========健康检查所有区块链节点========="
        echo "--------orderer0 节点---------"
        curl -X GET http://192.168.0.21:5443/healthz -w "\n"
        echo "--------orderer1 节点---------"
        curl -X GET http://192.168.0.21:6443/healthz -w "\n"
        echo "--------orderer2 节点---------"
        curl -X GET http://192.168.0.21:7443/healthz -w "\n"
        echo "--------peer0 节点---------"
        curl -X GET http://192.168.0.21:8443/healthz -w "\n"
        echo "--------peer1 节点---------"
        curl -X GET http://192.168.0.21:9443/healthz -w "\n"
        sleep 2
    done

