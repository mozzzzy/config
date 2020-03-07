package config

/*
 * Module Dependencies
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/mozzzzy/config/json/configOption"
)

/*
 * Types
 */

type Config struct {
	options []configOption.Option
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Package Private Functions
 */

func getMaxLen(strs []string) int {
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}

func (conf Config) findOptByKey(key string) *configOption.Option {
	for index := 0; index < len(conf.options); index++ {
		if conf.options[index].Key == key {
			return &conf.options[index]
		}
	}
	return nil
}

func getAllKeys(opts *[]configOption.Option) []string {
	var keys []string
	for _, opt := range *opts {
		keys = append(keys, opt.Key)
	}
	return keys
}

func (conf *Config) parseOneLayer(kvs map[string]interface{}, parentKey string) error {
	// Get keys
	var keys []string
	for key, _ := range kvs {
		keys = append(keys, key)
	}
	// Sort keys
	sort.Strings(keys)

	for _, key := range keys {
		var absolutePath string
		if parentKey != "" {
			absolutePath = fmt.Sprintf("%v.%v", parentKey, key)
		} else {
			absolutePath = key
		}
		opt := conf.findOptByKey(absolutePath)
		// If matched option is not found, continue.
		if opt == nil {
			continue
		}
		// If found option has already set, return error.
		if opt.IsSet() == true {
			return errors.New(fmt.Sprintf("Duplicate definition of %v", key))
		}
		// If found option require value, set value.
		switch opt.ValueType {
		case "":
		case "nil":
		case "array":
			ary, ok := kvs[key].([]interface{})
			if !ok {
				return errors.New(fmt.Sprintf(
					"Invalid array value for %v \"%v\".", key, kvs[key]))
			}
			if err := opt.SetValue(ary); err != nil {
				return err
			}
		case "float64":
			flt64, ok := kvs[key].(float64)
			if !ok {
				return errors.New(fmt.Sprintf(
					"Invalid float64 value for %v \"%v\".", key, kvs[key]))
			}
			if err := opt.SetValue(flt64); err != nil {
				return err
			}
		case "int":
			flt64, ok := kvs[key].(float64)
			if !ok {
				return errors.New(fmt.Sprintf(
					"Invalid float64 value for %v \"%v\".", key, kvs[key]))
			}
			integer := int(flt64)
			if err := opt.SetValue(integer); err != nil {
				return err
			}
		case "object":
			if err := opt.SetValue(0); err != nil {
				return err
			}
			nextKvs, ok := kvs[key].(map[string]interface{})
			if !ok {
				return errors.New(fmt.Sprintf(
					"Invalid object value for %v \"%v\".", key, kvs[key]))
			}
			if err := conf.parseOneLayer(nextKvs, absolutePath); err != nil {
				return err
			}
		case "string":
			str, ok := kvs[key].(string)
			if !ok {
				return errors.New(fmt.Sprintf(
					"Invalid string value for %v \"%v\".", key, kvs[key]))
			}
			if err := opt.SetValue(str); err != nil {
				return err
			}
		}
	}
	return nil
}

func sortOptsByKey(opts []configOption.Option) []configOption.Option {
	var sortedOpts []configOption.Option

	keys := getAllKeys(&opts)
	sort.Strings(keys)

	for _, key := range keys {
		for _, opt := range opts {
			if opt.Key == key {
				sortedOpts = append(sortedOpts, opt)
				break
			}
		}
	}
	return sortedOpts
}

/*
 * Public Functions
 */

func (conf *Config) AddOption(opt configOption.Option) error {
	validatedOpt, err := configOption.New(opt)
	if err != nil {
		return err
	}
	keyElems := strings.Split(validatedOpt.Key, ".")
	if len(keyElems) > 1 {
		for keyCount := 0; keyCount < len(keyElems)-1; keyCount++ {
			joinedKey := strings.Join(keyElems[:keyCount+1], ".")
			searchedOpt := conf.findOptByKey(joinedKey)
			if searchedOpt == nil {
				return errors.New(fmt.Sprintf("Parent option \"%v\" is not found.", searchedOpt))
			}
		}
	}
	conf.options = append(conf.options, *validatedOpt)
	return nil
}

func (conf *Config) AddOptions(opts []configOption.Option) error {
	sortedOpts := sortOptsByKey(opts)
	for _, opt := range sortedOpts {
		if err := conf.AddOption(opt); err != nil {
			return err
		}
	}
	return nil
}

func (conf Config) Get(key string) (interface{}, error) {
	// Find key from keys
	opt := conf.findOptByKey(key)
	// If requested key is not found, return error.
	if opt == nil {
		return nil, errors.New(fmt.Sprintf("Required key \"%v\" is not found.", key))
	}

	if opt.ValueType == "object" {
		var childConf interface{}
		var err error
		childConf, err = conf.GetObject(opt.Key)
		return childConf, err
	}

	// If requested option and its default value are not set, return error
	return opt.GetValue()
}

func (conf Config) GetFloat64(key string) (float64, error) {
	var zeroVal float64
	value, err := conf.Get(key)
	if err != nil {
		return zeroVal, err
	}
	flt64, ok := value.(float64)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf(
			"Value of option \"%v\" is not float64. Its type is %T.", key, value))
	}
	return flt64, nil
}

