package cache

import (
	"testing"
	"util/json"
)

type Obj struct {
	Name  string
	Value int
}

func TestExist(t *testing.T) {
	exist, err := Exist("key")
	t.Logf("exist = %v, err %v.", exist, err)
}

func TestSet(t *testing.T) {
	obj := &Obj{
		Name:  "Adam",
		Value: 101,
	}
	err := Set("key", json.ToString(obj))
	t.Logf("set error %v.", err)
}

func TestGet(t *testing.T) {
	str := ""
	err := Get("key", &str)
	t.Logf("value = %v, err %v.", str, err)
}

func TestSetI(t *testing.T) {
	obj := &Obj{
		Name:  "Adam",
		Value: 100,
	}
	err := SetI("key2", obj)
	t.Logf("seti error %v.", err)
}

func TestGetI(t *testing.T) {
	obj := &Obj{}
	err := GetI("key", obj)
	t.Logf("value = %v, err %v.", json.ToString(obj), err)
}
