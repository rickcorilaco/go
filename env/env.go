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
	return
}

func MustFromFile(filePath string) {
	err := FromFile(filePath)
	if err != nil {
		panic(err)
	}
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

func TryString(key string) (value string) {
	value, _ = String(key)
	return
}

func Bool(key string) (value bool, err error) {
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	value, ok := v.(bool)
	if !ok {
		err = fmt.Errorf("invalid bool to key: %s", key)
	}

	return
}

func MustBool(key string) (value bool) {
	value, err := Bool(key)
	if err != nil {
		panic(err)
	}

	return
}

func TryBool(key string) (value bool) {
	value, _ = Bool(key)
	return
}
