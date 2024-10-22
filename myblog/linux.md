[toc]

# ubuntu

## 自带文本编辑器gedit

使用gedit可以像widows文本编辑器那样打开一个文件

```shell
sudo gedit xxx
```

## fish-超好用的shell工具

```shell
sudo apt-get install fish
#在终端中输入fish即可使用
```

fish自带命令补全还能判断命令是否有错，而且还会记录历史命令，使用linux瞬间变得丝滑了许多

不过刚学linux的初学者还是别装了，许多命令还是要多打才会熟练。除了fish之外zsh也很有名，不过安装配置较为麻烦。

## 网络代理（QV2ray）

在打开科学上网软件后还需要对网络代理进行设置，如下图所示：

<img src="/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240116233542598.png" style="zoom:50%;" /> 

<img src="F:\myblog\typora-user-images\linux\image-20240116233542598.png" alt="image-20240116233542598" style="zoom:50%;" /> 

其中127.0.0.1代表本地回环地址，所有流量都将被定向到本地机器上，8090是端口号，这种设置通常用于实现VPN的全局代理，确保所有网络请求都通过VPN连接发送和接收数据。

如果将网络代理设置为192.168.x.x或10.x.x.x时，流量将被转发到具有该IP地址的设备上，该设备可能是**本地机器上运行的代理服务器，或者是局域网中的其他设备**。

补充：如果要对终端进行加速可以打开Qv2ray中首选项，然后开启系统代理。

## 修改网络ip

```shell
cd /etc/netplan/
vim 01-network-manager-all.yaml
```

https://www.cnblogs.com/liujiaxin2018/p/16287463.html

## debian包管理形式

### dpkg和apt

dpkg主要用于已经下载的deb文件但需要自己安装依赖，apt-get则用于从源下载软件但可以自动解决依赖。

**当然使用aptitude命令工具可以很好解决dpkg依赖的问题**，aptitude工具本质上是apt工具和dpkg的前端。dpkg是软件包管理系统工具，而aptitude则是完整的软件包管理系统。

## aptitude使用

dpkg工具的一个前端是aptitude，它提供了处理dpkg格式软件包的简单命令行选项

```shell
sudo apt install aptitude		#先安装aptitude
aptitude		#即可打开该工具
```

- 查找所需软件包（采用模糊匹配）

  i 表示已经安装到系统上了

  p或v，说明这个包可用，但还没安装

```shell
$ aptitude search mysql
p   mysql-client                             - MySQL database client (metapackage depending on th
p   mysql-client-8.0                         - MySQL database client binaries                    
p   mysql-client-8.0:i386                    - MySQL database client binaries                    
p   mysql-client-core-8.0                    - MySQL database core client binaries               
p   mysql-client-core-8.0:i386               - MySQL database core client binaries               
p   mysql-common                             - MySQL database common files, e.g. /etc/mysql/my.cn
v   mysql-common:i386                        -                                                   
v   mysql-common-5.6                         -                                                   
p   mysql-router                             - route connections from MySQL clients to MySQL serv
p   mysql-router:i386                        - route connections from MySQL clients to MySQL serv
p   mysql-sandbox                            - Install and set up one or more MySQL server instan                    
```

- 查找所需要的包mysql-client,软件包相关的详细信息来自于软件仓库。

```shell
$ aptitude show mysql-client
软件包：mysql-client             
版本号：8.0.36-0ubuntu0.22.04.1
状态: 未安装
优先级：可选
部分：database
维护者：Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
体系：all
未压缩尺寸：35.8 k
依赖于: mysql-client-8.0
描述：MySQL database client (metapackage depending on the latest version)
 This is an empty package that depends on the current "best" version of mysql-client (currently
 mysql-client-8.0), as determined by the MySQL maintainers.  Install this package if in doubt
 about which MySQL version you want, as this is the one considered to be in the best shape by the
 Maintainers.
主页：http://dev.mysql.com/
```

- 安装软件

```shell
sudo aptitude install mysql-client
```

- 更新软件

```shell
aptitude safe-upgrade
```

- 删除软件

```shell
sudo aptitude purge mysql-cilent
```

## 删除软件

- 终端里直接使用apt-get安装的软件，即sudo apt-get install 软件名，卸载方法如下：

```shell
sudo apt-get remove 软件名				#可能有些配置没有清理干净
sudo apt-get remove --purge 软件名		#添加上–purge可以卸载并清除配置，完美
```

- 使用.deb包进行安装的软件，即sudo dpkg -i *.deb，此时，可使用sudo dpkg -info *.deb 查看软件包信息，然后卸载方法如下：

```shell
sudo dpkg -r *.deb				#这种方法也是一样，清除不干净！!
sudo dpkg -r --purge *.deb		#可以很好的连同配置文件一起删除！
```

注意：若是像查看系统中已安装软件包信息，可以使用命令 dpkg -l

## 修复安装

```shell
sudo apt -f install			
```

## 系统备份

初学ubuntu很容易不小心就将系统弄坏了，所以一定提前备份一下，尤其是本身系统就是ubuntu，虚拟机玩家直接vm里添加快照就好了。

https://www.guyuehome.com/34859

## ubuntu扩容

https://www.bilibili.com/video/BV1Cc41127B9?p=21

## ubuntu安装搜狗输入法

https://blog.csdn.net/kobayashiii/article/details/136297842?spm=1001.2014.3001.5502

## ubuntu防火墙管理

### ufw、firewalld及iptables

UFW是Debian系列的默认防火墙，firewall 是红帽系列7及以上的防火墙（如CentOS7.x）

首先，iptables是最底层、最古老的防火墙系统，所有系统都会存在此防火墙，但一般而言只需保证该防火墙处于完全开放状态即可，其他不用管他，更不需要复杂的配置。而ufw和firewall都是较新linux系统上的替代iptables的工具，当他们同时安装在服务器上时，两者之间就会存在冲突。

firewall和ufw可共同影响服务器，任一防火墙开启都会使端口无法连接

### ufw具体使用

#### 查看防火墙状态

 ```shell
 sudo ufw status
 ```

#### 防火墙重启

```shell
sudo ufw reload  
```

#### 开启/关闭防火墙自启动

```shell
sudo ufw enable/disable
```

#### 开放/关闭端口

```shell
sudo ufw allow 21		#开启21端口
sudo ufw allow 21/ftp		#开启21端口的ftp协议
sudo ufw allow from 192.168.121.1 	#开启指定ip的所有协议 
sudo ufw allow from 192.168.121.2 to any port 3306	#开起指定ip的指定端口
 
sudo ufw delete allow 21		#关闭，其余的都同理
```

## 闲聊

### 为什么snap不被大家所待见

https://cloud.tencent.com/developer/article/2017496

而且我觉得ubuntu上的snap商店一点都不好用.........

# Linux

## 三种网络连接模式：

**1、NAT（网络地址转换模式）--多用于家庭环境**

安装好虚拟机后，它的默认网络模式就是NAT模式。

原理：通过宿主机的网络来访问公网。虚拟局域网内的虚拟机在对外访问时，使用的则是宿主机的IP地址，这样从外部网络来看，只能看到宿主机，完全看不到新建的虚拟局域网。

优势：虚拟系统接入互联网非常简单，只需宿主机器能访问互联网即可, 不需要进行任何手工配置。

**2、Bridged(桥接模式）--多用于办公环境 **

类似局域网中的一台独立的主机，它可以访问内网任何一台机器，但是它要和宿主机器处于同一网段，这样虚拟系统才能和宿主机器进行通信【主机防火墙开启会导致ping不通】

设置：

（1）默认存在自动获取ip机制，只需要将虚拟机设置为Bridged(桥接模式），虚拟机会自动获取新的ip，保证ip地址与宿主机在同一个网段。

（2）如果是手工配置机制，那么为了保持虚拟机与宿主机在同一个网段，其中涉及人工配置ip，比较麻烦。

使用场景：如果想利用VMWare在局域网内新建一个虚拟服务器，为局域网用户提供网络服务，就应该选择桥接模式。

**3、Host-only(主机模式)  -- 用得比较少**

 在某些特殊的网络环境中，要求将真实环境和虚拟环境隔离开，这时你就可采用host-only模式。

在这种模式下宿主机上的所有虚拟机是可以相互通信的，但虚拟机和真实的网络（物理机网络）是被隔离开的。

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/1111.png) 

![1111](F:\myblog\typora-user-images\linux\1111.png)

#### hosts文件

hosts是一个文本文件，用来记录IP和主机名的映射关系

## 目录结构

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240103190541275.png) 

![image-20240103190541275](F:\myblog\typora-user-images\linux\image-20240103190541275.png) 

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/derictory.png) 

![image-20240103190541275](F:\myblog\typora-user-images\linux\derictory.png) 

## vim简单介绍

[精通 VIM ，此文就够了 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/68111471)

## 常用命令

#### shutdown关机

```shell
shutdown -h now
```

#### reboot重启

```shell
reboot
```

#### sync将内存数据同步到磁盘

```shell
sync
```

## 用户管理

#### gorupadd组的创建

