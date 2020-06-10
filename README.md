# gconfig 配置中心
## 支持存储
- consul
##使用方法
```
import (
	"github.com/ZR233/gconfig"
	"github.com/ZR233/gconfig/consul"
) 
db, err := consul.NewDBConsul(nil)
if err != nil {
    panic(err)
}

cfg := gconfig.NewConfig(db)
PostgreSql, err := cfg.GetPostgreSQL()
if err != nil {
    panic(err)
}
Redis, err := cfg.GetRedis()
if err != nil {
    panic(err)
}
Hbase, err := cfg.GetHbase()
if err != nil {
    panic(err)
}
Zookeeper, err := cfg.GetZookeeper()
if err != nil {
    panic(err)
}
//自定义配置
var MyConfig struct{
    Name string
    Detail string
}
configPath := "test/myconfig"
err = db.Unmarshal(configPath, &MyConfig)
if err != nil {
    panic(err)
}

```