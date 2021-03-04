package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var mEnv = map[string]interface{}{}

func find(key string, m map[string]interface{}) (value interface{}, err error) {
	if key == "" {
		err = errors.New("key was not informed")
		return
	}

	subkeys := strings.Split(key, ".")
	subkey := subkeys[0]
	exists := false

	value, exists = m[subkey]
	if !exists {
		err = fmt.Errorf("key %s is not found", subkey)
		return
	}

	if len(subkeys) > 1 {
		subm, isMap := value.(map[string]interface{})
		if !isMap {
			err = fmt.Errorf("value of key %s is not a map", subkey)
			return
		}

		return find(strings.Join(subkeys[1:], "."), subm)
	}

	return
}

func FromFile(filePath string) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &mEnv)
	fmt.Println(mEnv)
	return
}

func String(key string) (value string, err error) {
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	value = fmt.Sprint(v)
	return
}

func MustString(key string) (value string) {
	value, err := String(key)
	if err != nil {
		panic(err)
	}

	return
}