```shell
groupadd xxx		#组名为xxx
					#存放位置是/etc/group
```

#### useradd添加用户

```shell
useradd 用户名  #位置在/home/用户名，使用cd命令直接跳转到该目录
useradd -g 组名 用户名 	#在对应组下面创建用户
```

#### passwd设置密码

```shell
passwd 用户名  #给指定用户设置密码
```

- 1、删除用户

```shell
userdel -r 用户名
```

- 2、删除用户但是保留目录

```shell
urerdel 用户名
```

#### id查询用户信息

```shell
id 用户名
```

```shell
$ id kobayashi 
$ uid=1000(kobayashi) gid=1000(kobayashi) 组=1000(kobayashi),4(adm),24(cdrom),27(sudo),30(dip)
  uid:用户 gid:用户主组 groups:用户主组及其附属组
```

#### hostname主机名

```shell
/etc/hostname			#该文件中存放了主机名，可以通过修改该文件修改主机名
```

#### su切换用户

```shell
su - 用户名
```

#### 显示登录到系统的账户

```shell
who am i
```
#### chmod修改权限

```shell
  rwx | rwx | rwx
Owner  Group  Other Users
```

权限分为三级 : 文件所有者（Owner）、用户组（Group）、其它用户（Other Users）。

只有文件所有者和超级用户可以修改文件或目录的权限。可以使用绝对模式（八进制数字模式），符号模式指定文件的权限   

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/rwx-standard-unix-permission-bits.png) 

![](F:\myblog\typora-user-images\linux\rwx-standard-unix-permission-bits.png) 

#### chown修改所有者

```shell
chown 用户组 文件名
chown -R 用户组 目录		#-R表示递归
```

#### chgrp修改文件/目录所在组

```shell
chgrp 组名 文件名
chgrp -R 组名 目录
```

#### usermod修改用户所在组

```shell
usermod -g 新目录 用户名
```



## 帮助指令

- man获得帮助信息

```shell
man [命令或配置文件]
```

例如：![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240116221428879.png)

![image-20240116221428879](F:\myblog\typora-user-images\linux\image-20240116221428879.png)

- help

```shell
help 命令
```

例如：![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240116222422960.png)

![image-20240116222422960](F:\myblog\typora-user-images\linux\image-20240116222422960.png)

#### 二者区别

- help 是内部命令的帮助，比如cd 
- man 是外部命令的帮助，比如ls

**内部命令**：内部命令实际上是shell程序的一部分，其中包含的是一些比较简单的Linux系统命令，常驻内存，写在bashy源码里面，其执行速度比外部命令快，因为解析内部命令shell不需要创建子进程。比如：exit，history，cd，echo等。

**外部命令**：外部命令是Linux系统中的实用程序部分，因为实用程序的功能通常都比较强大，所以其包含的程序量也会很大，<font color='red'>在系统加载时并不随系统一起被加载到内存中，而是在需要时才将其调用内存</font>。通常外部命令的实体并不包含在shell中，<font color='red'>但是其命令执行过程是由shell程序控制的</font>。shell程序管理外部命令执行的路径查找、加载存放，并控制命令的执行。外部命令是在bash之外额外安装的，通常放在/bin，/usr/bin，/sbin，/usr/sbin......等等。可通过“echo $PATH”命令查看外部命令的存储路径，比如：ls、vi等。
**用type命令可以分辨内部命令与外部命令**

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240116230424184.png) 

![image-20240116230424184](F:\myblog\typora-user-images\linux\image-20240116230424184.png) 

## 文件目录类
#### pwd 显示当前目录（绝对路径）

```shell
pwd
```

#### tree 梳理目录结构

```shell
tree /opt			#显示/opt目录的所有文件结构
tree /opt -L 1		#只显示一级目录
```

#### ls  查看指定目录信息

常用：ls -alh

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240127151606753.png) 

![image-20240127151606753](F:\myblog\typora-user-images\linux\image-20240127151606753.png) 

#### **mkdir **创建目录（默认一级目录）

```shell
mkdir [选项] 要创建的目录
```

- -p 创建多级目录

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/image-20240116231858695.png) 

![image-20240116231858695](F:\myblog\typora-user-images\linux\image-20240116231858695.png) 

#### **rm **删除目录

```shell
rm -rf 要删除的目录/文件
```

- -i 删除前逐一询问确认。

- <font color='red'>-f （force 强制）</font>：即使原档案属性设为唯读，亦直接删除，无需逐一确认。
- <font color='red'>-r（recursive 递归）</font>：将目录及以下之档案亦逐一删除。

#### touch 修改文件

用于修改文件或者目录的时间属性，包括存取时间和更改时间。若文件不存在，系统会建立一个新的文件。

```shell
touch 文件名                   #创建一个空文件
touch testfile                #修改文件"testfile"的时间属性为当前系统时间 
```

#### cp 复制文件

```shell
cp [选项] 源文件 目标文件
```
- <font color='red'> -r （recursive 递归）：若要复制的是文件夹则要加上-r 选项表示将该文件夹里所以文件递归复制到新文件夹</font>
- -a：此选项通常在复制目录时使用，它保留链接、文件属性，并复制目录下的所有内容。其作用等于 dpR 参数组合。
- <font color='red'>-d：复制时保留链接。这里所说的链接相当于 Windows 系统中的快捷方式。</font>
 - -i 或  --interactive ：在复制前提示确认，如果目标文件已存在，则会询问是否覆盖，回答 **y** 时目标文件将被覆盖。
 - <font color='red'> -u 或  --update ：仅复制源文件中更新时间较新的文件。</font>
- -v 或  --verbose：显示详细的复制过程。
- -p 或 --preserve：保留源文件的权限、所有者和时间戳信息。
- <font color='red'> -f 或 --force：强制复制，即使目标文件已存在也会覆盖，而且不给出提示。</font>

#### mv 为文件或目录改名、或将文件或目录移入其它位置。

目标目录与原目录一致，指定了新文件名，效果就是仅仅重命名。

```shell
mv  /home/ffxhd/a.txt   /home/ffxhd/b.txt    
```

目标目录与原目录不一致，没有指定新文件名，效果就是仅仅移动。

```shell
mv  /home/ffxhd/a.txt   /home/ffxhd/test/ 
```

目标目录与原目录一致, 指定了新文件名，效果就是：移动 + 重命名。

```shell
mv  /home/ffxhd/a.txt   /home/ffxhd/test/c.txt
```

#### cat 用于连接文件并打印到标准输出设备上

```shell
cat [选项] 要查看的文件
```

- -n ：显示行号

将源文件的内容输入目标文件里 (替换)
```shell
cat [选项] 源文件 > 目标文件
```

将源文件的内容添加到目标文件末尾

```shell
cat [选项] 源文件 >> 目标文件
```

将 "Hello,world"输入到hello.txt文件里

```shell
echo "Hello, World!" | cat > hello.txt
echo "This is a new line" | cat >> hello.txt   #在末尾添加
```

-  \> 表示输出重定向
-  \>\> 表示追加

#### 输入输出重定向

输出重定向在cat已经介绍了

输入重定向：输入重定向将文件的内容重定向到命令

```shell
$ wc < test.txt
2	11	60
```

wc命令可以对对数据中的文本进行计数。默认情况下，它会输出3个值：

- 文本的行数
- 文本的词数
- 文本的字节数

```shell
$ wc << EOF
> test string 1
> test string 2
> test string 3
> EOF
3	9	42
```

**综合应用**

向test.sh中追加

7777
8888

```shell
$ cat << EOF >>test.sh
> 7777
> 8888
> EOF
```

- EOF是END Of File的缩写,表示自定义终止符

#### 管道

管道命令`|`仅能处理前一个命令传来的**标准输出**信息，而对于标准错误信息并没有直接处理能力

Linux系统会同时运行管道的几个命令，在系统内部将它们连接起来。在第一个命令产生输出的同时，输出会被立即送给第二个命令，以此类推。数据
传输不会用到任何中间文件或缓冲区。

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/v2-08aaea8db0fed11502948e77dc9c14cf_r.jpg) 

![img](F:\myblog\typora-user-images\linux\v2-08aaea8db0fed11502948e77dc9c14cf_r.jpg) 

在每个管道后面接的第一个数据必定是**命令**，而且**这个命令必须要能够接受标准输出的数据**才行，这样的命令才可为管道命令。例如`less`、`grep`、`sed`、`awk`等都是可以接受标准输入的管道命令，而`ls`、`cp`、`mv`就不是管道命令，因为它们并不会接受来自stdin的数据。

#### more 基于vi的文本过滤器

```shell
$ more 要查看的文件 
```

| 空格键（space） | 代表向下翻一页     |
| ----------- | :----------------- |
| **Enter**   | **代表向下翻一行** |
| **q**       | **退出**           |
| **b**       | **向前查看一屏**   |
| **=** | **输出当前行的行号** |

#### less 分屏查看文件内容

less在显示文件内容时，并不是一次将整个文件加载后才显示，而是根据显示需要加载内容，对于大型文件有较高的效率

