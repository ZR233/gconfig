package test

import (
	"fmt"
	"github.com/ZR233/gconfig"
	"github.com/ZR233/gconfig/consul"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	db, err := consul.NewDBConsul(nil)
	if err != nil {
		t.Error(err)
		return
	}

	cfg := gconfig.NewConfig(db)
	p, err := cfg.GetPostgreSQL()
	println(p, err)
	r, err := cfg.GetRedis()
	println(r, err)
	z, err := cfg.GetZookeeper()
	println(z, err)
	h, err := cfg.GetHbase()
	println(h, err)
}
func TestNewConsul_Watch(t *testing.T) {
	db, err := consul.NewDBConsul(nil)
	if err != nil {
		t.Error(err)
		return
	}
	cfg := gconfig.Zookeeper{}
	db.Watch("common/zookeeper", &cfg, func(err error) {
		t.Log(fmt.Sprintf("%v", cfg))
	})

	time.Sleep(time.Minute * 10)
}
