package entity

import (
	"gopkg.in/mgo.v2/bson"
	"gymongo"
	"gylogger"
	"strings"
	"errors"
)

const (
	ColAppVersion = "AppVersion"
)

type AppVersion struct {
	Id        bson.ObjectId `bson:"_id"`
	Platform  string `bson:"platform"`
	Version   string `bson:"version"`
	Sub       string `bson:"sub"`
	Iteration string `bson:"iteration"`
	Channel   string `bson:"channel"`
	Valid     bool `bson:"valid"`
	FileId    bson.ObjectId `bson:"file_id"`
}

func ParseAppVersionStr(versionStr string) (version, sub, iteration, channel string, err error) {
	parts := strings.Split(versionStr, ".")
	len := len(parts)
	logger.Infof("len = %d", len)
	if !(len == 3 || len == 4) {
		err = errors.New("Invalid version format.")
		logger.Errorf("Parse version error.")
		return
	}

	version = parts[0]
	sub = parts[1]
	iteration = parts[2]
	if len == 4 {
		channel = parts[3]
	} else {
		channel = ""
	}
	return
}

func InitAppVersion() {
	collection := mongo.GetCollection(ColAppVersion)
	countAndroid, _ := collection.Find(bson.M{"platform":"android"}).Count()
	countIos, _ := collection.Find(bson.M{"platform":"ios"}).Count()
	if countAndroid <= 0 {
		androidAppVersion := AppVersion{
			Id:bson.NewObjectId(),
			Platform:"android",
			Version:"1",
			Sub:"0",
			Iteration:"2",
			Valid:true}
		err := collection.Insert(androidAppVersion)
		logger.Infof("Default android app version inited, err %v.", err)
	}
	if countIos <= 0 {
		iosAppVersion := AppVersion{
			Id:bson.NewObjectId(),
			Platform:"ios",
			Version:"1",
			Sub:"0",
			Iteration:"2",
			Valid:true}
		err := collection.Insert(iosAppVersion)
		logger.Infof("Default ios app version inited, err %v.", err)
	}
}