import sys
from fabric.api import run, settings,cd
import utils.py
reload(sys)
sys.setdefaultencoding('utf8')

def remove_node(type):
    with settings(warn_only=True):
        if type == "all":
            run("docker rm -f $(docker ps -aq)")
            run("docker images |grep 'dev\-peer'|awk '{print $3}'|xargs docker rmi -f")
        else:
            run("docker ps -a | grep %s | awk '{print $1}' | xargs docker rm -f"%type)
            if type == "peer":
                run("docker images |grep 'dev\-peer'|awk '{print $3}'|xargs docker rmi -f")

        run("rm -rf ~/fabtest")

def remove_client():
    with settings(warn_only=True):
        run("docker rm -f $(docker ps -aq)")
        run("rm -rf ~/fabtest")
        utils.kill_process("eventserver")




