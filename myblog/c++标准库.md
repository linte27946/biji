# c++标准库

# 字符串操作

## string类

### 初始化

- 初始化

```c++
string s="hello word";
```

- 重复字符

```c++
string s1(5,'a');
```

**注：拷贝构造是复制值，而移动构造类似于剪切数值，原先变量将变为空**

- 拷贝构造

```c++
格式：空代表默认全部
string s(字符串,起始位置,字符个数)
```

```c++
string s2(s1);
string s2(s1,4,3)
```

- 移动构造

```c++
string s3(move(s1));
```

- 字符串拼接

```c++
string s4=s1+s2;
string s4=s1+"hello word"
```

### 元素访问

- 类似于数组访问

```c++
string s="hello word"
cout<<s[4]<<endl	//结果为o
cout<<s[-1]<<endl;//没有越界警告,显示为空
```

- at 按字节进行访问，因为一个汉字占3字节如果使用at会输出乱码

```c++
s.at(1)='a';
s.at(100)='a';	//进行检查，会报错
cout<<s<<endl;	//结果为hallo word
```

- front（开头）/back（结尾）

```c++
cout<<s.front()<<s.back()<<endl;
```

### 容量操作

- 判断空串 empty()

```c++
string s;
cout<<s.empty()<<endl;	//s为空返回0，不为空返回1
```

- 求字节数 size()
```c++
string s="你好";	//中文占3字节
cout<<s.size()<<endl;	//输出结果为6
```

- resize()重新分配大小

```c++
//函数原型
resize(size_type __n, _CharT __c)
第一个参数表面重新分配的大小，第二个表明若超出长度则用字母填充
//举例：
string s="hello word("
s.resize(5,'A');	//结果hello
s.resize(20,'A');	//结果hello wordAAAAAAAAAA
```



- capacity()显示已经分配的空间大小

```c++
string s="hello world"
cout<<s.capacity()<<endl;
```

- reserve() 重新分配空间

```c++
s.reserve(200);
cout<<s.capacity()<<endl;	//输出200
```

- shrink_to_fit()

```c++
将多余分配的空间释放
s.shrink_to_fit();
```

 ### 迭代器

auto会自动判断元素的类型。

如果要修改element的值需要使用&符号，否则可以不使用

```c++
for (auto& element : container){}
```

- 正向迭代器

​		begin() ---> end()

​		string::iterator itr;	

```c++
string s="hello word"
string::iterator itr;	
for(itr=s.begin();itr!=s.end();itr++)
{
    cout<<*itr<<endl;
}
```

- 反向迭代器:使用同正向迭代器

​		rbegin() ---> rend()

​		string::reverse_iterator ritr;

![range-rbegin-rend.svg](/media/kobayashi/新加卷/myblog/typora-user-images/c++标准库/range-rbegin-rend.svg)

### 字符串比较

```
string s="hdllo world";
string s1="hello worldAAAAAA";
```

```c++
cout << boolalpha << (s == s1) << endl;	//相同输出1不同输出0
```

```c++
cout<<s.compare(s1)<<endl;	//从第一个字母开始比较ASCII码输出相差的值
//结果：1
```

### append追加

```c++
#include <iostream>
#include <string>
 
int main()
{
    std::basic_string<char> str = "string";
    const char* cptr = "C-string";
    const char carr[] = "Two and one";
 
    std::string output;
 
    // 1) 后附 char 3 次。
    // 注意，这是仅有的接受 char 的重载。
    output.append(3, '*');
    std::cout << "1) " << output << "\n";
 
    // 2) 后附整个字符串
    output.append(str);
    std::cout << "2) " << output << "\n";
 
    // 3) 后附字符串的一部分（下例后附最后 3 个字母）
    output.append(str, 3, 3);   //字符串，开始，个数
    std::cout << "3) " << output << "\n";
 
    // 4) 后附 C 字符串的一部分
    // 注意，因为 `append` 返回 *this，我们能一同链式调用
    output.append(1, ' ').append(carr, 4);
    std::cout << "4) " << output << "\n";
 
    // 5) 后附整个 C 字符串
    output.append(cptr);
    std::cout << "5) " << output << "\n";
 
    // 6) 后附范围
    output.append(std::begin(carr) + 3, std::end(carr));
    std::cout << "6) " << output << "\n";
 
    // 7) 后附初始化式列表
    output.append({' ', 'l', 'i', 's', 't'});
    std::cout << "7) " << output << "\n";
}
```

输出：

```c++
1) ***
2) ***string
3) ***stringing
4) ***stringing Two 
5) ***stringing Two C-string
6) ***stringing Two C-string and one
7) ***stringing Two C-string and one list
```

