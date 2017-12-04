import sys
from fabric.api import run

reload(sys)
sys.setdefaultencoding('utf8')

def remove_node(type):
    if type == "all":
        run("docker ps -a | awk '{print $1}' | xargs docker rm -f")
    else:
        run("docker ps -a | grep %s | awk '{print $1}' | xargs docker rm -f"%type)


