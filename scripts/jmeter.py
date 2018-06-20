#!/bin/python

import sys
import os
from fabric.api import cd, put, lcd, local, run, get
import utils
reload(sys)
sys.setdefaultencoding('utf8')

#cp jmeter config to remote
def start_jmeter(file_name, config_dir):
    dir_name = "jmeter_config"
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %sjmeter.jmx" % (file_name, file_name))
        #remote yaml
        run("rm -rf ~/fabtest/%s"%dir_name)
        run("mkdir -p ~/fabtest/%s"%dir_name)
        put("%s.tar.gz" % file_name, "~/fabtest/%s" % dir_name)
        local("rm %s.tar.gz" % file_name)
    with cd("~/fabtest/%s"%dir_name):
        run("tar zxvfm %s.tar.gz" % file_name)
        run("rm %s.tar.gz" % file_name)
        utils.kill_process("jmeter")
        run("screen -d -m ~/jmeter/apache-jmeter-3.2/bin/jmeter -n -t %sjmeter.jmx -l %sjmeter.jtl"%(file_name,file_name), pty=False)

#get jmeter log from remote
def get_jmeter_log(yaml_name, config_dir):
    dir = "%s/event_logs/%s"%(config_dir,log_dir)
    local("mkdir -p %s"%dir)
    file = '%s/%sjmeter.jtl'%(dir,yaml_name)
    if os.path.exists(file):
        local("rm -rf %s"%file)
    get('~/fabtest/jmeter_config/%sjmeter.jtl'%yaml_name,'%s'%file)
    with lcd(dir):
        local("~/jmeter/apache-jmeter-3.2/bin/jmeter -g %s -e -o ./jmeterReport"%file)

#get eventserver log from remote
def get_eventserver_log(yaml_name, config_dir, log_dir):
    dir = "%s/event_logs/%s"%(config_dir,log_dir)
    local("mkdir -p %s"%dir)
    file = '%s/%s_eventserver.log'%(dir,yaml_name)
    if os.path.exists(file):
        local("rm -rf %s"%file)
    get('~/fabtest/event_server/eventserver.log','%s'%file)
    #echo  empty log
    run("cat /dev/null > ~/fabtest/event_server/eventserver.log")

# remote
def start_haproxy(file_name, config_dir):
    dir_name = "haproxy_config"
    with lcd(config_dir):
        local("rm -rf haproxy_config/haproxy.cfg")
        local("cp %shaproxy.cfg haproxy_config/haproxy.cfg"%file_name)
        local("tar -zcvf %shaproxyconfig.tar.gz haproxy_config" %file_name)
        #remote yaml
        run("rm -rf ~/fabtest/%s"%dir_name)
        run("mkdir -p ~/fabtest/")
        put("%shaproxyconfig.tar.gz" %file_name, "~/fabtest/")
        local("rm %shaproxyconfig.tar.gz" %file_name)
    with cd("~/fabtest"):
        run("tar zxvfm %shaproxyconfig.tar.gz" %file_name)
        run("rm %shaproxyconfig.tar.gz" % file_name)
    with cd("~/fabtest/haproxy_config"):
        run("docker-compose -f docker-compose.yaml up -d")
