## 7、项目的部署(负责人林特)

**7.1硬件环境：**

**1）**     **虚拟机：VMware Workstation Pro**

**2）**     **操作系统：CenterOS 7.6**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image002.gif)**

**7.2 软件环境：**

**1）docker、Django、mysql**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image004.gif)**

**7.3****部署的步骤**

**1****）安装****docker**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image006.gif)**

**2****）拉取****mysql****镜像**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image008.gif)**

**3****）****Django****容器的构建**

**(a)** **项目的目录结构：**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image010.gif)**

**(b)** **编写****requirements.txt**

**requirements.txt****中包含了这个项目所需要用到的包**

**在****python****中使用****pipreqs****自动生成****requirements.txt**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image012.gif)**

**(c)** **编写****Dockerfile**

**Dockerfile****中对系统环境，镜像目录等进行设置，通过****Dockerfile****去执行设置好的操作命令，保证通过****Dockerfile****的构建镜像是一致的。**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image014.gif)**

**(d)** **编写****docker-compose.yml****文件**

**Compose** **是一个用于定义和运行多容器** **Docker** **的工具。借助** **Compose****，可以使用** **YAML** **文件来配置应用程序的服务。**

**先安装****docker-compose**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image016.gif)**

**docker-compose.yml****文件的具体内容：**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image018.gif)**

**(e)** **注意在****setting.py****中的设置要与****docker-compose.yml****文件一样，尤其是****HOST****部分不能写****127.0.0.1****，否则会报错无法连接上数据库。**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image020.gif)**

**(f)** **创建镜像，然后再启动服务。**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image022.gif)**

**也可以创建完镜像容器再分别启动**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image024.gif)**

**(g)****进行数据库的迁移**

**使用命令****docker exec -it databaseteemwork_db_1 bash****进入数据库容器，登录账号后创建名为****ustb2021****的数据库**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image026.gif)**

**然后进入****django****容器中****docker exec -it databaseteemwork_dj_web_1 bash**

**执行****python manage.py makemigrations****和****python manage.py migrate****将数据迁移过去**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image028.gif)**

**再进入****mysql****容器中执行****sql****语句创建数据信息****,****由于篇幅只展示部分**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image030.gif)**

**最后验证一下：**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image032.gif)**

**7.4** **测试**

**最后在主机上使用虚拟机的****ip****加端口号访问该项目，****如果要部署到服务器也是同样的步骤。**

**登录****http://192.168.37.138:8090/logreg/login/****可以看到成功访问到了该项目****,****且数据库里的内容能够正确的访问到。**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image034.gif)**

**![img](file:///C:/Users/KOBAYA~1/AppData/Local/Temp/msohtmlclip1/01/clip_image036.gif)**