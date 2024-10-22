# c++STL 输入输出流

## 输入流

|    iostream(数据)     |     fstream(文件)      |          sstream(字符串)          |
| :-------------------: | :--------------------: | :-------------------------------: |
| istream从流中读取数据 | ifstream从文件读取数据 | istringstream从字符串对象读取数据 |
| ostream向流中写入数据 | ofstream向文件写入数据 | ostringstream向字符串对象写入数据 |
|   iostream 读/写流    |    fstream读/写文件    |    stringstream读/写字符串对象    |



## 输出流

**每个输出流都管理一个缓冲区，用于暂存读、写的数据。有了缓冲机制，操作系统就可以将程序的多个输出操作合并成一个单一的系统级写操作，从而带来性能上的提升。**

- **每次将数据写到输出设备或文件中的操作被称为缓冲刷新**

  **以下是三种操作符介绍：**

```c++
cout<<"hello"<<endl;   //输出hello后换行，刷新缓冲区
cout<<"hello"<<flush;  //输出hello后直接刷新缓冲区
cout<<"hello"<<ends;   //输出hello和一个空字符后，刷新缓冲区
```

# 输入输出流iostream

## **一、流对象get函数**

```c++
#include<iostream>
using namespace std;
int main()
{
    char ch[20];
    //test1
    cout<<"输入字符串:";
	int n=cin.get();
	cout<<"ASCII码为："<<n<<endl;
	cin.getline(ch,20);
	cout<<"该字符串是："<<ch<<endl;
    
	//test2
    cout<<"输入字符串:";
	cin.getline(ch,20,'\\');
	cout<<"该字符串是："<<ch<<endl;
	return 0; 
}
```

**结果：**

```
输入字符串:2023 hello world!!
ASCII码为：50
该字符串是：023 hello world!!
输入字符串:1234\7890
该字符串是：1234
```

- **int get( );    读取一个字符 (包含空白字符)，返回该字符的ASCII码 **

  **int n=cin.get();  将n的值赋予2的ACSCII码值50**

- **istream&get(char* ch,int n, char c='\n')   读取n-1个字符（包含空白字符），存入数组中，遇到c则提前结束**

- **istream&getline(char* ch,int n, char c='\n')    与get类似，<font color='red'>区别在于get遇到终止符时，读位置将停留在终止符前，下次从该位置读取；而getline则跳过终止符</font>。二者都返回输入流对象。   **

- **如果将getline函数替换为get结果将变为,具体原因如上红色部分**

```
输入字符串:2023 hello world!!
ASCII码为：50
该字符串是：023 hello world!!
输入字符串:1234\7890
该字符串是：
1234
```

## 二、利用流错误信息

 本来要输入int数据，结果却包含字符，该如何处理错误呢？

- bool eof( ):  当到达流末尾时返回true
- <font color='red'>bool fail( );  I/O操作失败，遇到非法数据时返回true，流可以继续使用。</font>
- bool bad( );  遇到致命错误，流不能继续使用时返回true

**如何发现错误**

```c++
#include<iostream>
using namespace std;
int main()
{
    int a;
    while(1)
    {
    	cin>>a;
    	if(cin.fail()) 
    	{
    		//--------输入非法数据-----------
			cout<<"输入有误"<<endl;
			cin.clear();        //清空流状态位
			cin.sync() ;        //清空流缓冲区
		}
	}
	return 0; 
}
```

# 文件I/O流

## 一、打开文件

- 创建流对象的同时打开文件

```c++
ifstream in("test.dat",iso::in);  //缺省路径
```

- 先创建流对象再打开文件

```c++
ofstream out1,out2;
out1.open("d:\\test.dat",iso::out);  //绝对路径
out1.open("..\\test.dat",iso::out);  //相对路径
```

## 二、文件模式

| 文件模式标志 | 含义                                                         |
| :----------: | :----------------------------------------------------------- |
|   ios::app   | 追加:在文件尾部写入                                          |
|   ios::ate   | 最后:打开文件后定位到文件末尾                                |
| ios::binary  | 以二进制方式读取或写入文件                                   |
|   ios::in    | 只读方式打开文件。如果文件不存在，打开将失败                 |
|   ios::out   | 写入方式打开文件。如果文件不存在，则创建一个给定名称的空文件 |
|  ios::trunc  | 截断: 如果打开的文件存在，其内容将被丢弃，其大小被截断为零   |

可以使用二元运算符“ | ”进行组合 

例如 ios::in |ios::out |ios::binary 表示以二进制方式打开文件并对文件进行读/写

## 三、文件关闭

c++中同时打开的文件数量是有限的，所以请及时关闭文件

```c++
ofstream out("demo.text",iso::out); //将out与demo.txt文件关联
.....
out.close();   //用完后关闭
out.open("text.dat",ios::out | ios::binary) //将out与test.dat文件关联
```

## 四、写文本示例

使用流提取符>>和流插入符<<完成对文本的读写，基本操作方法和流对象cin、cout一样

```c++
#include <iostream>
#include <fstream>
using namespace std;

struct student
{
    char name[8];
    int grade;
};

int main()
{
    ofstream out;
    out.open("D:\\datastruct\\a.txt");

    struct student s1 = {"张三", 80};
    struct student s2 = {"李四", 80};

    out << s1.name << "\t" << s1.grade << endl;
    out << s2.name << "\t" << s2.grade << endl;

    out.close();

    return 0;
}
```

执行完后D:\\datastruct\\a.txt文件内容会被改写为<font color='red'>（原先内容将被删除）</font>

张三	80
李四	80

## 五、以二进制方式读写文件

打开文件时设置ios::binary将会以二进制方式打开文件（默认是以文本方式），任何文件都可以二进制方式打开，在读写时不会像文本方式一样进行编码转换，所以不能使用<<和>>运算符函数，应该使用如下方式

```c++
iostream& read((char*) buffer,streamsize n);
iostream& write((char*) buffer,streamsize n);
```

流对象可以调用read( )成员函数将来自文件的n个字节读取到buffer（内存块首地址）中，也可以通过调用write（）成员函数将buffer中的n个字节写入到文件流所绑定的文件中

示例：

```c++
#include <iostream>
#include <fstream>
using namespace std;

struct student
{
    char name[8];
    int grade;
};

bool readfile(string &file, student *st);
void writefile(string &file, student *st);

bool readfile(string &file, student *st)
{
    ifstream in(file);
    if (!in)
    {
        return false;
    }
    in.read((char *)&st[0], sizeof(student));
    in.read((char *)&st[1], sizeof(student));
    in.close();
    return true;
}

void writefile(string &file, student *st)
{
    ofstream out(file);
    if (!out)
    {
        cout << "打开文件失败！";
        return;
    }
    out.write((char *)&st[0], sizeof(student));
    out.write((char *)&st[1], sizeof(student));
    out.close();
}

int main()
{
    student s[2] = {{"张三", 80}, {"李四", 80}};
    string filename = "D:\\datastruct\\a.txt";
    writefile(filename, s);
    if (!readfile(filename, s))
    {
        cout << "打开文件失败！";
        return 0;
    }
    cout << s[0].name << "\t" << s[0].grade << endl;
    cout << s[1].name << "\t" << s[1].grade << endl;
    return 0;
}

```



