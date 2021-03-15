package env

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
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

// FromFile opens the informed file (json) and loads the values into memory.
func FromFile(filePath string) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &mEnv)
	// if err != nil {
	// 	return
	// }

	// for key, value := range mEnv {
	// 	var envValue string

	// 	switch v := value.(type) {
	// 	case string:
	// 		envValue = v
	// 	case int:
	// 		envValue = strconv.Itoa(v)
	// 	case int64:
	// 		envValue = strconv.FormatInt(v, 64)
	// 	case float64:
	// 		envValue = strconv.FormatFloat(v, 64, 0, 64)
	// 	case bool:
	// 		envValue = strconv.FormatBool(v)
	// 	case map[string]interface{}:
	// 		// todo: recursion...
	// 	default:
	// 		err = errors.New("")
	// 		return
	// 	}

	// 	err = os.Setenv(key, envValue)
	// 	if err != nil {
	// 		return
	// 	}
	// }

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
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	value = fmt.Sprint(v)
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

// TryString tries to get a string value, if not found then return the default value of type
func TryString(key string) (value string) {
	value, _ = String(key)
	return
}

// Int fetches the value by the key and returns the value converted to a int.
func Int(key string) (value int, err error) {
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	return strconv.Atoi(fmt.Sprint(v))
}

// MustInt does the same as the Int function, but generates panic in case of an error.
func MustInt(key string) (value int) {
	value, err := Int(key)
	if err != nil {
		panic(err)
	}

	return
}

// TryInt tries to get a int value, if not found then return the default value of type
func TryInt(key string) (value int) {
	value, _ = Int(key)
	return
}

// Int64 fetches the value by the key and returns the value converted to a int64.
func Int64(key string) (value int64, err error) {
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	return strconv.ParseInt(fmt.Sprint(v), 10, 64)
}

// MustInt64 does the same as the Int64 function, but generates panic in case of an error.
func MustInt64(key string) (value int64) {
	value, err := Int64(key)
	if err != nil {
		panic(err)
	}

	return
}

// TryInt64 tries to get a int64 value, if not found then return the default value of type
func TryInt64(key string) (value int64) {
	value, _ = Int64(key)
	return
}

// MustFromFile does the same as the FromFile function, but generates panic in case of an error.
func Float64(key string) (value float64, err error) {
	v, err := find(key, mEnv)
	if err != nil {
		return
	}

	return strconv.ParseFloat(fmt.Sprint(v), 64)
}

// MustFloat64 does the same as the Float64 function, but generates panic in case of an error.
func MustFloat64(key string) (value float64) {
	value, err := Float64(key)
	if err != nil {
		panic(err)
	}

	return
}

// TryFloat64 tries to get a float64 value, if not found then return the default value of type.
func TryFloat64(key string) (value float64) {
	value, _ = Float64(key)
	return
}

// Bool fetches the value by the key and returns the value converted to a boolean.
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

// MustBool does the same as the Bool function, but generates panic in case of an error.
func MustBool(key string) (value bool) {
	value, err := Bool(key)
	if err != nil {
		panic(err)
	}

	return
}

// TryBool tries to get a bool value, if not found then return the default value of type.
func TryBool(key string) (value bool) {
	value, _ = Bool(key)
	return
}
