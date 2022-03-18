# Todo List 备忘录

go开发练手项目，感谢[CocaineCong/TodoList](https://github.com/CocaineCong/TodoList)

## 接口文档

[接口文档链接](https://www.apifox.cn/apidoc/shared-0d6400a1-f53b-4226-81af-734cef97ddb0)

### 配置文件

在`conf/config.ini`下：

```
[service]
AppMode = debug
HttpPort = :3000

[mysql]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassword = root
DbName = todo_list_db
```

### 项目运行

1. 需要修改相应的配置，并新建`todo_list_db`数据库

2. 下载依赖

```go mod download```

3. 运行项目

```go run main.go```