```shell
$ less 要查看的文件
```

| 空白键           | 向下翻一页                                               |
| ---------------- | -------------------------------------------------------- |
| **[ pagedown ]** | **向下翻动一页**                                         |
| **[ pageup ]**   | **向上翻动一页**                                         |
| **/ 字符串**     | **向下搜寻【字串】的功能；n表示向下查找；N表示向上查找** |
| **q**            | **退出less程序**                                         |

例：ps查看进程信息并通过less分页显示

```
ps -ef |less
```

#### $符的作用

在Linux中，使用`$`符号可以引用变量的值。当你希望输出变量的值时，需要在变量名前加上`$`符号。这样，`echo`命令会将变量的值替换为实际的内容，并输出到终端。

#### echo输出内容到控制台

```shell
echo [选项] [输出内容]
```

- 输出文本：

  ```shell
  echo "Hello, World!"
  ```

- 输出变量的值：

  ```shell
  name="John"
  echo $name
  ```

- 输出到文件：

  ```shell
  echo "Hello, World!" > output.txt
  echo "Additional line" >> output.txt   #追加
  ```

- 输出命令执行结果：

  ```shell
  echo $(ls)
  ```

- 使用\ 输出一些特殊字符，例如双引号、单引号、反斜杠等 

  ```shell
  echo "This is a \"quoted\" text."
  ```


#### head显示文件开头部分内容（默认10行   ）

```shell
head [参数] [文件]
```

- -n<行数> 显示的行数。

#### tail显示文件尾部部分内容，或实时监控文件

```shell
tail [参数] [文件]
```

- -n<行数> 显示的行数。
- -f 循环读取

例：跟踪名为 notes.log 的文件的增长情况

```shell
tail -f notes.log
```

此命令显示 notes.log 文件的最后 10 行。当将某些行添加至 notes.log 文件时，tail 命令会继续显示这些行。 显示一直继续，直到您按下（Ctrl-C）组合键停止显示。

#### ln链接

- **软链接（相当于快捷方式）**

```shell
 ln -s [源文件或目录][目标文件或目录]
```

- **硬链接（无法对目录进行链接）**

  以文件副本的形式存在。但不占用实际空间。

```shell
ln [源文件] [目标文件]
```

#### history查看历史命令

存储于 ~/.bash_history文件中

```shell
history    		# 查看所有执行过的指令
history 10  	# 查看最近执行的10条指令
!n          	# 再次执行历史中第n号指令
```

## 时间日期类

#### data 显示当前日期

```shell
data    	
```

结果：2024年 01月 26日 星期五 21:58:56 CST

```shell
date +%F
```

结果：2024-01-26

## 搜索查找类

#### find查找文件

- 按文件名查找

```shell
find /home -name *.txt  	#查找/home下面的所有txt文件
```

- 按拥有者查找

```shell
find /home -user kobayashi	#在/home下查找用户名为kobayashi的文件
```

- 按大小查找

  **+表示大于 ；—表示小于 ；单位有k，M，G**

```shell
find /home -size +200M		#在/home目录下查找大小大于200Mb的文件
```

#### locate查找文件

在ubuntu上先使用命令安装软件包

```shell
sudo apt install plocate
```
```shell
locate 文件名
```
![image-20240127152808195](F:\myblog\typora-user-images\linux\image-20240127152808195.png) 


#### find与locate的区别

- 使用find指令时从硬盘上查找文件
- locate命令使用数据库来定位文件且在使用前要更新数据库

```shell
sudo update
```

#### grep过滤查找

- 只保留包含hello的行
- -n表示显示匹配行以及行号
- -v（invert-match）表示查找不匹配的

```shell
cat /home/a.txt | grep -n "hello"	#（配合管道"|"使用）
grep -n "hello" /home/a.txt		#
```

## 压缩和解压类

#### 只能压缩文件

- **.gz文件（注意解压以及压缩均会将原文件删除）**

```shell
gzip /home/hello.txt		#将/home下的hello.txt文件进行压缩
gunzip /home/hello.txt.gz	#将/home下的hello.txt.gz文件进行解压
```

#### 文件以及目录均可压缩

- **.zip文件**

**压缩：**

```shell
zip [选项] 压缩包名 源文件或源目录列表

zip hello.zip /home/hello.txt	#将/home/hello.txt文件压缩到当前文件夹下并命名为hello.zip
zip -r myhome.zip /home 		#将/home下所有文件压缩命名为myhome.zip
```

-r 递归压缩即压缩目录

-P 使用指定密码加密（注意P为大写）但是密码可能会被其他用户通过查看历史命令窥探到

通常使用-e来对文件进行加密，这样更加安全

**解压：**

```shell
unzip [选项] 缩包名地址

unzip -d /home/myhome.zip		#将myhome.zip解压缩
```

-o 不必先询问用户，unzip执行后覆盖原有文件。

- **.tar.gz文件**

```shell
tar [选项] [归档文件] [文件或目录...]
```

- `-z`: 表示使用gzip压缩算法进行压缩或解压缩。英文单词是"gzip"。

- `-x`: 表示提取（解压缩）归档文件。英文单词是"extract"。

- `-v`: 表示在命令执行过程中显示详细信息。英文单词是"verbose"。

- `-f`: 表示指定压缩后的文件名。英文单词是"file"。

- `-c`: 表示创建归档文件。英文单词是"create"。

  ​		**归档就是把一堆文件和目录放到一个新的文件里**

- `-C(大写)`：切换到指定的目录，然后执行操作。

案例一：

压缩多个文件，将/home/test1.txt 和 /home/test2.txt 压缩成 test.tar.gz

```shell
tar -zcvf test.tar.gz /home/test1.txt /home/test2.txt
```

案例二：

将/home/test.tar.gz解压到/opt目录下

```shell
tar -zxvf /home/test.tar.gz -C /opt 	#注意C是大写的
```

## 任务调度

#### crond定时任务

系统调度文件

```shell
vim /etc/crontab		#这是一个系统范围的crontab文件，通常用于包含系统级别的定时任务。
```

用户调度文件

```shell
crontab -e				
```
格式：

- crontab表达式，前面四项的关系之间为and的关系，需要同时满足才能执行；
  但星期那一项与前面月份日期是or的关系，只需满足其一即执行；

```shell
 *   *   *   *   * command
分钟 小时 日期 月份 星期 command
```

| 特殊符号 | 含义                                                         |
| -------- | ------------------------------------------------------------ |
| *        | 所有时间                                                     |
| ,        | 代表不连续的时间，0 8,12,16 * * * 代表每天的8点0分，12点0分，16点0分 |
| -        | 代表连续的时间范围，0 8-12 * * * 代表每天的8-12点            |
| */n      | 代表多久执行一次，* */2 * * * 代表每隔2h执行一次             |

终止任务调度

```shell
crontab -r
```

列出当前有哪些任务

```shell
crontab -l
```

重启任务调度

```shell
service crond restart
```

#### at单次任务

**基本介绍：**

**at命令是一次性计划任务，atd在后台每60秒检查作业队列，有作业时检查其时间如果匹配上就运行。**

- 检测atd是否运行

```shell
kobayashi@LEGION:~$ ps -ef | grep atd
kobayas+    4758    4744  0 15:03 pts/0    00:00:00 grep --color=auto atd
```

- **基本语法**

```shell
at [选项] [日期时间]
```

- **选项**

```shell
-f：指定包含具体指令的任务文件
-q：指定新任务的队列名称
-l：显示待执行任务的列表
-d：删除指定的待执行任务
-m：任务执行完成后向用户发送 E-mail
```

- **时间**

```shell
at now + 3 minutes
或者具体时间
at 04:00 2024-02-19
```

- 查看当前任务以及删除任务

```shell
kobayashi@LEGION:~$ atq
4	Mon Feb 19 04:00:00 2024 a kobayashi
kobayashi@LEGION:~$ atrm 4
```

## 设备分区

#### IDE硬盘与SCSI硬盘

- **IDE硬盘驱动器标识符为`hdx～`**

- **SCSI硬盘标识符为`sdx~`**

  其中x为盘号（a为基本盘，b为基本从属盘，c为辅助主盘，d为辅助从属盘）

  ～代表分区（前4个分区用数字1-4表示，代表主分区或扩展分区，往后都是逻辑分区）

#### 主分区、扩展分区、逻辑分区

- **主分区**

  主分区也就是包含操作系统启动所必需的文件和数据的硬盘分区，要在硬盘上安装操作系统，则该硬盘必须得有一个主分区。

- **扩展分区**

  扩展分区也就是除主分区外的分区，但它不能直接使用，必须再将它划分为若干个逻辑分区才行。

- **逻辑分区**

  逻辑分区也就是我们平常在操作系统中所看到的D、E、F等盘。

- **新硬盘建立分区顺序**

  建立主分区→建立扩展分区→建立逻辑分区→激活主分区→格式化所有分区。

#### lsblk查看分区

```shell
$lsblk
NAME        MAJ:MIN			RM  		SIZE 	RO 		TYPE 	MOUNTPOINTS
设备名称  主设备号和次设备号  是否可移动设备	大小  只读标志  设备类型  挂载点
```

