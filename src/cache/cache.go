package cache

import (
	"github.com/BurntSushi/toml"
	"github.com/mediocregopher/radix.v3"
	"util/json"
	"errors"
)

var pool *radix.Pool

var config struct {
	Url      string
	PoolSize int
}

func init() {
	loadConfig()
	var err error
	pool, err = radix.NewPool("tcp", config.Url, config.PoolSize)
	if err != nil {
		panic(err)
	}
	return
}

func loadConfig() {
	_, err := toml.DecodeFile("./conf/redis.conf", &config)
	if err != nil {
		panic(err)
	}
}

func Exist(key string) (exist bool, err error) {
	cmd := radix.Cmd(&exist, "EXISTS", key)
	err = pool.Do(cmd)
	return
}

func Set(key string, value interface{}) (err error) {
	cmd := radix.FlatCmd(nil, "SET", key, value)
	err = pool.Do(cmd)
	return
}

func Get(key string, value interface{}) (err error) {
	cmd := radix.Cmd(value, "GET", key)
	err = pool.Do(cmd)
	return
}

func SetI(key string, value interface{}) (err error) {
	return Set(key, json.ToString(value))
}

func GetI(key string, value interface{}) (err error) {
	if value == nil {
		err = errors.New("empty input")
		return
	}
	var str string
	err = Get(key, &str)
	if err != nil {
		return
	}
	err = json.FromString(str, value)
	return
}

func Del(key string) (err error) {
	cmd := radix.Cmd(nil, "DEL", key)
	err = pool.Do(cmd)
	return
}

//
//func GetBytes(key string) (val []byte, err error) {
//	conn := Conn()
//	defer conn.Close()
//	val, err = redis.Bytes(GetStr(key))
//	return
//}
//
//func GetString(key string) (string, error) {
//	conn := Conn()
//	defer conn.Close()
//	val, err := redis.String(GetStr(key))
//	return val, err
//}
//
//func GetInt(key string) (val int, err error) {
//	val, err = redis.Int(GetStr(key))
//	return
//}
//
//func Expire(key string, minutes int) (interface{}, error) {
//	conn := Conn()
//	defer conn.Close()
//	return conn.Do("EXPIRE", key, minutes*60)
//}
//
//func Persist(key string) (interface{}, error) {
//	conn := Conn()
//	defer conn.Close()
//	return conn.Do("PERSIST", key)
//}
//
//func HExist(key string, field interface{}) (bool, error) {
//	conn := Conn()
//	defer conn.Close()
//	val, err := redis.Int(conn.Do("HEXISTS", key, field))
//	return val == 1, err
//}
//
//func HSet(key string, field interface{}, value interface{}) (err error) {
//	conn := Conn()
//	defer conn.Close()
//	_, err = conn.Do("HSET", key, field, value)
//	return
//}
//
//func HGet(key string, field interface{}) (val interface{}, err error) {
//	conn := Conn()
//	defer conn.Close()
//	val, err = redis.String(conn.Do("HGET", key, field))
//	return
//}
//
//func HDel(key string, field interface{}) (interface{}, error) {
//	conn := Conn()
//	defer conn.Close()
//	return conn.Do("HDEL", key, field)
//}
