package config

/*
 * Module Dependencies
 */

import (
	"testing"

	"github.com/mozzzzy/config/json/configOption"
	"github.com/mozzzzy/config/validator"
	"github.com/mozzzzy/testUtil"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

const ONE_FLOAT64_JSON string = "testData/one_float64.json"
const ONE_FLOAT64_ARRAY_JSON string = "testData/one_float64_array.json"
const ONE_INT_JSON string = "testData/one_int.json"
const ONE_INT_ARRAY_JSON string = "testData/one_int_array.json"
const ONE_INT64_JSON string = "testData/one_int64.json"
const ONE_INT64_ARRAY_JSON string = "testData/one_int64_array.json"
const ONE_OBJECT_JSON string = "testData/one_object.json"
const ONE_STRING_JSON string = "testData/one_string.json"
const ONE_STRING_ARRAY_JSON string = "testData/one_string_array.json"
const ALL_IN_ONE_JSON string = "testData/all_in_one.json"

/*
 * Functions
 */

func TestAddOption(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("float64", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("float64 with default value", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:          "float64",
				ValueType:    "float64",
				Description:  "some description.",
				DefaultValue: 1.2,
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("int", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("int with default value", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:          "int",
				ValueType:    "int",
				Description:  "some description.",
				DefaultValue: 10,
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("object", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("string", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("string with default value", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:          "string",
				ValueType:    "string",
				Description:  "some description.",
				DefaultValue: "some default value",
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("required", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
				Required:    true,
			},
		)
		testUtil.NoError(t, err)
	})

	t.Run("invalid (required and has default value)", func(t *testing.T) {
		var conf Config
		err := conf.AddOption(
			configOption.Option{
				Key:          "int",
				ValueType:    "int",
				Description:  "some description.",
				DefaultValue: 10,
				Required:     true,
			},
		)
		testUtil.WithError(t, err)
	})
}

func TestAddOptions(t *testing.T) {
	t.Run("multi options", func(t *testing.T) {
		var conf Config
		err := conf.AddOptions([]configOption.Option{
			{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
				Required:    true,
			},
			{
				Key:          "float64",
				ValueType:    "float64",
				Description:  "some description.",
				DefaultValue: 1.2,
			},
			{
				Key:          "int",
				ValueType:    "int",
				Description:  "some description.",
				DefaultValue: 10,
			},
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:          "string",
				ValueType:    "string",
				Description:  "some description.",
				DefaultValue: "some default value",
			},
		})
		testUtil.NoError(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Run("get array", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_ARRAY_JSON)
		testUtil.NoError(t, parseErr)

		expect := []string{"some", "value"}
		actual, getErr := conf.Get("array")
		testUtil.NoError(t, getErr)

		castedActual, ok := actual.([]interface{})
		testUtil.Match(t, true, ok)
		testUtil.Match(t, 2, len(castedActual))

		firstElem, firstElemOk := castedActual[0].(string)
		testUtil.Match(t, true, firstElemOk)
		testUtil.Match(t, expect[0], firstElem)

		secondElem, secondElemOk := castedActual[1].(string)
		testUtil.Match(t, true, secondElemOk)
		testUtil.Match(t, expect[1], secondElem)
	})

	t.Run("get float64", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "float64",
				ValueType:   "float64",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_FLOAT64_JSON)
		testUtil.NoError(t, parseErr)

		var expect float64 = 1.2
		actual, getErr := conf.Get("float64")
		testUtil.NoError(t, getErr)
		castedActual, ok := actual.(float64)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expect, castedActual)
	})

	t.Run("get int", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_JSON)
		testUtil.NoError(t, parseErr)

		var expect int = 10
		actual, getErr := conf.Get("int")
		testUtil.NoError(t, getErr)
		castedActual, ok := actual.(int)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expect, castedActual)
	})

	t.Run("get int64", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "int64",
				ValueType:   "int64",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT64_JSON)
		testUtil.NoError(t, parseErr)
		var expect int64 = 2147483648
		actual, getErr := conf.Get("int64")
		testUtil.NoError(t, getErr)
		castedActual, ok := actual.(int64)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expect, castedActual)
	})

	t.Run("get object", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.key",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_OBJECT_JSON)
		testUtil.NoError(t, parseErr)

		childConf, getErr := conf.Get("object")
		testUtil.NoError(t, getErr)

		castedChildConf, ok := childConf.(Config)
		testUtil.Match(t, true, ok)

		actual, getStringErr := castedChildConf.Get("key")
		testUtil.NoError(t, getStringErr)

		castedActual, ok := actual.(string)
		testUtil.Match(t, true, ok)

		expect := "value"
		testUtil.Match(t, expect, castedActual)
	})

	t.Run("get string", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.NoError(t, parseErr)

		var expect string = "some value"
		actual, getErr := conf.Get("string")
		testUtil.NoError(t, getErr)
		castedActual, ok := actual.(string)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expect, castedActual)
	})

	t.Run("get multi", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
			{
				Key:         "float64",
				ValueType:   "float64",
				Description: "some description.",
			},
			{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
			{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.array",
				ValueType:   "array",
				Description: "some description.",
			},
			{
				Key:         "object.float64",
				ValueType:   "float64",
				Description: "some description.",
			},
			{
				Key:         "object.int",
				ValueType:   "int",
				Description: "some description.",
			},
			{
				Key:         "object.string",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)

		// Check array
		expectAry := []string{"some", "value"}
		actualAry, getErr := conf.Get("array")
		testUtil.NoError(t, getErr)

		castedActualAry, ok := actualAry.([]interface{})
		testUtil.Match(t, true, ok)
		testUtil.Match(t, 2, len(castedActualAry))

		firstElem, firstElemOk := castedActualAry[0].(string)
		testUtil.Match(t, true, firstElemOk)
		testUtil.Match(t, expectAry[0], firstElem)

		secondElem, secondElemOk := castedActualAry[1].(string)
		testUtil.Match(t, true, secondElemOk)
		testUtil.Match(t, expectAry[1], secondElem)

		// Check float64
		expectFlt64 := 1.2
		actualFlt64, getErr := conf.Get("float64")
		testUtil.NoError(t, getErr)
		castedActualFlt64, ok := actualFlt64.(float64)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectFlt64, castedActualFlt64)

		// Check int
		expectInt := 10
		actualInt, getErr := conf.Get("int")
		testUtil.NoError(t, getErr)
		castedActualInt, ok := actualInt.(int)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectInt, castedActualInt)

		// Check string
		expectStr := "some value"
		actualStr, getErr := conf.Get("string")
		testUtil.NoError(t, getErr)
		castedActualStr, ok := actualStr.(string)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectStr, castedActualStr)

		// Check array in object
		expectAry2 := []string{"some", "value", "in", "object"}
		actualAry2, getErr := conf.Get("object.array")
		testUtil.NoError(t, getErr)

		castedActualAry2, ok := actualAry2.([]interface{})
		testUtil.Match(t, true, ok)
		testUtil.Match(t, 4, len(castedActualAry2))

		firstElem2, firstElemOk := castedActualAry2[0].(string)
		testUtil.Match(t, true, firstElemOk)
		testUtil.Match(t, expectAry[0], firstElem2)

		secondElem2, secondElemOk := castedActualAry2[1].(string)
		testUtil.Match(t, true, secondElemOk)
		testUtil.Match(t, expectAry2[1], secondElem2)

		thirdElem2, thirdElemOk := castedActualAry2[2].(string)
		testUtil.Match(t, true, thirdElemOk)
		testUtil.Match(t, expectAry2[2], thirdElem2)

		fourthElem2, fourthElemOk := castedActualAry2[3].(string)
		testUtil.Match(t, true, fourthElemOk)
		testUtil.Match(t, expectAry2[3], fourthElem2)

		// Check float64 in object
		expectFlt642 := 2.2
		actualFlt642, getErr := conf.Get("object.float64")
		testUtil.NoError(t, getErr)
		castedActualFlt642, ok := actualFlt642.(float64)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectFlt642, castedActualFlt642)

		// Check int in object
		expectInt2 := 20
		actualInt2, getErr := conf.Get("object.int")
		testUtil.NoError(t, getErr)
		castedActualInt2, ok := actualInt2.(int)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectInt2, castedActualInt2)

		// Check string in object
		expectStr2 := "some value in object"
		actualStr2, getErr := conf.Get("object.string")
		testUtil.NoError(t, getErr)
		castedActualStr2, ok := actualStr2.(string)
		testUtil.Match(t, true, ok)
		testUtil.Match(t, expectStr2, castedActualStr2)
	})
}

