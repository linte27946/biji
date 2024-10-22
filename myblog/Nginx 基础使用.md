# Nginx下载安装

#### Nginx开源版

 http://nginx.org/

#### 前提

安装gcc 

```
yum install -y gcc
```

安装perl库

```
 yum install -y pcre pcre-devel
```

安装zlib库

```
yum install -y zlib zlib-devel
```

#### 编译安装

```
./configure --prefix=/usr/local/nginx 

make 

make install 
```

#### 启动Nginx 

进入安装好的目录 /usr/local/nginx/sbin

一些相关指令

```
./nginx 启动
./nginx -s stop 快速停止
./nginx -s quit 优雅关闭，在退出前完成已经接受的连接请求
./nginx -s reload 重新加载配置
```

#### 关闭防火墙 

```
systemctl stop firewalld.service
```

#### 禁止防火墙开机启动

```
systemctl disable firewalld.service
```

#### 放行端口

```
firewall-cmd --zone=public --add-port=80/tcp --permanent
```

#### 安装成系统服务

```
vi /usr/lib/systemd/system/nginx.service
```

```
[Unit]
Description=nginx - web server
After=network.target remote-fs.target nss-lookup.target
[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
ExecQuit=/usr/local/nginx/sbin/nginx -s quit
PrivateTmp=true
[Install]
WantedBy=multi-user.target
```

#### 重新加载系统服务

```
systemctl daemon-reload
```

#### 启动服务 

```
systemctl start nginx.service
```

####  开机启动 

```
systemctl enable nginx.service
```

# Nginx 基础使用

#### nginx常用命令

开启服务

```
start nginx
```

停止服务：nginx停止命令stop与quit参数的区别在于stop是快速停止nginx，可能并不保存相关信息，quit是完整有序的停止nginx  ，`并保存相关信息。nginx启动与停止命令的效果都可以通过Windows任务管理器中的进程选项卡观察。

```
nginx -s stop
nginx -s quit
```

其他命令重启、关闭nginx
```
ps -ef | grep nginx
```

查看nginx状态

```
ps -ef | grep nginx
```

快速停止Nginx

```
kill -TERM 主进程号
```

强制停止Nginx
```
pkill -9 nginx
```

重启服务：

```
nginx -s reload
```



![image-20230926152649004](C:\Users\小林\AppData\Roaming\Typora\typora-user-images\image-20230926152649004.png)

### 目录结构

进入Nginx的主目录我们可以看到这些文件夹

```
client_body_temp conf fastcgi_temp html logs proxy_temp sbin scgi_temp uwsgi_temp
```

其中这几个文件夹在刚安装后是没有的，主要用来存放运行过程中的临时文件

```
client_body_temp fastcgi_temp proxy_temp scgi_tem
```

#### conf

 用来存放配置文件相关 

#### html 

用来存放静态文件的默认目录 html、css等 

#### sbin 

nginx的主程序 

#### 基本运行原理

![image-20230926152932314](C:\Users\小林\AppData\Roaming\Typora\typora-user-images\image-20230926152932314.png)

注：master负责协调多个子进程以及读取并校验配置文件

# Nginx配置与应用场景

#### 最小配置文件

```
worker_processes  1;
# 默认为1，表示开启一个业务进程，最好与服务器cpu内核个数相同

events {
    worker_connections  1024;
}
#单个业务进程可接受连接数

http {
    include       mime.types;
    #引入http mime类型，告诉浏览器服务器返回哪种类型的文件，mime.types可以对应到具体的文件后缀
    default_type  application/octet-stream;
    # 如果mime类型没匹配上，默认使用二进制流的方式传输
```

<img src="C:\Users\小林\AppData\Roaming\Typora\typora-user-images\image-20230926155026395.png" alt="image-20230926155026395" style="zoom: 67%;" />

在配置文件中可以查看修改

```
    sendfile        on;
#高效网络传输，也就是数据0拷贝。nginx不加载该文件直接将其发送

    keepalive_timeout  65;
    
#虚拟主机vhost
    server {
        listen       80;
        #监听端口号80

        server_name  localhost;
        #主机名，可以在此配置域名
        
        location / {
            root   html;
            index  index.html index.htm;
        }
         #uri
        
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

# 虚拟主机与域名解析

#### 出现原因

主机资源过剩，在一台主机上配置多个项目，使用不同域名访问同一个IP（即同一台主机），nginx处理后不同的域名指向不同的资源。

#### 虚拟主机 

原本一台服务器只能对应一个站点，通过虚拟主机技术可以虚拟化成多个站点同时对外提供服务 

#### servername匹配规则 

我们需要注意的是servername匹配分先后顺序，<u>*写在前面的匹配上就不会继续往下匹配了。*</u>

若没有匹配上则连接到第一个站点

支持正则表达式

#### 完整匹配 

我们可以在同一servername中匹配多个域名

```
server_name www.example.com www1.example.com;
```

#### 通配符匹配

```
server_name *.example.com;
```

#### 通配符结束匹配

```
server_name www.example.*;
```

# 负载均衡

<img src="https://pic1.zhimg.com/70/v2-8c1cfe007a2b5b64d221a20a335a9333_1440w.image?source=172ae18b&biz_tag=Post" alt="搞懂“负载均衡”，一篇就够了" style="zoom:80%;" />

#### 正向代理

为了从原始服务器取得内容，用户A向代理服务器Z发送一个请求并指定目标(服务器B)，然后代理服务器Z向服务器B转交请求并将获得的内容返回给客户端

![img](https://pic1.zhimg.com/80/v2-57472d50305b1525ecdc871cd811aa20_720w.webp)

#### 反向代理

nginx拿到请求的域名把他交给应用服务器(tomcat)然后tomcat提取二级域名并在database里面查找相关信息返回给nginx

![image-20230927195622424](C:\Users\小林\AppData\Roaming\Typora\typora-user-images\image-20230927195622424.png)

使用proxy_pass将不会加载后面的静态文件

```
 location / {
             proxy_pass http://baidu.com/;
           # root   html;
           # index  index.html index.htm;
        }
```

#### 基于反向代理的负载均衡

```
    sendfile        on;

    keepalive_timeout  65;
    
   upstream var {
   server 192.168.44.102:80;
   server 192.168.43.103:80;
}

    server {
        listen       80;
        server_name  localhost;
        
        location / {
             proxy_pass http://var/;
           # root   html;
           # index  index.html index.htm;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }
}
```

# 负载均衡策略 

## 轮询 

默认情况下使用轮询方式，逐一转发，这种方式适用于无状态请求。

####  weight(权重) 

指定轮询几率，weight和访问比率成正比，用于后端服务器性能不均的情况。

```
upstream httpd {
server 127.0.0.1:8050 weight=10 down;
server 127.0.0.1:8060 weight=1;
server 127.0.0.1:8060 weight=1 backup;
}
```

- **down：表示当前的server暂时不参与负载** 

- **weight：默认为1.weight越大，负载的权重就越大。**
- **backup： 其它所有的非backup机器down或者忙的时候，请求backup机器。**

#### 不常用：

轮询存在问题：无法保持会话

ip_hash ：根据客户端的ip地址转发同一台服务器，可以保持回话。

least_conn ：最少连接访问 

url_hash ：根据用户访问的url定向转发请求 

fair ：根据后端服务器响应时间转发请求

