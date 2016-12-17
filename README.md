# Design

![](https://raw.githubusercontent.com/beautytop/dudu/master/doc/img/1.jpg)


![](https://raw.githubusercontent.com/beautytop/dudu/master/doc/img/2.jpg)


# USE

```
go get -u -v github.com/beautytop/dudu
go run initsql.go
source /data/www/web/go/src/github.com/beautytop/dudu/doc/sql/uk_category.sql
```

# Result

![](https://raw.githubusercontent.com/beautytop/dudu/master/doc/img/3.png)

# ERROR
1.Redis

q1:

```
MISCONF Redis is configured to save RDB snapshots, but is currently not able to persist on disk. 
Commands that may modify the data set are disabled. Please check Redis logs for details about the error.
panic: REDIS ERROR

```

solve it

```
redis command
config set stop-writes-on-bgsave-error no
```

q2:

```
redis: connection pool timeout
```

solve it

```
config/config.json set 

"Redispoolsize":2000
```

2.Mysql

q1:

```
mysql error:Error 1016: Can't open file: './uk_smart_base/A296.frm' (errno: 24 - Too many open files)
```

solve it

```
mysql command
SHOW VARIABLES LIKE 'open_files_limit';
SHOW GLOBAL STATUS LIKE 'Open_files';

edit my.cnf

[mysqld]
open_files_limit = 65535

vim /etc/systemd/system/mysql.service
add  LimitNOFILE=65535

systemctl daemon-reload
service mysqld restart

vim /etc/security/limits.conf,add
 
*  soft    nofile  65536
*  hard    nofile  65536

ulimit -n
```

q2:

```
mysql error:Error 1040: Too many connections
```

solve it

```
SHOW VARIABLES like 'max_%';

set GLOBAL max_connections=5000
```


q3

```
DNS skip
[mysqld]
skip-name-resolve
```


# 中文介绍

‌Golang全自动亚马逊全网分布式爬虫（美国，日本，德国和英国）

2016.12.10
亚马逊爬虫支持

1. 列表页和详情页可选择代理方式
2. 多浏览器保存cookie机制
3. 机器人检测达到阈值自动换代理
4. 检测日期过期自动停止程序
5. IP池扫描周期填充代理IP
6. 支持分布式跨平台抓取
7. 高并发进程设置抓取
8. 默认网页爬取去重
9. 日志记录功能
10. 配套可视化网站，支持多角度查看数据，小类数据，大类数据，Asin数据和类目数据，支持查看每件Asin商品的历史记录，如排名，价格，打分，reviews变化。部分数据支持导出，且网站支持RBAC权限，可分配每部分数据的查看和使用权限。
11. 网络端监控爬虫，可查看爬虫当前时段数据抓取状态，爬取的进度，IP的消耗程度。   **可支持网络端启动和停止爬虫，彻底成为Saas**（待做）
12. 可自定义填入IP，如塞入其他代理IP网站API获取的IP
13. 可选择HTML文件保存本地

分布式，高并发，跨平台，多站点，多种自定义配置，极强的容错能力是这个爬虫的特点。机器数量和IP代理足够情况下，每天每个站点可满足抓取几百万的商品数据。

2016.12.15

1. 数据库初始化脚本优化
2. IP配置文件优化
3. 网络WEB端自定义IP

# 中文使用

安装Go环境，MYSQL和Redis

一.Go安装
Go包安装 https://yun.baidu.com/s/1jHKUGZG 选择1.6后缀msi安装

环境变量设置

```
Path G:\smartdogo\bin
GOBIN G:\smartdogo\bin
GOPATH G:\smartdogo
GOROOT C:\Go\
```

然后

```
go get -u -v github.com/beautytop/dudu
```

二. Mysql安装

https://yun.baidu.com/s/1hrF0QC8 找到mysql文件夹，下面的5.6.17.0.msi根据说明安装。

三.Redis安装

https://yun.baidu.com/s/1jHKUGZG 选择redis64bit或32bit，解压 ，然后Shift+右键 在此次打开命令窗口 输入“redis-server.exe redis.conf ”，在redis.conf设置密码smart2016,已经设置好

四.修改配置文件

可以修改GOPATH，即后缀为?_config.json的配置（默认不需要改，下面注释说明需要删除）

```
{
  "Type": "USA",     //美国站类型，有四种usa,jp,uk,de
  "Datadir": "/data/db/usa",   // 文件保存位置，可选择保存，/代表在本盘下
  "Proxymaxtrytimes": 88,     // 机器人错误最大次数，超过换IP
  "Rank": 500000,               // 只保存排名在这个数字之前的商品
  "Listtasknum": 30,        // 抓列表页进程数，建议不要太大，并发最多设置50
  "Asintasknum": 30,      // 抓详情页进程数，建议不要太大，并发最多设置50
  "Localtasknum": 150,  // 本地文件处理进程数，建议不要太大，并发最多设置50，可不管
  "Proxypool": "USAIPPOOL",   // Redis IP代理池名字
  "Proxyhashpool": "USAIPPOLLHASH",  // Redis IP已用池名字
  "Proxyloophours": 4,        // 重导IP时间（小时,Redis IP池用完）
  "Proxyasin": true,         // 详情页使用代理？
  "Proxycategory": false,    //列表页使用代理？
  "Proxyinit": false,   // IP池程序每次启动是否追加，可不管
  "Urlpool": "USAURLPOOL",  //列表页待抓池名字
  "Urldealpool": "USAURLDEALPOOL", //列表页处理中池
  "Urlhashpool": "USAURLHASHPOOL",  //列表页已抓池
  "Asinpool": "USAAsinPOOL",       // 同理
  "Asindealpool": "USAAsinDEALPOOL",
  "Asinhashpool": "USAAsinHASHPOOL",
  "Otherhashpool": "USAOtherHashPOOL",  // 小类数据额外redis池，方便填充大类数据，开关在ExtraFromRedis,如果关，大类数据填充查找小类数据库，大数据下会导致慢
  "Asinautopool": true,   //列表页抓取数据后自动把Asin数据塞在Asinpool,如果设置为false，需要手动运行asinpool.exe
  "ExtraFromRedis": true,  //搭配Otherhashpool
  "Asinlocalkeep": false,   //保存详情页在Datadir
  "Categorylocalkeep": false, //保存列表页在Datadir
  "Urlsql": "SELECT distinct url,id,bigpid ,name,bigpname,page FROM smart_category where isvalid=1 order by bigpid limit 100000",  //抓取那些列表页，可改
  "Asinsql": "SELECT distinct asin as id FROM `{?}` order by bigname limit 1000000", //抓取哪些Asin，{?}是程序预带占位符，被今天日期替代，可去掉
  "Spidersleeptime": 3, // 无用
  "Spidertimeout": 20,  //链接抓取超时时间
  "Spiderloglevel": "INFO",  //爬虫日志记录，可不管
  "Redisconfig": {  // redis配置
    "Host": "45.41.88.189:6379",  //主机
    "Password": "smart2016",   //密码
    "DB": 0
  },
  "Redispoolsize": 100,  // redis程序库连接池最大数量，应该比Listtasknum和Asintasknum都大
  "Basicdb": {   // 基础数据库配置
    "Ip": "45.41.88.187",
    "Port": "3306",
    "Username": "root",
    "Password": "smart2016",
    "Dbname": "smart_base"
  },
  "Hashdb": {   //历史数据库配置
    "Ip": "45.41.88.187",
    "Port": "3306",
    "Username": "root",
    "Password": "smart2016",
    "Dbname": "smart_hash"
  },
  "Hashnum": 500,   //历史数据库按hashcode分表，分表数量
  "Datadb": {     // 日期数据库，按天分表
    "Ip": "45.41.88.187",
    "Port": "3306",
    "Username": "root",
    "Password": "smart2016",
    "Dbname": "smartdb"
  },
  "Ipuse": {   //要用的IP组
    "d": {    //端口和密码，密码可留空，组名所在的IP在下面
      "Port": "808",
      "Secret": "smart:smart2016"
    },
    "e": {
      "Port": "808",
      "Secret": "smart:smart2016"
    },
    "f": {
      "Port": "808",
      "Secret": "smart:smart2016"
    },
    "h": {
      "Port": "808",
      "Secret": "smart:smart2016"
    }
  },
  "Ips": {
    "d": [   //组名为d的IP们
      "146.148.149.203-254",   // 连续Ip,也可以不连续，如146.148.149.203
      "146.148.150.194-254",
      "146.148.157.194-254"
    ],
    "e": [
      "104.143.159.233-254",
      "107.190.231.232-254",
      "104.128.118.170-190",
      "104.128.119.169-190",
      "104.128.121.41-62",
      "104.128.122.41-62",
      "104.128.123.41-62",
      "104.128.124.105-126"
    ],
    "f": [
      "146.148.178.234-254",
      "146.148.179.168-190",
      "104.149.46.234-254",
      "104.149.64.105-126",
      "104.149.66.41-62",
      "104.149.71.41-62",
      "146.148.133.41-62",
      "146.148.176.169-190"
    ],
    "h": [
      "146.148.217.105-126",
      "146.148.220.232-254",
      "146.148.238.232-254",
      "146.148.240.232-254",
      "146.148.246.41-62",
      "146.148.247.232-254",
      "146.148.215.169-190",
      "146.148.216.41-62"
    ],
    "i": [
      "104.216.7.235-254",
      "104.216.18.233-254",
      "104.216.76.169-190",
      "104.216.77.41-62",
      "104.216.84.233-254",
      "104.216.91.105-126",
      "104.216.137.40-62",
      "104.216.157.168-190"
    ],
    "j": [
      "104.216.135.233-254",
      "104.216.161.168-190",
      "104.216.167.104-126",
      "104.216.178.104-126",
      "104.128.112.168-190",
      "107.178.161.232-254",
      "104.143.128.232-254",
      "104.143.129.232-254"
    ]
  }
}
```

运行程序

在spiders文件夹下,进去各个站点，运行go build *.go会得到exe文件

1. 点击initsql.exe初始化数据库
2. SQL文件导入数据库，列表URL
3. 点击ippool.exe填充IP
4. 点击urlpool.exe填充URL
5. 点击listmain.exe抓取列表页
6. 点击asinmain.exe抓取详情页

7. 如果配置中Asinautopool设置为false，那么需要自己导Asin,运行asinpool.exe