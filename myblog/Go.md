# Go

## 一、环境搭建

1、安装SDK

https://go.dev/dl/

在官网找到自己电脑对应的go版本下载，我电脑下载的是go1.22.2 linux/amd64

```bash
 sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```

2、设置GoPath(默认工作目录，将自动存放项目所依赖的库等) **也可以不用设置，之后使用gomod进行项目管理**

ubuntu中在默认目录下编辑 .profile

```shell
#修改~/.profile
vim ~/.profile
 
#添加Gopath路径
export GOROOT="/usr/local/go"
export GOPATH="/home/kobayashi/gopath"
export PATH=$PATH:/usr/local/go/bin
 
# 激活配置
source ~/.profile
```

3、配置vsode

- 安装go插件

#### 如何下载第三方库

- go get下载第三方库

```go
go get -u github.com/go-sql-driver/mysql 	//-u表示依赖更新到最新
```

- 查看已安装的库

```go
go list -m all
```

- 删除第三方库

首先查看GOPATH：

```go
go clean
```

## 二、GO语言基本语法

### 1、hello world

```go
package main

import "fmt"

func sayhelloworld() {
	fmt.Println("hello world")
}

func main() {
	sayhelloworld()
}
```

- 编译链接：

```go
go build ./hello.go 
```

- 直接运行：

```go
go run ./hello.go 
```

### 2、初始化项目

在终端代码目录下输入如下命令，这时会在当前目录生成go.mod文件，这样就不用设置gopath了。

```go
go mod init xxx		//xxx为项目的名称
```

2.1、初始化项目以后就可以进行多文件编写了

```go
-----------------------hello.go-----------------
package main

import "fmt"

func sayhelloworld() {
	fmt.Println("hello world")
}
----------------------main.go-------------------
package main

func main() {
	sayhelloworld()
}
```

```bash
go run /home/kobayashi/gopath		#这里写代码所在文件夹的绝对路径
```

2.2、在跨包调用函数

项目目录，packet为包所在的目录

├── day1.md
├── main.go
└── packet
    └── packet1.go

```go
..............packet1.go........................
package packet1

import "fmt"

func Sayhelloworld() {
	fmt.Println("hello world")
}
...............main.go.......................
package main

import packet1 "helloworld/packet"

func main() {
	packet1.Sayhelloworld()
}
```

### 3、变量

#### 声明变量

格式：var 变量名 变量类型

```go
var a int			//a为整数类型
var b string		//b为字符串类型
var c []float32	//c为32位浮点切片类型变量，是由多个多个浮点类型组成的数据结构
var d func() bool	//返回值为bool类型的函数变量
vare struct{		//结构体
	x int
}
var(				//也可以全部一起定义
    a int
    b string
    c []float32
    d func() bool
    e struct{
        x int
    }
)
```

#### 初始化变量

格式：var 变量名 类型 = 表达式

```go
例如：
var hp int = 100
var hp = 100		//右值推导出hp的类型
hp :=100			//更加简便的写法
```

注意：如果左值全是声明过的变量然后使用:=会报错

```go
var hp int
//hp :=100		//报错no new variables on left side of :=	

conn1,err :=net.Dial("tcp","127.0.0.1:8080")
conn2,err :=net.Dial("tcp","127.0.0.1:8080")		//由于:=左边不完全相同则不会报错
```

#### 多个变量同时赋值

- 交换变量的值

  ```go
  var (
  		a int = 100
  		b int = 200
  	)
  	b, a = a, b
  ```


- 匿名变量

在使用多重赋值的时候如果不需要在左值中接受变量，可以使用匿名变量 _ 代替

```go
func GetData()(int,int){
	return 100,200
}
a,_ :=GetData()
_,b :=GetData()
fmt.Println(a,b)
```

#### 全局变量

如果想要跨包调用这个变量只需要将变量的首字母大写就好了非常方便

如果想要只在包内调用变量首字母要小写

函数也是同意的道理

### 3、常量

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

格式：

```go
const identifier [type] = value
```

特殊常量iota：

iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始。

### 4、变量逃逸(Escape Analysis)

go语言会自动根据变量的类型进行内存位置的分配，

分配在栈中更快，而对堆上的内存操作会更慢。所以函数中的局部变量一般都分配在栈中。

编译器会根据变量是否被外部引用来决定是否逃逸：

1. 如果函数外部没有引用，则优先放到栈中；
2. 如果函数外部存在引用，则必定放到堆中；
3. 如果需要的内存过大则放入堆中；
4. 如果考察到在函数返回后，此变量不会被引用，那么还是会被分配到栈上。套个取址符，就想骗补助？Too young！

使用`go run -gcflags "-m -N" main.go`命令可以来观察变量逃逸情况

- -gcflags表示编译参数，-m表示进行内存分配分析，-N表示避免程序优化 -l表示禁止内联

举例：

```go
package main

import (
	"fmt"
)

func Escape() *int {
	var c int = 10
	return &c
}

func main() {
	fmt.Println("返回值是", *Escape())
}
```

```shell
$ go run -gcflags "-m -N -l" main.go
# command-line-arguments
./main.go:8:6: moved to heap: c
./main.go:13:13: ... argument does not escape
./main.go:13:14: "返回值是" escapes to heap
./main.go:13:30: *Escape() escapes to heap			
返回值是 10
```

- Escape()是一个函数返回值，而*Escape()是一个表达式，所以进行分析时会有\*Escape() escapes to heap

- 在调用`fmt.Println`时，字符串文字"返回值是"作为参数之一，需要被传递给`fmt.Println`。编译器为了确保在`fmt.Println`执行期间字符串仍然有效，可能会选择将其分配到堆上。所以这里 "返回值是" escapes to heap

### 5、字符串

#### 计算字符串长度

- ASCII字符串长度使用len()函数
- Unicode字符串使用utf8.RuneCountInString()函数

在Go中，如果你想逐个获取字符串 `s` 的每个字符，可以将字符串转换为 `rune` 切片，或者直接使用 `for range` 循环遍历 `s`，因为 `for range` 会自动将字符串分解为 `rune` 类型的字符。

这里有两种常见的方法：

#### 方法1：直接使用 `for range` 遍历
```go
package main

import "fmt"

func main() {
    s := "Hello, 世界"
    
    for i, r := range s {
        fmt.Printf("字符 %d: %c, Unicode码点: %U\n", i, r, r)
    }
}
```

**解释**：`for range` 会将字符串按字符分割，每个字符存储在 `r` 中，`i` 是字符的字节起始位置。这种方法可以正确处理多字节的Unicode字符。

如果是全英文的字符串直接使用s[n]数组的方式获取对应字符（注意不要越界）

#### 方法2：转换为 `rune` 切片
```go
r := []rune(s)
```

**解释**：将字符串 `s` 转换为 `rune` 切片后，每个元素都是一个 `rune`，可以表示一个Unicode字符。这种方法适合需要按字符（而不是字节）索引访问字符串时使用。比如访问第一个元素使用r[0]就可以了，最后使用string(r)将其转回string类型

举例：给定字符串11*3+4，要求提取出数字和符号

```go
func extractTokens(expression string) ([]int, []rune) {
	var numbers []int    // 保存提取的数字
	var operators []rune // 保存提取的运算符

	i := 0
	for i < len(expression) {
		c := rune(expression[i])

		if unicode.IsDigit(c) {
			// 解析多位数字
			start := i
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				i++
			}
			// 将多位数字转换为整数并保存
			num := 0
			for _, digit := range expression[start:i] {
				num = num*10 + int(digit-'0')
			}
			numbers = append(numbers, num)
			continue
		}

		// 如果是运算符，直接保存
		if c == '+' || c == '-' || c == '*' || c == '/' {
			operators = append(operators, c)
		}
		i++
	}

	return numbers, operators
}
```

#### 方法 3：使用 strings.Split 将字符串转换为数组
可以先用 strings.Split 函数将字符串拆分为字符数组，然后遍历该数组，将每个字符转换为整数：
```go
strDigits := strings.Split(num, "")
digits := make([]int, len(strDigits))

for i, s := range strDigits {
	digit, err := strconv.Atoi(s)
	if err != nil {
	return nil, err
	}
}
```

### 6、defer延迟调用

语法：defer + 函数/方法可以实现在当前goroutine结束前调用该函数/方法，即使发生了panic，该方法也会被调用

```go
defer 函数/方法
```

**核心原理：**当执行到defer语句时，会调用deferproc函数，每一个goroutine都对应着一个结构体g，deferproc函数新建的_defer结构会被存储到该goroutine的\_defer结构的链表中(堆上),新加入的\_defer会被放到链表的头部，来保证最后执行时是后进先出的顺序。

**优化：**在某些情况下，编译器会对 `defer` 进行优化以减少开销，例如：

- **快速路径（Fast-path）优化**：简单的 `defer` 调用直接在栈上记录，从而减少内存分配和垃圾回收的压力。这意味着一些简单的 `defer` 调用不会涉及到链表管理，可以直接在栈上处理，从而大幅提高性能。
- **内联优化**：编译器可能会直接将其转化为一个普通的尾部调用，而不使用链表，这样可以省去放置到_defer链表以及遍历链表的时间。但是如果在if块中调用defer，只能等到执行时才能知道其是否成立，为了解决这一问题go采用了deferBits位图来进行判断，这里就不细说了。

### 7、panic异常与异常捕获



## 三、容器：存储和组织数据

### 1、数组与切片

#### 数组

使用方法与c语言非常相似

- **初始化**

```go
//.......方式一...........
var team [3]string
team[0]="jack"
team[1]="tom"
tea,[2]="anny"
//......方式二............
var team = [3]string{"jack","tom","anny"}		
var team = [...]string{"jack","tom","anny"}		//编译器自动推导大小
```

- **遍历数组**

方法一：使用c语言方式对数组进行操作

方法二：for 键,值 range 数组名 (当然也可以使用 _ 来表示缺省值)

```go
for i,v:=rang team{
	fmt.Println(i,v)
}
```

#### 切片

go语言中的切片其实就是动态数组

- **初始化一个切片**

```go
//................从数组或切片生成新的切片........................
arry := [5]int{1, 2, 3, 4, 5}
var s1 []int = arry[0:4]	//这里的arry[]是一个左闭右开区间
//这里的切片s1是对数组arry的引用，如果改变s1指向的内容，arry也会变。

...................声明切片.............................
var s2 []int
//切片是一个引用类型，默认值为nil，只是声明切片则不会分配内存空间
var s2 = []int{}
//声明了一个空切片，会分配内存空间
```

- **make函数构造切片**

```go
格式：
make( []T, size, cap)
```

**T：切片的元素类型**

**size：就是为了这个类型分配多少个元素**

**cap：预分配的数量，这个值设定后不影响size，只是提前分配空间，避免多次分配内存空间导致性能下降**

举例：

```go
s2 = make([]int, 3, 5) //预分配3个元素，存储空间分配5个元素，可以用cap(s2)显示存储空间
fmt.Println(s2)
//使用make函数生成的切片一定发生了内存分配操作，而之前的给定开始与结束位置的切片只是一个引用类型且不会发生内存分配。
```

**如何删除切片中的某个元素**

```
for	
```

**make和new的区别：**

**`new`**：

1、用于任何类型的内存分配，返回该类型的零值的指针。

2、适用于需要引用类型零值的情况，如 `new(int)` 返回 `*int` 类型的指针。

**`make`**：

1、仅适用于切片、映射和通道等引用类型的初始化。

2、返回初始化后的引用类型实例，不是指针。

- **append函数添加元素**

  当空间不能容纳足够多的元素时，切片会进行扩容，规则是按当前容量的2倍进行扩充

```go
格式：
切片 = append(切片,添加的元素...)
```

举例：

```go
s1 = append(s1, 6, 7, 8) //因为s1是对静态数组的引用，所以创建了一个新动态数组
//如果s1是使用make创建的动态数组，则和c语言的malloc一样
```

- **copy函数可以将一个切片的数据复制到另外一个切片空间中**

```
格式：
copy( destSlice, srcSlice []T) int
```

举例：

```go
s:[0 2 3 4] s2:[0 0 0 6 7 8]
copy(s1[1:], s2[3:])
fmt.Println(s1)
s1:[0 6 7 8]
```

- **删除元素**

go语言并没有提供切片元素的删除函数，但是我们可以使用append将删除位置前的元素和删除位置后的元素连接起来

```go
//s2：[0 0 0 6 7 8]
index := 2 //删除第二个位置
s2 = append(s2[:index], s2[index+1:]...)
fmt.Println(s2)
//结果：s2：[0 0 6 7 8]
```

### 2、map映射

使用散列表hash实现----时间复杂度O(1)~O(n)

- **定义**

```go
//KeyType为键类型，ValueType是键对应的值类型
map [KeyType] ValueType
```

**举例：**

map是一个内部实现的类型，需要使用make创建

key可以使用除函数以外的任意类型

当查找不存在的键时会返回ValueType的默认值

```go
//............创建一个map映射................
scene :=make(map[string]int)
scene["route"] = 66
fmt.Println(scene["route"])
fmt.Println(scene["route2"])
//.........查询是否存在某个键..................
v,ok:=scene["route"]  	//多取一个ok变量可以判断route是否存在map中
//.........声明时填充.........................
m:=map[string]string{
    "w":"forward",
    "A":"left",
    "B":"right",
}
```

**map[string]interface{}**

在 Go 语言中，`map[string]interface{}` 是一个映射类型（map type），其中：

- `string` 是键的类型，表示映射的键为字符串类型。
- `interface{}` 是值的类型，表示映射的值可以是任意类型的数据。

- **遍历**

使用for range循环完成

```go
scene :=make(map[string]int)

scene["route"]=66
scene["brazil"]=4
scene["china"]=960
for k,v :=range scene{    //如果只需要值可以使用匿名变量
    fmt.Println(k,v)
}		
```

如果需要特定顺序的遍历结果，可以将map中的数据遍历复制到切片中，然后使用sort函数进行排序，最后输出

- **删除**

使用delete函数可以删除一组键值对

```go
格式：
delete(map实例,键)
举例：
delete(scene, "route")
```

- **清空map**

go语言中没有为清空map提供函数，所以不用管旧的map直接新建一个就好了。

不用担心垃圾回收的效率，golang中的并行垃圾回收效率比写一个清空函数高效多了

- **并发时的map----sync.Map**

golang中的map在并发时只读线程是安全的，同时读写不安全，进程之间为了对map进行读写而发生竞争。

**使用sync.Map可以解决这个问题**

//TODO

- **表明某个键是否存在，而不需要存储任何值:。**

map[\*ListNode]struct{}`：这是声明一个键为指向 `ListNode` 类型的指针（`*ListNode），值为 struct{}的映射。struct{} 是一种空结构体，它不占用任何空间。

`struct{}{}`：表示一个值为 **空结构体实例** 的表达式。这个空结构体是一种零值类型，没有字段，也不占内存。用 `struct{}` 作为值的目的是在 map 中表示集合（set）这种数据结构的方式。

为什么使用 `struct{}` 作为 map 的值类型而不是 `bool` 或其他类型呢？

**1、节省内存**：因为 `struct{}` 是零大小的类型，不占用额外的内存。

**2、语义清晰**：`struct{}` 表示 "我只关心这个键是否存在"，而不关心具体的值，这使得它非常适合用作集合。

```go
words := make(map[string]struct{})
for _, w := range wordDict {
    words[w] = struct{}{}
}
```

### 3、列表list

列表list是一个可以快速增删的非连续空间

- **初始化**

list的初始化有两种方法：new和声明

```go
//通过new初始化list
变量名 := list.New()
//通过声明初始化list
var 变量名 list.List
```

与切片和map不同的是列表并没有具体元素类型的限制



## 四、流程控制

### 1、if-else控制流

- 写法和c语言相同

```go
ten := 1
if ten > 0 {
    fmt.Println("ten大于0")
} else if ten == 0 {
    fmt.Println("ten=0")
} else {
    fmt.Println("ten小于0")
}
```

- 特殊写法

```go
if err := Escape(); err != nil {
    fmt.Println(err)
}
```

这么写的好处是将返回值的作用域限制在了if-else语句中，提升代码的稳定性

### 2、for循环

格式：直到条件表达式返回false时结束，或者break、goto、return、panic强制退出

```go
for 初始代码;条件表达式；结束语句{
循环体代码
}
```

举例：

```go
//普通循环
step := 2
for ; step > 0; step-- {
    fmt.Printlen(step)
}
// 无限循环
var i int
for {
    if i > 10 {
        break
    }
    i++
}
// 单个条件的循环
i := 0
for i < 10 {
    i++
}
```

#### **for range**

- 遍历切片与数组-----获取索引和元素值

```go
for key, value := range []int{1,2,3,4}{
	fmt.Printf("key:%d value:%d\n",key,value)
}
```

输出结果：

```go
key:0 value:1
key:1 value:2
key:2 value:3
key:3 value:4
```

- 遍历字符串获取字符

```go
var str = "hello 世界"
for key, value :=range str{
	fmt.Printf("key:%d value:0x%x\n",key,value)
}
```

输出结果：

```go
key:0 value:0x68
key:1 value:0x65
key:2 value:0x6c
key:3 value:0x6c
key:4 value:0x6f
key:5 value:0x20
key:6 value:0x4f60		//汉字占3bit
key:9 value:0x597d
```

### 3、switch

go的switch与c语言的switch语法基本相同。

不同点：

- 相比于c语言golang的switch条件可以基于**表达式**进行判断,**在这种情况下switch后就不需要跟判断变量**

- go语言switch的case与case之间是独立的代码块，所以不需要使用break

```go
举例：
//一分多支
var a = "mom"
switch a{
    case "mom","dad":
    fmt.Println("family")
}
//分支表达式
var r int = 11
switch{
    case r > 10 && r < 20:
    fmt.Println(r)
}
```

### 4、break、continue

用法与c语言相同

不同点：continue和break可以在语句后面打上相应流程控制的标签来break、continue相应的代码块

```go
BreakTag:
    for i := 0; i < 5; i++ {
        BreakTag2:
            for j := 0; j < 4; j++ {
                fmt.Println(i, j)
                switch i {
                    case 1:
                        break BreakTag
                    case 3:
                        continue BreakTag2
                }
            }
    }

```

## 五、函数

形式如下：

```go
func 函数名 (参数列表) (返回参数列表){
函数体
}
```

- 函数名由字母数字下划线组成，不能以数字开头
- 函数的调用方法和c语言相同，与c语言不同的是go可以返回多个值



## 六、结构体

#### 1、定义结构体

```go
type 类型名 struct{
    字段1 字段1类型
    字段2 字段2类型
}
```

举例：

```go
type Point struct{
    x int
    y int
    R,G,B byte
}
```

#### 2、实例化

----为其分配内存空间并初始化

- **基本的实例化**

```go
type Player struct{
	Name string
    Health_Point int
}
var jack Player  		
jack.Name="Jack"
jack.Health_Point=100
```

- **结构体指针**

```go
jack :=new(Player)
jack.Name="Jack"
jack.Health_Point=100
```

jack的类型是*Player，属于指针，但是go语言使用了语法糖将jack.转化为(\*jack).

- **取地址实例化**

```go
jack:=&Player{
    Name: "jack",
    Health_Point : 100,
}		//{}内可以填入初始化的值，不填就是默认值
```

#### 3、构造函数

golang语言的类型或结构体没有构造函数对功能，结构体的初始化可以使用函数封装来实现。

```go
type Cat struct{
	Color string
	name string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

```

#### 4、嵌入(c++中的派生)

为了方便理解我这里借用了c++的概念

解释：分别为基类Person和派生类Employee、Manager写了一个构造函数，并且在派生类中会调用基类的构造函数

- **显式嵌入**：因为有字段名，所以不会和外部结构体中的其他字段名冲突。
- **匿名嵌入**：如果嵌入的结构体和外部结构体有相同的字段名，会导致命名冲突，需要特别小心处理。

