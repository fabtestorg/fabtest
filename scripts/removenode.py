import sys
from fabric.api import run, settings,cd,sudo
import utils
reload(sys)
sys.setdefaultencoding('utf8')

def remove_node(type):
    with settings(warn_only=True):
        if type == "all":
            run("docker ps -a | awk '{print $1}' | xargs docker rm -f")
            run("unset GREP_OPTIONS && docker images |grep 'dev\-peer'|awk '{print $3}'|xargs docker rmi -f")
        else:
            run("unset GREP_OPTIONS && docker ps -a | grep %s | awk '{print $1}' | xargs docker rm -f"%type)
            if type == "peer":
                run("unset GREP_OPTIONS && docker images |grep 'dev\-peer'|awk '{print $3}'|xargs docker rmi -f")
        run("docker network prune -f")
        sudo("rm -rf ~/fabtest")
        sudo("rm -rf /data/*")

def remove_client():
    with settings(warn_only=True):
       run("docker ps -a | awk '{print $1}' | xargs docker rm -f")
       run("docker network prune -f")
       sudo("rm -rf ~/fabtest")
       sudo("rm -rf ~/fabTestData")
       utils.kill_process("eventserver")

def remove_jmeter():
    with settings(warn_only=True):
        utils.kill_process("jmeter")



