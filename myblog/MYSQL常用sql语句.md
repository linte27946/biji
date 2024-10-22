# MYSQL数据库使用sql语句

- **查看有哪些数据库**

```sql
SHOW DATABASES;
```

如果需要查看某个数据库的详细信息（如包含的表），你可以使用以下查询：

```sql
USE database_name;  -- 切换到你想查看的数据库
SHOW TABLES;        -- 显示该数据库中的所有表
```

## mysql使用

以下是一些常用的MySQL SQL语句，分为数据定义语言（DDL）、数据操作语言（DML）和数据控制语言（DCL）：

### 数据定义语言（DDL）

1. **创建数据库**

   ```sql
   CREATE DATABASE database_name;
   ```

2. **删除数据库**

   ```sql
   DROP DATABASE database_name;
   ```

3. **使用数据库**

   ```sql
   USE database_name;
   ```

4. **创建表**

   ```sql
   CREATE TABLE table_name (
       column1 datatype PRIMARY KEY,
       column2 datatype NOT NULL,
       column3 datatype DEFAULT value,
       ...
   );
   ```

5. **修改表**

   ```sql
   ALTER TABLE table_name ADD column_name datatype;
   ALTER TABLE table_name MODIFY column_name new_datatype;
   ALTER TABLE table_name DROP COLUMN column_name;
   ```

6. **删除表**

   ```sql
   DROP TABLE table_name;
   ```

### 数据操作语言（DML）

1. **插入数据**

   ```sql
   INSERT INTO table_name (column1, column2, column3, ...)
   VALUES (value1, value2, value3, ...);
   ```

2. **查询数据**

   ```sql
   SELECT column1, column2, ...
   FROM table_name
   WHERE condition
   ORDER BY column ASC|DESC
   LIMIT number;
   ```

3. **更新数据**

   ```sql
   UPDATE table_name
   SET column1 = value1, column2 = value2, ...
   WHERE condition;
   ```

4. **删除数据**

   ```sql
   DELETE FROM table_name
   WHERE condition;
   ```

### 数据控制语言（DCL）

1. **创建用户**

   ```sql
   CREATE USER 'username'@'host' IDENTIFIED BY 'password';
   ```

2. **删除用户**

   ```sql
   DROP USER 'username'@'host';
   ```

3. **授予权限**

   ```sql
   GRANT ALL PRIVILEGES ON database_name.* TO 'username'@'host';
   FLUSH PRIVILEGES;  -- 刷新权限
   ```

4. **撤销权限**

   ```sql
   REVOKE ALL PRIVILEGES ON database_name.* FROM 'username'@'host';
   FLUSH PRIVILEGES;  -- 刷新权限
   ```

### 其他常用SQL语句

1. **显示表结构**

   ```sql
   DESCRIBE table_name;
   ```

2. **显示数据库中的所有表**

   ```sql
   SHOW TABLES;
   ```

3. **显示创建表的语句**

   ```sql
   SHOW CREATE TABLE table_name;
   ```

这些SQL语句涵盖了MySQL中的基本操作，可以帮助你管理和操作数据库。根据具体需求，可以灵活运用这些语句。