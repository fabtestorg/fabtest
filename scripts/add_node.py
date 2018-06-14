#!/bin/python

from fabric.api import cd,put,lcd,local,run,settings
import sys
import os
import utils
reload(sys)
sys.setdefaultencoding('utf8')

def load_images(type,images_path):
    filter = type
    if type == "ca":
        filter = "fabric-ca"
    result = run('unset GREP_OPTIONS && docker images | grep -e "%s" | wc -l'%filter)
    if result == "0":
        with settings(warn_only=True):
            run("mkdir -p ~/images")
        print "check local image package is exist"
        local("ls %s/%s.tar.gz"%(images_path,type))
        put("%s/%s.tar.gz"%(images_path,type),"~/images/")
        with cd("~/images/"):
            #load image
            run("tar zxvfm %s.tar.gz"%type)
            run("rm %s.tar.gz"%type)
            run("docker load -i %s.tar"%type)
    else:
        sys.stdout.write("%s image is exsit"%type)

def start_node(type, node_id, yaml_name, config_dir):
    dir_name = type + node_id
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.yaml"%(yaml_name,yaml_name))
        #remote yaml
        run("mkdir -p ~/fabtest/%s"%dir_name)
        put("%s.tar.gz"%yaml_name,"~/fabtest/%s"%dir_name)
        local("rm %s.tar.gz"%yaml_name)

    #start container
    with cd("~/fabtest/%s"%dir_name):
        run("tar zxvfm %s.tar.gz"%yaml_name)
        run("rm %s.tar.gz"%yaml_name)
        run("docker-compose -f %s.yaml up -d"%yaml_name)

def start_api_event(peer_id, org_id, config_dir, clitype):
    name = "peer" + peer_id + "org" + org_id
    yamlname = name + "%sclient"%clitype
    parent_path  = os.path.dirname(config_dir)
    #apiserver or eventserver
    with lcd(config_dir):
        local("cp %s.yaml %s_server/client_sdk.yaml"%(yamlname,clitype))
    with lcd(parent_path):
        run("cp /etc/hosts /etc/hosts)
        local("tar -zcvf %s_server.tar.gz %s_server"%(clitype,clitype))
        #remote yaml
        run("mkdir -p ~/fabtest/")
        put("%s_server.tar.gz"%clitype,"~/fabtest")
        local("rm %s_server.tar.gz"%clitype)
        utils.kill_process("%sserver"%clitype)
    with cd("~/fabtest"):
        run("tar zxvfm %s_server.tar.gz"%clitype)
        run("rm %s_server.tar.gz"%clitype)
    with cd("~/fabtest/%s_server"%clitype):
        run("tar zxvfm %sserver.tar.gz"%clitype)
        run("chmod +x %sserver"%clitype)
        run("rm -rf %sserver.log"%clitype)
        run("$(nohup ./%sserver >> %sserver.log 2>&1 &) && sleep 1"%(clitype,clitype))
        run("cat /dev/null > %sserver.log"%clitype)
        #run("docker-compose -f docker-compose.yaml up -d")

def stop_node(type, node_id, yaml_name):
    dir_name = type + node_id
    #start container
    with cd("~/fabtest/%s"%dir_name):
        run("docker-compose -f %s.yaml stop"%yaml_name)

def restart_node(type, node_id, yaml_name):
    dir_name = type + node_id
    #start container
    with cd("~/fabtest/%s"%dir_name):
        run("docker-compose -f %s.yaml start"%yaml_name)
