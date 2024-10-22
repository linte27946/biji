# Django常用命令

## 创建项目

```python
django-admin startproject bysms
```

 `bysms` 就是项目的根目录名，执行上面命令后，就会创建 如下的目录结构：

## 运行 Django web服务

```
python manage.py runserver
```

## 创建项目app

```
python manage.py startapp sales 
```

`sales`就是app的名字

## 多级路由表

![image-20240121224543128](/media/kobayashi/新加卷/myblog/typora-user-images/Django/image-20240121224543128.png) 

- 使用include来包含子表，凡是以sales/开头的都交由子路由表进行查找

**子表**

![image-20240121224621743](/media/kobayashi/新加卷/myblog/typora-user-images/Django/image-20240121224621743.png) 

## 创建数据库表

第一步：

```
python manage.py startapp common 
```

创建一个目录名为 common， 对应 一个名为 common 的app，里面包含了如下自动生成的文件。

```
├── admin.py
├── apps.py
├── __init__.py
├── migrations
│   └── __init__.py
├── models.py
├── tests.py
└── views.py
```

第二步：

在models.py模块中写入要添加的表，CharField对应 varchar类型的数据库字段， `max_length`  指明了该 varchar字段的 最大长度。

```
from django.db import models

class Customer(models.Model):
    # 客户名称
    name = models.CharField(max_length=200)

    # 联系电话
    phonenumber = models.CharField(max_length=200)

    # 地址
    address = models.CharField(max_length=200)
```

第三步：

在settings.py中INSTALLED_APPS里面添加

```
'common.apps.CommonConfig',
```

第四步：更新表

```
python manage.py makemigrations 
```

## 官方字段类型

https://docs.djangoproject.com/en/2.0/ref/models/fields/#model-field-types

## 更新数据库

```
python manage.py migrate
```

## 创建管理员账号

```
python manage.py createsuperuser
```

