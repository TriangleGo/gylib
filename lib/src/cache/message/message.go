package message

import (
	"cache"
	"encoding/json"
	"reflect"
	"uuid"
	"logger"
	"util"
)

func CheckExists(msgKey string) (bool, error) {
	return cache.Exist(msgKey)
}

func CacheMsg(msg interface{}) (key string, err error) {
	msgBytes, err := json.Marshal(msg)
	key = uuid.Rand().Hex()
	err = cache.Set(key, msgBytes)
	if err == nil {
		cache.Expire(key, 1)
	}
	return
}

func GetMsg(key string, obj interface{}) (err error) {
	logger.Debugf("msg key = %s, type = %s", key, reflect.TypeOf(obj))
	msgBytes, err := cache.GetBytes(key)
	logger.Debugf("got cached message %s with error %v", string(msgBytes), err)
	if err != nil {
		return
	}
	cache.Del(key)
	err = json.Unmarshal(msgBytes, obj)
	logger.Debugf("unmarshal message byte to obj %s with err %v", test.ToJsonString(obj), err)
	return
}