#### 挂载分区

**一、增加硬盘**

在电脑上插入新的硬盘或U盘

**二、对硬盘进行分区**

1、找到新的硬盘

```shell
$lsblk			
```

2、sad为硬盘名

```shell
$fdisk /dev/sda
```

3、输入n	（add a new partition）按照提示继续往后做，最后w退出，不保存用q

![fidik](/media/kobayashi/新加卷/myblog/typora-user-images/linux/fdisk.png)

![fdisk](F:\myblog\typora-user-images\linux\fdisk.png) 

**三、进行格式化**

```shell
mkfs -t ext4 /dev/sda1
```

**四、进行挂载**

```shell
mount /dev/sda1 /xxx
```

- xxx为要挂载的目录名字
- sda1为分区

**挂载之后xxx目录下的文件都存放在新加的硬盘上**

**五、卸载**

```shell
umount /dev/sda1	#卸载会导致该挂载点中的内容不再可见，但并不会删除这些内容
```

#### 开机默认挂载

修改配置文件

```shell
vim /etc/fstab
```

默认格式

```shell
UUID=xxxxxxxxxxxxxxx		/home   		ext4    	defaults    0    2
要挂载的文件系统的UUID	指定挂载点的目标目录	文件系统的类型	挂载选项（默认）0表示不应该检查，1表示应该在启动时检查，2和1一样但如果检查失败，系统不会强制进行修复。
```

#### 查看文件系统磁盘使用情况

```shell
df -h 		#-h表示带计量单位
df -h /opt	#指定/opt目录
```

#### 查看目录的磁盘使用情况

```shell
du -h --max-depth=1 /opt	#查询/opt目录的磁盘占用情况，深度为1
```

## 进程管理

#### ps查看当前进程

**ps是显示瞬间进程的状态，并不动态连续，如果想对进程进行实时监控应该用top命令（ubuntu中可以直接打开系统监视器查看）**

```shell
ps [选项]
```

| 选项 |             功能             |
| :--: | :--------------------------: |
|  -A  |        列出所有的进程        |
|  -w  |  显示加宽可以显示较多的资讯  |
| -au  |       显示较详细的资讯       |
| -aux | 显示所有包含其他使用者的进程 |

- USER: 行程拥有者
- PID: pid
- %CPU: 占用的 CPU 使用率
- %MEM: 占用的内存使用率
- VSZ: 占用的虚拟内存大小
- RSS: 占用的存储空间大小
- TTY: 该进程是在那个终端机上面运行 (minor device number of tty)
- STAT: 该行程的状态:
  - D:  无法中断的休眠状态 (通常 IO 的进程)
  - R: 正在执行中
  - S: 休眠状态
  - T: 暂停执行
  - Z: 不存在但暂时无法消除，造成zombie(疆尸)程序的状态
  - W: 等待状态，等待内存分配
  - <: 高优先级的行程
  - N: 低优先级的行程
  - L: 有记忆体分页分配并锁在记忆体内 (实时系统或捱A I/O)
- START: 进程开始时间
- TIME: 执行的时间
- COMMAND:所执行的指令

#### zombie进程

概念：子进程死亡后，它的父进程会接收到通知去执行一些清理操作，如释放内存之类。然而，若父进程并未察觉到子进程死亡，子进程就会进入到“ *僵尸(zombie)*”状态。

https://cloud.tencent.com/developer/article/1903722

#### 查找进程

以全格式显示当前所有进程 

```shell
ps -ef | grep 进程关键字		#-e显示所有进程 ，-f全格式
```

#### pstree以树状图显示进程

```shell
pstree -p		#-p表示显示pid
```

#### kill/killall终止进程

**注：Linux 的 kill 命令是向进程发送信号，kill 不是杀死的意思，-9 表示无条件退出，但由进程自行决定是否退出，这就是为什么 kill -9 终止不了系统进程和守护进程的原因。**

```shell
kill [选项] 进程号		#一般用-9来强行终止进程
killall 进程名称
```

### 服务管理

#### 守护进程

定义：脱离于终端并且在后台运行的进程，一般为服务进程。

守护进程经常以超级用户（root）权限运行，因为它们要使用特殊的端口（1-1024）或访问某些特殊的资源。

#### 运行级别

| init运行级别 | 作用                                                         |
| :----------: | ------------------------------------------------------------ |
|      0       | 系统默认运行级别不能设置为0，否则无法正常启动系统（一开机就自动关机） |
|      1       | 也称为救援模式，root权限，用于系统维护，禁止远程登陆，类似Windows下的安全模式登录。 |
|      2       | 多用户状态，没有NFS网络支持。                                |
|      3       | 完全的多用户状态，有NFS，登陆后进入控制台命令行模式。        |
|      4       | 系统未使用（保留）                                           |
|      5       | 登陆后进入图形GUI模式或GNOME、KDE图形化界面，如X Window系统。 |
|      6       | 系统正常关闭并重启，默认运行级别不能设为6，否则无法正常启动系统。 |

1、服务器通常为3,个人电脑为5

```shell
$ runlevel
N 5
```

2、修改系统启动时默认的运行级别

```shell
sudo systemctl set-default multi-user.target		#重启后就是命令行界面
 
sudo systemctl set-default graphical.target			#默认运行为图型界面

sudo systemctl start lightdm	/    init 5		#之后如果想回到图形界面
```

#### systemctl命令

systemctl命令管理的服务在/usr/lib/systemd/system中可以查看

```shell
systemctl list-unit-files 		#查看服务开机自启状态,可结合管道查询指定服务
systemctl enable 服务名		#设置服务开机自启动
systemctl disable 服务名		#关闭服务开机自启动
```

#### chkconfig设置服务在各个运行级别下是否自启动

chkconfig用于一些较早的Linux发行版，现在都采用systemctl了

```shell
chkconfig --level x 服务名 on/off
```

### 动态监控进程

#### top动态监控

```shell
top 		#默认3秒更新
```

| 操作    | 功能                |
| ------- | ------------------- |
| P(大写) | 以cpu使用率排序     |
| M       | 以内存排序          |
| N       | 以pid排序           |
| q       | 退出                |
| u       | 查找特定用户        |
| k       | 结束对应pid的进程-9 |

### 监控网络状况

#### netstat查看网络情况

```shell
netstat -anp
```

- -an 按一定顺序排列输出
- -p显示哪个进程在调用

## 日志管理rsyslogd

```shell
#系统日志文件存放在/var/log中

#存放了验证授权方面的信息，比如ssh登录，su切换，sudo授权等
/var/log/auth.log 		#ubuntu
/var/log/secure			#centeros

#存放了系统启动日志
/var/log/boot.log		

#存放了系统大部分重要信息
/var/log/messages

```

#### 配置文件 /etc/rsyslog.conf

```
/etc/rsyslog.d/50-default.conf
```

重启服务

```shell
systemctl restart rsyslog.service
```

文件格式* . * 其中第一个 * 代表日志类型，第二个 * 代表日志级别，如果使用*代表全部

| 日志类型（部分）     | 说明                                |
| -------------------- | ----------------------------------- |
| auth                 | pam产生的日志                       |
| authpriv             | ssh、ftp等登录信息的验证信息        |
| corn                 | 时间任务相关                        |
| kern                 | 内核                                |
| lpr                  | 打印                                |
| mail                 | 邮件                                |
| mark(syslog)-rsyslog | 服务内部的信息，时间标识            |
| users                | 用户程序产生的相关信息              |
| uucp                 | unix to nuix copy主机之间相关的通信 |

**日志级别从高到低排序 ERROR、WARN、INFO、DEBUG，什么都不记录用none**

#### 日志文件

格式：

```
事件产生的时间		产生事件的服务器主机名		产生事件的服务名或程序名	事件的具体信息
```

#### 日志轮替

#### logrotate配置文件

logrotate 的配置文件 /etc/logrotate.conf 为全局配置文件。

也可以把某个日志文件轮替规则单独指定，位置在/etc/logrotate.d

```shell
# rotate log files weekly
weekly

# use the adm group by default, since this is the owning group
# of /var/log/syslog.
su root adm

# keep 4 weeks worth of backlogs
rotate 4

# create new (empty) log files after rotating old ones
create

# use date as a suffix of the rotated file
dateext

# uncomment this if you want your log files compressed
#compress

# packages drop log rotation information into this directory
include /etc/logrotate.d

# system-specific logs may also be configured here.

```

#### logrotate.d单独指定

格式：

```shell
/var/log/ppp-connect-errors {
	weekly
	rotate 4
	missingok
	notifempty
	compress
	nocreate
}
```

| 参数                 | 说明                               |
| -------------------- | ---------------------------------- |
| daily/weekly/monthly | 日志轮替周期天/周/月               |
| rotate 数字          | 保留日志个数                       |
| missingok            | 如果日志不存在忽略该日志的警告信息 |
| notifempty           | 如果日志为空则不进行日志轮替       |
| compress             | 日志轮替时对旧日志进行压缩         |
| ......               | .......                            |

