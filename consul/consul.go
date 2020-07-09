package consul

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
func (d *DBConsul) Get(keyPath string) (data []byte, version uint64, err error) {
	p, m, err := d.client.KV().Get(keyPath, nil)
	if err != nil {
		return
	}
	version = m.LastIndex
	data = p.Value
	return
}
func (d *DBConsul) Watch(keyPath string, o interface{}, onChanged func(err error)) (err error) {
	p, m, err := d.client.KV().Get(keyPath, nil)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(p.Value, o)
	if err != nil {
		return
	}
	lastVersion := m.LastIndex
	go func() {
		for {
			p, m, err = d.client.KV().Get(keyPath, nil)
			if err != nil {
				onChanged(err)
				return
			}
			err = yaml.Unmarshal(p.Value, o)
			if err != nil {
				onChanged(err)
				return
			}
			if lastVersion != m.LastIndex {
				lastVersion = m.LastIndex
				onChanged(nil)
			}
		}
	}()

	return
}

func NewDBConsul(cfg *api.Config) (DB *DBConsul, err error) {
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
