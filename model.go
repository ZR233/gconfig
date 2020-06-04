package gconfig

type SqlCluster struct {
	Write Sql   `yaml:"write"`
	Read  []Sql `yaml:"read"`
}
type Sql struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
}

type Redis struct {
	Password   string   `yaml:"password"`
	Addrs      []string `yaml:"addrs"`
	Mastername string   `yaml:"mastername"`
}

type Zookeeper struct {
	Hosts []string `yaml:"hosts"`
}
type Hbase struct {
	Thrift  string `yaml:"thrift"`
	Thrift2 string `yaml:"thrift2"`
}
