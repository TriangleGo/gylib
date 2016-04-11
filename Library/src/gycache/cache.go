package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"gylogger"
	"github.com/stvp/go-toml-config"
)

var RedisClient *redis.Pool

var (
	url string
	pwd string
)

func InitCache(path string) (err error) {
	loadConfig(path)

	//url := "uat-sz.lansion.cn:7500"
	option := redis.DialPassword(pwd)
	var conn redis.Conn
	RedisClient = &redis.Pool{
		MaxIdle:10,
		MaxActive:30,
		IdleTimeout:180 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err = redis.Dial("tcp", url, option)
			if err != nil {
				logger.Error("connect to redis error", err)
				return nil, err
			}
			return conn, nil
		},
	}

	if err == nil {
		logger.Debugf("Cache redis server connected %s.", url)
	} else {
		logger.Errorf("Cache redis servier connect err %v", err)
	}
	return
}


func loadConfig(path string) {
	cacheConfig := config.NewConfigSet("cacheConfig", config.ExitOnError)
	cacheConfig.StringVar(&url, "url", ":7500")
	cacheConfig.StringVar(&pwd, "password", "abcdefg_m")
	err := cacheConfig.Parse(path)
	if err != nil {
		logger.Warnf("load cache config error, %v", err)
	} else {
		logger.Info("loaded cache config")
	}
}

func Exist(key string) (bool, error) {
	conn := Conn()
	defer conn.Close()
	val, err := redis.Int(conn.Do("EXISTS", key))
	return val == 1, err
}

func Set(key string, value interface{}) (err error) {
	conn := Conn()
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	return
}

func Get(key string) (interface{}, error) {
	conn := Conn()
	defer conn.Close()
	return conn.Do("GET", key)
}

func GetBytes(key string) (val []byte, err error) {
	conn := Conn()
	defer conn.Close()
	val, err = redis.Bytes(Get(key))
	return
}

func GetString(key string) (string, error) {
	conn := Conn()
	defer conn.Close()
	val, err := redis.String(Get(key))
	return val, err
}

func GetInt(key string) (val int, err error) {
	val, err = redis.Int(Get(key))
	return
}

func Del(key string) (interface{}, error) {
	conn := Conn()
	defer conn.Close()
	return conn.Do("DEL", key)
}

func Expire(key string, minutes int) (interface{}, error) {
	conn := Conn()
	defer conn.Close()
	return conn.Do("EXPIRE", key, minutes * 60)
}

func Persist(key string) (interface{}, error) {
	conn := Conn()
	defer conn.Close()
	return conn.Do("PERSIST", key)
}

func HExist(key string, field interface{}) (bool, error) {
	conn := Conn()
	defer conn.Close()
	val, err := redis.Int(conn.Do("HEXISTS", key, field))
	return val == 1, err
}

func HSet(key string, field interface{}, value interface{}) (err error) {
	conn := Conn()
	defer conn.Close()
	_, err = conn.Do("HSET", key, field, value)
	return
}

func HGet(key string, field interface{}) (val  interface{}, err error) {
	conn := Conn()
	defer conn.Close()
	val, err = redis.String(conn.Do("HGET", key, field))
	return
}

func HDel(key string, field interface{}) (interface{}, error) {
	conn := Conn()
	defer conn.Close()
	return conn.Do("HDEL", key, field)
}

func Conn() redis.Conn {
	return RedisClient.Get()
}