#                                                               makefile基础入门

## makefile基本介绍

makefile相当于脚本程序，make程序是makefile的解析器，它定义了各种关键字、语法结构、函数变量，甚至可以用include关键字包含其他makefile。

make只是在makefile中找出哪些文件更新了，然后调用其他东西进行编译执行，所以命令规则里的命令都是shell命令。

## make如何找出文件是否更新

在linux中每个文件都有3种时间，分别是atime、mtime、ctime

**（1）atime（access time）**

- 每次读取文件数据部分时就会更新atime，比如cat、less命令就会更新时间ls则不会

**（2）ctime（change time）**

- 当文件的属性、数据被修改时就会更新ctime

**（3）mtime（modify time）**

- 当文件数据被修改就会更新时间，除此之外ctime也会更新。

## make时执行的文件

- 可以使用make -f [目标文件名] 指定想要make的文件

- 直接执行make时寻找文件优先级为GNUmakefile > makefile >Makefile

## makefile基本语法

- **基本语法：**

  目标文件：依赖文件

  【TAB】命令

- **伪目标**

  即不考虑mtime直接执行的语句

  不能与真实的文件同名，也不需要依赖文件，一般情况下用.PHONY来修饰

  格式为:

  .PHONY:伪目标名

  举例：用clean来删除编译过程中的.o文件

  ```makefile
  .PHONY:clean
  clean:
  	rm ./build/*.o
  ```

- **make：递归推导目标**

  在makefile中的目标是以递归的方式逐层向上查找目标的

  此时目录下只有test1.c和test2.c文件

  ```makefile
  test2.0:test2.c
  	gcc -c -o test2.o test2.c
  test1.o:test1.c
  	gcc -c -o test1.o test1.c
  test.bin:test1.o test2.o
  	gcc -o test.bin test1.o test2.o
  ```

  执行make test.bin，make发现其需要test1.o和test2.o但是并没有相应文件于是去寻找以其为目标文件的规则

  分别在第一行与第三行找到，通过执行其相应操作得到了test1.o和test2.o，然后返回test.bin执行。

- **make [函数名]**

  当make中有许多函数时使用make [函数名]可以指定想要执行的函数而不执行其他的

  直接make则会将makefile里面所有脚本都会执行

- **自定义变量**

  与上面的例子相同，这里定义了一个变量object用来指代test1.o和test2.o

  于是在之后的语句中就可以使用object来取代test1.o test2.o     

  <font color='red'>注意使用变量时格式为$(变量名)</font>

  ```makefile
  test2.0:test2.c
  	gcc -c -o test2.o test2.c
  test1.o:test1.c
  	gcc -c -o test1.o test1.c
  object = test1.o test2.o	
  test.bin:$(object)
  	gcc -o test.bin $(object)
  ```

- **@符号**

  make时会把执行的命令也打印出来，在其前面加个@就不会打印命令了。

  例如

  ```makefile
  test:
  	@echo "makefile"
  ```

  结果输出：makefile

- **$@与$^符号**

  <font color='red'>$@</font>符号表示该语句的目标文件，此处为<font color='red'>$(TAGRET)</font>即hello

  <font color='red'>$^</font>符号表示该语句的所有依赖文件，此处为<font color='red'>$(OBJ)</font> 即main.o print.o factorial.o

  <font color='red'>$ <</font>符号表示该语句的第一个依赖文件

  <font color='red'>-Wall</font>：选项可以打印出编译时所有的错误或者警告信息。

  这个选项很容易被遗忘，编译的时候，没有错误或者警告提示，以为自己的程序很完美，其实，里面有可能隐藏着许多陷阱。变量没有初始化，类型不匹配，或者类型转换错误等警告提示需要重点注意，错误就隐藏在这些代码里面。

  ```makefile
  #变量的定义
  CXX = g++
  TARGET = hello
  
  #相当于 OBJ = main.o print.o factorial.o
  SRC = $(wildcard *.cpp)  #把当前目录下所有.cpp文件全放进去
  OBJ = $(patsubst %.cpp, %.o, $(SRC)) #路径替换，把SRC里面的cpp文件替换成.o
  #编译选项
  CXXFLAGS = -c -Wall
  
  $(TAGRET):$(OBJ)
  	$(CXX) -o $@ $^
  
  #统一的规则，将所有cpp生成.o文件
  %.o:%.cpp
  	$(CXX) $(CXXFLAGS) $< -o $@	
  ```

  附加说明：

  ```makefile
  SRC = $(wildcard *.cpp)  #把当前目录下所有.cpp文件全放进去
  OBJ = $(patsubst %.cpp, %.o, $(SRC)) #路径替换，把SRC里面的cpp文件替换成.o
  ```

  相当于 OBJ = main.o print.o factorial.o

  实现了自动编译，以后再往当前目录中加入cpp文件时不需要再修改makefile

## 通用makefile模板

不同项目只需要修改定义源文件就可以了

