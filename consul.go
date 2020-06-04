package gconfig

import (
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

type DBConsul struct {
	client *api.Client
}

func (d *DBConsul) Unmarshal(keyPath string, o interface{}) (err error) {
	p, _, err := d.client.KV().Get(keyPath, nil)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(p.Value, o)
	return
}

func NewDBConsul(cfg *api.Config) (DB DB, err error) {
	if cfg == nil {
		cfg = api.DefaultConfig()
		cfg.Address = "localhost:8500"
	}
	client, err := api.NewClient(cfg)
	if err != nil {
		return
	}
	DB = &DBConsul{client: client}
	return
}
