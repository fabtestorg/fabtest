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
    result = run('docker images | grep -e "%s" | wc -l'%filter)
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

def start_api(peer_id, org_id, config_dir):
    name = "peer" + peer_id + "org" + org_id
    apiyamlname = name + "apiclient"
    parent_path  = os.path.dirname(config_dir)
    #apiserver
    with lcd(config_dir):
        local("cp %s.yaml fft_apiserver/"%apiyamlname)
        local("cp %s.yaml fft_apiserver/client_sdk.yaml"%apiyamlname)
    with lcd(parent_path):
        local("tar -zcvf fft_apiserver.tar.gz fft_apiserver")
        #remote yaml
        run("mkdir -p ~/fabtest/")
        put("fft_apiserver.tar.gz","~/fabtest")
        local("rm fft_apiserver.tar.gz")
        utils.kill_process("apiserver")
    with cd("~/fabtest"):
        run("tar zxvfm fft_apiserver.tar.gz")
        run("rm fft_apiserver.tar.gz")
    with cd("~/fabtest/fft_apiserver"):
        run("chmod +x apiserver")
        run("rm -rf apiserver.log")
        run("$(nohup ./apiserver >> apiserver.log 2>&1 &) && sleep 1")
        run("cat /dev/null > apiserver.log")
        #run("docker-compose -f docker-compose.yaml up -d")

def start_event(peer_id, org_id, config_dir):
    name = "peer" + peer_id + "org" + org_id
    eventyamlname = name + "eventclient"
    #eventserver
    with lcd(config_dir):
        #remote yaml
        local("tar -zcvf %s.tar.gz %s.yaml"%(eventyamlname,eventyamlname))
        put("%s.tar.gz"%eventyamlname,"~/event_server")
        put("eventserver.tar.gz","~/event_server")
        put("current.info","~/event_server")
        local("rm %s.tar.gz"%eventyamlname)
        utils.kill_process("eventserver")
    with cd("~/event_server"):
        run("tar zxvfm %s.tar.gz"%eventyamlname)
        run("tar zxvfm eventserver.tar.gz")
        run("cp  %s.yaml client_sdk.yaml"%eventyamlname)
        run("rm %s.tar.gz"%eventyamlname)
        run("$(nohup ./eventserver >> eventserver.log 2>&1 &) && sleep 1")

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