```go
// ............嵌入也就是派生...............
// 定义基础结构体
type Person struct {
	Name string
	Age  int
}

// 定义包含基础结构体的结构体
type Employee struct {
	person     Person // 嵌入 Person 结构体 ---- 写法1 显示嵌入
	EmployeeID string
}

type Manager struct {
	Person    // 嵌入 Person 结构体  ------写法2 匿名嵌入
    		  //字段名就是他的类型名
	ManagerID string
}

func CreatPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// 显示嵌入
func CreateEmployee(name string, age int, id string) *Employee {
	return &Employee{
		person:     *CreatPerson(name, age),
		EmployeeID: id,
	}
}

// 匿名嵌入
func CreateManager(name string, age int, id string) *Manager {
	return &Manager{
		Person: Person{
			Name: name,
			Age:  age,
		},					//进行了一个初始化嵌套
		ManagerID: id,
	}
}
func Insert() {
	// 创建 Employee 实例
	emp := CreateEmployee("jack", 37, "114514")
	Mag := CreateManager("Tom", 37, "114514")

	//调用时两种方法的不同点
	fmt.Println(emp.person.Age) //访问Person结构体中的字段和方法时，需要通过字段名
	fmt.Println(Mag.Age)        //可以直接访问嵌入的结构体中的字段和方法。
	
}
```

#### 5、方法

在 Golang 中，方法（methods）是与特定类型相关联的函数。用c++的话来说就是类里边的成员函数。

该成员函数与普通函数唯一不同点在于多了个接收器，**每个方法只能有一个接收器**

接收器变量名一般采用结构体名的第一个小写字母，而不是self、this之类的命名。

```go
func (b *Bag)Insert(item int)
        ^
        |
      接收器
```

**举例：**

```go
// ......................方法(成员函数)..........................
type Bag struct {
	items []int
}

func (b *Bag) Insert(item int) {
	b.items = append(b.items, item)
}
func (b *Bag) Cout() {
	for _, value := range b.items {
		fmt.Println(value)
	}
}
func (b *Bag) Num() int {
	return len(b.items)
}

func Method() {
	bag := new(Bag)
	bag.Insert(100)
	bag.Cout()
	fmt.Println(bag.Num())
}
```

**非指针类型的接收器**

当使用非指针类型的接收器时，只能够进行读取，修改无效

```go
func (b Bag)Insert(item int) 	//类似于这种形式	
```

var del func(int)=c.Do

## 七、接口

Go 语言提供了另外一种**数据类型**即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

Go 语言中的接口是隐式实现的，也就是说，如果一个类型实现了一个接口定义的所有方法，那么它就自动地实现了该接口。因此，我们可以通过将接口作为参数来实现对不同类型的调用，从而实现多态。

#### 1、声明接口

**格式**：

```go
type 接口类型名 interface{
	方法名1( 参数列表1 ) 返回值列表1
	方法名2( 参数列表2 ) 返回值列表2
}
```

- **接口类型名：自定义的类型名一般会在单词后面加上er，比如写操作接口较Writer**
- **方法名：如果首字母大写，可以被其他包访问**

**举例**：

```
type Writer interface{
	Write(p []byte) (n int,err error)
}
```

#### 2、实现接口

只有该结构体实现了接口所有的方法才能够赋值给相应的接口

- **对象实例赋值接口**

```go
var 名称 接口类型名称
mane = new(结构体名)
名称 = name
名称.方法()
```

举例：

```go
// 定义了动物行为的接口
type AnimalBehavior interface {
	say()
}

// 定义了猫和狗的结构体
type Cat struct {
	name string
}
type Dog struct {
	name string
}

// 猫和狗的构造函数
func CreatCat(N string) *Cat {
	return &Cat{
		name: N,
	}
}
func CreatDog(N string) *Dog {
	return &Dog{
		name: N,
	}
}

// 实现了猫行为的方法
func (C Cat) say() {
	fmt.Println("mimi")
}

// 实现了狗行为的方法
func (D Dog) say() {
	fmt.Println("wolf")
}
func (D Dog) eat() {
	fmt.Println("shit")
}

// 多态调用接口
func Polymorphisms() {
	var AB AnimalBehavior //定义接口
	//创建实例
	c := CreatCat("咪咪")
	d := CreatDog("旺财")
	//猫实例赋值接口
	AB = c
	AB.say()
	//狗实例赋值接口
	AB = d
	AB.say()
	d.eat()
}
```

实现多态要满足的三个条件

1、有包含方法的接口

2、有结构体实现该接口

3、将结构体的实例赋值给接口类型的实例(指针)

- **接口赋值**

接口C是接口P的子集所以可以将P直接赋值给C，反过来则不行

```
type Phoner interface{
	call()
	play()
}
type Caller interface{
	call()
}
type MyPhone struct{
	.......
}
func main(){
	var P Phoner
	var c Caller
	phone = new(MyPhone)
	P=phone
	c=P
}
```

- **接口嵌入**

```go
type Caller interface{
	Call()
}
type Player interface{
	play()
}
type Phone interface{
	Caller		//接口嵌入
	Player		
}
```

```go
等同于：
type Phone interface{
	call()		
	play()		
}
```

#### 3、接口详解

接口值是一个两字节长的数据结构。第一个字节包含一个指向内部表的指针，这个内部表叫做iTable，包含所存储的值的类型信息-----类型+方法集

第二个字节是一个指向所存储值的指针。将类型信息和指针组合在一起，就将这两个值组成立一种特殊的关系。

```go
----------------
| iTable的地址 |-----------> c的类型、方法集(itable)
| C值的地址    |-----------> 存储的值Cat
---------------
var a AnimalBehavior
c:=new(Cat)
a=c
```

#### 4、类型断言

- **空接口**

```go
type A interface{
}
```

空接口可以保存任意的数据类型

举例：

```go
type Element interface{}	//空接口
type Person struct{
    name string
    age int
}
func CreatPerson(N string,A int)*Person{
    return &Person{
        name: N,
        age: A,
    }
}
list:=make([]Element,3)		//包含3个元素的空切片
list[0]=1
list[1]="hello"
list[2]=CreatPerson("张三",37)
```

- **类型断言**

格式：

```go
value, ok := x.(T)
```

`x` 是一个接口类型的变量。

`T` 是你想要断言的具体类型。

`value` 是断言成功后得到的变量。

`ok` 是一个布尔值，表示断言是否成功。

**例1、**：

```go
switch 接口变量.(type){
	case 类型1:
		//类型1时候的处理
    case 类型2:
    	//类型2的处理
    ...
    default:
    	//不是case列举类型的处理
}
```

**例2、**

```go
// ..................类型断言...........................
type Element interface{} //空接口
type Person struct {
	name string
	age  int
}

func CreatPerson(N string, A int) *Person {
	return &Person{
		name: N,
		age:  A,
	}
}
func Allert() {
	list := make([]Element, 4) //包含3个元素的空切片
	list[0] = 1
	list[1] = "hello"
	list[2] = CreatPerson("张三", 37)
	list[3] = Person{"李四", 18}
	for index, element := range list {
		if Value, ok := element.(int); ok {
			fmt.Println(index, Value)
		} else if Value, ok := element.(string); ok {
			fmt.Println(index, Value)
		} else if Value, ok := element.(*Person); ok {
			fmt.Println(index, Value.name, Value.age)
		} else if Value, ok := element.(Person); ok {
			fmt.Println(index, Value.name, Value.age)
		}
	}
}
```

**注意事项：**

1、**类型断言失败时不会引发恐慌（panic）**：使用带有两个返回值的形式时（如上例），你可以检查 `ok` 来判断断言是否成功。如果使用单返回值形式，当断言失败时会引发恐慌。

2、**适用于接口类型**：类型断言只能用于接口类型的变量。如果 `x` 不是接口类型，编译器将报错。

3、**类型转换和类型断言的区别**：类型转换用于基本类型之间的转换，而类型断言用于从接口类型到具体类型的转换。

### type用法

在 Go (Golang) 语言中，`type` 关键字用于定义新的类型。它可以用于定义结构体、别名、函数类型等。`type` 在 Go 中非常重要，帮助创建复杂的数据结构以及实现类型安全。

1. **定义结构体 (Struct)**：
   结构体是一种复合类型，可以将多个字段组合在一起。
   
   ```go
   type Person struct {
       Name string
       Age  int
   }
   
   func main() {
       p := Person{Name: "Alice", Age: 30}
       fmt.Println(p.Name) // 输出: Alice
   }
   ```
   这里的 `Person` 是一种新类型，包含两个字段 `Name` 和 `Age`。
   
2. **类型别名 (Type Alias)**：
   类型别名为现有类型创建一个新的名称，但它们指向相同的底层类型。
   ```go
   type Age int
   var myAge Age = 25
   ```
   这里的 `Age` 是对 `int` 类型的别名，但 `Age` 类型和 `int` 类型是不同的。

   自定义类型（如 `Age`）与底层类型（如 `int`）是不同的，因为 Go 是一种强类型语言，它强调类型区分，确保语义和类型安全。
   
   这种机制有助于防止意外的类型混用，保证代码的安全性和一致性。
   
3. **定义函数类型**：
   可以使用 `type` 来定义函数类型，用于传递和调用函数。
   ```go
   type Adder func(int, int) int
   
   func add(a, b int) int {
       return a + b
   }
   
   func main() {
       var f Adder = add
       fmt.Println(f(3, 4)) // 输出: 7
   }
   ```
   在这里，`Adder` 是一个新的函数类型，代表具有两个 `int` 参数并返回 `int` 的函数。

4. **自定义接口 (Interface)**：
   使用 `type` 可以定义接口，接口是 Go 的一种抽象类型。
   ```go
   type Speaker interface {
       Speak() string
   }
   
   type Dog struct{}
   
   func (d Dog) Speak() string {
       return "Woof!"
   }
   
   func main() {
       var s Speaker = Dog{}
       fmt.Println(s.Speak()) // 输出: Woof!
   }
   ```
   `Speaker` 接口要求实现 `Speak` 方法，`Dog` 类型实现了该接口。

5. **定义切片和映射类型**：
   你可以通过 `type` 来定义新的切片或映射类型。
   ```go
   type IntSlice []int
   type StringMap map[string]string
   
   func main() {
       var nums IntSlice = []int{1, 2, 3}
       var dict StringMap = map[string]string{"a": "apple", "b": "banana"}
   
       fmt.Println(nums) // 输出: [1 2 3]
       fmt.Println(dict) // 输出: map[a:apple b:banana]
   }
   ```

6. **定义自定义类型的接收器方法**：
   通过 `type` 定义的自定义类型，可以为其定义方法。
   ```go
   type Rectangle struct {
       Width, Height float64
   }
   
   func (r Rectangle) Area() float64 {
       return r.Width * r.Height
   }
   
   func main() {
       rect := Rectangle{Width: 5, Height: 10}
       fmt.Println(rect.Area()) // 输出: 50
   }
   ```

## 错误处理

1、error接口

```go
type error interface{
	Error() string
}
```

go函数支持返回多个值，这些结果会伴随一个错误变量error。一般将其与nil值进行比较，nil值表示没有发生错误，而非nil值则表明出现了错误，然后打印出错误

2、处理错误的方式

- Sentinel errors



## 九、并发

### atomic原子操作

**原子操作**是不可中断的操作，通常用于实现计数器、自旋锁、状态标记等。

`sync/atomic` 提供对以下基本类型的原子操作支持：

- 整型（`int32`, `int64`, `uint32`, `uint64`）
- 指针（`unsafe.Pointer`）
- 布尔值（`uint32` 用作替代）

#### **常用操作函数**

(1) **整数操作**

- `atomic.AddInt32` / `atomic.AddInt64`：对整数执行原子加法。
- `atomic.LoadInt32` / `atomic.LoadInt64`：原子读取整数值。
- `atomic.StoreInt32` / `atomic.StoreInt64`：原子写入整数值。
- `atomic.SwapInt32` / `atomic.SwapInt64`：原子交换整数值。
- `atomic.CompareAndSwapInt32` / `atomic.CompareAndSwapInt64`：比较并交换整数值。

(2) **布尔值操作**

布尔值可以用 `uint32` 模拟，常用方法是 `atomic.CompareAndSwapUint32` 和 `atomic.StoreUint32`。

(3) **指针操作**

- `atomic.LoadPointer`：原子读取指针。
- `atomic.StorePointer`：原子写入指针。
- `atomic.SwapPointer`：原子交换指针。
- `atomic.CompareAndSwapPointer`：比较并交换指针。

举例：

```go
type Account struct {
	Balance int64 //余额
	InTx    bool  //是否在操作
}

func transfer(amount int64, accountFrom, accountTo *Account) bool {
	bal := atomic.LoadInt64(&accountFrom.Balance) //原子操作，读取
	if bal < amount {
		return false
	}
	atomic.AddInt64(&accountTo.Balance, amount)//原子操作，加
	atomic.AddInt64(&accountFrom.Balance, -amount)

	return true
}
```

### Goroutine

go语言中的并发指的是能让某个函数独立于其他函数运行的能力

在go中独立的任务叫做goroutine，goroutine是go中最基本的执行单元。事实上每一个go程序至少有一个goroutine，即主goroutine，生命周期同main函数

#### 启动goroutine

使用go关键字就可以创建goroutine,当存在多个goroutine，其执行顺序是随机的

- **匿名函数**

  这段代码会形成闭包，导致输出结果不符合预期。闭包捕获了变量 `i` 和 `val`，而不是它们在每次迭代时的值。这会导致所有的 goroutine 在执行时使用的是循环结束时的 `i` 和 `val` 的值。

  ```go
  func main() {
  	var values = [5]int{1, 2, 3, 4, 5}
  	for i, val := range values {
  		go func() {
  			fmt.Printf("索引%d的结果是%d\n", i, val)
  		}()
  	}
  	time.Sleep(1 * time.Second)
  }
  
  ```

  为了避免这种情况，建议在 goroutine 内部传递变量的副本。如下

  ```go
  func main() {
  	var values = [5]int{1, 2, 3, 4, 5}
  	for i, val := range values {
  		go func(i int) {
  			fmt.Printf("索引%d的结果是%d\n", i, val)
  		}(i)
  	}
  	time.Sleep(1 * time.Second)
  }
  ```

- **go+函数**

  ```go
  func main() {
  	for i := 0; i < 5; i++ {
  		go sleeptime(i)
  	}
  	time.Sleep(4 * time.Second)
  }
  func sleeptime(i int) {
      time.Sleep(3 * time.Second)
      fmt.Println("... snore... ", i)
  }
  ```

### 闭包

闭包（Closure）是一个函数，它引用了其词法作用域中的变量。换句话说，闭包不仅仅是一个函数，它还保留了定义它时所处的作用域中的变量。闭包可以在它的词法作用域外被调用，即使这些变量已经超出了它们的作用域，闭包仍然能够访问它们。

在 Go 中，闭包是通过匿名函数实现的。匿名函数可以访问它们外部函数中的变量，即使这些变量在外部函数执行完毕后依然存在。以下是一个简单的例子来解释闭包的概念：

```go
package main

import "fmt"

func main() {
    // 声明一个变量
    x := 10

    // 定义一个匿名函数，并赋值给变量 closure
    closure := func() {
        // 这个匿名函数引用了变量 x
        fmt.Println(x)
    }

    // 即使 x 的作用域在 main 函数内，closure 仍然可以访问 x
    closure() // 输出 10

    // 修改 x 的值
    x = 20

    // closure 仍然可以访问修改后的 x
    closure() // 输出 20
}
```

在这个例子中，匿名函数 `closure` 引用了变量 `x`，并且即使 `x` 的作用域在 `main` 函数中，当调用 `closure` 时，它仍然可以访问并打印 `x` 的值。这个匿名函数和它引用的变量一起构成了一个闭包。

### channel

传统的进程间通信是通过内存共享进行的，不同进程之间共享同一片内存需要进行上锁，而golang的进程间通信是通过数据传递实现的。

通道可以在多个gorutine之间安全的传递值，通道可以作为变量、函数参数、结构体字段

- **创建使用通道**

  chan是引用类型，空值为nil需要配合使用make函数，并且指定传输数据的类型。

  ```go
  c:=make(chan int）	//无缓存
  ```

  默认创建的都是无缓存的channel，读写都是即时阻塞。这时候可以使用有缓存的channel

  只有在通道中没有要接受的值时，接受才会阻塞；只有通道没有缓冲区域时，发送才会阻塞

  ```go
  c:=make(chan int,3)	 //有缓存
  ```

- **<-发送与接收**

  向通道发送值c<-99

  从通道接收r:=<- c

  **发送操作会等待直到另一个goroutine尝试对该通道进行接收操作为止，执行发生的goroutine在等待期间将无法进行其他操作**

  **执行接收的goroutine将等待直到另一个goroutine尝试向该通道进行发送操作为止**

  ```go
  //创建channel，异步执行的等待队列
  	c := make(chan int)
  	for i := 0; i < 5; i++ {
  		go func(i int, c chan int) {
  			time.Sleep(3 * time.Second)
  			fmt.Println("...", i, "snore ...")
  			c <- i
  		}(i, c)
  	}
  //接收channel发送过来的值
  	for i := 0; i < 5; i++ {
  		gopherID := <-c
  		fmt.Println("gopher", gopherID, "has finished sleeping")
  	}
  }
  ```

  多返回判断

  ```go
  data,ok:=<-ch
  ```

  data表示接受到的数据，ok的值表示是否接收到数据，通过ok可以判断channel是否被关闭，如果通道关闭ok值为false

- **关闭通道**

  ```go
  defer close(channel)
  ```

- **select处理多个通道**

  select包含的每一个case都有一个通道，用来发送或者接收数据（与switch有些类似）

  select会等待直到某一个case分支的操作就绪，然后就会执行该case分支

  `time.After` 是 Go 语言 `time` 包中的一个函数，它创建一个定时器，并在指定的时间段后向返回的通道发送当前时间。

  ```
  func After(d Duration) <-chan Time
  ```

  **参数**

  - `d Duration`：定时器的持续时间。在这个时间段之后，通道会接收到当前时间。

  **返回值**

  - `<-chan Time`：一个只读的 `Time` 通道。定时器到期后，会向该通道发送当前时间。

  **举例：**

  ```go
  //创建channel，异步执行的等待队列
  timedout := time.After(2 * time.Second)
  c := make(chan int)
  for i := 0; i < 5; i++ {
      go func(i int, c chan int) {
          time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
          fmt.Println("...", i, "snore ...")
          c <- i
      }(i, c)
  }
  
  //两秒之内运行上面分支，之后时间到了运行下面分支运行结束
  for i := 0; i < 5; i++ {
      select {
      case gopherID, ok := <-c:
          if !ok {
              fmt.Println("通道关闭")
          }
          fmt.Println("gopher", gopherID, "has finished sleeping")
      case <-timedout:
          fmt.Println("time run out")
          return
      }
  }
  ```

注意：即使已经停止等待goroutine，但只要main函数还没有返回仍在运行的goroutine将会继续占用内存

- **nil通道**

  如果不使用make初始化通道，那么通道变量的值就是nil，对nil通道进行发送或者接收不会引起panic，但会导致永久阻塞


- **循环接受数据**

  for range循环自动判断通道是否关闭，自动break循环。

  1）在遍历时，如果channel没有关闭，则会出现死锁的错误

  2）在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完成后退出

  ```go
  func main() {
  	ch := make(chan int, 3) //有缓冲的通道
  	ch1 := make(chan bool)
  
  	// 创建用于发送数据的 goroutine
  	go func(ch chan int, ch1 chan bool) {
  		for i := 0; i < 3; i++ {
  			ch <- i
  		}
  		fmt.Println("发送数据结束")
  		defer close(ch) // 发送完数据后关闭通道
  		ch1 <- true
  	}(ch, ch1)
  
  	// 等待发送数据的 goroutine 完成
  	<-ch1
  
  	// for range 循环自动判断通道是否关闭，自动 break 循环
  	for value := range ch {
  		fmt.Println("读取数据", value)
  	}
  	fmt.Println("读取结束")
  }
  ```

  

