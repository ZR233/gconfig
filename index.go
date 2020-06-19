package gconfig

import (
	"encoding/json"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path"
	"time"
)

type DB interface {
	Unmarshal(keyPath string, o interface{}) (err error)
}

type Config struct {
	db DB
}

func NewConfig(db DB) *Config {
	return &Config{db: db}
}

func (c *Config) GetPostgreSQL() (cfg *SqlCluster, err error) {
	cfg = &SqlCluster{}
	err = c.db.Unmarshal("common/postgresql", cfg)
	return
}
func (c *Config) GetRedis() (cfg *Redis, err error) {
	cfg = &Redis{}
	err = c.db.Unmarshal("common/redis", cfg)
	return
}
func (c *Config) GetZookeeper() (cfg *Zookeeper, err error) {
	cfg = &Zookeeper{}
	err = c.db.Unmarshal("common/zookeeper", cfg)
	return
}
func (c *Config) GetHbase() (cfg *Hbase, err error) {
	cfg = &Hbase{}
	err = c.db.Unmarshal("common/hbase", cfg)
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