### 插入与删除

https://zh.cppreference.com/w/cpp/string/basic_string/insert

```c++
string&insert(size_t pos,const string &str);
//插入的位置，字符串
string& erase(size_t pos=0,size_t len=npos);
//删除的起始位置，个数
```

### 替换

replace(位置，要替换的个数，字符串)

```c++
string s="hello word";
string s1="你好世界"；
cout<<s.replace(2,3,s1);
//输出he你好世界 word
```

basic_string& replace( size_type pos, size_type count,

​            const basic_string& str,

​            size_type pos2, size_type count2 = npos );

```c++
cout<<s.replace(2,3,s1,0,3);
//输出he你 word
```



### find查找

```c++
#include <iomanip>
#include <iostream>
#include <string>
 
void print(int id, std::string::size_type n, std::string const& s)
{
    std::cout << id << ") ";
    if (std::string::npos == n)
        std::cout << "没有找到！n == npos\n";
    else
        std::cout << "在位置 n = " << n << " 找到，substr(" << n << ") = "
                  << std::quoted(s.substr(n)) << '\n';
}
 
int main()
{
    std::string::size_type n;
    std::string const s = "This is a string";  /*
                             ^  ^  ^
                             1  2  3           */
 
    // 从首个位置开始搜索
    n = s.find("is");
    print(1, n, s);
 
    // 从位置 5 开始搜索
    n = s.find("is", 5);
    print(2, n, s);
 
    // 寻找单个字符
    n = s.find('a');
    print(3, n, s);
 
    // 寻找单个字符
    n = s.find('q');
    print(4, n, s);
}
```

## C 标准库 <ctype.h>

这些函数用于测试字符是否属于某种类型，这些函数接受 **int** 作为参数，它的值必须是 EOF 或表示为一个无符号字符。

如果参数 c 满足描述的条件，则这些函数返回非零（true）。如果参数 c 不满足描述的条件，则这些函数返回零。

