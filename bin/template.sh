#!/bin/bash

# single render
gomplate -d zookeeper=data/zookeeper.json < ./templates/zookeeper/zookeeper1.tpl

# batch render
gomplate -d zookeeper=data/zookeeper.json  --input-dir=templates/zookeeper --output-dir=testscripts