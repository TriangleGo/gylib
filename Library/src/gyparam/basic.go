package param

import "strconv"

func GetStringWithKey(req map[string]interface{}, key string) (val string, ok bool) {
	tempVal, ok := req[key]
	if ok {
		val = tempVal.(string)
	}
	return
}

type Op string

const (
	Lte Op = "<="
	Lt Op = "<"
	Eq Op = "=="
	Gt Op = ">"
	Gte Op = ">="
)

func GetStringWithKeyWithLengthLimitation(req map[string]interface{}, key string, op Op, length int) (val string, ok bool) {
	val, ok = GetStringWithKey(req, key)
	switch op {
	case Lte:
		ok = len(val) <= length
	case Lt:
		ok = len(val) < length
	case Eq:
		ok = len(val) == length
	case Gt:
		ok = len(val) > length
	case Gte:
		ok = len(val) >= length
	default:
		ok = false
	}
	return
}

func GetFloat64WithKey(req map[string]interface{}, key string) (val float64, ok bool) {
	tempVal, ok := req[key]
	if ok {
		strVal := tempVal.(string)
		var err error
		val, err = strconv.ParseFloat(strVal, 64)
		if err != nil {
			ok = false
		}
	}
	return
}

func GetIntWithKey(req map[string]interface{}, key string) (val int, ok bool) {
	tempVal, ok := req[key]
	if ok {
		strVal := tempVal.(string)
		var err error
		val, err = strconv.Atoi(strVal)
		if err != nil {
			ok = false
			val = 0
		}
	}
	return
}

func GetIntWithKeyDefaultVal(req map[string]interface{}, key string, def int) (val int) {
	val, ok := GetIntWithKey(req, key)
	if !ok {
		val = def
	}
	return
}

func GetBoolWithKey(req map[string]interface{}, key string) (val bool, ok bool) {
	tempVal, ok := req[key]
	if ok {
		val = (tempVal.(string) == "true")
	}
	return
}