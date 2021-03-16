package env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func Set(key string, value interface{}) (err error) {
	if m, isMap := value.(map[string]interface{}); isMap {
		for subKey, subValue := range m {
			fullKey := key + "." + subKey

			err = Set(fullKey, subValue)
			if err != nil {
				return
			}
		}
	}

	err = os.Setenv(key, fmt.Sprint(value))
	return
}

func Get(key string) (value string) {
	return os.Getenv(key)
}

// FromFile opens the informed file (json) and loads the values into memory.
func FromFile(filePath string) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	v := map[string]interface{}{}

	err = json.Unmarshal(file, &v)
	if err != nil {
		return
	}

	for key, value := range v {
		err = Set(key, value)
		if err != nil {
			return
		}
	}

	return
}

// MustFromFile does the same as the FromFile function, but generates panic in case of an error.
func MustFromFile(filePath string) {
	err := FromFile(filePath)
	if err != nil {
		panic(err)
	}
}

// String fetches the value by the key and returns the value converted to a string.
// Example:
//	{
// 		"repository": {
// 			"username": "root"
// 		}
// 	}
//
// 	username, err := String("repository.username")
// 	if err != nil {
// 		// ...
// 	}
func String(key string) (value string, err error) {
	value = Get(key)
	return
}

// MustString does the same as the String function, but generates panic in case of an error.
func MustString(key string) (value string) {
	value, err := String(key)
	if err != nil {
		panic(err)
	}

	return
}

// Int fetches the value by the key and returns the value converted to a int.
func Int(key string) (value int, err error) {
	return strconv.Atoi(fmt.Sprint(Get(key)))
}

// MustInt does the same as the Int function, but generates panic in case of an error.
func MustInt(key string) (value int) {
	value, err := Int(key)
	if err != nil {
		panic(err)
	}

	return
}

// Int64 fetches the value by the key and returns the value converted to a int64.
func Int64(key string) (value int64, err error) {
	return strconv.ParseInt(fmt.Sprint(Get(key)), 10, 64)
}

// MustInt64 does the same as the Int64 function, but generates panic in case of an error.
func MustInt64(key string) (value int64) {
	value, err := Int64(key)
	if err != nil {
		panic(err)
	}

	return
}

// MustFromFile does the same as the FromFile function, but generates panic in case of an error.
func Float64(key string) (value float64, err error) {
	return strconv.ParseFloat(fmt.Sprint(Get(key)), 64)
}

// MustFloat64 does the same as the Float64 function, but generates panic in case of an error.
func MustFloat64(key string) (value float64) {
	value, err := Float64(key)
	if err != nil {
		panic(err)
	}

	return
}

// Bool fetches the value by the key and returns the value converted to a boolean.
func Bool(key string) (value bool, err error) {
	return strconv.ParseBool(Get(key))
}

// MustBool does the same as the Bool function, but generates panic in case of an error.
func MustBool(key string) (value bool) {
	value, err := Bool(key)
	if err != nil {
		panic(err)
	}

	return
}