### 同步操作

- **channel阻塞实现同步**

  channel默认是阻塞的，当数据被发送到channel时会发生阻塞，直到有其他的goroutine从该channel中读取数据。

  当从channel中读取数据时，读取也会阻塞，直到其他的goroutine将数据写入该channel。

  `<-ch` 语句可以用于同步，而不需要接收具体的数据。

  ```go
  func main(){
  	ch1 := make(chan int)
  	go func() {
  		ch1 <- 1
  	}()
  	<-ch1		//在向ch1发送数据之前都会阻塞
  }
  ```
  
- **WaitGroup来实现同步等待**

  举例：

  ```go
  func worker(id int, wg *sync.WaitGroup) {
      defer wg.Done() // 标记此 goroutine 完成
  
      fmt.Printf("Worker %d starting\n", id)
      time.Sleep(time.Second) // 模拟工作
      fmt.Printf("Worker %d done\n", id)
  }
  func main() {
      var wg sync.WaitGroup
  
      for i := 1; i <= 5; i++ {
          wg.Add(1) // 每启动一个 goroutine，计数器增加 1
          go worker(i, &wg)
      }
      wg.Wait() // 阻塞，直到所有 goroutine 完成
      fmt.Println("All workers done")
  }
  ```

  确保每个 `Add` 对应一个 `Done`：每次调用 `Add` 增加计数器后，确保在相应的 goroutine 中调用 `Done`。否则，`Wait` 将永远阻塞。

  不要在 `Wait` 之后再调用 `Add`： 调用 `Wait` 表示你已经完成了所有任务的添加操作，不应该再调用 `Add` 增加计数器，否则会导致不一致的行为。

  如果计数器在调用 `Add` 时传递了负数或者 `Done` 被多次调用，导致计数器变为负数，程序会立即触发 panic。这是因为 `WaitGroup` 计数器不允许为负数。

### 竞争状态与互斥锁

**竞争条件：**在多线程环境下由于操作顺序的不确定导致执行结果不确定。比如两个线程对同一个变量进行操作，我们无法明确他们的执行顺序。

**数据竞争：**在多线程环境下由于操作顺序的不确定导致数据不一致。比如两个线程同时对一个变量进行修改，并且没有任何同步机制。

- **竞争**

  如果在多个goroutine没有同步的情况下访问同一个资源就会处于竞争状态，如下所示：

  ```go
  var Number int		//共享访问的资源
  var wait sync.WaitGroup		//实现同步等待
  
  func addNumber(addnumber int) {
  	for i := 0; i < addnumber; i++ {	//for循环对资源进行操作
  		Number++
  	}
  	wait.Done()
  }
  
  func main() {
  	wait.Add(3)
  	go addNumber(10000)	//启动3个goroutine进行竞争
  	go addNumber(20000)
  	go addNumber(30000)
  	wait.Wait()
  	fmt.Println(Number)
  }
  ```

  最后结果：44375	因为重复赋值导致了覆盖，所以小于60000


- **检查程序中的数据竞争**

  在测试或者运行时使用-race开启数据竞争检测器。

  ```shell
  $ go test -race mypack	//测试mypack包
  $ go run -race mysrc.go //运行时测试
  $ go build -race mycmd	//编译时测试
  ```

- **互斥锁**

  共享资源的使用是互斥的，即一个goroutine使用时会将其上锁，其他goroutine想要使用必须等它解锁。如果被锁那么想要获取该资源的goroutine会被阻塞进入睡眠状态，直到该资源解锁后才会被唤醒

  var mutex sync.Mutex   声明了一个类型为Mutex的变量，默认使用零值表示锁未被任何goroutine所持有，也没有等待获取锁的goroutine

  ```go
  var Number int          //共享访问的资源
  var wait sync.WaitGroup //实现同步等待
  var mutex sync.Mutex    //定义锁
  
  func addNumber(addnumber int) {
  	for i := 0; i < addnumber; i++ { //for循环对资源进行操作
  		mutex.Lock() //加锁
  		Number++
  		mutex.Unlock() //解锁
  	}
  	wait.Done()
  }
  
  func main() {
  	wait.Add(3)
  	go addNumber(10000) //启动3个goroutine进行竞争
  	go addNumber(20000)
  	go addNumber(30000)
  	wait.Wait()
  	fmt.Println(Number)
  }
  ```

  

### contex

`context`可以用来在`goroutine`之间传递上下文信息，相同的`context`可以传递给运行在不同`goroutine`中的函数，上下文对于多个`goroutine`同时使用是安全的



## 十、反射



## 十一、网络编程

### 1、TCP套接字

- **流式套接字(SOCK_STREAM)：面向连接的服务，针对TCP服务应用**
- **数据报式套接字(SOCK_DGRAM)：无连接的套接字，针对UDP服务应用**

**实现步骤：服务器监听、客户端请求、连接确认**

**golang中有关套接字的api都保存在net包中，Dial函数连接服务器，listen函数监听，accept函数接受连接**

1.1、Conn接口，这个接口是服务端与客户端进行交互的接口

```go
type Conn interface {
    // Read 从连接中读取数据到 b 中
    // Read 会阻塞直到有数据可读，或者发生错误
    // Read 返回读取的字节数和可能发生的错误
    Read(b []byte) (n int, err error)

    // Write 将 b 中的数据写入连接
    // Write 会阻塞直到所有数据写入完成，或者发生错误
    // Write 返回写入的字节数和可能发生的错误
    Write(b []byte) (n int, err error)

    // Close 关闭连接
    // 关闭连接后，所有未完成的读写操作都会被取消，并返回错误
    Close() error

    // LocalAddr 返回本地网络地址
    LocalAddr() Addr

    // RemoteAddr 返回远程网络地址
    RemoteAddr() Addr

    // SetDeadline 设置连接的读写超时时间
    // 读写操作在超时时间内没有完成将会返回错误
    SetDeadline(t time.Time) error

    // SetReadDeadline 设置连接的读取超时时间
    // 读取操作在超时时间内没有完成将会返回错误
    SetReadDeadline(t time.Time) error

    // SetWriteDeadline 设置连接的写入超时时间
    // 写入操作在超时时间内没有完成将会返回错误
    SetWriteDeadline(t time.Time) error
}

```

1.2、Dial函数，用于客户端连接服务器

```go
func Dial(network, address string) (Conn, error)
```

network用于指定网络类型如TCP、TCP4/6，UDP、UDP4/6，IP、IP4/6

address用于指定连接的地址格式为host : port，对于IPV6来说要用\[ ]比如：[ : : 1] : 80 ，：80表示本地系统。

1.3、Listen函数

用于创建监听，定义如下：

```go
func Listen(network, address string)(Listenner,error)
```

network表示网络类型，指定面向流的网络可选tcp、tcp4、tcp6、unix、unicpacket

address用于表示端口host: port ，若省略host表示监听所有本地地址，返回一个Listenner对象

1.4、Listenner的定义

```go
type Listenner interface{
	Accept()(Conn,error)
	Close()error
	Addr()Addr
}
```

**总结：所有带有close函数的方法的对象都需要使用defer调用close函数比如Listenner对象、conn对象**

举例：server服务端

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	//创建tcp监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//结束时关闭套接字
	defer listenner.Close()
	//等待连接
	conn, err := listenner.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	IpAddr := conn.RemoteAddr().String()
	fmt.Println(IpAddr, "连接成功")
	recvMsg := make([]byte, 1024)
	for {
		//阻塞等待用户发送数据
		n, err := conn.Read(recvMsg) //n为代码接受数据的长度
		if err != nil {
			fmt.Println(err)
			return
		}
		result := recvMsg[:n]
		fmt.Println("接收到", IpAddr, "的数据", string(result))
		if string(result) == "exit" {
			fmt.Println(IpAddr, "退出连接")
			return
		}
		conn.Write([]byte(string(result)))
	}

	//获取数据，发送数据

}
```

cilent客户端：

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	sendMsg := make([]byte, 1024)
	for {
		fmt.Println("请输入要发送的内容：")
		fmt.Scan(&sendMsg)

		conn.Write(sendMsg)
		//阻塞等待服务器返回消息
		n, err := conn.Read(sendMsg)
		if err != nil {
			fmt.Println(err)
		}
		result := sendMsg[:n]
		fmt.Println("接收到消息", string(result))
	}
}
```

TCP客户端和TCP服务器之间连接过程

```go
						net.Listen()
   			阻塞等待用户连接	 |
net.Dial()  ------------->  Accept()
   |		 数据（请求）		   |
  Write()  -------------->   Read()
   |		数据(应答)			|处理请求
  Read()   <--------------  Write()
  	|						   |
  Close()					Close()
 （TCP客户端）  				  （TCP服务端）
```

### 2、HTTP网络

golang使用内置的net/http包就可以实现GET与POST方式请求数据


#### **2.1、交互流程：**

cilent---->Request----->Multiplexer(路由)---->handle(处理函数)---->Response---->Cilent

客户端发送请求到服务器，服务器接受请求后，分配路由，然后选择相应的handle处理请求，最后将相应的信息返回客户端

#### **2.2、函数介绍：**

Http包的关键类型为Handler接口、ServerMux接口、HandlerFunc函数和Server方法

**2.2.1Handler接口：**

```go
type Handler interface{
	ServeHTTP(ResponseWriter, *Request)	//路由具体实现
}
```

`Handler` 接口只有一个方法 `ServeHTTP`，该方法接收两个参数：

- `ResponseWriter`：用于构建 HTTP 响应。
- `*Request`：表示客户端的 HTTP 请求。

只要实现了ServeHTTP这个方法的结构体都可以称为handle对象

**ResponseWriter:**

`http.ResponseWriter` 接口用于构建和发送 HTTP 响应。它提供了一组方法，可以用来设置响应头、写入响应体和设置响应状态码。

```go
type ResponseWriter interface {
	Header() Header

	Write([]byte) (int, error)

	WriteHeader(statusCode int)
}
```

- write方法接受一个byte切片作为参数，然后把他写入到http响应的body里面。如果在Writer方法被调用时，header里面没有设定content type，那么数据的前512字节就会被用来检测content type

  举例：

  ```go
  w.Write([]byte("hello world"))
  ```

- WriteHeader方法接受了一个整数类型(HTTP状态码)作为参数，并把他作为HTTP响应的状态码返回，如果该方法没有显示调用，那么在第一次调用Write方法前，会隐式地调用WriteHeader(http.StatusOK)，显示调用主要是用来发送错误类的HTTP状态码，在调用完WriteHeader方法后，仍然可以写入到ResponseWriter上，但无法再修改header了

  举例：

  ```go
  w.WriteHeader(302)
  ```

- Header方法返回headers的map，可以进行修改，修改后的headers将会体现在返回客户端的HTTP响应里面

  以下方法都是 `http.Header` 类型的方法，用于操作 HTTP 头部字段

  **func (h Header) Add(key string, value string)**

  - **用途**：
    
    - `Add` 方法用于向头部添加指定键（key）的值（value）。如果头部已经存在该键，则会在现有值的末尾添加新值，多个值之间用逗号分隔。
    
  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "application/json")
    h.Add("Cache-Control", "max-age=3600")
    ```

  **func (h Header) Clone() Header**

  - **用途**：
    - `Clone` 方法用于复制当前的头部。返回一个新的 `http.Header`，包含与原始头部相同的所有键值对。

  - **示例**：
    
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "text/html")
    
    cloned := h.Clone()
    ```

  **func (h Header) Del(key string)**

  - **用途**：
    - `Del` 方法用于删除指定键的所有值。

  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "text/plain")
    
    h.Del("Content-Type")
    ```

  **func (h Header) Get(key string) string**

  - **用途**：
    
    - `Get` 方法用于获取指定键的第一个值。
    
  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "application/json")
    
    contentType := h.Get("Content-Type")
    ```

  **func (h Header) Set(key string, value string)**

  - **用途**：
    - `Set` 方法用于设置指定键的值。如果键已存在，则覆盖原有的值。

  - **示例**：
    ```go
    h := make(http.Header)
    h.Set("Content-Type", "text/plain")
    ```

  **func (h Header) Values(key string) []string**

  - **用途**：
    - `Values` 方法返回指定键的所有值的切片。

  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Accept-Encoding", "gzip")
    h.Add("Accept-Encoding", "deflate")
    
    values := h.Values("Accept-Encoding")
    ```

  **func (h Header) Write(w io.Writer) error**

  - **用途**：
    - `Write` 方法将头部以 HTTP/1.1 格式写入到指定的 `io.Writer` 中。

  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "text/html")
    
    var buf bytes.Buffer
    err := h.Write(&buf)
    ```

  **func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error**

  - **用途**：
    
    - `WriteSubset` 方法类似于 `Write` 方法，但允许排除一些特定的头部字段。
    
  - **示例**：
    ```go
    h := make(http.Header)
    h.Add("Content-Type", "text/plain")
    h.Add("Cache-Control", "max-age=3600")
    
    exclude := map[string]bool{"Cache-Control": true}
    var buf bytes.Buffer
    err := h.WriteSubset(&buf, exclude)
    ```

  

**http.Request**

`*http.Request` 是一个结构体，表示客户端的 HTTP 请求。它包含了所有与请求相关的信息，例如请求方法、URL、头信息、请求体等。

- Method

```
Method string
```

请求的方法，例如 "GET", "POST", "PUT", "DELETE" 等。

```go
if r.Method == http.MethodGet {
    // 处理 GET 请求
}
```

- URL

```go
URL *url.URL
```

请求的 URL，包括路径和查询字符串。

```go
path := r.URL.Path
query := r.URL.Query()
```

- Header

```go
Header http.Header
```

请求头信息，是一个 `http.Header` 类型的对象。

```go
userAgent := r.Header.Get("User-Agent")
```

- Body

```go
Body io.ReadCloser
```

请求体，用于读取客户端发送的数据。需要注意的是，读取完 Body 后需要关闭它。

```go
body, err := ioutil.ReadAll(r.Body)
if err != nil {
    log.Fatal(err)
}
defer r.Body.Close()
```

- FormValue

```go
func (r *http.Request) FormValue(key string) string
```

获取 POST 或 URL 查询参数的值。如果有多个同名参数，只返回第一个。

```go
value := r.FormValue("username")
```

**2.2.2、ListenandServe函数**

http.ListenAndServe( )

- 第一个参数是网络地址

  如果为""，那么就是所有网络接口的80端口

- 第二个参数是handler

  如果为nil就表示DefaultServeMux

**2.2.3 创建Web Server**

创建webserver可以使用两种方法，其中mh是自定义的Handler

- http.Server()

```go
server := http.Server{
    Addr:    "localhost:8080",
    Handler: &mh,
}
server.ListenAndServe()
```

- http.ListenAndServer()

```go
http.ListenAndServe("localhost:8080", &mh)
```

这两种方法都会调用func (srv *Server) ListenAndServe() error 这个函数。

完整代码：

```go
package main

import "net/http"

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func main() {
	mh := myHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
    //http.ListenAndServe("localhost:8080", &mh)
}
```

举例：

```go
package main

import (
    "fmt"
    "net/http"
)

// 定义一个结构体
type MyHandler struct{}

// 实现 ServeHTTP 方法
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from MyHandler!")
}

func main() {
    // 实例化自定义 Handler
    handler := &MyHandler{}

    // 使用自定义 Handler
    http.Handle("/myhandler", handler)

    // 启动服务器
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

在这个例子中，我们定义了一个结构体 `MyHandler`，并为其实现了 `ServeHTTP` 方法。然后，我们使用 `http.Handle` 将路径 `/myhandler` 与这个处理器关联。这样，当访问 `http://localhost:8080/myhandler` 时，服务器会调用 `MyHandler` 的 `ServeHTTP` 方法并返回响应。

http.Handle用于将 URL 路径与 HTTP 处理器（实现了 `http.Handler` 接口的对象）关联起来。当服务器收到对该路径的请求时，会调用相应的处理器来处理请求。

**2.2.4、添加多个路由**

有两种方法一种是使用http.Handle	第二种使用HandlerFunc函数

但是在这之前我们要使用ServeMux接口创建路由，并确保所有路由都绑定到同一个 `ServeMux` 实例上，然后将该 `ServeMux` 传递给你的自定义服务器。

**其实ServeMux也是一个Handler,因为在源码中实现了这个接口，他不是直接处理请求和相应的，而是找到路由注册的handler然后间接调用它保存的muxEntry中保存的handler处理器的ServeHTTP()方法，所以后面例子可以看到将mux赋值给了Handler。**

**方法一、http.Handle**

```go
//创建新路由
mux := http.NewServeMux()
//使用Handle添加路由
mux.Handle("/", &mh)
mux.Handle("/hello", &hello)
//创建服务器时使用自定义的ServeMux，然后启动服务
server := http.Server{
    Addr:    "localhost:8080",
    Handler: mux,
}
server.ListenAndServe()
//启动服务的第二种方法
//http.ListenAndServe("localhost:8080", mux)
```

**方法二、HandlerFunc**

Go 语言标准库还提供了一个 `HandlerFunc` 类型，它是一个适配器，用于将普通函数转换为 `Handler` 接口。这使得我们可以更简洁地定义处理函数。

HandlerFunc 的定义如下：

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

举例：这里第二个参数为handler可以直接写，也可以写一个handler函数然后将函数赋值过去

```go
http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("world"))
})
//............................................................................
func welcome(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("welcome"))
}
http.HandleFunc("/welcome",welcome)
```

**若还是使用http.Handle这个函数，第二个参数可以使用http.HandlerFunc来将函数适配成Handler**

```go
func welcome(w http.ResponseWriter,r *http.Request){
    w.Write([]byte("welcome"))
}
http.Handle("/welcome",http.HandlerFunc(welcome))
//这是一个函数类型，该函数有个方法实现了Handler，所以http.HandlerFunc()就是一个Handler
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}

```

#### 2.3、内置handler

- http.NotFundHandler( )

```go
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
```

返回一个Handler，他会给每个请求都返回404 page not found

- http.RedirectHandler

```go
func RedirectHandler(url string, code int) Handler {
	return &redirectHandler{url, code}
}
```

返回一个handler，它把每个请求用给定的状态码跳转到指定的URL

- http.StripPrefix

```go
func StripPrefix(prefix string, h Handler) Handler {}
```

返回一个handler，他从请求URL中去掉指定的前缀，然后调用另一个handler，如果请求的url不符那么404

- http.TimeoutHandler

  ```go
  func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler {
  	return &timeoutHandler{
  		handler: h,
  		body:    msg,
  		dt:      dt,
  	}
  }
  ```

  返回一个handler，它用来在指定时间内运行传入的h

  - h，将要被修饰的handler
  - dt，第一个handler允许的处理时间
  - msg，如果超时，那么就把msg返回给请求，表示响应时间过长

- http.FileServer（提供了文件上传的接口）

  ```go
  func FileServer(root FileSystem) Handler {
  	return &fileHandler{root}
  }
  ```

  返回一个handler，使用基于root的文件系统来响应请求

  ```go
  type FileSystem interface {
  	Open(name string) (File, error)
  }
  ```

  使用时需要用到操作系统的文件系统,所以还需要委托给：
  
  type Dir string
  
  func  ( d Dir ) Open ( name string ) ( File , error )
  
  **举例：**
  
  ```go
  func (m *hellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  	// 使用文件服务器提供静态文件服务
  	fileServer := http.FileServer(http.Dir("cmd/day7/static"))
      fileServer.ServeHTTP(w, r)
  }
  ```
  
  `http.FileServer` 是一个处理器（handler），用于提供静态文件服务。
  
  `http.Dir("cmd/day7/static")` 指定了文件服务器的根目录，即 `"cmd/day7/static"` 目录。所有在该目录中的文件都可以通过 HTTP 请求访问。
  
  
  
  `ServeHTTP` 是 `http.FileServer` 处理器的方法，它接收两个参数：`http.ResponseWriter` 和 `*http.Request`。
  
  `w` 是 HTTP 响应写入器，用于向客户端写入响应。
  
  `r` 是 HTTP 请求，包含了客户端请求的所有信息。
  
  通过调用 `fileServer.ServeHTTP(w, r)`，文件服务器处理器会检查请求的 URL 路径，并尝试在指定的目录中找到相应的文件。如果找到文件，则读取文件内容并作为 HTTP 响应发送回客户端。

