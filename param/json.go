package param

import (
	"encoding/json"
	"fmt"
)

func GetJsonArrayWithKey(req map[string]interface{}, key string) (val[] string, ok bool) {
	fmt.Printf("input %v, key %s", req, key)
	tempVal, ok := req[key]
	var inputStr string
	if ok {
		inputStr = tempVal.(string)
		fmt.Printf("input str: %s", inputStr)
	}
	err := json.Unmarshal([]byte(inputStr), &val)
	if err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(val)
	} else {
		fmt.Print(err.Error())
	}

	return
}

