package test
import "encoding/json"

func ToJsonString(any interface{}) (string) {
	bytes, err := json.Marshal(any)
	if err == nil {
		return string(bytes)
	} else {
		return ""
	}
}

func ToJsonByte(any interface{}) (result []byte) {
	result, err := json.Marshal(any)
	if err != nil {
		result = nil
	}
	return
}