package uniqueId

import "github.com/sluu99/uuid"

func GetUUIDWithKey(req map[string]interface{}, key string) (val string, ok bool) {
	tempVal, ok := req[key]

	if ok {
		val = tempVal.(string)
		_, err := uuid.FromStr(val)
		if err != nil {
			val = ""
			ok = false
		} else {
			ok = true
		}
	}
	return
}