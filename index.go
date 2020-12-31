package gconfig

import (
	"encoding/json"
	"fmt"
	"github.com/ZR233/gconfig/v2/cerr"
	"github.com/ZR233/gconfig/v2/consul"
	"github.com/go-zookeeper/zk"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
	"path"
	"time"
)

const (
	Key_Postgresql = "common/postgresql"
	Key_Redis      = "common/redis"
	Key_Zookeeper  = "common/zookeeper"
	Key_Hbase      = "common/hbase"
)

type DB interface {
	Unmarshal(keyPath string, o interface{}) (err error)
	Get(keyPath string) (data []byte, version uint64, err error)
	Watch(keyPath string, ValuePtr interface{}, onChanged func(err error)) (err error)
	Set(keyPath string, in interface{}) (err error)
}

type Config struct {
	db       DB
	savePath string
}

func (c *Config) UseConsul(config *api.Config) *Config {
	var err error
	c.db, err = consul.NewDBConsul(config)
	if err != nil {
		panic(err)
	}
	return c
}

func NewConfig(savePath string) *Config {
	return &Config{
		savePath: savePath,
	}
}

func (c *Config) Watch(onChanged func(err error)) (err error) {
	err = c.db.Watch(Key_Postgresql, &SqlCluster{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(Key_Redis, &Redis{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(Key_Zookeeper, &Zookeeper{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(Key_Hbase, &Hbase{}, onChanged)
	return
}

func (c *Config) GetPostgreSQL() (cfg *SqlCluster, err error) {
	cfg = &SqlCluster{}
	err = c.db.Unmarshal(Key_Postgresql, cfg)
	return
}
func (c *Config) GetRedis() (cfg *Redis, err error) {
	cfg = &Redis{}
	err = c.db.Unmarshal(Key_Redis, cfg)
	return
}
func (c *Config) GetZookeeper() (cfg *Zookeeper, err error) {
	cfg = &Zookeeper{}
	err = c.db.Unmarshal(Key_Zookeeper, cfg)
	return
}
func (c *Config) GetHbase() (cfg *Hbase, err error) {
	cfg = &Hbase{}
	err = c.db.Unmarshal(Key_Hbase, cfg)
	return
}

func (c *Config) GetKafkaAddrs(zkCfg *Zookeeper) (addrs []string, err error) {
	conn, _, err := zk.Connect(zkCfg.Hosts, time.Second*5)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ids, _, err := conn.Children("/brokers/ids")
	if err != nil {
		panic(err)
	}
	var data []byte
	type KafkaNode struct {
		Host string
		Port int
	}

	for _, id := range ids {
		data, _, err = conn.Get(path.Join("/brokers/ids", id))
		if err != nil {
			panic(err)
		}
		node := &KafkaNode{}
		err = json.Unmarshal(data, node)
		if err != nil {
			panic(err)
		}
		addrs = append(addrs, fmt.Sprintf("%s:%d", node.Host, node.Port))
	}

	return
}
func (c *Config) Unmarshal(value interface{}) (err error) {
	data, _, err := c.db.Get(c.savePath)
	if len(data) > 0 {
		err = yaml.Unmarshal(data, value)
		if err != nil {
			return
		}
	} else {
		err = c.db.Set(c.savePath, value)
		if err != nil {
			return
		}
		err = cerr.ErrConfigNotInit
	}
	return
}
