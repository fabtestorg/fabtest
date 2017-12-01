#!/bin/python

import sys
from fabric.api import cd, put, lcd, local, run, get

reload(sys)
sys.setdefaultencoding('utf8')

#cp jmeter config to remote
def cp_jmeter_config(yaml_name, config_dir):
    dir_name = "jmeter_config"
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.jmx"%(yaml_name,yaml_name))
        #remote yaml
        run("mkdir -p ~/fabtest/%s"%dir_name)
        put("%s.tar.gz"%yaml_name,"~/fabtest/%s"%dir_name)
        local("rm %s.tar.gz"%yaml_name)
    with cd("~/fabtest/%s"%dir_name):
        run("tar zxvfm %s.tar.gz"%yaml_name)
        run("rm %s.tar.gz"%yaml_name)

#start remote jmeter
def start_jmeter(yaml_name):
    dir_name = "jmeter_config"
    with cd("~/fabtest/%s"%dir_name):
        run("jmeter -n -t %s.jmx -l %s.jtl"%(yaml_name,yaml_name))

#get jmeter log from remote
def get_jmeter_log(yaml_name, config_dir):
    dir_name = "jmeter_config"
    with cd("~/fabtest/%s"%dir_name):
        get('%s.jtl'%yaml_name, '%s/jmeter_logs/%s.jtl'%(config_dir,yaml_name))