| 序号 | 函数 & 描述                                                  |
| ---- | ------------------------------------------------------------ |
| 1    | [int isalnum(int c)](https://www.runoob.com/cprogramming/c-function-isalnum.html) 该函数检查所传的字符是否是字母和数字。 |
| 2    | [int isalpha(int c)](https://www.runoob.com/cprogramming/c-function-isalpha.html) 该函数检查所传的字符是否是字母。 |
| 3    | [int iscntrl(int c)](https://www.runoob.com/cprogramming/c-function-iscntrl.html) 该函数检查所传的字符是否是控制字符。 |
| 4    | [int isdigit(int c)](https://www.runoob.com/cprogramming/c-function-isdigit.html) 该函数检查所传的字符是否是十进制数字。 |
| 5    | [int isgraph(int c)](https://www.runoob.com/cprogramming/c-function-isgraph.html) 该函数检查所传的字符是否有图形表示法。 |
| 6    | [int islower(int c)](https://www.runoob.com/cprogramming/c-function-islower.html) 该函数检查所传的字符是否是小写字母。 |
| 7    | [int isprint(int c)](https://www.runoob.com/cprogramming/c-function-isprint.html) 该函数检查所传的字符是否是可打印的。 |
| 8    | [int ispunct(int c)](https://www.runoob.com/cprogramming/c-function-ispunct.html) 该函数检查所传的字符是否是标点符号字符。 |
| 9    | [int isspace(int c)](https://www.runoob.com/cprogramming/c-function-isspace.html) 该函数检查所传的字符是否是空白字符。 |
| 10   | [int isupper(int c)](https://www.runoob.com/cprogramming/c-function-isupper.html) 该函数检查所传的字符是否是大写字母。 |
| 11   | [int isxdigit(int c)](https://www.runoob.com/cprogramming/c-function-isxdigit.html) 该函数检查所传的字符是否是十六进制数字。 |

标准库还包含了两个转换函数，它们接受并返回一个 "int"

| 序号 | 函数 & 描述                                                  |
| ---- | ------------------------------------------------------------ |
| 1    | [int tolower(int c)](https://www.runoob.com/cprogramming/c-function-tolower.html) 该函数把大写字母转换为小写字母。 |
| 2    | [int toupper(int c)](https://www.runoob.com/cprogramming/c-function-toupper.html) 该函数把小写字母转换为大写字母。 |

```c++
//只保留字符串中的字母，然后判断是否为回文
//示例 输入: s = "A man, a plan, a canal: Panama"
//	  输出：true
//    解释："amanaplanacanalpanama" 是回文串。

class Solution {
public:
    bool isPalindrome(string s) {
        string sgood;
        for (char ch: s) {
            if (isalnum(ch)) {
                sgood += tolower(ch);
            }
        }
        string sgood_rev(sgood.rbegin(), sgood.rend());
        return sgood == sgood_rev;
    }
};
```

## arry 容器

### 初始化

```c++
array<int, 5> ar = {1, 2, 3, 4, 5};
.....................................................
array<int, 5> ar;
ar.fill(111);
```

### 操作数组元素

```c++
ar[1];			//同数组
ar.at(1)		//at进行边界检查访问
array<int, 5>::iterator it=ar.begin();		//使用迭代器
//可以使用auto代替arry<int, 5>::iterator
```

### 自定义类型数组

```c++
#include <iostream>
#include <string>
#include <array>
using namespace std;

class student
{
private:
    string m_name;
    int m_age;

public:
    student();
    student(string name, int age);
    void show();
};
student::student() : m_name("null"), m_age(-1) {}
student::student(string name, int age) : m_name(name), m_age(age)
{
}
void student::show()
{
    cout << "名字:" << m_name << " 年龄:" << m_age << endl;
}
int main()
{
    array<student, 3> stu = {student("xiaom", 15), student("张三", 16)};
    auto it = stu.begin();	//直接将其当指针使用
    for (; it != stu.end(); it++)
    {
        it->show();
    }
}
```

## vector容器

### 初始化

```c++
vector<int> vec;	//无参构造
vector<int>size_type count;	//元素个数为count，值为0
vector<int> vec(3,100);		//元素个数为3，值为100

vec.reserve(10);	//重新分配内存大小为10字节
```

### push_back队尾增加元素

每次push_back会分配两倍插入之前的元素的内存，以方便后续的插入。

```c++
vector<int> vec(5);
vec.push_back(10);
//size为6，capacity为10，vec内容为0，0，0，0，0，10
```

### insert

方式同string。

```c++
vec.insert(vec.begin()+1,USER("xiaom",15));	//其返回为迭代器类型
```

### 原位构造

在使用push_back，insert插入元素时，都是拷贝构造或者移动构造，这里经历了两步，如果想要一步达成，即直接对vec元素赋值，就要用到原位构造

```c++
vec.emplace_back("xiaoming",12);	//代替push_back
//对比之前的insert
vec.insert(vec.begin()+1,USER("xiaom",15));
vec.emplace(vec.begin()+1,"xiaom",15);
```



## set集合

### 简单介绍：

头文件：#include \<set>

set会对其中的元素根据键值进行排序、去重，

排序实现原理：红黑树。

去重原理：默认是用比较函数!comp(a,b) && !comp(b,a)
### 初始化

```c++
set<int> st{1, 2, 3, 3, 7};		//方式1直接赋值
........................................................
vector<int> vec{1, 1, 4, 5, 1, 4, 3, 7};	//方式2使用迭代器赋值
set<int> st1(vec.begin(), vec.end());
...........................................................
set<int> st2(st);			//拷贝构造
set<int> st3=move(st);		//移动构造
```


### 迭代器

- 使用解引用运算符“*”取值，不用 first 和 second。
- 所有迭代器都是常量迭代器，不能用来修改所指元素。

注意在set中的迭代器是const类型的，因为在 C++ 的 `std::set` 中，元素是不可修改的，因为 `std::set` 是一个基于红黑树的容器，要保持元素的排序和唯一性。

### 基本操作

set的基本操作包括增、删、查、是否为空、求元素个数、交换两个容器的内容。

| 成员方法  | 功能介绍                                                     |
| --------- | :----------------------------------------------------------- |
| insert( ) | 向 set 容器中插入元素。                                      |
| erase( )  | 删除 set 容器中存储的元素。                                  |
| find( )   | 在 set 容器中查找值为 val 的元素，如果成功找到，则返回指向该元素的双向迭代器；反之，则返回和 end() 方法一样的迭代器。 |
| end( )    | 返回指向容器最后一个元素（注意，是已排好序的最后一个）所在位置后一个位置的双向迭代器 |
| begin( )  | 返回指向容器中第一个（注意，是已排好序的第一个）元素的双向迭代器。 |
| size( )   | 返回当前set中元素的个数                                      |
| swap( )   | 交换 2 个 set 容器中存储的所有元素。这意味着，操作的 2 个 set 容器的类型必须相同。 |

### 自定义的存储类型

如果set中使用的是自定义的类型，比如自定义的结构体，此时需要重载一下比较器或者自定义一个比较器。

举例如下：按字符串大小存储

```c++
#include <string>
#include <set>
#include <iostream>
class MyType
{
private:
    std::string data;
    int num;

public:
    MyType(const std::string &str, const int number) : data(str), num(number) {}

    const std::string &getData() const
    {
        return data;
    }
    const int &getnum() const
    {
        return num;
    }
    void print() const
    {
        std::cout << data << num << std::endl;
    }
};

// 自定义的比较器类
class MyTypeComparator
{
public:
    bool operator()(const MyType &lhs, const MyType &rhs) const
    {
        if (lhs.getData() != rhs.getData())
        {
            return lhs.getData() < rhs.getData();
        }
        return lhs.getnum() < rhs.getnum();
    }
};

void Print(std::set<MyType, MyTypeComparator> Myset)
{
    for (auto t = Myset.begin(); t != Myset.end(); t++)
    {
        t->print();
    }
}

int main()
{
    std::set<MyType, MyTypeComparator> mySet;

    // 插入一些元素
    mySet.insert(MyType("banana", 10));
    mySet.insert(MyType("banana", 5));
    mySet.insert(MyType("orange", 3));

    // 输出...
    Print(mySet);

    // 删除一些元素
    mySet.erase(MyType("banana", 5));
    // 输出...
    Print(mySet);

    // 查找
    auto a = mySet.find(MyType("orange", 3));
    if (a != mySet.end())
    {
        std::cout << "存在该元素" << std::endl;
    }

    // 查找元素具体位置
    int rank = std::distance(mySet.begin(), a) + 1; // 因为std::distance返回的是从起始位置到当前位置的距离，所以要+1
    std::cout << rank << std::endl;

    // 交换元素
    std::set<MyType, MyTypeComparator> yourSet;
    yourSet.swap(mySet);
    // 输出...
    std::cout << "输出yourset" << std::endl;
    Print(yourSet);
    std::cout << "输出myset" << std::endl;
    Print(mySet);

    return 0;
}
```

## map容器

### 简单介绍：

头文件 #include\<map>

map容器存储的都是pair对象(键值对)，各个键值对的键和值可以使任意数据类型(一般选用string字符串作为键的类型)。

map会根据键的大小对值进行排序(和set一样)。

map 容器中存储的各个键值对不仅键的值独一无二，键的类型也会用 const 修饰，这意味着只要键值对被存储到 map 容器中，其键的值将不能再做任何修改。

通过first和second访问键与值。

```c++
pair<const k, T>;
```

### 初始化

```c++
map<string, int> m = {{"张三", 15}, {"王5"}, 250};		//直接赋值
map<string, int> m1(m);		//拷贝复制
map<string, int> m2 = move(m);		//移动构造，构造完以后m为空
......................................................................
map<int, string> tempMap;
tempMap[1] = "One";			//通过[]来直接赋值，不存在则直接添加。
tempMap[2] = "Two";
```

### 基本操作

与set基本相同

```c++
m1.insert(pair<string, int>("张三", 15));		//使用insert插入时需要使用pair表明,当插入已经存在的键时则该语句失效同时vsode会报错但能继续运行。
m1.insert_or_assign("张三", 14);		//insert_or_assign当不存在时插入，存在则修改
...................................................................................
map<string, string, greater<string>>myMap;	//改变排序顺序
```

## C++ 标准库 `<stack>`

在 C++ 中，标准库提供了多种容器和算法来帮助开发者更高效地编写程序。

`<stack>` 是 C++ 标准模板库（STL）的一部分，它实现了一个后进先出（LIFO，Last In First Out）的数据结构。这种数据结构非常适合于需要"最后添加的元素最先被移除"的场景。

`<stack>` 容器适配器提供了一个栈的接口，它基于其他容器（如 `deque` 或 `vector`）来实现。栈的元素是线性排列的，但只允许在一端（栈顶）进行添加和移除操作。

### 基本操作

- `push()`: 在栈顶添加一个元素。
- `pop()`: 移除栈顶元素。
- `top()`: 返回栈顶元素的引用，但不移除它。
- `empty()`: 检查栈是否为空。
- `size()`: 返回栈中元素的数量。



## 算法库

头文件#include\<algorithm>



## 多线程

### thread线程库

一、启动一个线程

格式：thread  线程名称(启动的函数)；

- 使用lambda表达式

```c++
thread s1([&sum1, iter, iter_end]()
              {
                sum1 = visit(iter,iter_end,caculate);
                std::cout<<"thread 2 end"<<endl; });
s1.join();
```

- 直接启动

```
thread s1(function);
s1.join();
```

- 使用类方法中的函数

在`std::thread`的构造函数中，参数默认会被拷贝传递给线程函数，因此需要使用`std::ref`来明确指定引用传递。这点与函数参数引用不同。

```c++
class A
{
    void th(int &num)
    {}
	.........
}
A a;
int num=0;
thread th(&A::fn,&a,ref(num);	//thread 线程名（函数指针,对象,参数）
```

### mutex互斥量

一、包含头文件

```c++
#include<mutex>
```

二、lock()锁与unlock()解锁

对于需要互斥访问的变量，在多个线程中可以进行如下操作

```c++
mutex mtx；	//全局
mtx.lock();
......
mtx.unlock();
```

三、trylock尝试锁住

```c++
bool try_lock();	//返回bool类型
```

```c++
std::mutex mutex;
while (true)
{
    // 尝试锁定 mutex 以修改 'job_shared'，如果锁失败了就继续运行else，不会等待
    if (mutex.try_lock())
    {
        std::cout << "共享任务 (" << job_shared << ")\n";
        mutex.unlock();
        return;
    }
    else
    {
        // 不能获取锁以修改 'job_shared'
        // 但有其他工作可做
        ++job_exclusive;
        std::cout << "排他任务 (" << job_exclusive << ")\n";
        std::this_thread::sleep_for(interval);
    }
}
```

四、unique_lock()锁管理工具

```c++
mutex mtx；
unique_lock<mutex> lock(mtx,std::xxx);
```

| xxx             | 效果                               |
| :-------------- | :--------------------------------- |
| `defer_lock_t`  | 不获得互斥体的所有权               |
| `try_to_lock_t` | 尝试获得互斥体的所有权而不阻塞     |
| `adopt_lock_t`  | 假设调用方线程已拥有互斥体的所有权 |

五、call_once()只调用一次

多用于初始化操作，用于多次调用但只进行一次的操作。

```c++
once_flag fg;
call_once(fg,函数名称,参数)；
```

### condition_variable条件变量（**线程同步**）

一、引入头文件

```c++
#include<condition_variable>
```

二、wait等待

函数定义：

```c++
template<typename _Predicate>
      void
      wait(unique_lock<mutex>& __lock, _Predicate __p)
      {
	while (!__p())
	  wait(__lock);
      }
```

如果`__p`只是一个简单的布尔值，那么`__p()`实际上就是一个变量的值。而且，无论这个变量的值如何，`!__p()`都将会产生一个布尔值。这就导致了一个问题：无法在等待期间动态地检查条件，因为谓词只会在第一次调用时被评估，而之后每次调用都会得到同样的结果。

所以__p为一个可以调用的函数对象，将条件封装为一个可以调用的函数对象，以便在每次调用时都可以动态地评估条件。所以在这里使用lambda表达式返回bool值

```c++
condition_variable cv;
bool sub_run=false;
unique_lock<mutex> lock(mtx);
cv.wait(lock, [&]
                { return !sub_run; });
```

三、notify_all()唤醒

```c++
//在wait中等待的线程唤醒，
cv.notify_all();	//之后wait函数会再一次执行wait函数中的lambda表达式
```

举例：

```cpp
#include <iostream>
#include <thread>
#include <condition_variable>
#include <mutex>
#include <chrono>
using namespace std;

int i = 0;
mutex mtx;
condition_variable cv;
bool sub_run = false;

void tf()
{
    while (i < 10)
    {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [&]
                { return sub_run; });
        cout << "子线程：" << i << endl;
        i++;
        this_thread::sleep_for(chrono::milliseconds(10));
        sub_run = false;
        cv.notify_all();
    }
}

int main()
{
    thread th(tf);
    while (i < 10)
    {
        unique_lock<mutex> lock(mtx);
        cv.wait(lock, [&]
                { return !sub_run; });
        cout << "主线程：" << i << endl;
        i++;
        this_thread::sleep_for(chrono::milliseconds(10));
        sub_run = true;
        cv.notify_all();
    }
    th.join();
}
```

如果将cv.notify_all();注释掉

结果将会阻塞住

```c++
主线程：0
子线程：1
```

正常结果：

```c++
主线程：0
子线程：1
主线程：2
子线程：3
主线程：4
子线程：5
主线程：6
子线程：7
主线程：8
子线程：9
主线程：10
```

四、notify_one()唤醒一个线程

```c++
notify_one()		//随机唤醒一个线程
```

### future

一、头文件

```c++
#include<future>
```

二、应用场景：主线程为了获取子线程中的变量

```

```



### atomic原子操作

一、头文件

```c++
#include<atomic>
```

二、

atomic_int total(0);
