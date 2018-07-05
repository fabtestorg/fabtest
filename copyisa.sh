#!/bin/bash
set -x
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@orderer0
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@peer10
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@peer20
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@api10
ssh-copy-id -i ~/.ssh/id_rsa.pub ubuntu@api20