func TestGetAllKeys(t *testing.T) {
	t.Run("some keys exist", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option {
			{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
			{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		keys := conf.GetAllKeys()
		testUtil.Match(t, 2, len(keys))
		testUtil.Match(t, "array", keys[0])
		testUtil.Match(t, "string", keys[1])
	})
}

func TestGetFloat64Array(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_FLOAT64_ARRAY_JSON)
		testUtil.NoError(t, parseErr)

		expect := []float64{1.2, 2.2}
		actual, getErr := conf.GetFloat64Array("array")
		testUtil.NoError(t, getErr)

		testUtil.Match(t, 2, len(actual))
		if len(actual) == 2 {
			testUtil.Match(t, expect[0], actual[0])
			testUtil.Match(t, expect[1], actual[1])
		}
	})
}

func TestGetIntArray(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_ARRAY_JSON)
		testUtil.NoError(t, parseErr)

		expect := []int{10, 20}
		actual, getErr := conf.GetIntArray("array")
		testUtil.NoError(t, getErr)

		testUtil.Match(t, 2, len(actual))
		if len(actual) == 2 {
			testUtil.Match(t, expect[0], actual[0])
			testUtil.Match(t, expect[1], actual[1])
		}
	})
}

func TestGetInt64Array(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT64_ARRAY_JSON)
		testUtil.NoError(t, parseErr)

		expect := []int64{-2147483649, 2147483648}
		actual, getErr := conf.GetInt64Array("array")
		testUtil.NoError(t, getErr)

		testUtil.Match(t, 2, len(actual))
		if len(actual) == 2 {
			testUtil.Match(t, expect[0], actual[0])
			testUtil.Match(t, expect[1], actual[1])
		}
	})
}