详细信息：https://c.biancheng.net/view/1106.html

## 环境变量

**概念：**

**bash shell用一个叫作环境变量（environment variable）的特性来存储有关shell会话和工作环境的信息（这也是它们被称作环境变量的原因）**
**这项特性允许你在内存中存储数据，以便程序或shell中运行的脚本能够轻松访问到它们。这也是存储持久数据的一种简便方法。**

#### 全局/局部环境变量

**全局环境变量对于shell会话和所有生成的子shell都是可见的。局部变量则只对创建它们的shell可见**

查看全局变量

```shell
printenv		#显示全局变量
```

局部变量没有命令单独展示

```shell
set				#命令会显示出全局变量、局部变量以及用户定义变量。它还会按照字母顺序对结果进行排序
```

#### 局部用户定义变量

export将自定义的变量设为全局变量

```shell
$ my_variable="I am Global now"
$ export my_variable
$ echo $my_variable
I am Global now
$ bash
$ echo $my_variable
I am Global now
$ exit
exit
$ echo $my_variable
I am Global now
```

unset删除全局变量

```shell
$ echo $my_variable
I am Global now
$ unset my_variable
$ echo $my_variable

$
```

注意：

- 修改子shell中全局环境变量并不会影响到父shell中该变量的值，子shell甚至无法使用export命令改变父shell中全局环境变量的值。
- 如果你是在子进程中删除了一个全局环境变量，这只对子进程有效。该全局环境变量在父进程中依然可用

#### 设置path环境变量

当你在shell命令行界面中输入一个外部命令时，shell必须搜索系统来找到对应的程序。

PATH环境变量定义了用于进行命令和程序查找的目录，PATH中的目录使用冒号分隔。

```shell
$ echo $PATH
/home/kobayashi/.cargo/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/snap/bin:/home/kobayashi/go/gopath/bin
```

如果命令或者程序的位置没有包括在PATH变量中，那么如果不使用绝对路径的话，shell是没法找到的，这时会产生错误信息

```bash
$ myprog
-bash: myprog: command not found
```

#### 环境变量持久化

之前几节设置的环境变量只能持续当前shell或者下次开机前

最好是在/etc/profile.d目录中创建一个以.sh结尾的文件。把所有新的或修改过的全局环境变量设置放在这个文件中。

在大多数发行版中，存储个人用户永久性bash shell变量的地方是$HOME/.bashrc文件

# shell脚本编程

参考这本书：这本书写的非常详细

![shell](/media/kobayashi/新加卷/myblog/typora-user-images/linux/shell.png) 

## 构建基本脚本

### 格式

#!/bin/bash 是为了告诉shell这时一个脚本文件，除此之外所有的#语句都是注释

```shell
#!/bin/bash
...........
```

### 命令替换

var1=$(pwd)

```shell
#!/bin/bash
testing=$(date)
echo "The date and time are: " $testing
.......................................................
$ ./test
The date and time are: Mon Jan 31 20:23:25 EDT 2014
```

### 数学运算

- bash shell数学运算符只支持整数运算

- 在bash中，在将一个数学运算结果赋给某个变量时，可以用美元符和方括号  $[ operation ]  将数学表达式围起来

```shell
$ cat test7
#!/bin/bash
var1=100
var2=50
var3=5
var4=$[$var1 / ($var2 - $var3)]
#var4=$(($var1 / ($var2 - $var3)))		也可以使用双括号框起来
echo The final result is $var4
..................................................................................
$ ./test7
The final result is 2
```

如果想在脚本中使用浮点数运算可以使用bc计算器

#### bc计算器

在bash中直接输入bc即可打开

```shell
kobayashi@LEGION:~$ bc
bc 1.07.1
Copyright 1991-1994, 1997, 1998, 2000, 2004, 2006, 2008, 2012-2017 Free Software Foundation, Inc.
This is free software with ABSOLUTELY NO WARRANTY.
For details type `warranty'. 
12*3.7
44.4
quit
```

在脚本中结合命令替换就能使用bc计算浮点数

```shell
$ cat test12
#!/bin/bash
var1=10.46
var2=43.67
var3=33.2
var4=71
var5=$(bc << EOF
scale = 4					#保留4位小数
a1 = ( $var1 * $var2)
b1 = ($var3 * $var4)
a1 + b1
EOF
)
echo The final answer for this mess is $var5
```

#### 退出状态码

默认情况下，shell脚本会以脚本中的最后一个命令的退出状态码退出。

- 成功结束的命令

```shell
$ echo $?
0
```

- 无效命令

```shell
$ echo $?
127
```

有些时候可以根据状态码来判断命令错误类型，不过大部分情况下命令出错bash会自动显示问题所在

| 状态码 | 描述                       |
| ------ | -------------------------- |
| 0      | 命令成功结束               |
| 1      | 一般性未知错误             |
| 2      | 不适合的shell命令          |
| 126    | 命令不可执行               |
| 127    | 没找到命令                 |
| 128    | 无效的退出参数             |
| 128+x  | 与Linux信号x相关的严重错误 |
| 130    | 通过Ctrl+C终止的命令       |
| 255    | 正常范围之外的退出状态码   |

可以使用exit命令来人为设置退出码（之后会用到）

```shell
#!/bin/bash
var=10
exit $var
$ ./test
$ echo $?
10
```

## 结构化命令

### if - then - elif / else

if语句会运行条件命令。如果该命令的退出状态码是0（该命令成功运行），位于then部分的命令就会被执行。如果该命令的退出状态码是其他值，then部分的命令就不会被执行，此时如果有else将会去执行else的命令。

和c语言中的if-else语句非常像，当然也可以使用if嵌套，这里不过多赘述。

- 以if开头以fi结尾
- then和else这些命令是一个代码块，均会执行

使用else语句格式如下

```shell
if command1
then
commands
else
commands
fi
```

使用elif可以对条件进行嵌套查询

记住在 elif 语句中，紧跟其后的 else 语句属于 elif 代码块。它们并不属于之前的if-then代码块。

```shell
if command1
then
commands
elif command2
then
more commands
else
commands
fi
```

#### test命令

如果想在脚本中和其他语言一样使用if else，可以使用test命令，即if命令不检测状态码

如果test命令中列出的条件成立，test命令就会退出并返回退出状态码0

如果不写test命令的condition部分，它会以非零的退出状态码退出，并执行else语句块

```shell
#if test condition
if [ condition ]			#在代码中我们通常使用[]来代替test命令，condition与[]之间一定要用空格隔开
then
commands
fi
```

test命令可以判断三类条件：
数值比较、 字符串比较、文件比较

- 数值比较

![截图 2024-03-01 15-01-14](/media/kobayashi/新加卷/myblog/typora-user-images/linux/num_cmp.png) 

举例：

```shell
#!/bin/bash
value1=10						#
if [ $value1 -gt 5 ]
then
echo "The test value $value1 is greater than 5"
fi
```

- 字符串比较

![截图 2024-03-01 15-03-34](/media/kobayashi/新加卷/myblog/typora-user-images/linux/string_cmp.png) 

注意使用 < > 时要加上转义符\不然会被识别为重定向符

举例：

```shell
#!/bin/bash
val1=baseball
val2=hockey
if [ $val1 \> $val2 ]
then
echo "$val1 is greater than $val2"
else
echo "$val1 is less than $val2"
fi
```

- 文件比较

![截图 2024-03-01 15-19-53](/media/kobayashi/新加卷/myblog/typora-user-images/linux/file_cmp.png) 

```shell
#!/bin/bash
jump_directory=/home/arthur
if [ -d $jump_directory ]
then
echo "The $jump_directory directory exists"
cd $jump_directory
ls
else
echo "The $jump_directory directory does not exist"
fi
```

#### 复合条件

if-then语句允许你使用布尔逻辑来组合测试（同c语言）。有两种布尔运算符可用：
[ condition1 ] && [ condition2 ]
[ condition1 ] || [ condition2 ]

#### 双括号

应用情况：较为复杂的数学运算

![截图 2024-03-01 15-36-44](/media/kobayashi/新加卷/myblog/typora-user-images/linux/double parenthesis.png) 

举例：

```shell 
#!/bin/bash
val1=10
if (( $val1 ** 2 > 90 ))
then
(( val2 = $val1 ** 2 ))
echo "The square of $val1 is $val2"
fi
```

### case命令

```shell
case variable in
pattern1 | pattern2) commands1;;
pattern3) commands2;;
*) default commands;;
esac
```

```shell
#!/bin/bash
case $(pwd) in
"/opt" | "/etc")
echo "Welcome, $USER"
echo "you are visit /opt or /etc";;
"/home")
echo "you are in /home";;
*)
echo "Sorry, you are miss";;
esac
```

### for循环（shell）

```shell
for var in list
do
commands
done
```

$test变量的值会在shell脚本的剩余部分一直保持有效。它会一直保持最后一次迭代的值（除非你修改了它）

使用" "可以很好的表示出我们想要的单词

```shell
#!/bin/bash
for test in "bei'jing" shanghai "nan jing"
do
	echo "The next 省份 is $test"