### 3、http消息

#### **3.1、HTTP概念**

**a)URL与URI**

URI局部环境做唯一标识，URL全局做唯一标识

**b)HTTP详解**

HTTP报文分为请求报文和响应报文两种类型。

1. **请求报文**：由客户端发送给服务器，用于请求服务器发送特定的资源。请求报文包含了请求行、请求头和请求体三部分。例如，当我们在浏览器中访问一个网页时，浏览器会发送一个GET请求报文给服务器，请求服务器返回该网页的内容。
2. **响应报文**：由服务器发送给客户端，用于响应客户端的请求。响应报文包含了响应行、响应头和响应体三部分。例如，当服务器收到一个GET请求后，它会返回一个响应报文给客户端，响应报文中包含了请求的资源内容（如HTML页面）、状态码（如200表示成功）以及其他附加信息。

HTTP报文由请求行（或响应行）、请求头（或响应头）和请求体（或响应体）三部分组成。

1. **请求行（或响应行）**：位于报文的首行，包含了HTTP方法（如GET、POST等）、URL和HTTP版本等信息。请求行用于描述客户端的请求，而响应行则用于描述服务器对请求的响应。
2. **请求头（或响应头）**：位于请求行（或响应行）之后，包含了关于客户端、服务器或请求/响应的附加信息。例如，`Host`头指定了请求的目标主机，`User-Agent`头包含了客户端的类型和版本信息，`Content-Type`头则指定了请求体（或响应体）的媒体类型等。
3. **请求体（或响应体）**：位于请求头（或响应头）之后，包含了实际传输的数据。对于GET请求，请求体通常为空；而对于POST请求，请求体则包含了提交给服务器的数据。响应体则包含了服务器返回给客户端的数据，如HTML页面、图片等。

<img src="/media/kobayashi/新加卷/myblog/typora-user-images/GO/http0.jpg" style="zoom:50%;" /><img src="/media/kobayashi/新加卷/myblog/typora-user-images/GO/http3.jpg" style="zoom:50%;" /> 

<img src="/media/kobayashi/新加卷/myblog/typora-user-images/GO/http2.jpg" style="zoom:50%;" /> <img src="/media/kobayashi/新加卷/myblog/typora-user-images/GO/http1.jpg" style="zoom:50%;" />  

#### 3.2、请求request

request（是一个struct），代表客户端发送的HTTP请求消息

重要字段：URL、Header、Body、Form、PostForm、Multipartform

也可以通过Request的方法访问请求中的cookie、URL、User Agent等信息

- **URL**

**URL字段表示请求第一行里面的部分内容，URL字段是一个指向url.URL类型的指针，而url.URL是一个结构体**

```go
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port (see Hostname and Port methods)
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

URL 表示已解析的 URL（从技术上讲，是 URI 引用）。

所表示的一般形式为：

```go
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
举例：
https://username:password@www.example.com:8080/path/to/resource?query1=value1&query2=value2#section1
```

不以斜杠开头的 URL 被解释为：

```go
scheme:opaque[?query][#fragment]
举例：
http://www.example.com/path/to/resource?query1=value1&query2=value2
```

URL中的`#fragment`字段（也称为片段标识符或锚点）用于定位资源中的特定部分。它通常在浏览器中使用，用于在加载网页后跳转到页面的某个部分。`#fragment`在HTTP请求中不会被发送到服务器，而是由客户端（如浏览器）处理。

**Query字段**

```go
type HelloHandler struct{}
func (m *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	query := url.Query()
	id := query["id"]
	log.Println(id)
	name := query.Get("name")
	log.Panicln(name)
}
```

r.URL.Query()会提供查询字符串对应的map[string]\[]string、

通过map索引可以获得所有值，使用Get方法可以获取第一个对应的值

- **Header（首部行）**

请求和响应的headers是通过header类型来描述的，它是一个map，用来表述HTTP Header中的Key-Value对

Header map的key是string类型，value是[]string，设置key的时候会创建一个空的[]string作为value，value里面第一个元素就是新header的值

举例如下：自定义了一个Header

```go
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)							//r.Header返回整个请求头map
	fmt.Fprintln(w, r.Header["Accept-Encoding"])		//r.Header["Accept-Encoding"]根据键找到对应的值
	fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))	//r.Header.Get("Accept-Encoding")返回字符串
}

```

- **Body**

  请求和响应的bodies都是使用Body字段来表示的

  Body是一个io.ReadCloser接口-----Reader和Closer

  - Reader接口定义了一个Open方法
    - 参数：[]byte
    - 返回：byte的数量、可选的错误

  - Closer接口定义了一个Close方法
    - 没有参数返回可选的错误信息

<span style="font-size:30px">**Form表单**</span>

- **通过表单发送请求**

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>GET Form Example</title>
</head>
<body>
    <form action="/process?name=nick" method="post">
        <input type="text" name="name"><br>
        <input type="text" name="age"><br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
```

表单中的数据是以name-value对的形式，通过post请求发送出去的，内容存放在post请求的body里面，并且数据格式是由enctype指定的。

**enctype属性**

1、enctype属性的默认值是application/x-www-form-urlencoded

此时浏览器会将表单数据编码到查询字符串里面

2、如果enctype是一个multipart/form-data那么每一个name-value对都会被转化为一个MIME消息部分

每一部分都有自己的Content Type和Content Disposition

不同点：application/x-www-form-urlencoded只能上传文本格式的文件，multipart/form-data是将文件以二进制的形式上传，这样可以实现多种类型的文件上传。

**表单GET**

get请求都是通过URL的name-value对来发送的

- **Form**

  **代码：**

  ```go
  package main
  
  import (
  	"fmt"
  	"net/http"
  )
  
  type hellHandler struct{}
  
  func (m *hellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  	// 使用文件服务器提供静态文件服务
  	fileServer := http.FileServer(http.Dir("cmd/day7/static"))
  	fileServer.ServeHTTP(w, r)
  	// 解析表单数据
  	r.ParseForm()
  	// 获取表单字段值
  	fmt.Fprintln(w, r.Form["name"])
  
  }
  
  func main() {
  	server := http.Server{
  		Addr: "localhost:8080",
  	}
  
  	he := hellHandler{}
  	http.Handle("/", &he)
  	server.ListenAndServe()
  }
  ```

  **通过Request上的函数我们可以从URL或/和Body中提取数据**

  **Form里面的数据是key-value对，需要调用ParseForm或ParseMultipartForm来解析Request，然后访问Form、PostForm或MultipartForm字段**

  - **Form字段**

    ```go
    Form url.Values
    
    type Values map[string][]string
    ```

    因为Form是一个map所以可以使用r.Form["name"]来返回含有一个元素切片例如：["jack"]，如果不指定则会返回所有且是乱序的

    注：只支持application/x-www-form-urlencoded

  - **PostForm字段**

    如果表单和URL里面具有相同的key，那么他们都会放在同一个slice里，表单靠前，URL的值靠后，如果只是想要表单的key-value可以使用PostForm字段

    注：只支持application/x-www-form-urlencoded

  - **MultipartForm字段**

    要想使用MultipartForm这个字段首先需要调用**ParseMultipartForm这个方法**，该方法会在必要时调用ParseForm方法，参数是需要读取数据的长度

    MultipartForm只包含表单的key-value对，返回struct而不是map

    ```go
    type Form struct {
    	Value map[string][]string
    	File  map[string][]*FileHeader
    }
    ```
  
    举例：
  
    ```go
    func (m *hellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    	// 使用文件服务器提供静态文件服务
    	fileServer := http.FileServer(http.Dir("cmd/day7/static"))
    	fileServer.ServeHTTP(w, r)
    	// 解析表单数据
    	r.ParseMultipartForm(1024)
    	// 获取表单字段值
    	fmt.Fprintln(w, r.MultipartForm)
    }
    ```
  
  - **FormValue&PostFormValue方法**

​				FormValue方法会返回Form字段中指定key对应的第一个value，使用这个方法时无需调用ParseForm或ParseMultipartForm来解析Request。

​				PostFormValue方法也是一样，但是只能读取PostForm

​				因为他们都会调用ParseMultipartForm方法。

但如果表单的enctype设为multipart/form-data,那么只能使用PostFormValue获取想要的值，FormValue则只能获取url的值

总结：

```go
func (m *hellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//...............方法一...................................................
    // 解析表单数据
	r.ParseMultipartForm(1024)		//r.ParseForm()
	// 获取表单字段值
	fmt.Fprintln(w, r.MultipartForm)	//fmt.Fprintln(w, r.PostForm["name"])
//.............方法二....................................................
    fmt.Fprintln(w,r.FormValue("name"))
}
```

- **文件上传**

```go
func (m *hellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用文件服务器提供静态文件服务
	fileServer := http.FileServer(http.Dir("cmd/day7/static"))
	fileServer.ServeHTTP(w, r)

	// 获取表单字段值
	fmt.Fprintln(w, r.FormValue("name"))

	// 解析多部分表单数据
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Error parsing multipart form data", http.StatusBadRequest)
		return
	}
//....................方法一、...........................................................
	//上传文件
	file, handler, err := r.FormFile("uploaded")
	if err != nil {
		http.Error(w, "Error retrieving the uploaded file", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err == nil {
		fmt.Fprintln(w, string(data))
	}
	// 打印文件信息
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

//......................方法二、........................................................
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()
	// if err == nil {
	// 	data, err := io.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(data))
	// 	}

	// }
}
```

**第一步：r.ParseMultipartForm( )**

首先需要调用 `r.ParseMultipartForm` 方法来解析请求中的 multipart/form-data 表单数据。

```go
err := r.ParseMultipartForm(1024) 
	if err != nil {
		http.Error(w, "Error parsing multipart form data", http.StatusBadRequest)
		return
	}
```

注意：调用r.ParseMultipartForm一定要记得错误检查，解析模板时检查返回`error`值，而不是忽略它。不然可能会报错找不到文件或其他原因。

**第二步：使用r.FormFile( )**

**适用于处理单个文件上传。它封装了打开文件和获取文件头的步骤，更加简便。**

**r.FormFile** 是一个方法，它从 **multipart/form-data** 表单中解析并返回指定键的第一个文件

```go
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
```

**key string**：表单字段的名称，用于识别文件。

返回值：

**multipart.File**：一个接口，表示上传的文件。以通过 `file` 来读取文件内容。

**multipart.FileHeader**：包含文件的头信息，如文件名、文件大小和 MIME 类型。

**error**：错误信息，如果获取文件失败，会返回一个非 nil 的错误。

**第二步：使用r.MultipartForm.File[]**

**适用于处理多个文件上传。需要手动打开文件，但提供了更大的灵活性。**

```go
//方法一、r.FormFile( )
file, handler, err := r.FormFile("uploaded")
//方法二、r.MultipartForm.File[]
fileHeader := r.MultipartForm.File["uploaded"][0]
file, err := fileHeader.Open()
```

#### 3.3、响应ResponseWriter

从服务器向客户端返回响应时需要使用ResponseWriter

ResponseWriter是一个接口，handler用它来返回响应

真正支撑Responsewriter的幕后struct是非导出的http.response

**内置的Response如下**

- NotFound函数，包装一个404状态码和一个额外的信息

- ServeFile函数，从文件系统提供文件，返回给请求者

- ServeContent函数，他可以把实现了io.ReadSeeker接口的任何东西里面的内容返回给请求者

  还可以处理Range请求，如果只是请求了资源的一部分内容，那么就可以使用ServeContent，而不能用ServeFile或者io.Copy

- Redirect函数，告诉客户端重定向到另一个URL

### 4、模板

模板就是写在动态页面中不变的部分，服务器程序渲染可变部分生成动态的网页

```go
//以下包提供了模板语言
import (
	"html/template"
	"text/template"
)
```

**步骤如下：**

**1、创建和解析模板**：

- 使用`template.New`创建一个新的模板（可选项，如果没有使用这个方法，在template.ParseFiles中会自动new一个新模板）。

- 使用`template.ParseFiles`或`template.ParseGlob`解析模板文件。

  ```go
  t,_:=template.ParseFiles("tmpl.html")	//写法一
  
  t:=template.New("tmpl.html")		//写法二
  t,_:=t.ParseFiles("tmpl.html")
  ```

**2、执行模板**：

- 使用`Execute`或`ExecuteTemplate`方法将数据应用到模板并生成最终输出。ExecuteTemplate多了一个name参数用于填写模版的名称

注：可以使用template.Must( ) 确保模板在解析过程中没有错误，如果有错误则会在程序初始化时崩溃。这对于开发阶段来说非常有用，因为它可以让你在程序启动时立即发现模板解析的问题，而不是在运行时。

```go
func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("cmd/day8/index.gohtml"))
	tmp.Execute(w, "helloworld")
}
```

#### 4.1、函数解析：

**template.ParseFiles("文件名")**

源码内容：先把模板文件里的内容读成字符串，然后使用文件名新建一个模板，最后调用parse方法来解析模板里的字符串

parsefiles的参数数量可变，但是只返回一个模板，当解析多个文件时，第一个文件作为返回的模板，其余作为map供后续使用

**template.ParseGlob**

使用模式匹配来解析特定文件，使用第一个文件名作为模板明

```go
t,err:=template.ParseGlob("*.html")		//匹配文件夹下所有html文件
```

**tmpl.Parse(s)** 

主要作用是将模板字符串 `s` 解析为模板定义，并将其附加到现有模板 `tmpl` 上。解析过程中，模板字符串被转换为一个模板树结构，表示模板的逻辑和内容。最终，解析后的模板树被添加到模板对象中，使其可以用于渲染动态内容。

**Lookup方法**

通过模板名来寻找模板，如果没找到返回nil

#### 4.2、变量传入

- **struct** 

```go
type User struct {
	UserName string
	Age      int
}
func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("cmd/day8/index.gohtml"))
	U.UserName = "jack"
	U.Age = 18
	tmp.Execute(w, U)
    //第二种写法
    //date := User{
	//	UserName: "jack",
	//	Age:      18,
	//}
	//tmp.Execute(w, date)
}
```

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>GET Form Example</title>
</head>

<body>
    {{.Age}} </br>
    {{.UserName}}
</body>

</html>
```

对于struct变量的赋值，只需要在HTML中的模板变量与Struct中的字段相同就行

- **map[string]interface{}**

```go
func (U *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("cmd/day8/index.gohtml"))
	//使用map
	M := map[string]interface{}{
		"game":   "Genshin",
		"player": "Tom",
	}
	tmp.Execute(w, M)
}
```

#### 4.3、加载多个页面

```go
package main

import (
	"html/template"
	"net/http"
)

func main() {
	templates := loadTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[1:]
		t := templates.Lookup(filename)
		t.Execute(w, nil)
	})
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.ListenAndServe("localhost:8080", nil)
}

func loadTemplates() *template.Template {
	result := template.New("templates")
	template.Must(result.ParseGlob("template/*.html"))
	return result
}

```

1、使用`template.New`创建一个新的模板

2、result.ParseGlob("template/*.html")解析所有html文件到模板中

3、根据url使用lookup函数找到对应的模板，然后Execute

#### 4.4、action

action就是go模板中嵌入的命令，位于{{}}之间，Action可以分为5类：条件类、迭代/遍历类、设置类、包含类、定义类

- **条件类**

```go
{{if .number}}
number is true
{{else}}
number is false
{{end}}
```

```html
举例：
<body>
    {{if .number}}
    number is true
    {{else}}
    number is false
    {{end}}
</body>
```

- **迭代/遍历Action**

  用于遍历数组、slice、map或者channel等数据结构，{{.}}用于表示每次迭代循环中的元素

```go
{{range array}}
do something {{.}}
{{else}}
arry is empty
{{end}}
```

举例：

```go
//....................main.go......................................
day:=[]string{"mon","tue","wed","thu","fri","sat","sun"}
t.Excute(w,day)
//.....................html.......................................
<body>
    {{range .}}
    number is {{.}}
	{{else}}
	. is empty
    {{end}}
</body>
```

- **设置Action**

  在这个with范围内将当前上下文替换为arg

  ```go
  {{with arg}}
  do something
  {{end}}
  ```

- **包含Action**

  允许在模板中包含其他的模板

  ```go
  {{template "name"}}
  ```


**在action中设置变量**

举例如下：

```go
 {{range $key,$value:=.}}
 the key is {{$key}} and the value is {{$value}}
 {{end}}
```

#### 4.5、函数与管道

- **管道**

管道是按顺序连接到一起的参数、函数、方法

例如：{{p1 | p2 | p3 }}

p1,p2,p3要么是参数，要么是函数，管道允许我们把参数的输出发给下一个参数，下一个参数由管道分隔开。

- **函数**

**内置函数：**

define、template、block

html，js，urlquery。对字符串进行转义，防止安全问题，如果是web模板，那么就不会经常使用

index

print、printf、println

len

with

**自定义函数：**

1、创建一个FuncMap（map类型）

key是函数名，value是函数

2、把FuncMap附加到模板

举例：

```

```

#### 4.6、组合模板

//TODO

#### 4.7、json

使用json将前后端的数据进行交互

- map[string]interface{}可以存储任意json对象
- []interface{}可以存储任意的json数组

**解码器：**

dec:=json.NewDecoder(r.Body) 	参数需要实现Reader接口

在解码器上进行解码：dec.Decode(&query)

**编码器：**

enc:=json.NewEncoder(w) 	参数需要实现Writer接口

编码：enc.Encode(results)

举例：

```go
func registerComRoute() {
	http.HandleFunc("/company", compaines)
}

func compaines(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		company := new(Company)
		err := dec.Decode(company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
```

## 十二、Gin框架

### 项目创建基本步骤

1、连接数据库

2、创建一个默认的 Gin 路由器

3、加载HTML和静态资源

4、编写handler/中间件，设置路由，使用中间件

5、运行项目

### 示例：main.go

```go
package main

import (
	"Gin/config"
	"Gin/route"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接数据库
	db := config.ConnectDB()
	defer db.Close()
	//gin框架
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// 加载HTML模板
	r.LoadHTMLGlob("template/*")
	// 加载静态文件（如JS、CSS）
	r.Static("/static", "./static")

	// 路由设置
	route.SetUserRoute(r)

	r.Run(":8080")
}
```

### 1、连接数据库

```go
package config

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "123456"
	network  = "tcp"
	server   = "127.0.0.1"
	port     = "3306"
	database = "gin"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, server, port, database)
	var err error
	DB, err = sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("数据库连接失败:", err.Error())
		return nil
	}

	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		fmt.Println("数据库连接失败:", err.Error())
		return nil
	}

	return DB
}
```

具体参考第十三章

### 2、创建默认路由器

#### gin.Default()

`gin.Default()` 是创建 Gin 应用程序的默认入口，使用New()返回一个已经初始化好的 `Engine` 对象，自动加载了两个常用中间件：

- **Logger 中间件**：记录每个请求的日志信息，包括 HTTP 方法、路径、状态码、执行时间等。它是帮助你监控和调试应用的一个工具。
- **Recovery 中间件**：处理应用程序中的 `panic`（运行时崩溃）情况。它会自动恢复程序并返回 500 错误，而不会导致服务器崩溃。

<font color="red">在我们的示例中使用Default来创建一个默认的Gin路由器</font>

### 3、加载HTML等前端资源

#### LoadHTMLGlob

