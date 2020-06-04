package gconfig

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
