package gconfig

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	db, err := NewDBConsul(nil)
	if err != nil {
		t.Error(err)
		return
	}

	cfg := NewConfig(db)
	p, err := cfg.GetPostgreSQL()
	println(p, err)
	r, err := cfg.GetRedis()
	println(r, err)
	z, err := cfg.GetZookeeper()
	println(z, err)
	h, err := cfg.GetHbase()
	println(h, err)
}
