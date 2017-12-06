#!/bin/python

import sys
import os
from fabric.api import cd, put, lcd, local, run, get

reload(sys)
sys.setdefaultencoding('utf8')

#cp jmeter config to remote
def start_jmeter(file_name, config_dir):
    dir_name = "jmeter_config"
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.jmx" % (file_name, file_name))
        #remote yaml
        run("rm -rf ~/fabtest/%s"%dir_name)
        run("mkdir -p ~/fabtest/%s"%dir_name)
        put("%s.tar.gz" % file_name, "~/fabtest/%s" % dir_name)
        local("rm %s.tar.gz" % file_name)
    with cd("~/fabtest/%s"%dir_name):
        run("tar zxvfm %s.tar.gz" % file_name)
        run("rm %s.tar.gz" % file_name)
        run("screen -d -m ~/jmeter/apache-jmeter-3.2/bin/jmeter -n -t %s.jmx -l %s.jtl"%(file_name,file_name), pty=False)

#get jmeter log from remote
def get_jmeter_log(yaml_name, config_dir):
    dir_name = "jmeter_config"
    with cd("~/fabtest/%s"%dir_name):
        local("rm -rf  %s/jmeter_logs/%s.jtl"%(config_dir,yaml_name))
        local("rm -rf  %s/jmeter_logs/%s.log"%(config_dir,yaml_name))
        get('%s.jtl'%yaml_name, '%s/jmeter_logs/%s.jtl'%(config_dir,yaml_name))
        get('jmeter.log', '%s/jmeter_logs/%s.log'%(config_dir,yaml_name))

#get eventserver log from remote
def get_eventserver_log(yaml_name, config_dir, log_dir):
    dir = "%sevent_logs/%s"%(config_dir,log_dir)
    local("mkdir -p %s"%dir)
    file = '%s/%s_eventserver.log'%(dir,yaml_name)
    if os.path.exists(file):
        local("rm -rf %s"%file)
    get('~/event_server/eventserver.log','%s'%file)
    #echo  empty log
    run("cat /dev/null > ~/event_server/eventserver.log")