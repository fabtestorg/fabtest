#!/usr/bin/env bash

echo "=====根据node.json配置修改节点连接方式======="
DELAY="$1"
: ${DELAY:="10"}
while [ true ]
    do
        echo "==========健康检查所有区块链节点间隔${DELAY}S==========="
        echo "-----------orderer0 节点-----------"
        curl -X GET http://192.168.0.21:5443/healthz -w "\n"
        echo "-----------orderer1 节点-----------"
        curl -X GET http://192.168.0.21:6443/healthz -w "\n"
        echo "-----------orderer2 节点-----------"
        curl -X GET http://192.168.0.21:7443/healthz -w "\n"
        echo "------------peer0 节点-------------"
        curl -X GET http://192.168.0.21:8443/healthz -w "\n"
        echo "------------peer1 节点-------------"
        curl -X GET http://192.168.0.21:9443/healthz -w "\n"
        sleep $DELAY
    done