func TestGetStringArray(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_ARRAY_JSON)
		testUtil.NoError(t, parseErr)

		expect := []string{"some", "value"}
		actual, getErr := conf.GetStringArray("array")
		testUtil.NoError(t, getErr)

		testUtil.Match(t, 2, len(actual))
		if len(actual) == 2 {
			testUtil.Match(t, expect[0], actual[0])
			testUtil.Match(t, expect[1], actual[1])
		}
	})
}

func TestGetFloat64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "float64",
				ValueType:   "float64",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_FLOAT64_JSON)
		testUtil.NoError(t, parseErr)

		var expect float64 = 1.2
		actual, getErr := conf.GetFloat64("float64")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("valid in object", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.float64",
				ValueType:   "float64",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)

		var expect float64 = 2.2
		actual, getErr := conf.GetFloat64("object.float64")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("invalid (string)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.NoError(t, parseErr)

		_, getErr := conf.GetFloat64("string")
		testUtil.WithError(t, getErr)
	})
}

func TestGetInt(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_JSON)
		testUtil.NoError(t, parseErr)

		var expect int = 10
		actual, getErr := conf.GetInt("int")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("valid in object", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.int",
				ValueType:   "int",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)

		var expect int = 20
		actual, getErr := conf.GetInt("object.int")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("invalid (string)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.NoError(t, parseErr)

		_, getErr := conf.GetInt("string")
		testUtil.WithError(t, getErr)
	})
}

func TestGetInt64(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "int64",
				ValueType:   "int64",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT64_JSON)
		testUtil.NoError(t, parseErr)

		var expect int64 = 2147483648
		actual, getErr := conf.GetInt64("int64")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("valid in object", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.int",
				ValueType:   "int",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)

		var expect int = 20
		actual, getErr := conf.GetInt("object.int")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("invalid (string)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.NoError(t, parseErr)

		_, getErr := conf.GetInt("string")
		testUtil.WithError(t, getErr)
	})
}



func TestGetObject(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.key",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_OBJECT_JSON)
		testUtil.NoError(t, parseErr)

		childConf, getErr := conf.GetObject("object")
		testUtil.NoError(t, getErr)

		actual, getStringErr := childConf.GetString("key")
		testUtil.NoError(t, getStringErr)

		expect := "value"
		testUtil.Match(t, expect, actual)
	})
}

func TestGetString(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.NoError(t, parseErr)

		var expect string = "some value"
		actual, getErr := conf.GetString("string")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("valid in object", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.string",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)

		var expect string = "some value in object"
		actual, getErr := conf.GetString("object.string")
		testUtil.NoError(t, getErr)
		testUtil.Match(t, expect, actual)
	})

	t.Run("invalid (int)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_JSON)
		testUtil.NoError(t, parseErr)

		_, getErr := conf.GetString("int")
		testUtil.WithError(t, getErr)
	})
}

func TestParse(t *testing.T) {
	t.Run("valid multi", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOptions([]configOption.Option{
			{
				Key:         "array",
				ValueType:   "array",
				Description: "some description.",
			},
			{
				Key:         "float64",
				ValueType:   "float64",
				Description: "some description.",
			},
			{
				Key:         "int",
				ValueType:   "int",
				Description: "some description.",
			},
			{
				Key:         "string",
				ValueType:   "string",
				Description: "some description.",
			},
			{
				Key:         "object",
				ValueType:   "object",
				Description: "some description.",
			},
			{
				Key:         "object.array",
				ValueType:   "array",
				Description: "some description.",
			},
			{
				Key:         "object.float64",
				ValueType:   "float64",
				Description: "some description.",
			},
			{
				Key:         "object.int",
				ValueType:   "int",
				Description: "some description.",
			},
			{
				Key:         "object.string",
				ValueType:   "string",
				Description: "some description.",
			},
		})
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ALL_IN_ONE_JSON)
		testUtil.NoError(t, parseErr)
	})

	t.Run("invalid (type is int <-> value is string)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:         "string",
				ValueType:   "int", // this should be string
				Description: "some description.",
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_STRING_JSON)
		testUtil.WithError(t, parseErr)
	})
}

func TestValidate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:            "int",
				ValueType:      "int",
				Description:    "some description.",
				Validator:      validator.IntSmallerThan,
				ValidatorParam: 100,
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_JSON)
		testUtil.NoError(t, parseErr)
	})

	t.Run("invalid (validation error)", func(t *testing.T) {
		var conf Config
		addOptionErr := conf.AddOption(
			configOption.Option{
				Key:            "int",
				ValueType:      "int",
				Description:    "some description.",
				Validator:      validator.IntSmallerThan,
				ValidatorParam: 1,
			},
		)
		testUtil.NoError(t, addOptionErr)

		parseErr := conf.Parse(ONE_INT_JSON)
		testUtil.WithError(t, parseErr)
	})
}
