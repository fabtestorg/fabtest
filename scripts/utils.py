import sys
from fabric.api import run, settings,local

reload(sys)
sys.setdefaultencoding('utf8')

def kill_process(name):
    # kill the jmeter processes for unified order project
    run("unset GREP_OPTIONS && ps -ef | grep %s | grep -v 'grep' | awk '{print $2'} | xargs kill -9"%name)

def check_remote_file_exist(file):
    with settings(warn_only=True):
        result = run("ls %s"%file)
        if "cannot" in result:
            exist = "false"
        else:
            exist = "true"
        return exist

def check_remote_dir_exist(dir):
    with settings(warn_only=True):
        result = run("du -sh %s"%dir)
        if "cannot" in result:
            exist = "false"
        else:
            exist = "true"
        return exist

def check_container_exist(name):
    containers = run('unset GREP_OPTIONS && docker ps |grep "%s" | wc -l' %name)
    if containers != "0":
        result = "true"
    else:
        result = "false"
    return result

def set_domain_name(network_name,node_full_name,domain_ip,domain_name):
    set_cmd = "echo %s %s >> /etc/hosts"%(domain_ip,domain_name)
    yaml_file = "~/networklist/%s/%s/%s.yaml"%(network_name,node_full_name,node_full_name)
    run("docker-compose -f %s exec %s bash -c '%s'"%(yaml_file,node_full_name,set_cmd))

def get_domain_name(network_name,node_full_name,domain_name):
    get_cmd = "unset GREP_OPTIONS && cat /etc/hosts | grep -E %s | awk '{print \\\$1}'"%domain_name
    yaml_file = "~/networklist/%s/%s/%s.yaml"%(network_name,node_full_name,node_full_name)
    out = run('docker-compose -f %s exec %s bash -c "%s"'%(yaml_file,node_full_name,get_cmd))

def rm_local(path):
    local("rm -rf %s"%path)
