#!/bin/bash
echo orderer0
ssh root@orderer10 -o ConnectTimeout=1 "docker logs orderer0.ord1.example.com 2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer0.ord1.example.com --tail 1 2>&1"
echo orderer1
ssh root@orderer11 -o ConnectTimeout=1 "docker logs orderer1.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer1.ord1.example.com --tail 1 2>&1 "
echo orderer2
ssh root@orderer12 -o ConnectTimeout=1 "docker logs orderer2.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer2.ord1.example.com --tail 1 2>&1"
echo orderer3
ssh root@orderer13 -o ConnectTimeout=1 "docker logs orderer3.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer3.ord1.example.com --tail 1 2>&1"
echo orderer4
ssh root@orderer14 -o ConnectTimeout=1 "docker logs orderer4.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer4.ord1.example.com --tail 1 2>&1"
echo orderer5
ssh root@orderer15 -o ConnectTimeout=1 "docker logs orderer5.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer5.ord1.example.com --tail 1 2>&1"
echo orderer6
ssh root@orderer16 -o ConnectTimeout=1 "docker logs orderer6.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer6.ord1.example.com --tail 1 2>&1"
echo orderer7
ssh root@orderer17 -o ConnectTimeout=1 "docker logs orderer7.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer7.ord1.example.com --tail 1 2>&1"
exit
echo orderer8
ssh root@orderer18 -o ConnectTimeout=1 "docker logs orderer8.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer8.ord1.example.com --tail 1 2>&1"
echo orderer9
ssh root@orderer19 -o ConnectTimeout=1 "docker logs orderer9.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer9.ord1.example.com --tail 1 2>&1"
echo orderer10
ssh root@orderer110 -o ConnectTimeout=1 "docker logs orderer10.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer10.ord1.example.com --tail 1 2>&1"
echo orderer11
ssh root@orderer111 -o ConnectTimeout=1 "docker logs orderer11.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer11.ord1.example.com --tail 1 2>&1"
echo orderer12
ssh root@orderer112 -o ConnectTimeout=1 "docker logs orderer12.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer12.ord1.example.com --tail 1 2>&1"
echo orderer13
ssh root@orderer113 -o ConnectTimeout=1 "docker logs orderer13.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer13.ord1.example.com --tail 1 2>&1"
echo orderer14
ssh root@orderer114 -o ConnectTimeout=1 "docker logs orderer14.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer14.ord1.example.com --tail 1 2>&1"
echo orderer15
ssh root@orderer115 -o ConnectTimeout=1 "docker logs orderer15.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer15.ord1.example.com --tail 1 2>&1"
echo orderer16
ssh root@orderer116 -o ConnectTimeout=1 "docker logs orderer16.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer16.ord1.example.com --tail 1 2>&1"
echo orderer17
ssh root@orderer117 -o ConnectTimeout=1 "docker logs orderer17.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer17.ord1.example.com --tail 1 2>&1"
echo orderer18
ssh root@orderer118 -o ConnectTimeout=1 "docker logs orderer18.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer18.ord1.example.com --tail 1 2>&1"
echo orderer19
ssh root@orderer119 -o ConnectTimeout=1 "docker logs orderer19.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer19.ord1.example.com --tail 1 2>&1"
echo orderer20
ssh root@orderer120 -o ConnectTimeout=1 "docker logs orderer20.ord1.example.com  2>&1 | grep -e loopNum -e BroadcastClientSend;docker logs orderer20.ord1.example.com --tail 1 2>&1"