func (conf Config) GetFloat64Array(key string) ([]float64, error) {
	var zeroVal []float64
	value, err := conf.Get(key)
	if err != nil {
		return zeroVal, err
	}
	ifArray, ok := value.([]interface{})
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf(
			"Value of option \"%v\" is not []interface{}. Its type is %T.",
			key,
			value,
		))
	}
	var flt64Array []float64
	for _, i := range ifArray {
		flt64, ok := i.(float64)
		if !ok {
			return zeroVal, errors.New(fmt.Sprintf(
				"Element of option \"%v\" is not float64. Its type is %T.", key, i))
		}
		flt64Array = append(flt64Array, flt64)
	}
	return flt64Array, nil
}

func (conf Config) GetInt(key string) (int, error) {
	var zeroVal int
	value, err := conf.Get(key)
	if err != nil {
		return zeroVal, err
	}
	integer, ok := value.(int)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf(
			"Value of option \"%v\" is not int. Its type is %T.", key, value))
	}
	return integer, nil
}

func (conf Config) GetIntArray(key string) ([]int, error) {
	var zeroVal []int
	flt64Array, err := conf.GetFloat64Array(key)
	if err != nil {
		return zeroVal, err
	}

	var intArray []int
	for _, flt64 := range flt64Array {
		integer := int(flt64)
		intArray = append(intArray, integer)
	}
	return intArray, nil
}

func (conf Config) GetObject(key string) (Config, error) {
	var childConf Config

	// Find key from keys
	opt := conf.findOptByKey(key)
	// If requested key is not found, return error.
	if opt == nil {
		return childConf, errors.New(fmt.Sprintf("Required key \"%v\" is not found.", key))
	}

	for _, opt := range conf.options {
		if opt.Key != key && strings.HasPrefix(opt.Key, key) {
			childConf.options = append(childConf.options, opt)
		}
	}
	return childConf, nil
}

func (conf Config) GetString(key string) (string, error) {
	var zeroVal string
	value, err := conf.Get(key)
	if err != nil {
		return zeroVal, err
	}
	str, ok := value.(string)
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf(
			"Value of option \"%v\" is not string. Its type is %T.", key, value))
	}
	return str, nil
}

func (conf Config) GetStringArray(key string) ([]string, error) {
	var zeroVal []string
	value, err := conf.Get(key)
	if err != nil {
		return zeroVal, err
	}
	ifArray, ok := value.([]interface{})
	if !ok {
		return zeroVal, errors.New(fmt.Sprintf(
			"Value of option \"%v\" is not []interface{}. Its type is %T.",
			key,
			value,
		))
	}
	var stringArray []string
	for _, i := range ifArray {
		str, ok := i.(string)
		if !ok {
			return zeroVal, errors.New(fmt.Sprintf(
				"Element of option \"%v\" is not float64. Its type is %T.", key, i))
		}
		stringArray = append(stringArray, str)
	}
	return stringArray, nil
}

func (conf *Config) Parse(path string) error {
	/// Read config file and parse into "map[string]interface{}"
	raw, readFileErr := ioutil.ReadFile(path)
	if readFileErr != nil {
		return readFileErr
	}
	keyValues := make(map[string]interface{})
	if unmarshalErr := json.Unmarshal(raw, &(keyValues)); unmarshalErr != nil {
		return unmarshalErr
	}
	if err := conf.parseOneLayer(keyValues, ""); err != nil {
		return err
	}
	return conf.Validate()
}

func (conf Config) String() string {
	str := "Following options are avairable.\n"

	// Get max str length of keys
	keys := getAllKeys(&conf.options)
	maxLen := getMaxLen(keys)

	for index, opt := range conf.options {
		// key
		key := keys[index]
		str += "  " + key
		for index := 0; index < maxLen-len(key); index++ {
			str += " "
		}
		// description
		if opt.Description != "" {
			str += ": "
			str += opt.Description
		}
		// required
		if opt.Required == true {
			str += " (required)"
		}
		// default value
		if opt.ValueType != "" && opt.ValueType != "nil" {
			if opt.DefaultValue != nil {
				switch opt.ValueType {
				case "int":
					defaultValInt, ok := opt.DefaultValue.(int)
					if ok {
						str += fmt.Sprintf(" (default: %v)", defaultValInt)
					}
				case "string":
					defaultValStr, ok := opt.DefaultValue.(string)
					if ok {
						str += fmt.Sprintf(" (default: \"%v\")", defaultValStr)
					}
				}
			}
		}
		str += "\n"
	}
	return str
}

func (conf Config) Validate() error {
	for _, opt := range conf.options {
		if err := opt.Validate(); err != nil {
			return err
		}
	}
	return nil
}
