# 短链接项目

## 项目骨架

### 建库建表

新建发号器表

```sql
CREATE TABLE `sequence` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stub` varchar(1) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT = '序号表';
```

新建长链接短链接映射表：

```sql
CREATE TABLE `short_url_map` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',
    
    `lurl` varchar(2048) DEFAULT NULL COMMENT '长链接',
    `md5` char(32) DEFAULT NULL COMMENT '长链接MD5',
    `surl` varchar(11) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    INDEX(`is_del`),
    UNIQUE(`md5`),
    UNIQUE(`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '长短链映射表';
```

### Go-zero

#### 编写api文件

```go
/* 短链接项目
* author: q1mi
*/

type ConvertRequest {
    LongUrl string `json:"longUrl"`
}

type ConvertResponse {
    ShortUrl string `json:"shortUrl"`
}

type ShowRequest {
    ShortUrl string `json:"shortUrl"`
}

type ShowResponse {
    LongUrl string `json:"longUrl"`
}

service shortener-api {

    @handler ConvertHandler
    post /convert(ConvertRequest) returns(ConvertResponse)

    // q1mi.cn/lycsa1
    @handler ShowHandler
    get /:shortUrl(ShowRequest) returns(ShowResponse)

}
```

根据api文件生成go代码:

```bash
goctl api go -api shortener.api -dir .
```

#### 根据sql生成model相关代码

```bash
goctl model mysql datasource -url="root:root1234@tcp(127.0.0.1:3306)/db3" -table="short_url_map"  -dir="./model"

goctl model mysql datasource -url="root:root1234@tcp(127.0.0.1:3306)/db3" -table="sequence"  -dir="./model"
```

### Docker-compose 启动服务

```dockerfile
version: "3.7"
services:
  mysqllatest:
    image: "mysql:latest"
    ports:
      - "9306:3306"
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    environment:
      MYSQL_ROOT_PASSWORD: "root1234"
      MYSQL_DATABASE: "shortener"
      MYSQL_PASSWORD: "root1234"
    volumes:
      - ./model/sql/init.sql:/data/application/init.sql
  
  redis507:
    image: "redis:5.0.7"
    ports:
      - "9379:6379"
```

### 修改config结构体和配置文件

```yaml
Name: shortener-api
Host: 0.0.0.0
Port: 8888

ShortUrlDB:
    # 与docker-compose 对应
  DSN: root:root1234@tcp(127.0.0.1:9306)/shortener?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai


Sequence:
  DSN: root:root1234@tcp(127.0.0.1:9306)/shortener?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai


CacheRedis:
  - Host: 127.0.0.1:9379


BaseString: J0rs12O5TUV8IW7D9aBdXeCfghiMQj3klmop6qtuvbcwx4zAEFGHKLNnPRYSZy


ShortUrlBlackList: ["version", "convert", "show", "short", "url", "health", "api", "css", "fuck", "shit"]


ShortDomain: shortener.cn
```

## 参数校验

1. go-zero使用validator
<https://pkg.go.dev/github.com/go-playground/validator/v10>

导入依赖：

```bash
import "github.com/go-playground/validator/v10"
```

在api中为结构体添加validate tag，并添加校验规则

## 接口内容

### Convert

将long url转换成short url

* 参数校验，检查输入url是否为合法的long url
* 检查 long url 是否已被转换
* 转换并生成 short url
* 添加到 Bloom Filter
* 添加到 Redis 和 Mysql

### Show

查询short url是否有对应的long url

* Bloom Filter查询
* 从Redis缓存中查询
* 从Mysql数据库中查询