- **作用**: `LoadHTMLGlob` 用于加载符合指定通配符模式的 HTML 模板文件。它通常用于 HTML 模板渲染，支持模板文件的全局加载。

- **使用场景**: 当你有多个 HTML 文件并想要全局加载时，可以使用此函数。

```go
r.LoadHTMLGlob("templates/*")
```

#### LoadHTMLFiles

- **作用**: `LoadHTMLFiles` 用于加载指定路径的 HTML 模板文件。与 LoadHTMLGlob 不同，LoadHTMLFiles 需要手动列出每个文件路径。
- **使用场景**: 当你只想加载特定的模板文件时，可以使用此函数。

```go
r.LoadHTMLFiles("templates/index.tmpl", "templates/about.tmpl")
```

#### Static

```go
func (engine *Engine) Static(relativePath, root string)
```

**作用**: `Static` 用于将服务器上的某个目录作为静态文件服务。例如，指定某个目录作为静态文件的根目录，所有文件都可以直接通过 HTTP 请求访问。

**使用示例**:

```go
r := gin.Default()
r.Static("/assets", "./static") // 将 ./static 目录映射到 /assets 路径
r.Run()
```

访问 `/assets/css/style.css` 时，会加载 `./static/css/style.css` 文件。

### 4、路由、中间件

#### 什么是中间件

Gin 提供了 `Use()` 方法来注册中间件

在 Gin 框架中，中间件是请求处理链中的一部分，**在请求到达最终的处理函数之前，或在响应返回客户端之前，执行的一段可插拔的逻辑。**中间件的基本作用可以分为以下几个步骤：

- **拦截请求**：当请求到达服务器时，中间件会先于路由处理函数执行，拦截并预处理请求。
- **执行逻辑**：中间件可以执行特定的任务，比如身份验证、请求日志记录、参数校验等。
- **调用下一个处理函数**：中间件可以决定是否将请求传递给下一个中间件或路由处理函数。这个通过 `c.Next()` 方法实现。
- **修改响应**：在请求处理函数完成后，中间件可以修改响应的内容，比如添加响应头、处理错误响应等。

源码实际上会把中间件注入到当前从属的routeGroup中，通过append方法追加到group.Handlers切片中

**全局中间件**会应用到所有的路由中。你可以在创建路由器后立即注册：

```go
func main() {
    router := gin.Default()

    // 注册一个全局中间件
    router.Use(MyMiddleware())

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.Run(":8080")
}
```

你也可以在路由组上注册中间件，限制中间件只应用到该组中的路由：

```go
func main() {
    router := gin.Default()

    // 创建一个路由组并为其注册中间件
    v1 := router.Group("/v1")
    v1.Use(MyMiddleware())
    {
        v1.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "pong",
            })
        })
    }

    router.Run(":8080")
}
```

中间件也可以在单个路由上使用：

```go
func main() {
    router := gin.Default()
    
    // 为单个路由注册中间件
    router.GET("/ping", MyMiddleware(), func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.Run(":8080")
}
```

**中间件的执行顺序**

- Gin 中的中间件按注册顺序执行。也就是说，先注册的中间件会先执行。
- 当你在中间件中调用 `c.Next()` 时，控制权会传递给下一个中间件或处理函数。如果你不调用 `c.Next()`，请求将不会继续向下传递。

<font color="red">在我们的示例中use了一个自定义的middleware，在处理 GET 请求之前都会记录每次请求的 URL 和处理时间</font>

#### 路由注册

这些方法是 Gin 框架中 `RouterGroup` 结构体的常用快捷方法，用于注册处理不同 HTTP 请求方法的路由。每个方法都是对 `handle` 方法的简化封装，通过不同的 HTTP 方法（如 GET、POST、PUT 等）来定义路由。

**`POST` 方法**

```go
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodPost, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP POST 请求。POST 一般用于提交数据，如表单提交、文件上传等。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 POST 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodPost`。

示例：

```go
router.POST("/submit", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "submitted",
    })
})
```

**`GET` 方法**

```go
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodGet, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP GET 请求。GET 一般用于获取数据，不会修改服务器上的资源。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 GET 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodGet`。

**`DELETE` 方法**

```go
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodDelete, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP DELETE 请求。DELETE 请求一般用于删除服务器上的资源。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 DELETE 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodDelete`。

**`PATCH` 方法**

```go
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodPatch, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP PATCH 请求。PATCH 通常用于更新资源的一部分。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 PATCH 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodPatch`。

**`PUT` 方法**

```go
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodPut, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP PUT 请求。PUT 通常用于替换服务器上的资源或创建新资源。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 PUT 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodPut`。

**`OPTIONS` 方法**

```go
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodOptions, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP OPTIONS 请求。OPTIONS 请求主要用于探测服务器支持的 HTTP 方法，通常用于跨域请求中的预检请求。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 OPTIONS 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodOptions`。

**`HEAD` 方法**

```go
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) IRoutes {
    return group.handle(http.MethodHead, relativePath, handlers)
}
```
- **功能**: 用于处理 HTTP HEAD 请求。HEAD 请求与 GET 类似，但只返回响应头部，不包含响应体。通常用于检查资源是否存在。
- **参数**:
  - `relativePath`：相对于当前路由组的路径。
  - `handlers`：一个或多个处理函数，用于处理该路径上的 HEAD 请求。
- **底层实现**: 调用了 `group.handle()`，并将 HTTP 方法设置为 `http.MethodHead`。

**底层的 `handle` 方法**

所有这些快捷方法最终调用了 `handle()` 方法，这是 Gin 中注册路由的核心逻辑。`handle` 方法将不同的 HTTP 方法和处理函数关联到具体的路由上。

```go
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}
```
1. **参数解析**：

   - `httpMethod`：HTTP 方法（如 `GET`、`POST` 等），表示当前路由对应的 HTTP 请求类型。
   - `relativePath`：相对路径（如 `/login`），是注册路由时使用的路径。
   - `handlers`：一个 `HandlersChain`，包含了要处理这个路由的一系列中间件和处理函数。

2. **`absolutePath := group.calculateAbsolutePath(relativePath)`**：

   - 这一步将相对路径转换为绝对路径。`RouterGroup` 可以分组路由，而这个方法会结合当前路由组的基础路径（`group` 的基础路径）和传入的 `relativePath` 来生成完整的绝对路径。

3. **`handlers = group.combineHandlers(handlers)`**：

   - 这里将当前路由组的全局中间件和传入的 `handlers` 合并。Gin 支持为不同的路由组定义中间件，这一步是为了确保路由的处理链上包含所有的中间件和处理函数。

4. **`group.engine.addRoute(httpMethod, absolutePath, handlers)`**：

   - 这一行调用了 `engine.addRoute` 方法，将构建好的完整路由注册到 Gin 的路由器中。Gin 的 `engine` 负责维护整个路由树，在请求到来时会根据注册的路由进行匹配和处理。

5. **`return group.returnObj()`**：

   - 最后，返回当前 `RouterGroup`（或者实现了 `IRoutes` 接口的对象）。这种设计允许链式调用，例如：
     ```go
     r.GET("/login", loginHandler).POST("/register", registerHandler)
     ```

#### Egine引擎

`Engine` 实现的 `ServeHTTP` 方法是 Gin 框架中的核心方法之一，用于处理所有进入的 HTTP 请求。它实现了 `http.Handler` 接口，这是 Go 的标准库 `net/http` 中定义的接口，任何实现了该接口的结构体都可以作为 HTTP 服务器的处理器。

```go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}
```

1. **从对象池获取 `Context` 对象**: 
   `c := engine.pool.Get().(*Context)` 从 `sync.Pool` 对象池中获取一个 `Context` 对象。`Context` 是 Gin 框架中用于管理请求和响应的核心数据结构，每个请求都会关联一个 `Context` 对象。

   Gin 框架使用了 Go 标准库中的 `sync.Pool` 对象池来管理 `Context` 对象。对象池的目的是复用 `Context` 对象，减少内存分配的开销。当有请求到来时，`pool.Get()` 会从对象池中取出一个可用的 `Context` 对象。如果池中没有可用对象，会自动创建一个新的。

   `(*Context)` 是类型断言，表示从池中取出的对象必须是 `*Context` 类型

2. **初始化 `ResponseWriter`**: 
   `c.writermem.reset(w)` 重置 `Context` 的 `writermem`，它是 `ResponseWriter` 的包装器，用于写入 HTTP 响应。这里将传入的 `http.ResponseWriter` 赋值给 `Context` 的 `writermem`。

   `writermem` 是 `Context` 内部封装的 `ResponseWriter`，它是处理 HTTP 响应的核心部分。`reset(w)` 将传入的 `http.ResponseWriter` 进行封装，方便 Gin 在处理中对响应进行操作，如设置状态码、写入响应数据等。

3. **绑定请求对象**: 
   `c.Request = req` 将当前的 HTTP 请求 `req` 绑定到 `Context` 的 `Request` 字段，方便后续处理中使用请求的内容，如 URL、Header、Body 等。

4. **重置 `Context`**: 
   `c.reset()` 清空或初始化 `Context` 中的其他字段，使其可以重新用于新的请求处理。这是防止上一个请求的状态影响当前请求。

5. **调用核心的请求处理逻辑**: 
   `engine.handleHTTPRequest(c)` 是 Gin 框架中的核心请求处理函数。这个方法负责根据请求的 URL 和方法，匹配注册的路由和中间件，依次执行它们。其流程通常包括：

   - 匹配路由：找到与请求路径和方法相匹配的处理函数（路由）。
   - 执行中间件：如果当前路由注册了中间件，按照注册顺序执行它们。
   - 执行处理函数：最终执行路由处理函数，并将结果返回给客户端。

6. **将 `Context` 归还到对象池**: 
   `engine.pool.Put(c)` 将处理完的 `Context` 对象放回对象池中，以便后续请求复用。Gin 使用对象池可以避免频繁地创建和销毁 `Context` 对象，提高性能。

#### 路由组

在 Gin 框架中，**路由组**（**Router Group**）是一种组织和管理路由的方式，允许将多个路由按逻辑分组，尤其适用于有相同前缀或共享中间件的路由。路由组有助于简化和优化代码，尤其在处理复杂的项目时更为有用。

**1、路由组的基本使用**

Gin 的路由组可以通过 `Group` 方法创建。你可以为一组路由指定公共的 URL 前缀和中间件。

**示例：创建路由组**

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建路由组，前缀为 "/api"
	api := r.Group("/api")
	{
		// 所有路由都会有 "/api" 前缀
		api.GET("/users", getUsers)       // GET /api/users
		api.POST("/users", createUser)    // POST /api/users
		api.GET("/users/:id", getUserByID) // GET /api/users/:id
	}

	// 启动服务器
	r.Run(":8080")
}
```

2. **使用中间件与路由组**

你可以为某个路由组单独应用中间件，避免为每条路由手动应用。

**示例：在路由组中使用中间件**

```go
func main() {
	r := gin.Default()

	// 定义一个路由组
	admin := r.Group("/admin", AuthMiddleware())
	{
		// 使用了 AuthMiddleware 的路由组
		admin.GET("/dashboard", adminDashboard)
		admin.GET("/settings", adminSettings)
	}

	r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
```

在这个例子中，`/admin/dashboard` 和 `/admin/settings` 都使用了 `AuthMiddleware` 中间件来验证请求头中的 token。

3. **嵌套路由组**

Gin 还支持嵌套路由组，允许你在一个路由组中再创建子组。每个子组继承父组的前缀和中间件，同时还可以添加新的前缀和中间件。

**示例：嵌套路由组**

```go
func main() {
	r := gin.Default()

	// API 路由组
	api := r.Group("/api")
	{
		// 用户相关的路由组
		users := api.Group("/users")
		{
			users.GET("/", getUsers)          // GET /api/users
			users.GET("/:id", getUserByID)     // GET /api/users/:id
			users.POST("/", createUser)        // POST /api/users
		}
		// 文章相关的路由组
		articles := api.Group("/articles")
		{
			articles.GET("/", getArticles)      // GET /api/articles
			articles.GET("/:id", getArticleByID) // GET /api/articles/:id
			articles.POST("/", createArticle)   // POST /api/articles
		}
	}

	r.Run(":8080")
}
```

路由结构：

- `/api/users` - 用户相关路由
- `/api/articles` - 文章相关路由

 **4、动态和静态路由结合**

你可以通过路由组组合动态和静态路由，处理带参数的路径。

**示例**：Param可以获取参数

```go
func main() {
	r := gin.Default()

	// 路由组
	products := r.Group("/products")
	{
		// 静态路由
		products.GET("/", getAllProducts)         // 获取所有产品
		// 动态路由
		products.GET("/:id", getProductByID)      // 获取特定 ID 的产品
	}

	r.Run(":8080")
}

func getAllProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all products"})
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get product with ID " + id})
}
```

### 5、handler的编写

#### 映射前端传递的值

**1、ShouldBindJSON **

**前端发送数据**：

- 假设前端通过表单或Ajax发送注册数据，数据的格式通常是JSON。例如：

```json
{
  "username": "testuser",
  "password": "123456"
}
```

**结构体字段的映射**

```go
type User struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}
```

- **`json:"username"`**：表示当Gin处理JSON数据时，`username` 键的值会映射到 `Username` 字段。
- **`binding:"required"`**：这个标签告诉Gin框架，这个字段是必填的。如果请求体中缺少 `username` 或 `password`，Gin会自动返回错误。

**ShouldBindJSON **

```go
var newUser User
if err := c.ShouldBindJSON(&newUser); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
    return
}
```

- **`ShouldBindJSON`** 是Gin框架提供的一个函数，用于从请求体中解析JSON数据，并将其自动映射到结构体上。
- 参数 `&newUser` 代表 `User` 结构体的一个实例。Gin框架会根据 `User` 结构体的字段名称和绑定规则，将前端传递的JSON数据的键值对与结构体中的字段对应起来。在本例中newUser中的值就是前端传输过来的结果。
- **数据校验**：通过结构体标签（`binding:"required"`），Gin可以自动验证某个字段是否存在。如果 `username` 或 `password` 没有出现在请求的JSON数据中，那么 `ShouldBindJSON` 会返回错误，代码会响应 `400 Bad Request` 错误。



**2、PostForm**

`ctx.PostForm` 是 Gin 框架中用于从 **表单请求** 中获取数据的方法。与 `ShouldBindJSON` 不同，`PostForm` 是用于处理 **`application/x-www-form-urlencoded`** 或 **`multipart/form-data`** 类型的请求，通常用来接收来自表单提交的数据。

假设前端的HTML表单是这样：

```html
<form method="POST" action="/register">
    <input type="text" name="username" placeholder="Username">
    <input type="password" name="password" placeholder="Password">
    <button type="submit">Register</button>
</form>
```

后端可以使用 `ctx.PostForm` 获取到这些表单字段：

```go
func handleRegister(c *gin.Context) {
    // 使用 PostForm 获取表单中的 username 和 password 字段
    username := c.PostForm("username")
    password := c.PostForm("password")

    if username == "" || password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
        return
    }

    // 继续处理注册逻辑，例如检查用户名是否存在、插入数据库等
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
```

1. **`c.PostForm("username")`**：从 `POST` 请求体中获取表单字段 `username` 的值。  
2. **`c.PostForm("password")`**：类似地，从请求体中获取 `password` 字段。
3. 如果某个字段不存在或为空，返回一个 `400 Bad Request` 响应。

如果表单提交，数据会以以下格式发送：

```
POST /register HTTP/1.1
Content-Type: application/x-www-form-urlencoded

username=testuser&password=123456
```



**3 、两种方式的区别**

- **`ctx.PostForm`**：非常适合从标准表单中提取字段，主要用于处理表单提交的数据（`application/x-www-form-urlencoded` 或 `multipart/form-data`）。
- **`ShouldBindJSON`**：用于处理以 JSON 格式发送的数据（`application/json`）。

| 特性               | `ctx.PostForm`                                               | `ctx.ShouldBind`                     |
| ------------------ | ------------------------------------------------------------ | ------------------------------------ |
| **数据来源**       | 表单数据（`application/x-www-form-urlencoded` 或 `multipart/form-data`） | 表单、JSON、XML、文件上传等          |
| **数据处理**       | 手动提取每个表单字段                                         | 自动解析请求体，映射到结构体         |
| **适合场景**       | 简单的表单提交                                               | 更复杂的数据结构或需要自动绑定的场景 |
| **字段校验**       | 需要手动校验字段是否存在                                     | 通过 `binding` 标签自动校验          |
| **复杂性**         | 较低                                                         | 较高（但更强大）                     |
| **支持的数据类型** | 表单数据                                                     | JSON、表单、XML 等                   |

#### 返回前端的响应

方式一、

使用gin.H{}在加载html模版时就返回给前端进行加载。前端使用{{.字段}}来与参数进行对应

```go
// 返回商品详情页面，渲染 HTML 模板
c.HTML(http.StatusOK, "product_details.html", gin.H{
    "product": product,
})
```

```html
<div class="product-info">
        <h2>{{.product}}</h2>
</div>
```

方式二、前后端分离架构

如果你使用的是前后端分离的开发模式（例如，前端用 React、Vue 或 Angular，后端用 Gin 提供 REST API），那么后端就不需要负责渲染 HTML 页面，而是通过 API 返回 JSON 数据，前端根据这些数据动态渲染页面。

前端项目通过构建工具（例如 Webpack 或 Vite）打包成静态文件，Gin 只需提供静态文件服务即可：

```
go复制代码// 提供静态文件服务
router.Static("/", "./dist")

// Vue/React 等 SPA 项目路由重定向
router.NoRoute(func(c *gin.Context) {
    c.File("./dist/index.html")
})
```

这样 Gin 后端就只负责 API，前端页面通过静态文件直接提供。 

**使用 JavaScript 获取 JSON 并动态显示商品信息**

`details.html` 示例（使用 JavaScript 动态获取数据）：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Product Details</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <header>
        <h1>Product Details</h1>
    </header>

    <section class="product-details">
        <img id="product-image" src="" alt="Product Image" width="300">
        <h2 id="product-name"></h2>
        <p id="product-price"></p>
        <button>Add to Cart</button>
    </section>

    <footer>
        <p>&copy; 2024 ShopMaster. All rights reserved.</p>
    </footer>

    <script>
        // 获取当前页面 URL 中的商品 ID
        const productId = window.location.pathname.split("/").pop();

        // 使用 fetch 请求后端的 JSON 数据
        fetch(`/product/${productId}`)
            .then(response => response.json())
            .then(product => {
                document.getElementById('product-image').src = product.imagePath;
                document.getElementById('product-name').textContent = product.name;
                document.getElementById('product-price').textContent = `Price: $${product.prices}`;
            })
            .catch(error => console.error('Error fetching product data:', error));
    </script>
</body>
</html>
```

**解释**：

1. **`fetch` API**：这个前端代码会从后端的 `/product/:id` 路由中获取 JSON 数据，并使用 JavaScript 将商品信息动态插入到 HTML 页面中。
2. **动态获取 ID**：`window.location.pathname.split("/").pop()` 会从 URL 中提取商品的 `id`，然后发送 `fetch` 请求获取对应的商品数据。
3. **更新页面元素**：当 JSON 数据返回时，使用 JavaScript 更新商品图片、名称和价格。

### 6、cookies

```go
// 设置 Cookie, 有效期为1小时
    c.SetCookie("username", use.Username, 3600, "/", "localhost", false, true)
func ShowIndexPage(c *gin.Context) {
    // 从 Cookie 中获取用户名
    username, err := c.Cookie("username")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in"})
        return
    }

    // 渲染主页并传递用户名
    c.HTML(http.StatusOK, "index.html", gin.H{
        "username": username,
    })
}
```

### 7、文件上传

前端代码：

```html
<input type="file" id="avatarUpload" name="avatarUpload" accept="image/*" onchange="previewAvatar(event)" style="display: none;">            
```

