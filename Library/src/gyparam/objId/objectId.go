package objId

import "gopkg.in/mgo.v2/bson"

func GetObjectIdHexStringWithKey(req map[string]interface{}, key string) (val string, ok bool) {
	tempVal, ok := req[key]

	if ok {
		val = tempVal.(string)
		ok = bson.IsObjectIdHex(val)
		if !ok {
			val = ""
		}
	}
	return
}

func GetObjectIdWithKey(req map[string]interface{}, key string) (val bson.ObjectId, ok bool) {
	tempVal, ok := req[key]

	if ok {
		valStr := tempVal.(string)
		ok = bson.IsObjectIdHex(valStr)
		if ok {
			val = bson.ObjectIdHex(valStr)
		}
	}
	return
}