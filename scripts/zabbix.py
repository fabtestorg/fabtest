#!/bin/python

import sys
from fabric.api import cd, put, lcd, local, run

reload(sys)
sys.setdefaultencoding('utf8')

#cp zabbix config to remote
def cp_zabbix_config(config_name, config_dir):
    dir_name = "zabbix"
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.conf" % (config_name, config_name))
        #remote yaml
        # run("mkdir -p /etc/%s"%dir_name)/
        put("%s.tar.gz" %config_name, "/etc/%s/"%dir_name)
        local("rm %s.tar.gz" % config_name)
    with cd("/etc/%s"%dir_name):
        run("tar zxvfm %s.tar.gz" %config_name)
        run("rm %s.tar.gz" %config_name)
        run("mv %s.conf zabbix-agent.conf"%config_name)

#start remote zabbix
def start_zabbix():
    run("systemctl restart zabbix-agent")