后端：

```go
// 获取上传的文件
file, err := c.FormFile("avatarUpload")
//选择路径并保存
path := "./static/image/avatar/" + username + ".png"
c.SaveUploadedFile(file, path)
```

在 Gin 框架中，`c.FormFile` 和 `c.SaveUploadedFile` 是处理文件上传的两个重要函数。它们分别用于从请求中获取上传的文件和保存上传的文件。以下是详细解析：

`c.FormFile` 用于从请求中获取上传的文件。它的作用是提取上传的文件并将其作为 `*multipart.FileHeader` 类型返回。`*multipart.FileHeader` 包含了关于上传文件的各种信息，比如文件名、文件大小、MIME 类型等。

**函数签名**

```go
func (c *Context) FormFile(name string) (*multipart.FileHeader, error)
```

**参数**

- `name`: 表单字段的名称，通常是文件上传表单中的 `input` 元素的 `name` 属性。

**返回值**

- **`*multipart.FileHeader`**: 包含文件上传的信息，如文件名、文件大小等。
- **`error`**: 如果有错误（如未找到文件），会返回一个错误。

**示例**

假设你有一个表单字段名为 `avatar`，你可以这样获取上传的文件：

```go
file, err := c.FormFile("avatar")
if err != nil {
    // 处理错误
    c.String(http.StatusBadRequest, "File upload error: %s", err.Error())
    return
}
```



`c.SaveUploadedFile` 用于将 `*multipart.FileHeader` 中的文件保存到指定路径。这是将上传的文件存储到服务器上的实际操作。

**函数签名**

```go
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error
```

**参数**

- **`file`**: 要保存的文件，通常是从 `c.FormFile` 获取的 `*multipart.FileHeader`。
- **`dst`**: 文件保存的目标路径，包括文件名。

**返回值**

- **`error`**: 如果保存过程中出现错误，会返回一个错误。

**示例**

将文件保存到指定目录：

```go
// 获取文件
file, err := c.FormFile("avatar")
if err != nil {
    c.String(http.StatusBadRequest, "File upload error: %s", err.Error())
    return
}

// 保存文件
dst := "./static/uploads/" + file.Filename
if err := c.SaveUploadedFile(file, dst); err != nil {
    c.String(http.StatusInternalServerError, "File save error: %s", err.Error())
    return
}
```

**总结**

- **`c.FormFile(name string)`**: 从请求中提取文件，返回一个 `*multipart.FileHeader` 和一个 `error`。
- **`c.SaveUploadedFile(file *multipart.FileHeader, dst string)`**: 将文件保存到指定路径，返回一个 `error`。

**常见错误处理**

1. **文件未找到**: 确保 `name` 参数与上传表单中的 `input` 元素名称匹配。
2. **保存路径错误**: 确保 `dst` 指定的路径存在并且具有写权限。
3. **文件大小限制**: 如果文件过大，可能需要配置 `MaxMultipartMemory` 或其他限制参数。

通过这些函数，你可以轻松处理文件上传的过程，包括从请求中提取文件和将其保存到服务器上。

### 8、搜索





## 十三、数据库

### database/sql简介

我们可以把标准库中的database/sql包看作一个数据库操作抽象层，因为database/sql并不直接连接并操作数据库，而是为不同的数据库驱动提供了统一的API,对数据库的具体操作由对应数据库的驱动包来完成，访问不同的类型的数据库需要导入不同的驱动包，而所有驱动包都必须实现database/sql/driver包下的相关接口。这样做的好处在于，如果某一天我们想迁移数据库，比如从MySQL切换为Postgres，只需要更换驱动包+微调sql语句即可，而不需要把数据库操作的代码重写一遍。

### 1、连接数据库

```go
package main

import (
	"database/sql"
	"fmt"

	// 匿名导入
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "123456"
	network  = "tcp"
	server   = "127.0.0.1"
	port     = "3306"
	database = "testdb"
)

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, server, port, database) //root:123456@tcp(127.0.0.1:3306)/testdb
	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
	defer db.Close()
}
```

第一步：导入mysql的数据驱动

第二步：使用const定义好数据库的具体信息

第三步：使用open函数连接数据库

使用这里面的db来对数据库进行各种操作。

- **Open()函数**

database/sql包提供了Open()函数，定义如下

```go
func Open(driverName, dataSourceName string) (*DB, error) {}
```

open函数根据给定的数据驱动以及驱动专属的数据源打开一个数据库，驱动数据源一般包含数据库的名字以及相关链接信息

**driverName：**表示驱动名，指代数据驱动注册到database/sql时所使用的名字

**dataSourceName：**表示驱动特定的语法，它告诉驱动如何访问底层数据存储。包含数据库用户名，密码等信息

open函数并没有创建连接，他只是验证参数是否合法，然后开启一个单独的goroutine去监听是否需要建立新的连接。

### 2、CRUD操作

#### 单行查询

```go
func (db *DB) QueryRow(query string, args ...any) *Row {}
```

QueryRow方法：

- func(r *Row)Err()error

- func(r *Row)Scan(dest ...interface{})error

  将结果中的数据拷贝出来放到对应的变量中

举例：

```go
type Student struct {
	Id        int
	FirstName string
	LastName  string
	Birth     string
	Gender    string
}
func queryOne(id int) (stu Student, err error) {
	err = db.QueryRow("SELECT gender from students WHERE student_id=?", id).Scan(
		&stu.Gender)
	return
}
```

`Scan` 方法中的参数必须是指针（使用 `&` 符号），这样 `Scan` 方法才能将查询结果赋值给这些变量。

#### 多行查询

```go
func (db *DB) Query(query string, args ...any) (*Rows, error) {}
```

使用Query会返回Rows这个类型

Rows的方法如下：

在 Golang 中，`Rows` 类型是 `database/sql` 包中的一个类型，用于表示从数据库查询返回的多行结果集。`Rows` 类型有几个常用的方法，用于遍历和处理查询结果。以下是一些常用的方法及其解释：

1. **`Next`**

   ```go
   func (rs *Rows) Next() bool
   ```

   `Next` 方法用于将 `Rows` 的游标移动到下一行。如果有更多的行可读取，它返回 `true`；否则返回 `false`。通常在循环中使用此方法来遍历所有行。

2. **`Scan`**

   ```go
   func (rs *Rows) Scan(dest ...any) error
   ```

   `Scan` 方法用于将当前行的列值复制到传入的目标变量中。目标变量必须是指针，`Scan` 会将数据解码并存储在这些变量中。

3. **`Close`**

   ```go
   func (rs *Rows) Close() error
   ```

   `Close` 方法用于关闭 `Rows` 对象并释放与其相关的所有资源。使用完 `Rows` 对象后，应该总是调用此方法。

4. **`Err`**

   ```go
   func (rs *Rows) Err() error
   ```

   `Err` 方法返回迭代过程中遇到的错误。如果没有错误，返回 `nil`。在遍历完所有行之后，应该检查是否有错误发生。

5. **`Columns`**

   ```go
   func (rs *Rows) Columns() ([]string, error)
   ```

   `Columns` 方法返回当前行中所有列的列名列表。

6. **`ColumnTypes`**

   ```go
   func (rs *Rows) ColumnTypes() ([]*ColumnType, error)
   ```

   `ColumnTypes` 方法返回当前行中所有列的类型信息列表。

举例：

```go
// 多行查询
func queryMany(id int) (stu []Student, err error) {
	rows, err := db.Query("SELECT gender from students WHERE student_id>?", id)
	s := Student{}
	for rows.Next() {
		err = rows.Scan(&s.Gender)
		if err != nil {
			fmt.Println("错误")
		}
		stu = append(stu, s)		//将查询到的内容append到返回的数组中
	}
	return
}
```

#### 更新数据

EXEC函数介绍：

exec函数会执行写在里面的sql语句

这里编写了一个student结构体的方法，new一个结构体，将要修改的值赋给该结构体，然后调用这个方法，使用exec对数据库的内容进行更新

```go
// 更新数据,这是student结构体的一个方法
func (s *Student) UpdateDate() (err error) {
	result, err := db.Exec("UPDATE students set gender=? where student_id=?", s.Gender, s.Id)
	fmt.Println(result.RowsAffected())
	return
}
	//调用student的update方法更新数据
	gen.Gender = "Other"
	gen.Id = 1
	err = gen.UpdateDate()
	if err != nil {
		fmt.Println(err)
	}
```

## 十四、logrus日志框架

`logrus` 是一个流行的 Golang 日志库，它功能强大、易用且支持丰富的日志格式和日志级别，非常适合用在 Web 应用如你要开发的 Gin 购物平台中。下面我将逐步讲解如何使用 `logrus`，并展示一些常见的使用模式。

### 1. **安装 `logrus`**

首先，使用 Go 模块添加 `logrus` 依赖：
```bash
go get github.com/sirupsen/logrus
```

### 2. **基础使用**

最简单的用法是直接在程序中使用 `logrus` 进行日志记录：

```go
package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus 基本用法
	logrus.Info("This is an info message")
	logrus.Warn("This is a warning message")
	logrus.Error("This is an error message")
}
```

**输出**：
```
INFO[0000] This is an info message
WARN[0000] This is a warning message
ERRO[0000] This is an error message
```

### 3. **日志级别**

`logrus` 提供多种日志级别，可以根据重要性选择记录哪些类型的日志：

- `Trace`：最细粒度的日志，调试细节。
- `Debug`：调试信息，开发阶段使用。
- `Info`：一般信息。
- `Warn`：警告信息，表明可能存在问题。
- `Error`：错误信息，表明出现错误但不影响程序继续运行。
- `Fatal`：严重错误，程序会直接退出。
- `Panic`：类似于 `Fatal`，但是会抛出 panic。

```go
logrus.Trace("This is a trace message")
logrus.Debug("This is a debug message")
logrus.Info("This is an info message")
logrus.Warn("This is a warning message")
logrus.Error("This is an error message")
logrus.Fatal("This is a fatal message") // 会调用 os.Exit(1)
logrus.Panic("This is a panic message") // 会引发 panic
```

### 4. **设置日志级别**

默认情况下，`logrus` 的最低日志级别是 `Info`，你可以通过 `SetLevel` 函数调整日志输出级别，小于该级别将不会输出。

```go
logrus.SetLevel(logrus.DebugLevel) // 设置为输出 Debug 级别以上的日志
```

### 5. **日志格式**

`logrus` 支持多种日志格式，例如 JSON 格式和文本格式。默认情况下，它使用带有时间戳的纯文本格式。

- **JSON 格式化日志**：

```go
logrus.SetFormatter(&logrus.JSONFormatter{})

logrus.WithFields(logrus.Fields{
	"username": "jack",
	"action":   "login",
}).Info("User logged in")
```

**输出**：
```json
{"action":"login","level":"info","msg":"User logged in","time":"2024-09-12T14:23:12Z","username":"jack"}
```

- **自定义文本格式**：
  你可以自定义日志格式，例如移除时间戳或者修改日志字段顺序：

```go
logrus.SetFormatter(&logrus.TextFormatter{
	DisableTimestamp: true,
	FullTimestamp:    true,
})
```

### 6. **添加字段（Fields）**

`logrus` 支持在日志中添加自定义字段，非常适合记录上下文信息，比如用户名、请求 ID 等。

- **单条日志添加字段**：
  使用 `WithFields` 添加字段。常用于记录结构化日志信息。

```go
logrus.WithFields(logrus.Fields{
	"username": "alice",
	"module":   "authentication",
}).Info("User login attempt")
```

- **日志上下文添加字段**：
  如果你想为多个日志共享相同的字段，可以使用 `WithField` 创建带上下文的 `logrus` 实例。

```go
logger := logrus.WithField("request_id", "12345")
logger.Info("User accessed the profile page")
logger.Warn("User attempted an invalid action")
```

### 7. **日志输出**

默认情况下，`logrus` 将日志输出到标准输出 (`os.Stdout`)。你可以通过 `SetOutput` 改变输出位置，例如输出到文件。

```go
package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main() {
	// 将日志输出到文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Failed to open log file:", err)
	}
	logrus.SetOutput(file)

	logrus.Info("This log is written to a file")
}
```

### 8. **日志钩子（Hooks）**

`logrus` 提供了钩子机制，你可以在记录日志时执行自定义的操作。例如，可以将日志发送到远程服务器，或者根据日志级别写入不同的文件。

- **添加日志钩子**：
  可以通过实现 `logrus.Hook` 接口创建自定义钩子。也可以使用第三方钩子，如将日志发送到 ElasticSearch、Logstash、Kafka 等。

```go
type MyCustomHook struct{}

func (hook *MyCustomHook) Levels() []logrus.Level {
	// 钩子仅应用于 Error 级别及以上的日志
	return logrus.AllLevels
}

func (hook *MyCustomHook) Fire(entry *logrus.Entry) error {
	// 在此处实现钩子的功能，例如发送日志到远程服务
	// entry.Data 包含日志的所有字段
	fmt.Println("Custom Hook executed for log:", entry.Message)
	return nil
}

func main() {
	logrus.AddHook(&MyCustomHook{})
	logrus.Info("This is an info log")
	logrus.Error("This is an error log")
}
```

### 9. **日志分片**

使用 `lumberjack` 实现日志的分片（滚动日志）是个常见需求。将 `logrus` 和 `lumberjack` 结合使用，可以轻松实现日志文件自动切割。

```go
package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 配置日志分片
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "rolling.log",
		MaxSize:    10, // 日志文件最大尺寸 (MB)
		MaxBackups: 5,  // 保留旧日志文件个数
		MaxAge:     30, // 保留旧日志文件的最大天数
		Compress:   true, // 启用压缩
	})

	logrus.Info("This is a log with rolling file")
}
```

### 10. **结合 Gin 框架**

在 Gin 框架中，`logrus` 可以用作中间件来记录每个 HTTP 请求的详细信息：

```go
package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()

	// 使用 logrus 记录请求信息
	r.Use(func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录请求信息
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		logrus.WithFields(logrus.Fields{
			"status":   statusCode,
			"method":   method,
			"path":     path,
			"duration": duration,
			"ip":       clientIP,
		}).Info("Request completed")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
```

### 总结

- **设置日志级别**：使用 `SetLevel` 设置日志的输出级别。
- **选择日志格式**：`logrus` 支持 JSON、文本等格式。
- **添加上下文字段**：通过 `WithFields` 可以为日志添加自定义的上下文字段。
- **日志分片和输出**：`logrus` 可以轻松与文件、远程服务器等集成，同时可以使用 `lumberjack` 进行日志分片。
- **钩子**：通过钩子机制，可以扩展日志功能，如发送日志到远程系统。

这些特性让 `logrus` 非常适合在高并发 Web 应用中使用，尤其是对你的线上购物平台来说，提供灵活和高效的日志功能。

## 十五、标准库

### fmt

**`fmt.Scan()`**

- 功能：从标准输入中读取数据并将其存储到传递的变量中。`Scan()` 按空格换行分隔输入数据。
- 使用场景：适合读取多项输入，按空格分隔。

示例：

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("请输入名字和年龄：")
    fmt.Scan(&name, &age)
    fmt.Printf("名字：%s, 年龄：%d\n", name, age)
}
```

**`fmt.Scanf()`**

- 功能：从标准输入中读取数据，并根据提供的格式化字符串解析数据。它类似于 `scanf`，可以指定输入格式。
- 使用场景：适合需要对输入数据进行特定格式解析的场景，比如读取特定格式的日期或字符串。

示例：

```go
package main

import "fmt"

func main() {
    var year, month, day int
    fmt.Print("请输入日期（格式：YYYY-MM-DD）：")
    fmt.Scanf("%d-%d-%d", &year, &month, &day)
    fmt.Printf("输入的日期是：%d年%d月%d日\n", year, month, day)
}
```

**`fmt.Scanln()`**

- 功能：类似于 `fmt.Scan()`，但它读取输入时会以换行符为界限，而不是空格。它适合读取整行输入。
- 使用场景：当你希望每次输入一整行数据时，比如读取用户的一句话或完整的文件路径。

示例：

```go
package main

import "fmt"

func main() {
    var sentence string
    fmt.Print("请输入一句话：")
    fmt.Scanln(&sentence)
    fmt.Printf("你输入的是：%s\n", sentence)
}
```

**`fmt.Printf()`**

在Go语言中，`fmt.Printf` 函数用于格式化字符串并将其打印到标准输出。格式化字符串使用类似C语言的占位符，每个占位符以 `%` 开头，指定如何格式化相应的参数。以下是 `fmt.Printf` 中常见的占位符及其说明：

**通用占位符**

- `%v`：默认格式，适用于大多数类型。
- `%+v`：类似 `%v`，但会在结构体字段前添加字段名。
- `%#v`：值的 Go 语法表示。
- `%T`：值的类型的 Go 语法表示。
- `%%`：字面上的百分号（%）。

**布尔值**

- `%t`：布尔值（true 或 false）。

**整数**

- `%b`：二进制表示。
- `%c`：相应的 Unicode 码点所表示的字符。
- `%d`：十进制表示。
- `%o`：八进制表示。
- `%O`：带零前缀的八进制表示（如0o377）。
- `%q`：单引号围绕的字符字面值，必要时会进行转义。
- `%x`：十六进制表示，小写字母（a-f）。
- `%X`：十六进制表示，大写字母（A-F）。
- `%U`：Unicode格式：U+1234，等同于 "U+%04X"。
- `%#U`：Unicode格式，带字符：U+1234 'c'。

**浮点数和复数**

- `%b`：无小数部分、二进制指数的科学计数法，如 -123456p-78。
- `%e`：科学计数法，如 -1.234456e+78。
- `%E`：科学计数法，如 -1.234456E+78。
- `%f`：有小数点但无指数，如 123.456。
- `%F`：等价于 `%f`。
- `%g`：根据情况选择 `%e` 或 `%f` 以产生更紧凑的（无末尾零的）输出。
- `%G`：根据情况选择 `%E` 或 `%F` 以产生更紧凑的（无末尾零的）输出。

**字符串和字节切片**

- `%s`：字符串或字节切片的无解译的字节。
- `%q`：带双引号的字符串字面值，必要时会进行转义。
- `%x`：每个字节用两字符十六进制数表示（小写字母 a-f）。
- `%X`：每个字节用两字符十六进制数表示（大写字母 A-F）。

**指针**

- `%p`：十六进制表示，前缀 0x。

**宽度和精度**

- `%[index]`：参数索引，从1开始。
- `%[flags][width][.precision][verb]`：
  - `flags`：
    - `+`：总是输出数值的符号（对于数值类型）。
    - `-`：左对齐。
    - `#`：使用 `#` 号解释备用格式。
    - `0`：在数字左边填充0而不是空格。
    - ` `：空格前缀。
  - `width`：输出的最小宽度。
  - `precision`：数值的精度，浮点数的小数点后的位数，或者字符串的最大字符数。

**示例**

以下是一些示例代码，展示了不同的格式化占位符：

