package consul

import (
	"testing"
)

func TestDBConsul_Get(t *testing.T) {
	db, err := NewDBConsul(nil)
	if err != nil {
		panic(err)
	}
	s, h, err := db.Get("common/zookeeper")
	if err != nil {
		panic(err)
	}

	println(s, h)
}
