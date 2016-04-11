package entity

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	ColFirmwareVersion = "FirmwareVersion"
)

type FirmwareVersion struct {
	Id      bson.ObjectId `bson:"_id"`
	Version string `bson:"version"`
	Latest  bool `bson:"latest"`
	FileId  bson.ObjectId `bson:"file_id"`
}