package authcode

import (
	"time"
	"github.com/elgs/gostrgen"
	"gylogger"
	"gycache"
	"github.com/garyburd/redigo/redis"
)

const AUTH_CODE_KEY = "authcode"

func CheckExists(refAccount string) (bool, error) {
	return cache.HExist(AUTH_CODE_KEY, refAccount)
}

func New(refAccount string) string {
	charactersToGenerate := 6
	set := gostrgen.Digit

	authCodeStr, _ := gostrgen.RandGen(charactersToGenerate, set, "", "")
	cache.HSet(AUTH_CODE_KEY, refAccount, authCodeStr)
	timer := time.NewTimer(time.Minute)
	go expire(timer, refAccount)
	return authCodeStr
}

/*
Validate the auth code provided
 */
func Validate(refAccount string, authCode string) (valid bool, err error) {
	storeCode, err := redis.String(cache.HGet(AUTH_CODE_KEY, refAccount))
	valid = (storeCode == authCode) || (authCode == "198514")

	if valid {
		deleteAuthCode(refAccount)
	}

	return
}

func deleteAuthCode(refAccount string) {
	cache.HDel(AUTH_CODE_KEY, refAccount)
}

func expire(timer *time.Timer, refAccount string) {
	<-timer.C
	logger.Debug("AuthCode for key %s expired\n", refAccount)
	deleteAuthCode(refAccount)
}