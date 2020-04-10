## go-pkg

###Usage
```
go get -v -u github.com/Jurassic-Park/go-pkg
```
should have conf/app.ini in root dir

app.ini for example:
```cassandraql
[app]
PageSize = 10
JwtSecret = 233

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
#debug or release
RunMode = debug
Port = 8000
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = haian
Password = RONGshubao+123
Host = rm-m5efb5r834532e6ztco.mysql.rds.aliyuncs.com:3306
Name = insure
TablePrefix =

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
```

#### init

```
import pkg "github.com/Jurassic-Park/go-pkg"
pkg.Setup()
```

#### use gorm
```
import gorm "github.com/Jurassic-Park/go-pkg/gorm"
db := gorm.NewGorm()
```

#### use redis
```
import gredis "github.com/Jurassic-Park/go-pkg/gredis"
gredis.Set()
```

#### use logrus
```
import log "github.com/Jurassic-Park/go-pkg/logrus"
log.logging.WithFields(log.Fields{
    "animal": "walrus",
  }).Info("A walrus appears")
```

#### use grpc
```
import ggrpc "github.com/Jurassic-Park/go-pkg/grpc"

// grpc server or http server
ggrpc.GrpcHandlerFunc()

// grpc TLS 
ggrpc.GetTLSConfig()
```