```go
package main

import (
    "fmt"
)

func main() {
    // 通用占位符
    fmt.Printf("%v\n", 123)               // 123
    fmt.Printf("%+v\n", struct{ Name string }{"Alice"})  // {Name:Alice}
    fmt.Printf("%#v\n", struct{ Name string }{"Alice"})  // struct { Name string }{Name:"Alice"}
    fmt.Printf("%T\n", 123)               // int
    fmt.Printf("%%\n")                    // %

    // 布尔值
    fmt.Printf("%t\n", true)              // true

    // 整数
    fmt.Printf("%b\n", 123)               // 1111011
    fmt.Printf("%c\n", 123)               // {
    fmt.Printf("%d\n", 123)               // 123
    fmt.Printf("%o\n", 123)               // 173
    fmt.Printf("%O\n", 123)               // 0o173
    fmt.Printf("%q\n", 123)               // '\u007b'
    fmt.Printf("%x\n", 123)               // 7b
    fmt.Printf("%X\n", 123)               // 7B
    fmt.Printf("%U\n", 123)               // U+007B
    fmt.Printf("%#U\n", 123)              // U+007B '{'

    // 浮点数和复数
    fmt.Printf("%b\n", 123.456)           // 1111011.0111000000111110011010101001011110001101010011111011111001110111111111011p+6
    fmt.Printf("%e\n", 123.456)           // 1.234560e+02
    fmt.Printf("%E\n", 123.456)           // 1.234560E+02
    fmt.Printf("%f\n", 123.456)           // 123.456000
    fmt.Printf("%F\n", 123.456)           // 123.456000
    fmt.Printf("%g\n", 123.456)           // 123.456
    fmt.Printf("%G\n", 123.456)           // 123.456

    // 字符串和字节切片
    fmt.Printf("%s\n", "Hello, World!")   // Hello, World!
    fmt.Printf("%q\n", "Hello, World!")   // "Hello, World!"
    fmt.Printf("%x\n", "Hello")           // 48656c6c6f
    fmt.Printf("%X\n", "Hello")           // 48656C6C6F

    // 指针
    fmt.Printf("%p\n", &struct{}{})       // 0x<address>

    // 宽度和精度
    fmt.Printf("%6d\n", 123)              // "   123" (6个字符宽度)
    fmt.Printf("%.2f\n", 123.456)         // "123.46" (小数点后两位)
    fmt.Printf("%6.2f\n", 123.456)        // "123.46" (总宽度为6，小数点后两位)
    fmt.Printf("%-6.2f\n", 123.456)       // "123.46 " (左对齐，总宽度为6，小数点后两位)
    fmt.Printf("%06d\n", 123)             // "000123" (总宽度为6，前缀填充0)
}
```

以上示例展示了各种占位符的用法及其效果。根据需要选择合适的占位符，可以实现对不同类型数据的格式化输出。

### bufio.Scanner

`bufio.Scanner` 是 Go 标准库中的一种简便高效的输入读取工具，尤其适合从文本或输入流（如 `os.Stdin`）中逐行或按自定义的分隔符读取数据。它常用于处理输入行（例如从标准输入读取一行），逐词分割，或者处理较小的文本文件。

**使用步骤**：

1. 创建一个 `Scanner`，将其与需要读取的输入源关联。
2. 使用 `Scan()` 方法迭代输入，直到输入结束或发生错误。
3. 用 `Text()` 方法获取当前扫描的文本内容。
4. 检查是否有错误，并处理结果。

```go
scanner := bufio.NewScanner(os.Stdin) // 从标准输入创建一个 Scanner
```

`bufio.NewScanner` 可以接受任何实现了 `io.Reader` 接口的输入源，比如文件、标准输入、网络连接等。常用的输入源包括：

- `os.Stdin`: 从命令行标准输入读取。
- 打开的文件，如 `file, _ := os.Open("file.txt")`。

#### 2. **`Scan()`**
```go
for scanner.Scan() {
    // 获取扫描到的文本
    text := scanner.Text()
    fmt.Println(text)
}
```
`Scan()` 是一个重要的方法，它用于读取输入源中的下一段内容。每调用一次 `Scan()`，它会读取下一行或下一段文本，直到输入结束。通常用于循环，直到扫描完成。

#### 3. **`Text()`**
```go
text := scanner.Text() // 获取当前行的文本内容
```
`Text()` 用于返回 `Scan()` 之后扫描到的内容。默认情况下，`Scanner` 以换行符为分隔符，因此 `Text()` 返回的是每次扫描到的文本行。

#### 4. **`Err()`**
```go
if err := scanner.Err(); err != nil {
    fmt.Println("读取时出错:", err)
}
```
在 `Scan()` 完成或结束时，可以使用 `Err()` 方法检查是否发生了错误。对于正常输入结束，错误值为 `nil`。

#### 5. **设置自定义分割函数 `Split()`**
默认情况下，`Scanner` 按行分割输入。如果你需要自定义分割规则（例如按空格分割），可以通过 `Split()` 方法来实现。

```go
scanner.Split(bufio.ScanWords) // 按单词分割
```

常见的分割函数：
- `bufio.ScanLines`: 默认值，按行读取（每行以换行符分隔）。
- `bufio.ScanWords`: 按单词读取（以空格或其他空白字符分隔）。
- `bufio.ScanBytes`: 按字节读取。

你也可以自己定义分割函数。

#### 完整的示例代码

以下代码从标准输入读取多行，按行输出：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("请输入几行内容（Ctrl+D 结束输入）：")

	// 使用 scanner.Scan() 循环读取每一行输入
	for scanner.Scan() {
		text := scanner.Text() // 获取当前行文本
		fmt.Println("你输入的是:", text)
	}

	// 检查扫描是否出错
	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入时出错:", err)
	}
}
```

示例运行：

```
请输入几行内容（Ctrl+D 结束输入）：
Hello, world!
你输入的是: Hello, world!
This is Go programming.
你输入的是: This is Go programming.
Ctrl+D
```

#### 读取并按空格分割的示例

下面的代码演示了如何使用 `bufio.ScanWords` 来逐个单词扫描输入：

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // 将扫描器设置为按单词分割

	fmt.Println("请输入一些单词，按空格分割（Ctrl+D 结束输入）：")

	for scanner.Scan() {
		word := scanner.Text() // 获取当前单词
		fmt.Println("你输入的单词:", word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取输入时出错:", err)
	}
}
```

示例运行：

```
请输入一些单词，按空格分割（Ctrl+D 结束输入）：
Hello Go Language
你输入的单词: Hello
你输入的单词: Go
你输入的单词: Language
Ctrl+D
```

#### 使用自定义分割函数

你还可以创建自己的分割函数。以下示例将输入按逗号分割：

```go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	input := "Hello,Go,Language"

	scanner := bufio.NewScanner(strings.NewReader(input))

	// 自定义分割函数，按逗号分割
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 查找逗号的位置
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("扫描时出错:", err)
	}
}
```

示例运行：

```
Hello
Go
Language
```

#### 总结

- `bufio.Scanner` 是 Go 中处理输入流的强大工具，适用于从标准输入或文件中逐行、逐词、逐字节读取数据。
- 常用方法包括 `Scan()`、`Text()`、`Err()` 和 `Split()`。
- 可以通过自定义 `Split()` 来实现灵活的数据分割。

### log日志

Go 语言内置的 log 包实现了简单的日志服务 log 包为我们封装了一系列日志相关方法

`log` 包是Go语言标准库中的一个用于记录日志的包。它提供了简单而强大的日志记录功能，可以在控制台或文件中记录日志信息。以下是 `log` 包的常见用法和一些示例：

#### 常见用法

1. **记录普通日志**:
   - `log.Print`、`log.Println` 和 `log.Printf` 用于记录普通日志信息。

   ```go
   log.Print("This is a log message")
   log.Println("This is a log message with a new line")
   log.Printf("This is a formatted log message: %s", "Hello, World!")
   ```

2. **记录错误日志并退出**:
   - `log.Fatal`、`log.Fatalln` 和 `log.Fatalf` 在记录日志后会调用 `os.Exit(1)` 退出程序。

   ```go
   log.Fatal("This is a fatal error message")
   log.Fatalln("This is a fatal error message with a new line")
   log.Fatalf("This is a formatted fatal error message: %s", "Fatal Error")
   ```

3. **记录错误日志并引发 panic**:
   - `log.Panic`、`log.Panicln` 和 `log.Panicf` 在记录日志后会引发 panic。

   ```go
   log.Panic("This is a panic error message")
   log.Panicln("This is a panic error message with a new line")
   log.Panicf("This is a formatted panic error message: %s", "Panic Error")
   ```

4. **设置日志前缀**:
   - 使用 `log.SetPrefix` 设置日志消息的前缀。

   ```go
   log.SetPrefix("INFO: ")
   log.Println("This is an informational message")
   ```

5. **设置日志标志**:
   - 使用 `log.SetFlags` 设置日志的输出格式。常见的标志包括：
     - `log.Ldate`：日期 (2009/01/23)
     - `log.Ltime`：时间 (01:23:23)
     - `log.Lmicroseconds`：微秒级时间 (01:23:23.123123)
     - `log.Llongfile`：完整文件名和行号 (/a/b/c/d.go:23)
     - `log.Lshortfile`：文件名和行号 (d.go:23)
     - `log.LUTC`：使用 UTC 时间

   ```go
   log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
   log.Println("This is a log message with custom flags")
   ```

6. **将日志输出到文件**:
   - 使用 `log.SetOutput` 设置日志的输出目的地，例如文件。

   ```go
   file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
       log.Fatal(err)
   }
   defer file.Close()
   log.SetOutput(file)
   log.Println("This is a log message written to a file")
   ```

#### 示例

以下是一个完整的示例，展示了如何使用 `log` 包记录不同类型的日志信息并将日志写入文件：

```go
package main

import (
    "log"
    "os"
)

func main() {
    // 设置日志前缀
    log.SetPrefix("LOG: ")
    // 设置日志标志
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    // 将日志输出到文件
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    log.SetOutput(file)

    // 记录普通日志
    log.Println("This is an informational message")

    // 记录格式化的日志
    log.Printf("This is a formatted log message: %s", "Hello, World!")

    // 记录错误日志并退出
    log.Fatal("This is a fatal error message")
}
```

通过这个示例，你可以看到如何使用 `log` 包的各种功能来记录和管理日志信息。 `log` 包简单易用，但功能足够强大，适用于大多数日志记录需求。

### 标准库heap

Golang 标准库中的 `container/heap` 提供了堆（Heap）的实现，主要用于优先队列的构建和操作。它使用数组表示一个完全二叉堆，并支持最小堆或最大堆的操作。

#### 如何使用

1、想要使用heap构建最大堆，必须手动实现sort下的三个方法Len()、Swap()，Less()。以及heap对应的Push和Pop方法。以下是一个标准模版。

```go
// 定义一个类型来实现最大堆
type MaxHeap []int

// Len 返回堆中元素的数量
func (h MaxHeap) Len() int {
	return len(h)
}

// Less 反转比较逻辑，父节点需要比子节点大
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // 大于表示最大堆
}

// Swap 交换两个元素的位置
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 将元素添加到堆中
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop 移除并返回堆顶元素
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // 保存堆顶元素
	*h = old[0 : n-1] // 从切片中移除堆顶
	return x
}
```

2、构建最大堆：使用heap.Init()函数       

```go
func main() {
	h := &MaxHeap{3, 1, 4, 1, 5, 9, 2, 6, 5}
	heap.Init(h)            // 初始化堆（堆化）
}
```

3、向最大堆中插入或删除元素

```go
// 向堆中添加新元素
heap.Push(h, 10)
fmt.Println("添加元素 10 后的最大堆：", *h)

// 从堆中移除堆顶元素
max := heap.Pop(h).(int)
fmt.Println("移除堆顶元素：", max)
```

#### 源码分析

核心代码heap.Init()

```go
func Init(h Interface) {
    n := h.Len()
    for i := n/2 - 1; i >= 0; i-- {
        down(h, i, n)
    }
}
```

从堆的最后一个非叶子节点（`n/2 - 1`）开始，逐层向下调整，确保每个子树符合堆的性质。

使用 `down` 函数递归调整子树。

down函数就是常规的的生成堆的代码，将根节点与叶节点进行比较选取最大的与根交换，然后递归对交换了的节点进行同样的操作。

除了down还有一个up函数思路相同，不同点是up函数是从叶子结点开始向上调整，down是从根向叶子结点进行调整。

### 标准库 list



### 标准库sort

Go语言的 `sort` 包提供了对切片和自定义集合的排序功能。它支持基本类型（如整数、浮点数、字符串）的内置排序，同时也允许自定义排序逻辑。

#### `sort.Interface` 。

这个接口是sort包的核心内容其要求实现三个方法：

1. **`Len()`**：返回集合的长度。
2. **`Less(i, j int) bool`**：判断索引 `i` 的元素是否应该排在索引 `j` 的元素前面。
3. **`Swap(i, j int)`**：交换索引 `i` 和 `j` 的元素。

#### `sort.Ints()：`

- 将一个整数切片按升序排序。
- **示例：**
  
  ```go
  ints := []int{3, 1, 4, 1, 5, 9}
  sort.Ints(ints)
  fmt.Println(ints) // 输出: [1, 1, 3, 4, 5, 9]
  ```

#### `sort.Float64s()：`

- 将一个 `float64` 类型的切片按升序排序。
- **示例：**
  
  ```go
  floats := []float64{3.14, 2.71, 1.41, 1.62}
  sort.Float64s(floats)
  fmt.Println(floats) // 输出: [1.41, 1.62, 2.71, 3.14]
  ```

#### `sort.Strings()`：

- 将字符串切片按字典序（ASCII码顺序）进行升序排序。
- **示例：**
  
  ```go
  strs := []string{"banana", "apple", "cherry"}
  sort.Strings(strs)
  fmt.Println(strs) // 输出: [apple, banana, cherry]
  ```

#### `sort.Slice()`：

- 允许对任意类型的切片进行自定义排序，通过传递排序条件实现。这个函数特别灵活，可以用于自定义对象排序。
- 这个方法可以只实现Interface接口的Less()方法，第一个参数是slice类型：
- **示例：**
  
  ```go
  people := []struct {
      Name string
      Age  int
  }{
      {"Alice", 30},
      {"Bob", 25},
      {"Charlie", 35},
  }
  sort.Slice(people, func(i, j int) bool {
      return people[i].Age < people[j].Age
  })
  fmt.Println(people)
  // 输出: [{Bob 25} {Alice 30} {Charlie 35}]
  ```

#### `sort.Sort()`：

- 用于排序自定义类型。需要实现 `sort.Interface` 接口，该接口要求实现 `Len()`, `Less()`, 和 `Swap()` 三个方法。
- **示例：**
  ```go
  type Person struct {
      Name string
      Age  int
  }
  
  type ByAge []Person
  
  func (a ByAge) Len() int           { return len(a) }
  func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
  func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
  
  people := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
  sort.Sort(ByAge(people))
  fmt.Println(people)
  // 输出: [{Bob 25} {Alice 30} {Charlie 35}]
  ```

#### `sort.IsSorted()` 

在 Go 语言的 `sort` 包中，`sort.IsSorted()` 是一个用于检查一个集合是否已经排序的函数。它可以帮助你确定某个切片是否按升序排序。

语法：

```go
func IsSorted(data Interface) bool
```

- `data` 参数是实现了 `sort.Interface` 接口的类型。
- 返回值为 `true` 表示 `data` 已经排序；返回 `false` 表示没有排序。

要使用 `sort.IsSorted()`，传入的类型必须实现 `sort.Interface` 接口。这个接口要求实现三个方法：
1. **`Len()`**：返回集合的长度。
2. **`Less(i, j int) bool`**：判断索引 `i` 的元素是否应该排在索引 `j` 的元素前面。
3. **`Swap(i, j int)`**：交换索引 `i` 和 `j` 的元素。

`sort` 包中已有的基本类型（如 `[]int`, `[]float64`, `[]string`）有现成的 `sort` 函数，直接可以使用 `IsSorted`。

#### `sort.Search()` 

`sort.Search()` 是一个通用的二分查找算法，用于在有序的数据中快速定位符合条件的元素。

**语法：**

```go
func Search(n int, f func(int) bool) int
```

- `n` 是搜索范围的长度（即数组或切片的长度）。
- `f` 是一个回调函数，它会在 `Search` 的过程中被多次调用，`f` 接受一个索引值 `i`，并返回一个布尔值，表示是否找到了符合条件的元素。

`Search()` 返回符合 `f(i) == true` 的最小索引 `i`，如果找不到这样的 `i`，则返回 `n`。

举例：

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 5, 7, 9}
	target := 3

	// 使用 sort.Search 查找第一个 >= 目标值的元素
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	if index < len(nums) && nums[index] == target {
		fmt.Printf("找到目标值 %d，索引为 %d\n", target, index)
	} else {
		fmt.Printf("没有找到目标值 %d\n", target)
	}
}
```

### OS文件读写

```go
// FileInfo描述文件，由[Stat]返回。
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() any           // underlying data source (can return nil)
}
```

使用os.Stat(path)这个方法打开文件，然后就可以对其进行各种操作

`os.Stat` 是 Go 语言标准库中的一个函数，用于获取文件或目录的信息。它返回一个 `os.FileInfo` 接口，该接口提供了文件或目录的详细信息，如文件大小、权限、修改时间等。如果指定的路径不存在或无法访问，它会返回一个错误。

```go
//举例：
func main() {
	path := "/home/kobayashi/GoCode/cmd/std/picture.png"
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("数据类型是%T \n", fileInfo)
	fmt.Println("文件名是：", fileInfo.Name())
	fmt.Println("是否为目录", fileInfo.IsDir())
	fmt.Println("文件大小：", fileInfo.Size())
	fmt.Println("文件权限：", fileInfo.Mode())
	fmt.Println("文件最后修改时间：", fileInfo.ModTime())
	fmt.Println("文件基础资源:", fileInfo.Sys())
}
```

| 方法             | 作用               |
| ---------------- | ------------------ |
| filepath.IsAbs() | 判断是否是绝对路径 |
| filepath.Rel()   | 获取相对路径       |
| filepath.Abs()   | 获取绝对路径       |
| filepath.Join()  | 拼接路径           |

#### 创建目录

使用os.Mkdir("文件夹名",权限) 		权限是drw

如果是多级目录则要使用MkdirAll

```go
path1 := "filename"
err = os.Mkdir(path1, os.ModePerm)	//ModePerm表示0777
if err != nil {
    log.Println(err.Error())
}
```

文件权限由三个部分组成：所有者权限、组权限和其他用户权限。每个部分都可以包含读、写和执行权限：

- 读权限（4）
- 写权限（2）
- 执行权限（1）

这三个部分的权限通常用三个八进制数字表示，例如 `0755`，0开头表示这是8进制数

#### 改变目录

使用os.Chdir(path)到达相应目录，使用os.Getwd获取当前目录路径

```go
//改变目录到path1
os.Chdir(path1)
pwd, err := os.Getwd() //获取当前目录的路径
if err != nil {
    log.Println(err.Error())
}
fmt.Println(pwd)
```

#### 创建文件

使用os.Creat(filename)会返回*os.File, `*os.File` 是一个代表文件的结构体指针，提供了多种文件操作方法，比如读、写、关闭等。

其实creat是调用了openfile函数

```go
func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

举例：

```go
_, err = os.Create("log.txt")
if err != nil {
    log.Println(err.Error())
}
```

#### 打开文件和关闭文件

os.open()方法本质上也是调用了OpenFile这个方法，并且权限为o_RDONLY

```go
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```

这些关键字可以组合使用但是要用|分开

| 关键字   | 模式                                                         |
| -------- | ------------------------------------------------------------ |
| O_RDONLY | 以只读模式打开文件。                                         |
| O_WRONLY | 以只写模式打开文件。                                         |
| O_RDWR   | 以读写模式打开文件                                           |
| O_APPEND | 以追加模式打开文件。如果文件已存在，则在文件末尾写入数据。   |
| O_CREATE | 如果文件不存在，则创建一个新文件                             |
| O_EXCL   | 与 `os.O_CREATE` 一起使用，确保创建的文件不存在。如果文件已经存在，则返回错误。 |
| O_SYNC   | 以同步 I/O 的方式打开文件。                                  |
| O_TRUNC  | 如果文件已存在，则清空文件内容。                             |

#### 删除文件

**`os.Remove`**：仅用于删除文件或空目录。尝试删除非空目录将返回错误。

**`os.RemoveAll`**：递归删除目录及其所有内容，包括子目录和文件。它是删除整个目录树的有效方法，但要小心使用，以避免意外删除重要数据。

#### 读写文件

- **写入**

write将len(p)个字节从p中写入到基本数据流中。它返回从p中写入的字节数n，如果n<len(p)则会返回一个错误

