import sys
from fabric.api import run, settings

reload(sys)
sys.setdefaultencoding('utf8')

def remove_node(type):
    with settings(warn_only=True):
        if type == "all":
            run("docker ps | awk '{print $1}' | xargs docker rm -f")
        else:
            run("docker ps | grep %s | awk '{print $1}' | xargs docker rm -f"%type)

