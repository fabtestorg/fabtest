#!/bin/python

from fabric.api import cd,put,lcd,local,run,settings,sudo
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

def replace_images(type, config_path):
    if type == "order":
        type = "orderer"
    local("docker save hyperledger/fabric-%s:latest -o %s%s.tar"%(type,config_path,type))
    run("rm -rf ~/%s.tar"%type)
    put("%s%s.tar"%(config_path,type),"~")
    sudo("systemctl restart docker")
    run("docker rmi hyperledger/fabric-%s:latest"%type)
    run("docker load -i ~/%s.tar"%type)
    sudo("systemctl restart docker")

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


def start_docker():
    #start docker service
    sudo("systemctl restart docker")

def start_api(peer_id, org_id, config_dir):
    name = "peer" + peer_id + "org" + org_id
    apiclientname = name + "apiclient"
    apidockername = name + "apidocker"
    parent_path  = os.path.dirname(config_dir)
    #apiserver
    with lcd(parent_path):
        #remote yaml
        run("mkdir -p ~/fabtest/")
        run("rm -rf ~/fabtest/api_server")
        put("api_server.tar.gz","~/fabtest")
    with cd("~/fabtest"):
        run("tar zxvfm api_server.tar.gz")
        run("rm -rf api_server.tar.gz")
    with lcd(config_dir):
        put("%s.yaml"%apiclientname, "~/fabtest/api_server/client_sdk.yaml")
        put("%s.yaml"%apidockername, "~/fabtest/api_server/docker-compose.yaml")
    with cd("~/fabtest/api_server"):
        run("docker-compose -f docker-compose.yaml down")
        run("docker-compose -f docker-compose.yaml up -d")

def start_event(peer_id, org_id, config_dir, clitype):
    name = "peer" + peer_id + "org" + org_id
    yamlname = name + "%sclient"%clitype
    parent_path  = os.path.dirname(config_dir)
    #apiserver or eventserver
    with lcd(parent_path):
        put("/etc/hosts","~")
        #remote yaml
        run("mkdir -p ~/fabtest/")
        run("rm -rf ~/fabtest/%s_server"%clitype)
        put("%s_server.tar.gz"%clitype,"~/fabtest")
        utils.kill_process("%sserver"%clitype)
    with cd("~/fabtest"):
        run("tar zxvfm %s_server.tar.gz"%clitype)
        run("rm %s_server.tar.gz"%clitype)
    with lcd(config_dir):
        put("%s.yaml"%yamlname, "~/fabtest/%s_server/client_sdk.yaml"%clitype)
    with cd("~/fabtest/%s_server"%clitype):
        sudo("cp ~/hosts /etc/hosts")
        run("chmod +x %sserver"%clitype)
        run("rm -rf %sserver.log"%clitype)
        run("$(nohup ./%sserver >> %sserver.log 2>&1 &) && sleep 1"%(clitype,clitype))
        run("cat /dev/null > %sserver.log"%clitype)

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
