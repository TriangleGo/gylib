package message

import (
	"util/uuid"
	"cache"
)

func CacheMsg(msg interface{}) (key string, err error) {
	key = uuid.Rand().Hex()
	err = cache.SetI(key, msg)
	return
}

func GetMsg(key string, msg interface{}) (err error) {
	err = cache.GetI(key, msg)
	return
}

func GetMsgString(key string, msg *string) (err error) {
	err = cache.Get(key, msg)
	return
}

func DelMsg(key string) {
	cache.Del(key)
}