```makefile
# 设置编译器和编译选项
CXX = g++
CXXFLAGS = -Wall -Wextra -std=c++11

# 定义目标文件
TARGET = main

# 定义源文件
SRCS = main.cpp qsort.cpp heap.cpp 

# 定义目标文件夹
OBJ_DIR = obj

# 生成目标
$(TARGET): $(addprefix $(OBJ_DIR)/, $(SRCS:.cpp=.o))
	$(CXX) $(CXXFLAGS) -o $@ $^

# 生成目标的依赖关系
$(OBJ_DIR)/%.o: %.cpp | $(OBJ_DIR)
	$(CXX) $(CXXFLAGS) -c -o $@ $<

# 创建目标文件夹
$(OBJ_DIR):
	mkdir -p $(OBJ_DIR)

# 清理生成的文件
clean:
	rm -f $(TARGET) $(addprefix $(OBJ_DIR)/, $(SRCS:.cpp=.o))
```

对应的目录结构如下，obj用于存放中间生成的目标文件

│── algo.h
│── heap.cpp
│── heap.h
│── qsort.cpp
│── qsort.h
├── obj
├── main.cpp
└── makefile

# CMake基础

## CMake基本介绍 

CMake是一个代码构建工具，与平台和构建系统无关，让程序员可以只关注于代码的编写。

## 构建一个基本命令

1、基本目录结构如下
├── algo.h
├── bubble.cpp
├── bubble.h
├── build
│   ├── CMakeCache.txt
│   ├── CMakeFiles
│   ├── cmake_install.cmake
│   └── Makefile
├── CMakeLists.txt
└── main.cpp

2、CMakeLists.txt文件内容如下

```cmake
# TODO 1:设置cmake的最低版本要求为3.10
cmake_minimum_required(VERSION 3.10)

# TODO 2:创建一个名称为xxx的项目
project(main)

# TODO 3:为项目添加一个叫做main的可执行文件
# Hint：一定要指定源文件
add_executable(main main.cpp bubble.cpp)
```

3、创建一个文件夹用于存放cmake构建出来的内容

```shell
mkdir build
```

4、进入该文件夹下，并构建相应内容

```shell
cd build
cmake ..			# ..表示cmakelists.txt文件所在目录
cmake --build .		# cmake --build全平台项目通用 .表示在当前目录下
#make		# 也可以直接使用makefile文件构建
```

## 指定c++标准

set用于给变量设置一个值

CMAKE_CXX_STANDARD用于指定标准

```cmake
set(<变量名><变量值>)
# 示范案例：
set(CMAKE_CXX_STANDARD 11)		#将CMAKE_CXX_STANDARD的值设为11表示使用c++11作为标准
set(SRC_DIR /home/src)		#将SRC_DIR设置为/home/src
```

## 自定义一个库

1、创建对应的库

add_library(库名称 对应的cpp文件)

2、添加链接库

target_link_libraries(目标名称 PUBLIC 对应的库名称)

目录结构如下,自定义了一个排序算法库

├── algorithm
│   ├── algo.h
│   ├── bubble.cpp
│   ├── bubble.h
│   ├── CMakeLists.txt
│   ├── heap.cpp
│   ├── heap.h
│   ├── merge.cpp
│   ├── merge.h
│   ├── qsort.cpp
│   └── qsort.h
├── build
├── CMakeLists.txt
├── main.cpp
└── makefile

algorithm/CMakeLists.txt中的内容如下

```cmake
# TODO 1:创建一个叫algo的库
add_library(algo bubble.cpp heap.cpp merge.cpp qsort.cpp algo.h)
# TODO 2：指定文件搜索路径
target_include_directories(algo PUBLIC
    ${CMAKE_CURRENT_SOURCE_DIR}
)
```

主目录下的CMakeLists.txt

```cmake
# TODO 1:设置cmake的最低版本要求为3.10
cmake_minimum_required(VERSION 3.10)

# TODO 2:创建一个名称为xxx的项目
project(main)

# TODO 3:为项目添加一个叫做TUtorial的可执行文件
# Hint：一定要指定源文件tutorial.cxx
add_executable(main main.cpp)

## TODO 4：设置c++标准
set(CMAKE_CXX_STANDARD 11)

# TODO 5：添加子文件夹
add_subdirectory(algorithm)

# TODO 6：添加链接库
target_link_libraries(main PUBLIC algo)

# TODO 7:添加到头文件搜索路径
target_include_directories(main PUBLIC "${PROJECT_SOURCE_DIR}/algorithm") #不写就默认当前目录
```

## 命令

### option

`option` 命令用于定义一个开关选项，允许用户在生成项目构建系统时控制某些行为或功能的开启或关闭。`option` 命令的基本语法如下：

```cmake
option(变量 "描述信息" 选项初始值)		#默认初始值为off
```

在 CMakeLists.txt 文件中，我们可以根据这个选项的值来控制编译过程中的行为。

### if



## 附录：CMake预定义的变量

- ${PROJECT_SOURCE_DIR}表示cmakelists所在的目录，也就是整个项目的根目录。
