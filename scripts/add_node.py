#!/bin/python

from fabric.api import cd,put,lcd,local,run,settings
import sys
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
    eventyamlname = name + "eventclient"
    #apiserver
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz %s.yaml"%(apiyamlname,apiyamlname))
        #remote yaml
        run("mkdir -p ~/fabtest/%s"%apiyamlname)
        put("%s.tar.gz"%apiyamlname,"~/fabtest/%s"%apiyamlname)
        local("rm %s.tar.gz"%apiyamlname)
    with cd("~/fabtest/%s"%apiyamlname):
        run("tar zxvfm %s.tar.gz"%apiyamlname)
        run("rm %s.tar.gz"%apiyamlname)
        run("docker-compose -f %s.yaml up -d"%apiyamlname)

    #eventserver
    with lcd(config_dir):
        local("tar -zcvf %s.tar.gz eventserver"%eventyamlname)
        #remote yaml
        run("mkdir -p ~/fabtest/%s"%eventyamlname)
        put("%s.tar.gz"%eventyamlname,"~/fabtest/%s"%eventyamlname)
        local("rm %s.tar.gz"%eventyamlname)
    with cd("~/fabtest/%s"%eventyamlname):
        run("tar zxvfm %s.tar.gz"%eventyamlname)
        run("rm %s.tar.gz"%eventyamlname)
        run("nohup %s &"%eventyamlname)