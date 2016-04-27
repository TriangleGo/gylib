package param

import (
	"time"
	"strconv"
)

func GetTimeWithKey(req map[string]interface{}, key string) (date time.Time, ok bool) {
	interfaceVal, ok := req[key]
	if ok {
		strVal := interfaceVal.(string)
		var err error
		dateInt, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			ok = false
		} else {
			ok = true
			date = time.Unix(dateInt, 0)
		}
	}
	return
}

func GetTimeAsInt64WithKey(req map[string]interface{}, key string) (checkTime int64, ok bool) {
	interfaceVal, ok := req[key]
	if ok {
		strVal := interfaceVal.(string)
		var err error
		checkTime, err = strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			ok = false
			checkTime = 0
		}
	}
	return
}