done
echo $test
$ ./test.sh
The next 省份 is bei'jing
The next 省份 is shanghai
The next 省份 is nan jing
nanjing
```

#### 向列表尾部添加值

```shell
list="beijing shanghai"
list=$list" nanjing"
echo $list
```

#### 从命令读取值

```shell
#!/bin/bash
file="province.txt"
for list in $(cat $file)
do echo $list
done
```

#### 内部字段分隔符IFS

IFS环境变量定义了bash shell用作字段分隔符的一系列字符。默认情况下，bash shell会将下列字符当作字段分隔符：

- 空格
- 制表符
- 换行符 

如果你想修改IFS的值，使其只能识别换行符，那就必须这么做：
IFS=$'\n'

```shell
#!/bin/bash
file="province"
IFS=$'\n'
for list in $(cat $file)
do echo "list"
done
.......................................................
这样就能够以行为单位输出省份
```

### for循环（c语言）

与c语言唯一不同点在于条件处要打两个括号

```shell
#!/bin/bash
for (( i=1; i <= 3; i++ ))
do
echo "The next number is $i"
done
$ ./test8
The next number is 1
The next number is 2
The next number is 3
```

还可以使用多个变量，但条件只能有一个

```shell
#!/bin/bash
for (( a=1, b=10; a <= 10; a++, b-- ))
do
echo "$a - $b"
done
```

### while

格式：

```shell
while test command
do
other commands
done
```

注意：while命令的关键在于所指定的test command的退出状态码必须随着循环中运行的命令而改变。如果退出状态码不发生变化， while循环就将一直不停地进行下去。

举例：

```shell
#!/bin/bash
test=10
while [ $test -gt 0 ]
do echo"this is a test$test"
test=$test-1
done
```

小提示：可以使用while true ; do进行死循环

### until命令

until命令和while命令工作的方式完全相反。until命令要求你指定一个通常返回非零退出状态码的测试命令。只有测试命令的退出状态码不为0，bash shell才会执行循环中列出的命令。一旦测试命令返回了退出状态码0，循环就结束了

格式：

```shell
until test commands
do
other commands
done
```

举例：

```shell
#!/bin/bash
test=0
until [ $test -gt 2 ]
do
echo "this is $test"
test=$[$test+1]
done
echo "the last num is $test"
..................................................................
$ ./test.sh 
this is 0
this is 1
this is 2
the last num is 3
```

最后说明：以上的这些循环同c语言一样是可以嵌套使用的。

### break与continue

break的用法同c语言

使用break n可以指定跳出几级循环（默认为1）

举例：

```shell
#!/bin/bash
for (( a = 1; a < 4; a++ ))
do
echo "Outer loop: $a"
for (( b = 1; b < 100; b++ ))
do
if [ $b -gt 4 ]
then
break 2
fi
echo "
Inner loop: $b"
done
done
```

continue用法类比于break

举例：

```shell
#!/bin/bash
for (( a = 1; a <= 5; a++ ))
do
echo "Iteration $a:"
for (( b = 1; b < 3; b++ ))
do
if [ $a -gt 2 ] && [ $a -lt 4 ]
then
continue 2
fi
var3=$[ $a * $b ]
echo "
The result of $a * $b is $var3"
done
done
$ ./test22
Iteration 1:
The result of 1 * 1 is 1
The result of 1 * 2 is 2
Iteration 2:
The result of 2 * 1 is 2
The result of 2 * 2 is 4
Iteration 3:
Iteration 4:
The result of 4 * 1 is 4
The result of 4 * 2 is 8
Iteration 5:
The result of 5 * 1 is 5
The result of 5 * 2 is 10
```
### 循环的输出重定向

shell会将for命令的结果重定向到文件output.txt中，而不是显示在屏幕上。

```shell
for file in /home/rich/*
do
if [ -d "$file" ]
then
echo "$file is a directory"
elif
echo "$file is a file"
fi
done > output.txt
```

## 处理用户输入参数

### 命令行参数

传递方式 :  运行脚本时直接写在后面，中间用空格隔开

```shell
./test.sh 10 20
```

#### $0~$9

脚本内使用$0~$9来代表第n个参数。

注意：有些shell当参数个数超过9个时必须在变量数字周围加上花括号，比如 ${10}，不然shell会将其当成$1和0来处理

```shell
#!/bin/bash
total=$[ ${10} * ${11} ]
echo The tenth parameter is ${10}
echo The eleventh parameter is ${11}
echo The total is $total
............................................................
$ ./test4.sh 1 2 3 4 5 6 7 8 9 10 11 12
The tenth parameter is 10
The eleventh parameter is 11
The total is 110
```

#### $@与$*

$@变量会将命令行上提供的所有参数当作同一字符串中的多个独立的单词。这样你就能够遍历所有的参数值，得到每个参数。

所以可以使用$@来遍历输入

$*变量会将命令行上提供的所有参数当作一个单词保存。这个单词包含了命令行中出现的每一个参数值。

基本上$*变量会将这些参数视为一个整体，而不是多个个体。

```shell
#!/bin/bash
# 循环遍历所有传入的参数
for num in "$@"
do
	echo "\$@的值为$num"
done
for num1 in "$*"
do
	echo "\$*的值为$num"
done
...................................................................
$ ./test.sh beijing shanghai
$@的值为beijing
$@的值为shanghai
$*的值为beijing shanghai

```

#### $#表示参数的个数

注：${!#}可以用来代表最后一个参数值。

```shell
#!/bin/bash
echo There were $# parameters supplied.
$ ./test8.sh
There were 0 parameters supplied.
$ ./test8.sh 1 2 4 
There were 3 parameters supplied.
```

#### $0

$0可用来表示程序名 如：test.sh

注意事项：命令会和脚本名混在一起，出现在$0参数中，所以要使用basename命令会返回不包含路径的脚本名

不然就使用bash test.sh也能得到想要结果

```shell
#!/bin/bash
name=$(basename $0)
echo The script name is: $name
$ bash /home/Christine/test.sh
The script name is: test.sh
$ ./test5b.sh
The script name is: test.sh
```

#### shift移动变量

```shell
#!/bin/bash
count=1
while [ -n "$1" ]
do
echo "Parameter #$count = $1"
count=$[ $count + 1 ]
shift
done
$ ./test.sh rich barbara katie 
Parameter #1 = rich
Parameter #2 = barbara
Parameter #3 = katie
$
```

### 处理选项

（这部分懒得记了.....以后用到了再说）

## read获取用户输入

### 基本读取

echo命令使用了-n选项。该选项不会在字符串末尾输出换行符，允许脚本用户紧跟其后输入数据，而不是下一行。这让脚本看起来更像表单。

```shell
#!/bin/bash
echo -n "Enter your name: "
read name
echo "Hello $name, welcome to my program. "
$ ./test21.sh
Enter your name: Rich Blum
Hello Rich Blum, welcome to my program.
$
```

也可以指定多个变量

如果变量数量不够，剩下的数据就全部分配给最后一个变量。

```shell
#!/bin/bash
read -p "Enter your name: " first last		#read中使用-p选项可以直接代替echo输出提示语句
echo "Checking data for $last, $first…"
..........................................................
$ ./test23.sh
Enter your name: Rich Blum
Checking data for Blum, Rich...
```

如果不指定参数，read命令会将它收到的任何数据都放进特殊环境变量REPLY中。通过$REPLY就可以得到输入值

#### 判断是否有输入

```shell
#!/bin/bash
if read -p "Enter something: " input && [ -n "$input" ]
then
    echo "You entered: $input"
else
    echo "No input provided within 5 seconds or input is empty."
fi
```

#### -t设置时间

可以用-t选项来指定一个计时器。-t选项指定了read命令等待输入的秒数。当计时器过期后，read命令会返回一个非零退出状态码。

结合if语句可以实现以下功能

```shell
#!/bin/bash
if read -t 5 -p "Please enter your name: " name
then
echo "Hello $name, welcome to my script"
else
echo
echo "Sorry, too slow! "
fi
............................................................
$ ./test25.sh
Please enter your name:
Sorry, too slow!
```

#### -s密码隐藏读取

-s选项可以避免在read命令中输入的数据出现在显示器上（实际上，数据会被显示，只是read命令会将文本颜色设成跟背景色一样）

综合案例：密码设置脚本

如果用户输入的密码不匹配，将会输出"passwd set failure, please try again"，然后继续循环，密码也不允许为空，让用户可以再次输入密码进行匹配。当密码匹配时，将输出"passwd set success"并跳出循环。如果在规定时间内没有输入密码，将输出"time out"继续循环,

```shell
#!/bin/bash
while true
do
if read -s -t 5 -p "Enter your passwd:" passwd
then
        echo
        if [ -n "$passwd" ]
        then
                read -s -t 5 -p "Enter your passwd again:" passwd1
                if [ $passwd = $passwd1 ]
                then
                        echo
                        echo "passwd set sucess,and your passwd is $passwd"
                        echo "please remember it"
                        break
                else
                        echo
                        echo "passwd set failure,please try again"
                        continue
                fi
        else
                echo "no input is provided"
        fi
else
        echo
        echo "time out"
        continue
fi
done
```

### 从文件中读取

每次调用read命令，它都会从文件中读取一行文本。当文件中再没有内容时，read命令会退出并返回非零退出状态码。

实现方法：对文件使用cat命令，将结果通过管道直接传给含有read命令的while命令

```shell
#!/bin/bash
count=1
cat test.txt | while read line
do
echo "Line $count: $line"
count=$[ $count + 1]
done
echo "Finished processing the file"
```

test.txt文件

```tex
The quick brown dog jumps over the lazy fox.
This is a test, this is only a test.
O Romeo, Romeo! Wherefore art thou Romeo?
```

结果：

```shell
$ ./test28.sh
Line 1: The quick brown dog jumps over the lazy fox.
Line 2: This is a test, this is only a test.
Line 3: O Romeo, Romeo! Wherefore art thou Romeo?
Finished processing the file
```

## 呈现数据

### 错误重定向

| 文件描述符 | 缩写   | 描述     |
| ---------- | ------ | -------- |
| 0          | STDIN  | 标准输入 |
| 1          | STDOUT | 标准输出 |
| 2          | STDERR | 标准错误 |

```shell
$ ls -al test test2 test3 badtest 2> test6 1> test7
$ cat test6
ls: cannot access test: No such file or directory
ls: cannot access badtest: No such file or directory
$ cat test7
-rw-rw-r-- 1 rich rich 158 2014-10-16 11:32 test2
-rw-rw-r-- 1 rich rich 0 2014-10-16 11:33 test3
```

### exec自定义重定向

可以用exec命令来给输出分配文件描述符。一旦将另一个文件描述符分配给一个文件，这个重定向就会一直有效，直到你重新分配

```shell
#!/bin/bash
exec 3>testout
echo "this should be stored in the testout" >&3
......................................................
运行脚本后：
cat testout
this should be stored in the testout
```

## 控制脚本

### 信号量

| 信号 | 值      | 描述                           |
| ---- | ------- | ------------------------------ |
| 1    | SIGHUP  | 挂起进程                       |
| 2    | SIGINT  | 终止进程                       |
| 3    | SIGOUT  | 停止进程                       |
| 9    | SIGKILL | 无条件终止进程                 |
| 15   | SIGTERM | 尽可能终止进程                 |
| 17   | SIGSTOP | 无条件停止进程，但不是终止进程 |
| 18   | SIGSTP  | 停止或暂停进程，但不终止进程   |
| 19   | SIGONT  | 继续运行停止的进程             |

默认情况下bash shell会忽略收到的任何SIGQUIT (3)和SIGTERM (15)信号（正因为这样，交互式shell才不会被意外终止）但bash shell会处理其他的信号。shell会将产生的信号交给由该shell启动的进程处理。

而shell脚本的默认行为是忽略这些信号，它们可能会不利于脚本的运行。要避免这种情况，你可以脚本中加入识别信号的代码，并执行命令来处理信号。

#### 中断Ctrl+C

在运行脚本时我们可以使用Ctrl+C产生SIGINT (2) 信号来终止进程。

#### 暂停Ctrl+Z

在进程运行期间暂停进程，而无需终止它。尽管有时这可能会比较危险（比如，脚本打开了一个关键的系统文件的文件锁），但通常它可以在不终止进程的情况下使你能够深入脚本内部一窥究竟。

Ctrl+Z会生产SIGSTP (18) 信号，这可以停止在shell中运行的任何进程，停止进程会让程序继续保留在内存中，并能从上次停止的位置继续运行

#### trap捕获信号

trap命令允许你来指定shell脚本要监看并从shell中拦截的Linux信号。如果脚本收到了trap命令中列出的信号，该信号不再由shell处理，而是交由本地处理

```shell
格式：
trap "commands" signals
```

每次使用Ctrl+C组合键，脚本都会执行trap命令中指定的echo语句，而不是处理该信号并允许shell停止该脚本。 

```shell
#!/bin/bash
trap "echo ' Sorry! I have trapped Ctrl-C'" SIGINT
echo This is a test script
count=1
while [ $count -le 5 ]
do
echo "Loop #$count"
sleep 1
count=$[ $count + 1 ]
done
echo "This is the end of the test script"
...............................................................
$ ./test1.sh
This is a test script
Loop #1
Loop #2
^C Sorry! I have trapped Ctrl-C
Loop #3
Loop #4
Loop #5
This is the end of the test script
$
```

#### 在脚本退出时捕获

要捕获shell脚本的退出，只要在trap命令后加上EXIT信号就行。

```shell
#!/bin/bash
trap "echo Goodbye..." EXIT
count=1
while [ $count -le 3 ]
do
echo "Loop #$count"
sleep 1
count=$[ $count + 1 ]
done
............................................
$ ./test2.sh
Loop #1
Loop #2
Loop #3
Goodbye...
$
```

#### 修改或移除捕获

重新使用带有新选项的trap命令就可以实现修改;

移除捕获则使用 trap -- 信号即可。

```shell
#!/bin/bash
trap "echo ' Sorry... Ctrl-C is trapped.'" SIGINT
count=1
while [ $count -le 5 ]
do
echo "Loop #$count"
sleep 1
count=$[ $count + 1 ]
done
trap "echo ' I modified the trap!'" SIGINT
#trap -- SIGINT
count=1
while [ $count -le 5 ]
do
echo "Second Loop #$count"
sleep 1
count=$[ $count + 1 ]
done
$
```

### 后台运行脚本

只需要在运行时后边加个&就行了，并且shell允许启动多个后台作业

```shell
./test.sh &
```

注意：当后台进程运行时，它仍然会使用终端显示器来显示STDOUT和STDERR消息。最好是将后台运行的脚本的STDOUT和STDERR进行重定向，避免这种杂乱的输出。

### 作业控制

#### jobs查看

可以使用jobs查看正在运行的作业，前台后台运行的都可以查看。

```shell
jobs		
```

#### bg重启停止的作业

要重启ctrl+z停止的作业，可用bg命令加上作业号，如果只有单个作业直接bg就行。

```shell
bg [n]
```

### 修改优先级

调度优先级是个整数值，从-20（最高优先级）到+19（最低优先级）。
默认情况下，bash shell以优先级0来启动所有进程

#### nice命令

提高test.sh的优先级

```shell
nice -10 ./test.sh
```

#### renice命令

注意

- 只能对属于你的进程执行renice；
- 只能通过renice降低进程的优先级；
- root用户可以通过renice来任意调整进程的优先级。如果想完全控制运行进程，必须以root账户身份登录或使用sudo命令

```shell
renice -n 10 -p 5055
#不写参数数名称也行
renice 10 5055
```

## shell脚本编程进阶

### 函数

#### 创建函数

<font color='red'>注意：函数名与{ }之间一定要有空格，否则会报错。</font>

**方式一：**

name属性定义了赋予函数的唯一名称。脚本中定义的每个函数都必须有一个唯一的名称

commands是构成函数的一条或多条bash shell命令。在调用该函数时，bash shell会按命令在函数中出现的顺序依次执行，就像在普通脚本中一样。

```shell
function name {
	commands
}
```

**方式二：**

类似于C语言中的函数

```shell
name() {
	commands
}
```

**举例**

```shell
#!/bin/bash
function func1 {
echo "This is an example of a function"
}
#func1()
#{
#	echo "this is an example of function"
#}
count=1
while [ $count -le 3 ]
do
	func1
	count=$[ $count + 1 ]
done
echo "This is the end of the loop"
func1
echo "Now this is the end of the script"
```

#### 函数值返回

对于较小的正数可以使用return返回状态码的方式得到，而较大的数值可以将其输出到变量中

**1、return(0~256)**

```shell
#!/bin/bash
function dbl {
	read -p "Enter a value: " value
	echo "doubling the value"
	return $[ $value * 2 ]
}
dbl
echo "The new value is $?"
....................................................
$ ./test
Enter a value: 100
doubling the value
The new value is 200
```

**2、通过函数输出复制(大于256的整数、浮点数或者字符串)**

正如可以将命令的输出保存到shell变量中一样，你也可以对函数的输出采用同样的处理办法。可以用这种技术来获得任何类型的函数输出，并将其保存到变量中

```shell
#!/bin/bash
function dbl {
	read -p "Enter a value: " value
	echo $[ $value * 2 ]
}
result=$(dbl)
echo "The new value is $result"
.................................................
$ ./test
Enter a value: 200
The new value is 400
```

#### 传递参数

直接在函数后面加参数即可,不过函数内的$于脚本中的$不是同一个变量

```shell
#!/bin/bash
function addem {
echo $[ $1 + $2 ]
}
value=$(addem 1 2)
va=$[ $1+$2 ]
echo $value
echo $va
.....................................
$ ./test.sh 10 20
3
30
```

#### 作用域

默认情况下，你在脚本中定义的任何变量都是全局变量。

函数内部想要使用局部变量只要在变量声明的前面加上local关键字就可以了,你甚至可以使用同名变量

```shell
#!/bin/bash
var1=10
var2=20
function addem {
	local var1=1
	local var2=2
	echo $[ $var1 + $var2 ]
}
value=$(addem)
va=$[ $var1+$var2 ]
echo $value
echo $va
.....................................................
$ ./test.sh
3
30
```

#### 数组

注意：shell只有一维数组

使用方法同C语言

例子：

```shell
#!/bin/bash
A=1
my_array=($A "A" C)
my_array[1]=B
echo "第一个元素为: ${my_array[0]}"
echo "第二个元素为: ${my_array[1]}"
echo "第三个元素为: ${my_array[2]}"
...............................................
$ ./test.sh
第一个元素为: 1
第二个元素为: B
第三个元素为: C
```

#### 关联数组

使用declare -A来表示这是一个关联数组

```shell
declare -A site
site["google"]="www.google.com"
site["runoob"]="www.runoob.com"
site["taobao"]="www.taobao.com"
echo ${site["runoob"]}
```

也可以声明的同时赋值

```shell
declare -A site=(["google"]="www.google.com" ["runoob"]="www.runoob.com" ["taobao"]="www.taobao.com")
echo ${site["runoob"]}
......................................................................................
$ ./test.sh
www.runoob.com
```

- 使用 @ 或 * 可以获取数组中的所有元素例如：${my_array[*]}

- 感叹号 ! 可以获取数组的所有键，${!my_array[*]}

- #可以获取数组长度例如：${#my_array[*]}

#### 使用数组作为函数参数

如果只是使用`函数名 $myarray`则只会传递一个参数

要想传递整个数组必须将该数组变量的值分解成单个的值，然后将这些值作为函数参数使用。在函数内部，可以将所有的参数重新组合成一个新的变量

也就是必须使用${myarray[@]}来传递参数

**函数返回数组:**函数用echo语句来按正确顺序输出单个数组值，然后脚本再将它们重新放进一个新的数组变量中。

```shell
#错误示例
#!/bin/bash
function testit {
	echo "The parameters are: $@"
	thisarray=$1
	echo "The received array is ${thisarray[*]}"
}
myarray=(1 2 3 4 5)
echo "The original array is: ${myarray[*]}"
testit $myarray
...........................................................
$ ./test.sh
The original array is: 1 2 3 4 5
The parameters are: 1
The received array is 1

#正确案例
#!/bin/bash
function testit {
    local newarray
    newarray=("$@")
    for((i=0;i<$#;i++))
    {
	newarray[$i]=$[ ${newarray[i]}*2 ]	
    }
    echo ${newarray[*]}
	}
myarray=(1 2 3 4 5)
echo "The original array is ${myarray[*]}"
result=$(testit "${myarray[@]}")
echo ${result[@]}
....................................................................
$ ./test.sh
The original array is 1 2 3 4 5
2 4 6 8 10
```

注意：在使用数组时除了最开始定义，其他时候只要用数组都是采用：${数组[]}的格式访问数组元素

#### 函数递归

递归使用同C语言一样

举例：计算阶乘

```shell
#!/bin/bash
function factorial {
if [ $1 -eq 1 ]
	then
		echo 1
	else
		local temp=$[ $1 - 1 ]
		local result=$(factorial $temp)
		echo $[ $result * $1 ]
	fi
}
read -p "Enter value: " value
result=$(factorial $value)
echo "The factorial of $value is: $result"
...................................................................
$ ./test.sh
Enter value: 5
The factorial of 5 is: 120
```

#### 库文件

和环境变量一样，shell函数仅在定义它的shell会话内有效。如果你在shell命令行界面的提示符下运行myfuncs shell脚本，shell会创建一个新的shell并在其中
运行这个脚本。它会为那个新shell定义这三个函数，但当你运行另外一个要用到这些函数的脚本时，它们是无法使用的。

使用函数库的关键在于source命令。source命令会在当前shell上下文中执行命令，而不是创建一个新shell。

使用`../库名称`来引入，可以类比于C语言的include<>

```shell
function add {
    echo $[ $1 + $2 ]
}

mul() {
    echo $[ $1 * $2 ]
}

function div {
    echo $[ $1 / $2 ]
}
```

```shell
#!/bin/bash
. ./myfunc

for((i=0; i<3; i++))
do
	value1=10
	value2=5
	array=("add" "mul" "div")
	result=$(${array[i]} $value1 $value2)
	echo $result
done
$
$ ./test.sh
15
50
2
```

#### 在命令行下使用函数

和在shell脚本中将脚本函数当命令使用一样，我们可以在命令行中使用函数，一旦在shell中定义了函数，你就可以在整个系统中使用它了，无需担心脚本是不是在PATH环境变量里

在命令行上创建函数时要特别小心。如果你给函数起了个跟内建命令或另一个命令相同的名字，函数将会覆盖原来的命令。

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/截图 2024-03-09 14-38-36.png) 

![1111](F:\myblog\typora-user-images\linux\截图 2024-03-09 14-38-36.png) 

#### .bashrc文件

bash shell在每次启动时都会在主目录下查找这个文件，不管是交互式shell还是从现有shell中启动的新shell。

在命令行写下的函数当关闭时就消失了，如果下次还想使用的话可以将其写入.bashrc文件中。

或者采用间接的方式引入

用source命令（或者. ./库文件）将库文件中的函数添加到你的.bashrc脚本中也可以

```shell
$ cat .bashrc
# .bashrc
# Source global definitions
if [ -r /etc/bashrc ]; then
. /etc/bashrc
fi
. /home/rich/libraries/myfuncs
$
```

#### shtool函数

shtool库提供了一些简单的shell脚本函数，可以用来完成日常的shell功能,不过得自己手动安装

[shtool]: https://mirrors.aliyun.com/gnu/shtool/

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/截图 2024-03-09 16-55-56.png) 

### 图形化桌面

#### 文本菜单与selete命令

使用echo输出各个选项，配合case命令实现图像化界面

- -e选项echo命令只显示可打印文本字符。在创建菜单项时，非可打印字符通常也很有用，比如制表符和换行符。要在echo命令中包含这些字符，必须用-e选项

- -en选项会去掉末尾的换行符。这让菜单看上去更专业一些，光标会一直在行尾等待用户的输入

```shell
clear
echo
echo -e "\t\t\tSys Admin Menu\n"
echo -e "\t1. Display disk space"
echo -e "\t2. Display logged on users"
echo -e "\t3. Display memory usage"
echo -e "\t0. Exit menu\n\n"
echo -en "\t\tEnter option: "
```

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/截图 2024-03-09 16-20-47.png) 

**当然更推荐使用selete命令进行菜单构建**

格式如下：

```shell
select variable in list
do
	commands
done
```

list参数是由空格分隔的文本选项列表，这些列表构成了整个菜单。select命令会将每个
列表项显示成一个带编号的选项，然后为选项显示一个由PS3环境变量定义的特殊提示符。

举例：

```shell
#!/bin/bash
# using select in the menu1
function diskspace {
    clear
    df -k
}

function whoseon {
    clear
    who
}

function memusage {
    clear
    cat /proc/meminfo
}

PS3="Enter option: "
select option in "Display disk space" "Display logged on users" "Display memory usage" "Exit program"
do
    case $option in
        "Exit program")
            break ;;
        "Display disk space")
            diskspace ;;
        "Display logged on users")
            whoseon ;;
        "Display memory usage")
            memusage ;;
        *)
            clear
            echo "Sorry, wrong selection";;
    esac
done
clear
```

#### 制作窗口

```shell
#使用该命令安装
sudo apt-get install dialog
```

![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/截图 2024-03-09 17-05-10.png) 

 ![](/media/kobayashi/新加卷/myblog/typora-user-images/linux/截图 2024-03-09 17-07-35.png) 

格式：

widget是表中的部件名，parameters定义了部件窗口的大小以及部件需要的文本

```shell
dialog --widget parameters
```

Yes或OK按钮，dialog命令会返回退出状态码0。如果选择了Cancel或No按钮，dialog命令会返回退出状态码1。

可以用标准的$?变量来确定dialog部件中具体选择了哪个按钮。echo $?

如果部件返回了数据，比如菜单选择，那么dialog命令会将数据发送到STDERR。可以用标准的bash shell方法来将STDERR输出重定向到另一个文件或文件描述符中。

- **msgbox部件**

```shell
#格式：
dialog --msgbox text height width
#例子：
dialog --title Testing --msgbox "This is a test" 10 20
```

- **yesno部件**

```shell
dialog --title "Please answer" --yesno "Is this thing on?" 10 20
```

- **inputbox部件**

```shell
dialog --inputbox "Enter your age:" 10 20 2>age.txt
```

- **textbox部件**

```shell
dialog --textbox /etc/passwd 15 45
```

- **menu部件**

```shell
dialog --menu "Sys Admin Menu" 15 30 1 1 "Display disk space" 2 "Display users" 3 "Display memory usage" 4 "Exit" 2> test.txt
```

此处省略后续内容.........................，当以后要用到再说

