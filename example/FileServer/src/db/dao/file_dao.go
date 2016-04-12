package dao
import (
	"gopkg.in/mgo.v2/bson"
	"github.com/TriangleGo/gylib/mongo"
	"gopkg.in/mgo.v2"
	"errors"
	"github.com/TriangleGo/gylib/logger"
)

const category = "gridfs"

func LoadFile(fileId string) (name string, contentType string, content []byte, err error) {

	if (!bson.IsObjectIdHex(fileId)) {
		err = errors.New("invalid id hex string.")
		return
	}

	gridFS := mongo.GetGridFS(category)
	gridFile, err := gridFS.OpenId(bson.ObjectIdHex(fileId))

	logger.Debugf("find grid file %v, err %v:", gridFile, err)
	if err == mgo.ErrNotFound {
		return
	}
	name = gridFile.Name()
	logger.Debugf("read filename, ", name)
	contentType = gridFile.ContentType()
	content = make([]byte, gridFile.Size())
	gridFile.Read(content)
	return
}

func SaveFile(name string, contentType string, content []byte) (fileId string, err error) {
	gridFS := mongo.GetGridFS(category)
	file, err := gridFS.Create(name)
	if err != nil {
		return
	} else {
		fileId = file.Id().(bson.ObjectId).Hex()
		file.SetContentType(contentType)
		num, err := file.Write(content)
		logger.Debug(num, ":", err)
		file.Close()
	}
	return
}

func DeleteFile(fileId string) (err error) {
	gridFS := mongo.GetGridFS(category)
	if (!bson.IsObjectIdHex(fileId)) {
		err = errors.New("invalid id hex string.")
	} else {
		err = gridFS.RemoveId(bson.ObjectIdHex(fileId))
	}
	return
}