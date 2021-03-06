package cache

import (
	"reflect"
	"testing"
)

var (
	key   = "key12345"
	value = []byte("value")
)

func TestRead_file(t *testing.T) {
	ca := NewFileCache("", "")
	setup(ca, 1)
	data, err := ca.Read(key)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(data, value) {
		t.Errorf("wanted(%q)\n got(%q).", value, data)
	}
	ca.clear()
}

func TestRead_file_expired(t *testing.T) {
	ca := NewFileCache("", "")
	setup(ca, -1)
	data, err := ca.Read(key)
	if err != nil {
		t.Error(err)
	}
	if data != nil {
		t.Error("Cache record did not expire correctly.")
	}
	ca.clear()
}

func setup(ca Cache, duration int64) {
	_ = ca.Store(key, value, duration)	
}
