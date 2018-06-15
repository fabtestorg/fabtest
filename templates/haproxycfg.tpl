#全局配置, 用于设定义全局参数, 属于进程级的配置, 通常与操作系统配置有关.
global
    #定义全局日志, 配置在本地, 通过local0 输出, 默认是info级别，可配置两条
    log 127.0.0.1 local0 debug
    #定义日志级别【error warning info debug】
    #log 127.0.0.1 local1 info

    #设置每haproxy进程的最大并发连接数, 其等同于命令行选项“-n”; “ulimit -n”自动计算的结果参照此参数设定.
    maxconn 4096

    #后台运行haproxy
    daemon

    #设置启动的haproxy进程数量, 只能用于守护进程模式的haproxy;
    #默认只启动一个进程, 鉴于调试困难等多方面的原因, 一般只在单进程仅能打开少数文件描述符的场景中才使用多进程模式.
    nbproc 1
    #设置每进程所能够打开的最大文件描述符数目, 默认情况其会自动进行计算, 因此不推荐修改此选项.
    #ulimit-n 819200

    #调试级别, 一般只在开启单进程时调试, 且生产环境禁用.
    #debug
    #haproxy启动后不会显示任何相关信息, 这与在命令行启动haproxy时加上参数“-q”相同
    #quiet

defaults
    mode http
    log global
    option forwardfor
    option httpclose
    #如果产生了一个空连接，那这个空连接的日志将不会记录.
    option dontlognull

    #当与后端服务器的会话失败(服务器故障或其他原因)时, 把会话重新分发到其他健康的服务器上; 当故障服务器恢复时, 会话又被定向到已恢复的服务器上;
    #还可以用”retries”关键字来设定在判定会话失败时的尝试连接的次数
    option redispatch
    retries 3

    #当haproxy负载很高时, 自动结束掉当前队列处理比较久的链接.
    option abortonclose

    #默认http请求超时时间
    timeout http-request 10s
    #默认队列超时时间, 后端服务器在高负载时, 会将haproxy发来的请求放进一个队列中.
    timeout queue 1m
    #haproxy与后端服务器连接超时时间.
    timeout connect 5s
    #客户端与haproxy连接后, 数据传输完毕, 不再有数据传输, 即非活动连接的超时时间.
    timeout client 1m
    #haproxy与后端服务器非活动连接的超时时间.
    timeout server 1m
    #默认新的http请求连接建立的超时时间，时间较短时可以尽快释放出资源，节约资源.
    timeout http-keep-alive 10s
    #心跳检测超时时间
    timeout check 10s

    #最大并发连接数
    maxconn 2000

    #设置默认的负载均衡方式
    #balance source
    #balnace leastconn

frontend api_front
    bind 0.0.0.0:2222
    default_backend api_back

backend api_back
    balance roundrobin
    mode http
    cookie SERVERID
    server api1 {{.apiip1}}:5555 maxconn 1024 weight 3 check inter 1500 rise 2 fall 3
    server api2 {{.apiip2}}:5555 maxconn 1024 weight 3 check inter 1500 rise 2 fall 3
