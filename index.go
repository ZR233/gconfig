package gconfig

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path"
	"time"
)

const (
	key_Postgresql = "common/postgresql"
	key_Redis      = "common/redis"
	key_Zookeeper  = "common/zookeeper"
	key_Hbase      = "common/hbase"
)

type DB interface {
	Unmarshal(keyPath string, o interface{}) (err error)
	Get(keyPath string) (data []byte, version uint64, err error)
	Watch(keyPath string, ValuePtr interface{}, onChanged func(err error)) (err error)
}

type Config struct {
	db DB
}

func NewConfig(db DB) *Config {
	return &Config{db: db}
}

func (c *Config) Watch(onChanged func(err error)) (err error) {
	err = c.db.Watch(key_Postgresql, &SqlCluster{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(key_Redis, &Redis{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(key_Zookeeper, &Zookeeper{}, onChanged)
	if err != nil {
		return
	}
	err = c.db.Watch(key_Hbase, &Hbase{}, onChanged)
	return
}

func (c *Config) GetPostgreSQL() (cfg *SqlCluster, err error) {
	cfg = &SqlCluster{}
	err = c.db.Unmarshal(key_Postgresql, cfg)
	return
}
func (c *Config) GetRedis() (cfg *Redis, err error) {
	cfg = &Redis{}
	err = c.db.Unmarshal(key_Redis, cfg)
	return
}
func (c *Config) GetZookeeper() (cfg *Zookeeper, err error) {
	cfg = &Zookeeper{}
	err = c.db.Unmarshal(key_Zookeeper, cfg)
	return
}
func (c *Config) GetHbase() (cfg *Hbase, err error) {
	cfg = &Hbase{}
	err = c.db.Unmarshal(key_Hbase, cfg)
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
