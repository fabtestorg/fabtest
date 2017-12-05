#!/bin/python

import sys
from fabric.api import cd, put, lcd, local, run

reload(sys)
sys.setdefaultencoding('utf8')

#cp zabbix config to remote
def start_zabbix(config_name, config_dir):
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.conf" % (config_name, config_name))
        #remote yaml
        # run("mkdir -p /etc/%s"%dir_name)/
        put("%s.tar.gz"%config_name, "/etc/zabbix/")
        # local("rm %s.tar.gz" % config_name)
    with cd("/etc/zabbix/"):
        run("tar zxvfm %s.tar.gz" %config_name)
        # run("rm %s.tar.gz" %config_name)
        run("cp %s.conf zabbix-agent.conf"%config_name)
        run("systemctl restart zabbix-agent")