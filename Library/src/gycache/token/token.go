package token

import (
	"gyuuid"
	"gylogger"
	"gycache"
	"github.com/garyburd/redigo/redis"
)

const TOKEN_KEY = "token"
const CACHE_ACCOUNT = "refaccount"
const CACHE_ID = "refId"

func CheckExists(refAccount string) (bool, error) {
	return cache.Exist(refAccount)
}

func New(refAccount string, accountId string) string {
	tokenStr := uuid.Rand().Hex()
	logger.Debugf("SessionToken: New for Account %s, token %s.", refAccount, tokenStr)
	saveSessionToken(accountId, refAccount, tokenStr)
	return tokenStr
}

func Validate(tokenStr string) (accountId string, refAccount string, ok bool) {
	logger.Debug("TokenModule: Validate tokenStr:", tokenStr)
	refAccount, err := redis.String(cache.HGet(tokenStr, CACHE_ACCOUNT))
	if ok = (err == nil); ok {
		accountId, err = redis.String(cache.HGet(tokenStr, CACHE_ID))
		renew(refAccount, tokenStr)
	}
	return
}

// Renew session expire time, 60 minutes
func renew(refAccount string, tokenStr string) {
	logger.Debugf("TokenModule: Refresh tokenStr %s expire time.", tokenStr)
	cache.Expire(refAccount, 60)
	cache.Expire(tokenStr, 60)
}

// Save phoneNo:tokenStr -> refAccount
// Save tokenStr:refaccount -> refAccount
// Save tokenStr:refId -> refAccount
func saveSessionToken(accountId string, refAccount string, tokenStr string) error {
	var err error
	errAc := cache.HSet(tokenStr, CACHE_ACCOUNT, refAccount)
	if errAc != nil {
		logger.Debugf("TokenModule: Put session token [%s:%s], err %v", tokenStr, refAccount, errAc)
	}
	errId := cache.HSet(tokenStr, CACHE_ID, accountId)
	if errId != nil {
		logger.Debugf("TokenModule: Put session token [%s:%s], err %v", tokenStr, accountId, errId)
	}

	err = cache.Set(refAccount, tokenStr)
	if err != nil {
		cache.HDel(tokenStr, CACHE_ACCOUNT)
		logger.Debugf("TokenModule: Put session token [%s:%s], err %v", refAccount, tokenStr, err)
	}

	renew(refAccount, tokenStr)
	return err
}

// Delete token->account and account->token mapping in cache
func deleteSessionToken(refAccount string) {
	tokenStr, err := redis.String(cache.Get(refAccount))
	if err == nil {
		re, err := cache.Del(tokenStr)
		logger.Debug("delete tokenStr", tokenStr, re, err)
		re, err = cache.Del(refAccount)
		logger.Debug("delete refAccount", refAccount, re, err)
	}
}

// For calling deleteSessionToken function
func Del(refAccount string) {
	deleteSessionToken(refAccount)
}