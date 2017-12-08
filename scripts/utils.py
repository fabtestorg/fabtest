import sys
from fabric.api import run, settings

reload(sys)
sys.setdefaultencoding('utf8')

def kill_process(name):
    # kill the jmeter processes for unified order project
    run("ps -ef | grep %s | grep -v 'grep' | awk '{print $2'} | xargs kill -9"%name)

def check_remote_file_exist(file):
    with settings(warn_only=True):
        result = run("ls %s"%file)
        if result.find("No such file or directory") == -1:
            exist = "false"
        else:
            exist = "true"
        return exist

def set_domain_name(network_name,node_full_name,domain_ip,domain_name):
    set_cmd = "echo %s %s >> /etc/hosts"%(domain_ip,domain_name)
    yaml_file = "~/networklist/%s/%s/%s.yaml"%(network_name,node_full_name,node_full_name)
    run("docker-compose -f %s exec %s bash -c '%s'"%(yaml_file,node_full_name,set_cmd))

def get_domain_name(network_name,node_full_name,domain_name):
    get_cmd = "cat /etc/hosts | grep -E %s | awk '{print \\\$1}'"%domain_name
    yaml_file = "~/networklist/%s/%s/%s.yaml"%(network_name,node_full_name,node_full_name)
    out = run('docker-compose -f %s exec %s bash -c "%s"'%(yaml_file,node_full_name,get_cmd))
