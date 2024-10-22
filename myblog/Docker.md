## 下载Docker

更新yum包：yum update -y

安装所需要的软件包：yum install -y yum-utils device-mapper-persistent-data lvm2

设置yum源：yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

安装：yum install docker-ce

查看版本：docker -v

## Docker相关指令

### 一、服务相关命令

启动命令：systemctl start docker

停止docker: systemctl stop docker

查看docker服务状态：systemctl status docker

重启docker ：systemctl restart docker

开机启动docker:systemctl enable docker

### 二、镜像相关命令

查看本地有哪些镜像文件：docker images

搜索镜像：docker search xxx   //xxx为镜像名称

##### 镜像创建的三种方法：

1、基于官方提供的镜像

下载镜像：docker pull xxx        //默认为最新版本，如果选择版本直接加在后面docker pull xxx:5.0

2、使用Dockerfile

[Dockerfile创建自定义Docker镜像以及CMD与ENTRYPOINT指令的比较 - lienhua34 - 博客园 (cnblogs.com)](https://www.cnblogs.com/lienhua34/p/5170335.html)

示例![image-20230921101717994](C:\Users\小林\AppData\Roaming\Typora\typora-user-images\image-20230921101717994.png)

删除镜像：docker rmi xxx         //xxx为image ID

删除所有镜像：docker rmi `docker images -q`

### 三、容器相关命令

创建并启动容器：docker run -it --name=xxx centeros:7 /bin/bash              

//-i表示一直运行 ; -t分配伪终端用于接受命令 ; -name=xxx 取名字为xxx ；/bin/bash打开shell脚本

​                   docker run -id --name=lt centeros:7 /bin/bash  

//d表示后台运行容器，不打开容器

退出容器：exit

查看正在运行的容器：docker ps            //查看所有容器：docker ps -a

进入正在运行的容器：docker exec -it xxx /bin/bash      //xxx为容器名字

停止容器：docker stop xxx            //xxx为名字

删除容器：docker rm xxx              //xxx为名字

删除所有容器：docker rm `docker ps -aq`     //无法将正在运行的容器删除

查看容器信息：docker inspect xxx        //xxx为名字

## 数据卷相关知识

##### 一、数据卷概念

1、数据卷是宿主机中的一个文件或目录，绑定容器后对其修改会同步到容器上

2、一个数据卷可以挂载多个容器

3、一个容器可以被多个数据卷挂载

##### 二、配置数据卷

创建启动容器时使用-v参数 

docker run -it --name=xxx -v /root/data:/root/data_container centeros:7 /bin/bash              

// /root/data:/root/data_container表示将当前的data目录挂在到容器的/root/data_container下

##### 三、两个容器挂载同一个数据卷实现两个容器数据交换

##### 四、数据卷容器

将两个容器挂在到另一个容器中

1、创建并启动c3数据卷容器，使用-v参数设置数据卷

docker run -it --name=c3 -v /volume centos:7 /bin/bash

//-v /volume 没有指定宿主机的数据卷目录，docker会自动分配一个目录，可使用docker inspect c3查看 (mounts中 )

2、创建启动c1，c2容器，使用--vloumes-from参数设置数据卷

docker run -it --name=c1 --volumes-from c3  centos:7 /bin/bash

docker run -it --name=c2 --volumes-from c3  centos:7 /bin/bash

## Docker部署mysql

在Docker容器中部署mysql，并通过外部mysql客户端操作mysql Server

参考文档：https://blog.csdn.net/qq_42971035/article/details/127831101

1、搜索MySQL镜像

```
docker search mysql
```

2、拉取MySQL镜像

```
docker pull mysql:latest
```

3、创建容器

#在/root目录下创建MySQL目录用于存贮mysql数据信息

```
mkdir ~/mysql
cd ~/mysql
```

```
docker run -p 3306:3306 --name=mysql --restart=always --privileged=true \
-v /root/mysql/conf:/etc/mysql/conf.d \
-v /root/mysql/logs:/var/logs/mysql \
-v /root/mysql/data:/var/lib/mysql \
-v /etc/localtime:/etc/localtime:ro \
-e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest
```

-p 3306:3306：指定宿主机端口与容器端口映射关系

--name mysql：创建的容器名称

--restart=always：总是跟随docker启动

--privileged=true：获取宿主机root权限

-v /root/mysql/conf:/etc/mysql/conf.d ：映射配置目录，宿主机:容器

-v /root/mysql/logs:/var/logs/mysql：映射日志目录，宿主机:容器

-v /root/mysql/data:/var/lib/mysql：映射数据目录，宿主机:容器

-v /etc/localtime:/etc/localtime:ro：让容器的时钟与宿主机时钟同步，避免时区的问题，ro是read only的意思，就是只读。

-e MYSQL_ROOT_PASSWORD=123456：指定mysql环境变量，root用户的密码为123456

-d mysql:latest：后台运行mysql容器，版本是latest。

4、操作容器中的mysql（采用端口映射方式，类似于容器卷）

使用sqlyog远程连接

sqlyog文件地址：https://blog.csdn.net/qq_43541746/article/details/123186058
名称：cr173
证书密钥：ec38d297-0543-4679-b098-4baadf91f983

 



### 1. 在 Github 上下载最新版本的 Docker Compose：

```shell
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

### 2. 将下载的文件赋予执行权限：

```shell
sudo chmod +x /usr/local/bin/docker-compose
```

### 3. 创建软连接：

```shell
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

### 3. 验证 Docker Compose 是否安装成功：

![image-20231216141336060](C:\Users\kobayashi\AppData\Roaming\Typora\typora-user-images\image-20231216141336060.png)











