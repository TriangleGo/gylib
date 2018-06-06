package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ToBytes(any interface{}) (result []byte) {
	result, err := json.Marshal(any)
	if err != nil {
		result = nil
	}
	return
}

func FromBytes(input []byte, output interface{}) (err error) {
	err = json.Unmarshal(input, output)
	if err != nil {
		fmt.Errorf("Unmarshal error %v.\n", err)
	}
	return
}

func ToString(any interface{}) (string) {
	bytes, err := json.Marshal(any)
	if err == nil {
		return string(bytes)
	} else {
		return ""
	}
}

func FromString(input string, output interface{}) (err error) {
	err = json.Unmarshal([]byte(input), output)
	if err != nil {
		fmt.Errorf("Unmarshal from string error %v.\n", err)
	}
	return
}

func FromFile(path string, obj interface{}) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf("Read file error %v.\n", err)
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		fmt.Errorf("Unmarshal error %v.\n", err)
	}
	return
}

func ToFile(file *os.File, obj interface{}) (err error) {
	result := ToString(obj)
	fmt.Println("to file " + result)

	_, err = file.WriteString(result)
	fmt.Println(ToString(err))
	return
}

func ToMap(any interface{}) (result map[string]interface{}) {
	bytes := ToBytes(any)
	if bytes == nil {
		result = nil
		return
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil
	}
	return result
}
