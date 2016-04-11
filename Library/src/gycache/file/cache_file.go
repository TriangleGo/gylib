package file

import (
	"gylogger"
	"gyuuid"
	"gycache"
	"github.com/garyburd/redigo/redis"
	"path"
	"fmt"
)

const CACHE_FIELD_NAME = "name"
const CACHE_FIELD_CONTENT_TYPE = "content_type"
const CACHE_FIELD_CONTENT = "content"

func RefreshExpire(key string, minute int) (exists bool) {
	if minute > 0 {
		cache.Expire(key, minute)
	} else if minute == 0 {
		cache.Del(key)
	} else {
		cache.Persist(key)
	}
	exists, err := cache.Exist(key)
	exists = exists && (err == nil)
	return
}

func NewCacheFile(name string, contentType string, content []byte, minute int) (key string) {
	key = uuid.Rand().Hex()

	logger.Debugf("CacheFile for key %s, name %s.", key, name)
	suffix := path.Ext(name)
	cache.HSet(key, CACHE_FIELD_NAME, fmt.Sprintf("%s%s", key, suffix))
	cache.HSet(key, CACHE_FIELD_CONTENT_TYPE, contentType)
	cache.HSet(key, CACHE_FIELD_CONTENT, content)
	if minute > 0 {
		cache.Expire(key, minute)
	}
	return
}

func NewCacheFileWithKey(key string, name string, contentType string, content []byte, minute int) {
	logger.Debugf("CacheFile for key %s, name %s.", key, name)
	suffix := path.Ext(name)
	cache.HSet(key, CACHE_FIELD_NAME, fmt.Sprintf("%s%s", key, suffix))
	cache.HSet(key, CACHE_FIELD_CONTENT_TYPE, contentType)
	cache.HSet(key, CACHE_FIELD_CONTENT, content)
	if minute > 0 {
		cache.Expire(key, minute)
	}
}

func GetCacheFile(key string, delete bool) (name string, contentType string, content []byte, exists bool) {

	name, _ = redis.String(cache.HGet(key, CACHE_FIELD_NAME))
	contentType, _ = redis.String(cache.HGet(key, CACHE_FIELD_CONTENT_TYPE))
	content, _ = redis.Bytes(cache.HGet(key, CACHE_FIELD_CONTENT))
	exists = len(name) > 0 && len(content) > 0

	logger.Debugf("GetCacheFile, key %s, name:type:len(content) %s:%v:%d\n", key, name, contentType, len(content))

	if exists && delete {
		delCacheFile(key)
	}

	return
}

func delCacheFile(key string) (err error) {

	val, err := cache.Del(key)
	logger.Debugf("del cache file %v, result, %v, err %v", key, val, err)
	return
}