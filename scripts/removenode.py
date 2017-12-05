import sys
from fabric.api import run, settings,cd

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
        run("rm -rf ~/fabTestData")

def remove_client():
    with settings(warn_only=True):
        run("docker rm -f $(docker ps -aq)")
        run("rm -rf ~/fabTestData")
        run("rm -rf ~/fabtest")
        kill_process("eventserver")

def kill_process(name):
    # kill the jmeter processes for unified order project
    with cd('/tmp/'):
        pids = run("ps -ef | grep %s | grep -v 'grep' | awk '{print $2'}"%name)
        pid_list = pids.split('\r\n')
        for i in pid_list:
            run('kill -9 %s' % i)




