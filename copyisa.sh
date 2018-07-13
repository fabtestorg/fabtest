#!/bin/bash
set -x
ssh-copy-id -i ~/.ssh/id_rsa.pub root@zk0
ssh-copy-id -i ~/.ssh/id_rsa.pub root@zk1
ssh-copy-id -i ~/.ssh/id_rsa.pub root@zk2
ssh-copy-id -i ~/.ssh/id_rsa.pub root@kafka0
ssh-copy-id -i ~/.ssh/id_rsa.pub root@kafka1
ssh-copy-id -i ~/.ssh/id_rsa.pub root@kafka2
ssh-copy-id -i ~/.ssh/id_rsa.pub root@kafka3
ssh-copy-id -i ~/.ssh/id_rsa.pub root@orderer10
ssh-copy-id -i ~/.ssh/id_rsa.pub root@orderer11
ssh-copy-id -i ~/.ssh/id_rsa.pub root@peer10
ssh-copy-id -i ~/.ssh/id_rsa.pub root@peer11
ssh-copy-id -i ~/.ssh/id_rsa.pub root@api10
ssh-copy-id -i ~/.ssh/id_rsa.pub root@api11
