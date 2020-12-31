package test

import (
	"fmt"
	"github.com/ZR233/gconfig/v2"
	"github.com/ZR233/gconfig/v2/consul"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {

	cfg := gconfig.NewConfig("test").UseConsul(nil)
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
func TestNewConsul_Default(t *testing.T) {
	cfg := gconfig.NewConfig("test").UseConsul(nil)

	var Test struct {
		A string
		B string
	}
	err := cfg.Unmarshal(&Test)
	println(err)
	time.Sleep(time.Minute * 10)
}
