# c语言结构体字节对齐

## 一、字节对齐规则

<font color='red'>通常情况下32位操作系统是4字节 ，64位操作系统是8字节</font>

- 当结构体中的<font color='red'>所有成员</font>字节长度没有超过基本字节单位长度时，按照<font color='red'>结构体</font>最大字节单位长度对其。

  ```c
  struct st2
  {
      char a;  //char占1字节
      char b;
      char c;
  };
  //32位和64位下, sizeof(struct st2)都是3个字节
  ```

- 当结构体的<font color='red'>成员字节</font>长度超过基本字节单位长度时，那么就按照系<font color='red'>统字节单位</font>来对齐。

  ```c
  struct st1
  {
      char name;  //char占1字节
      double age; //double占8字节
      char sex;
  };
  //32位下 sizeof(struct st1) = 16
  //64位下 sizeof(struct st1) = 24
  ```

  

## 三、为什么要对齐（以64位举例）

<img src="C:\Users\kobayashi\Desktop\myblog\963535203e43e716af3d594041c4ae6.png" alt="963535203e43e716af3d594041c4ae6" style="zoom: 33%;" />

所以说，字节对齐的根本原因其实在于cpu读取内存的效率问题，对齐以后，cpu读取内存的效率会更快。但是这里有个问题，就是对齐的时候这七个字节是浪费的，所以字节对齐实际上也有那么点以空间换时间的意思，具体写代码的时候怎么选择，其实是看个人的。

## **四、手动设置对齐**

`__attribute__((packed));`

```c++
//用法如下
struct abc
{
   char a;
   int b;
}__attribute__((packed)) abc,*ptr;//直接按照实际占用字节来对齐，其实就是相当于按照1个字节对齐了
//这里计算sizeof(abc)=5
```