```go
func (f *File) Write(b []byte) (n int, err error)
func (f *File) WriteString(s string) (n int, err error)
```

举例：

```go
file1.Write([]byte("abc123\n"))			//这里的\n表示换行
file1.WriteString("你好")
```

- **读取**

```go
func (f *File) Read(b []byte) (n int, err error) 
```

- **重置位置**

在读取文件时的位置不正确，导致尝试从文件当前位置读取数据时已经到达了文件末尾。写入操作之后并没有将文件的读取位置重置，因此读取操作会从文件的末尾开始，导致立即遇到 `EOF`（End of File）。

使用 `file1.Seek(0, 0)` 将文件的读取位置重新设置到文件的起始位置。这样可以确保能够从文件的开头正确读取数据。

```go
// 重置文件读取位置到起始位置
_, err = file1.Seek(0, 0)
if err != nil {
    log.Println("重置文件读取位置时出错:", err)
    return
}
```

`sync.Pool` 是 Go 语言中的一个高效的临时对象缓存池，用于在需要频繁分配和释放对象的场景中复用对象，减少垃圾回收的压力和内存分配的开销。它特别适合短期、频繁创建和销毁的对象，能够提升性能。

### **sync.Pool **

`sync.Pool` 主要是用来存储和复用一组不需要长时间保存的对象。当程序需要使用某个对象时，可以从 `sync.Pool` 中获取，而不是每次都分配新的内存；使用完之后，放回 `sync.Pool`，以备下次复用。

`sync.Pool` 的核心机制：

- 当程序需要一个对象时，可以从池中获取。
- 如果池中有可用的对象，则返回该对象；如果没有，则调用自定义的 `New` 函数来创建新对象。
- 当对象使用完毕后，可以将它放回池中，而不是直接丢弃或由垃圾回收机制回收。

#### 如何使用 sync.Pool

`sync.Pool` 使用起来非常简单，下面是一个基本的使用流程：

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个对象池
	pool := &sync.Pool{
		// 定义当 Pool 为空时如何生成新对象
		New: func() interface{} {
			fmt.Println("Creating new object")
			return "new object"
		},
	}

	// 从池中获取对象
	obj1 := pool.Get()
	fmt.Println("Got:", obj1)

	// 放回对象
	pool.Put("reusable object")

	// 再次从池中获取对象
	obj2 := pool.Get()
	fmt.Println("Got:", obj2)

	// 池中再次获取
	obj3 := pool.Get()
	fmt.Println("Got:", obj3)
}
```

**输出结果**：

```
Creating new object
Got: new object
Got: reusable object
Creating new object
Got: new object
```

**解释**：

1. **创建 `sync.Pool`**：我们创建了一个 `sync.Pool` 对象，并且通过 `New` 函数定义了如何在池为空时创建新对象。在这个例子中，当池中没有对象可用时，`New` 函数会返回一个字符串 `"new object"`。
   
2. **从池中获取对象 (`pool.Get()`)**：
   - 第一次调用 `pool.Get()` 时，池是空的，因此会调用 `New` 函数创建一个新对象。
   - 第二次调用 `pool.Get()`，我们在之前通过 `pool.Put()` 放入了一个 `"reusable object"`，因此这次直接从池中取出该对象，而不需要重新创建。
   - 第三次调用 `pool.Get()` 时，由于池已经没有可用对象了，因此再次调用 `New` 函数来创建一个新对象。

3. **放回对象 (`pool.Put()`)**：使用完对象后，我们可以调用 `pool.Put()` 将对象放回池中，方便下次复用。

####  sync.Pool 的作用场景

`sync.Pool` 最适合以下场景：

- **高频次的临时对象**：当你有很多短生命周期的对象时，通过 `sync.Pool` 复用对象可以显著减少内存分配和垃圾回收的压力。
- **避免频繁 GC**：在高并发环境中，如果频繁分配和释放内存，会导致大量的内存分配和垃圾回收。`sync.Pool` 可以减少这种分配，使垃圾回收更少被触发。
- **局部对象缓存**：如果某个对象在一段时间内被频繁使用，并且可以重用，则可以用 `sync.Pool`。

####  **注意事项**

- **短期缓存**：`sync.Pool` 是为了缓存短期对象的，它不能保证对象会一直存活。如果 `sync.Pool` 中的对象长期没有被使用，Go 的垃圾回收器可能会回收其中的对象。因此，`sync.Pool` 适合临时对象缓存，而不是长期保存数据的方式。
  
- **线程安全**：`sync.Pool` 是线程安全的，你可以在并发环境中使用 `sync.Pool` 而不需要额外的锁定机制。多个 Goroutine 可以同时从池中获取和放回对象。

- **对象的创建和销毁**：`sync.Pool` 本身不会主动创建对象，只有当你调用 `Get` 且池中没有对象时，才会通过 `New` 函数来创建。并且，池中的对象在被 GC（垃圾回收）触发时可能会被清理掉。

#### **高级使用场景**

**性能优化：减少内存分配**

在高并发下，如果你频繁地分配和销毁对象，使用 `sync.Pool` 可以显著降低内存分配的频率，从而提高性能。一个常见的场景是处理大量临时对象，比如缓冲区、字符串构建等。

```go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func process(data string) {
	// 从池中获取 buffer
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf) // 使用完后放回池中

	buf.Reset() // 重置 buffer 以便复用
	buf.WriteString(data)
	fmt.Println(buf.String())
}

func main() {
	for i := 0; i < 10; i++ {
		go process(fmt.Sprintf("Request #%d", i))
	}

	// 给 goroutine 足够的时间运行
	select {}
}
```

**解释**：

1. **缓存 `bytes.Buffer` 对象**：在这个例子中，我们使用 `sync.Pool` 缓存 `bytes.Buffer` 对象，用于处理字符串拼接操作。
2. **对象复用**：每次处理请求时，我们从池中获取一个 `bytes.Buffer`，使用完后将它放回池中，这样下次可以复用该对象，避免频繁创建新的 `bytes.Buffer`。

#### **适合与不适合使用 `sync.Pool` 的场景**

**适合：**

- 高并发下频繁创建和销毁对象的场景，比如字符串缓冲区、临时对象等。
- 需要降低内存分配和 GC 开销的场景。

**不适合：**

- 长期存储对象的场景，`sync.Pool` 并不能保证对象不会被回收。
- 对象较为复杂且初始化成本较高的场景，在这种情况下，反而可能会因为多次初始化消耗更多性能。

####  **总结**

- `sync.Pool` 是一个用于存储和复用临时对象的对象池，可以减少内存分配和垃圾回收的压力。
- 适用于短生命周期的高频次对象，能够显著提升性能。
- 在并发环境中使用非常方便，`sync.Pool` 是线程安全的。
- 但 `sync.Pool` 不适合长期存储对象，且对象可能会在 GC 时被回收，因此需要谨慎使用。

### sync.Mutex

#### 上锁机制

Mutex存在两种模式：普通模式 / 饥饿模式

- **普通模式**

  在普通模式下goroutine采用自旋不断尝试获取锁，如果竞争很激烈，自旋达到一定次数，goroutine会被挂起，等待信号量的唤醒。

  竞争情况：多个goroutine都在自旋等待获取锁，大家各自占得一部分的cpu时间片，这时锁空闲出来，谁在时间片上“刚好”执行获取锁的操作，谁的机会更大。

- **饥饿模式**

  饥饿模式下goroutine将不再自旋，等待锁的 goroutine 按信号量队列顺序唤醒，确保等待最久的 goroutine 先获取锁。

- **模式转换**

  当阻塞队列中存在goroutine等待时间过长(>1ms)，这时由普通模式转变为饥饿模式

  当阻塞队列已经清空或者已获得锁的goroutine等待时间低于1ms时，返回正常模式。

- **自旋**

  **自旋** 是一个忙等待过程，goroutine 在循环中不断尝试获取锁，期间会频繁调用 CPU 指令（例如 `CAS` 操作）。

  在 Go 的调度模型中，多个 goroutine 共享有限数量的操作系统线程（M 个 goroutine 会绑定到 N 个线程，M > N）。这些线程运行在多个 CPU 核心上，goroutine 的执行会被分配到某个线程的 CPU 时间片。

  自旋的好处：自旋避免了挂起和唤醒的上下文切换开销，适合锁持有时间较短的场景。但是自旋在高竞争场景下通常无效，反而浪费 CPU。

#### 源码分析

- **数据结构**

  ```go
  type Mutex struct {
  	state int32
  	sema  uint32
  }
  ```

  **`state`**：

  - 锁中最核心的状态字段，不同 bit 位分别存储了 mutexLocked(是否上锁)、mutexWoken（是否有 goroutine 从阻塞队列中被唤醒）、mutexStarving（是否处于饥饿模式）的信息

  **`sema` 字段**：

  - 使用信号量实现锁的阻塞与唤醒机制。

  - 当 goroutine 尝试获取锁但失败时，会挂起自身并等待信号量。

- **常量**

  ```go
  const (
      mutexLocked = 1 << iota // mutex is locked
      mutexWoken
      mutexStarving
      mutexWaiterShift = iota
  
      starvationThresholdNs = 1e6
  )
  ```

  - mutexLocked = 1：state 最右侧的一个 bit 位标志是否上锁，0-未上锁，1-已上锁；
  - mutexWoken = 2：state 右数第二个 bit 位标志是否有 goroutine 从阻塞中被唤醒，0-没有，1-有；
  - mutexStarving = 4：state 右数第三个 bit 位标志 Mutex 是否处于饥饿模式，0-非饥饿，1-饥饿；
  - mutexWaiterShift = 3：右侧存在 3 个 bit 位标识特殊信息，分别为上述的 mutexLocked、mutexWoken、mutexStarving；
  - starvationThresholdNs = 1 ms：sync.Mutex 进入饥饿模式的等待时间阈值.

`sync.Mutex` 的 `lockSlow` 方法用于处理复杂的锁获取逻辑，当快速路径（`lockFast`）失败时才会进入这个慢路径。它的设计考虑到了多个场景，包括自旋、饥饿模式、挂起和唤醒。以下是逐行详细解析：

#### **1. 函数开头**

```go
var waitStartTime int64
starving := false
awoke := false
iter := 0
old := m.state
```

- **`waitStartTime`**：记录当前 goroutine 开始等待锁的时间，用于判断是否进入饥饿模式。
- **`starving`**：标记当前 goroutine 是否处于饥饿状态。如果等待时间过长，则会尝试进入饥饿模式。
- **`awoke`**：标记当前 goroutine 是否被其他 goroutine 唤醒。
- **`iter`**：记录当前 goroutine 自旋的次数。
- **`old`**：缓存当前锁的状态，用于后续逻辑判断。

#### **2. 自旋逻辑**

```go
if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
    if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
        atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
        awoke = true
    }
    runtime_doSpin()
    iter++
    old = m.state
    continue
}
```

- **条件判断**：
  - 当前锁处于普通模式 (`mutexStarving` 未设置) 且已被占用 (`mutexLocked` 设置)。
  - `runtime_canSpin(iter)` 判断是否允许继续自旋，通常根据 CPU 核心数和自旋次数来决定。
- **自旋逻辑**：
  1. 如果当前 goroutine 未被唤醒（`awoke == false`），并且没有其他 goroutine 被唤醒（`mutexWoken == 0`），尝试设置 `mutexWoken` 标志，表示当前 goroutine 将被唤醒。
  2. 调用 `runtime_doSpin()` 执行自旋操作。
  3. 更新 `old` 变量为最新的锁状态，然后继续下一次循环。

#### **3. 锁状态更新逻辑**

```go
new := old
if old&mutexStarving == 0 {
    new |= mutexLocked
}
if old&(mutexLocked|mutexStarving) != 0 {
    new += 1 << mutexWaiterShift
}
if starving && old&mutexLocked != 0 {
    new |= mutexStarving
}
if awoke {
    if new&mutexWoken == 0 {
        throw("sync: inconsistent mutex state")
    }
    new &^= mutexWoken
}
```

- **普通模式下尝试获取锁**：
  - 如果当前锁未进入饥饿模式（`mutexStarving == 0`），将 `mutexLocked` 置为已占用。
- **增加等待者计数**：
  - 如果锁已被占用或处于饥饿模式，增加 `waiter` 计数（`mutexWaiterShift`）。
- **进入饥饿模式**：
  - 如果当前 goroutine 已进入饥饿状态（`starving == true`），并且锁已被占用，尝试将锁标记为饥饿模式（`mutexStarving`）。
- **清除唤醒标志**：
  - 如果当前 goroutine 被唤醒（`awoke == true`），需要清除 `mutexWoken` 标志。

#### **4. CAS 尝试获取锁**

```go
if atomic.CompareAndSwapInt32(&m.state, old, new) {
    if old&(mutexLocked|mutexStarving) == 0 {
        break // locked the mutex with CAS
    }
    ...
    runtime_SemacquireMutex(&m.sema, queueLifo, 1)
    starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
    old = m.state
    ...
}
```

- **成功获取锁**：
  - 如果 CAS 成功，并且当前锁是未被占用的普通模式，直接退出循环（`break`）。
- **挂起等待**：
  - 如果锁已被占用或处于饥饿模式，调用 `runtime_SemacquireMutex` 挂起当前 goroutine，进入等待队列。
  - 如果等待时间超过饥饿阈值（`starvationThresholdNs`），则将 `starving` 标记为 `true`。

------

#### **5. 饥饿模式修复逻辑**

```go
if old&mutexStarving != 0 {
    if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
        throw("sync: inconsistent mutex state")
    }
    delta := int32(mutexLocked - 1<<mutexWaiterShift)
    if !starving || old>>mutexWaiterShift == 1 {
        delta -= mutexStarving
    }
    atomic.AddInt32(&m.state, delta)
    break
}
```

- **饥饿模式检查**：

  - 如果锁处于饥饿模式（

    ```
    mutexStarving
    ```

    ），需要修复锁的状态：

    1. 确保锁未被占用（`mutexLocked` 未设置）。
    2. 确保等待队列中有至少一个 goroutine。

- **退出饥饿模式**：

  - 如果等待队列中只有一个 goroutine，或当前 goroutine 不再饥饿，退出饥饿模式（清除 `mutexStarving` 标志）。

------

#### **6. 自旋重置**

```go
awoke = true
iter = 0
```

- 每次被唤醒后重置自旋计数器，以防止无限循环。

------

#### **7. 使用举例**

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex // 定义互斥锁
	value int        // 共享资源
}

func (c *Counter) Increment() {
	c.mu.Lock()         // 加锁
	defer c.mu.Unlock() // 确保解锁

	c.value++ // 修改共享资源
}

func (c *Counter) GetValue() int {
	c.mu.Lock()         // 加锁
	defer c.mu.Unlock() // 确保解锁

	return c.value // 访问共享资源
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{} // 用来等待所有goroutine完成

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() // 确保goroutine完成时调用Done

			counter.Increment()
		}()
	}

	wg.Wait() // 等待所有goroutine完成
	fmt.Println("Final Counter Value:", counter.GetValue())
}

```

### sync.Map

`sync.Map` 是 Go 提供的一种并发安全的映射（map）。它属于标准库中的 `sync` 包，为多线程环境下的并发读写提供了无锁支持。

以下是 `sync.Map` 的详细介绍：

#### **特点**

1. **并发安全**：`sync.Map` 专为高并发场景设计，无需额外加锁即可安全访问。
2. **性能优化**：适合读多写少的场景，内部采用了高效的分段锁机制。
3. **弱类型**：`sync.Map` 的键值对类型是 `interface{}`，需要手动类型断言。

#### **与普通 map 的区别**

| 特性           | 普通 map           | `sync.Map`             |
| -------------- | ------------------ | ---------------------- |
| 并发安全性     | 不支持，需要加锁   | 内置支持               |
| 类型安全       | 强类型，编译时检查 | 弱类型，需断言         |
| 遍历时的一致性 | 遍历时可能会 panic | 保证并发环境下稳定遍历 |

------

#### **常用方法**

**1. `Store(key, value)`**

存储键值对。如果键已存在，则覆盖其值。

```go
m.Store("name", "Alice")
```

**2. `Load(key)`**

根据键取值，返回两个值：

- 键对应的值
- 是否存在的布尔值

```go
value, ok := m.Load("name")
if ok {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Key not found")
}
```

**3. `LoadOrStore(key, value)`**

如果键存在，返回已有的值及 `true`；如果键不存在，则存储新值并返回新值及 `false`。

```go
actual, loaded := m.LoadOrStore("name", "Bob")
fmt.Println("Value:", actual, "Already Loaded:", loaded)
```

**4. `Delete(key)`**

删除指定键值对。

```go
m.Delete("name")
```

**5. `Range(func(key, value interface{}) bool)`**

遍历所有键值对。参数是一个函数，返回 `false` 时停止遍历。

```go
m.Range(func(key, value interface{}) bool {
    fmt.Println("Key:", key, "Value:", value)
    return true // 继续遍历
})
```



------

#### **完整示例**

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 创建一个 sync.Map
	var m sync.Map

	// 存储键值对
	m.Store("name", "Alice")
	m.Store("age", 25)

	// 读取键值对
	if value, ok := m.Load("name"); ok {
		fmt.Println("Name:", value)
	}

	// LoadOrStore
	actual, loaded := m.LoadOrStore("location", "New York")
	fmt.Println("Location:", actual, "Already Loaded:", loaded)

	// 遍历键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // 继续遍历
	})

	// 删除键值对
	m.Delete("age")
	if _, ok := m.Load("age"); !ok {
		fmt.Println("Age key deleted")
	}
}
```

输出结果：

```
Name: Alice
Location: New York Already Loaded: false
Key: name, Value: Alice
Key: location, Value: New York
Key: age, Value: 25
Age key deleted
```

------

#### **使用场景**

1. **缓存**：在高并发场景中缓存计算结果或资源。
2. **配置数据**：存储共享配置数据。
3. **状态管理**：如连接状态、用户会话。

------

#### **注意事项**

1. **性能考虑**：`sync.Map` 并非适用于所有场景，在读写比例均衡或写操作频繁的场景下，普通 `map` 加锁（如 `sync.RWMutex`）可能表现更好。
2. **类型安全**：由于 `sync.Map` 使用 `interface{}`，需要手动进行类型断言，这可能导致运行时错误。
3. **初始化**：`sync.Map` 无需显式初始化，直接声明即可使用。

------

推荐在高并发场景中使用 `sync.Map`，特别是在读多写少的情况下，能够极大提升性能并简化代码复杂度。

## GC机制

### GO V1.3之前（标记清除）

 

### GO V1.5（三色标记）



### GO V1.8（混合写屏障机制）



## 注意事项

### 1、自动加分号

```go
func GetData()(int,int)	//这样会报错，因为go会自动在这行加分号，导致错误
{
	return 100,200
}
func GetData()(int,int){	//这样就不会有错
	return 100,200
}
```

### 2、golang语法糖...

#### a)函数参数

用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。 

```go
func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
}
```

结果

```go
qwr
234
yui
cvbc
```

strss切片内部的元素数量可以是任意个，test1函数都能够接受。

#### b)打散切片

strss2的元素被打散一个个append进strss

```go
var strss= []string{
    "qwr",
    "234",
    "yui",

}
var strss2= []string{
    "qqq",
    "aaa",
    "zzz",
    "zzz",
}
strss=append(strss,strss2...) 
fmt.Println(strss)
```

### 3、规范

- 在导入多个包时建议按如下格式：

```go
import(
	"fmt"
	"strings"		//标准库包
	
	"myproject/models"
	"myproject/utils"	//程序内部包
	
	"github.com/gin-gonic/gin"	//第三方包
)
```

- 建议在每个包添加注释

例如：

```go
//util包，该包包含了项目公用的一些变量和函数
//创建人：LT
//最后修改时间：2024.06.07
```

- 包命名

包的名称尽量和目录保持一致，不和标准库冲突，且使用小写字母
