#!/bin/python

import sys
import os
from fabric.api import cd, put, lcd, local, run, get
import utils
reload(sys)
sys.setdefaultencoding('utf8')

#cp jmeter config to remote
def start_jmeter(config_dir):
    dir_name = "jmeter_config"
    with lcd(config_dir):
        local("tar -zcvf jmeterjmx.tar.gz jmeter.jmx")
        #remote yaml
        run("rm -rf ~/fabtest/%s"%dir_name)
        run("mkdir -p ~/fabtest/%s"%dir_name)
        put("jmeterjmx.tar.gz", "~/fabtest/%s"%dir_name)
        local("rm jmeterjmx.tar.gz")
    with cd("~/fabtest/%s"%dir_name):
        run("tar zxvfm jmeterjmx.tar.gz")
        run("rm jmeterjmx.tar.gz")
        utils.kill_process("jmeter")
        run("screen -d -m ~/jmeter/apache-jmeter-3.2/bin/jmeter -n -t jmeter.jmx -l jmeter.jtl", pty=False)

#get jmeter log from remote
def get_jmeter_log(config_dir, log_dir):
    dir = "%sevent_logs/%s"%(config_dir,log_dir)
    local("mkdir -p %s"%dir)
    file = '%s/jmeter.jtl'%dir
    if os.path.exists(file):
        local("rm -rf %s"%file)
    get('~/fabtest/jmeter_config/jmeter.jtl',file)
    with lcd(dir):
        local("~/jmeter/apache-jmeter-3.2/bin/jmeter -g %s -e -o ./jmeterReport"%file)
        local("rm -rf jmeter.log")

#get eventserver log from remote
def get_eventserver_log(yaml_name, config_dir, log_dir):
    dir = "%sevent_logs/%s"%(config_dir,log_dir)
    local("mkdir -p %s"%dir)
    file = '%s/%s_eventserver.log'%(dir,yaml_name)
    if os.path.exists(file):
        local("rm -rf %s"%file)
    get('~/fabtest/event_server/eventserver.log','%s'%file)
    #echo  empty log
    run("cat /dev/null > ~/fabtest/event_server/eventserver.log")

# remote
def start_haproxy(config_dir):
    dir_name = "haproxy_config"
    with lcd(config_dir):
        local("tar -zcvf haproxyconfig.tar.gz haproxy_config")
        #remote yaml
        run("rm -rf ~/fabtest/%s"%dir_name)
        run("mkdir -p ~/fabtest/")
        put("haproxyconfig.tar.gz", "~/fabtest/")
        local("rm haproxyconfig.tar.gz")
    with cd("~/fabtest"):
        run("tar zxvfm haproxyconfig.tar.gz")
        run("rm -rf haproxyconfig.tar.gz")
    with cd("~/fabtest/haproxy_config"):
        run("docker-compose -f docker-compose.yaml up -d")
