package configOption

/*
 * Module Dependencies
 */

import (
	"testing"

	"github.com/mozzzzy/testUtil"
	"github.com/mozzzzy/config/validator"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Package Private Functions
 */

/*
 * Public Functions
 */

func TestNew(t *testing.T) {
	t.Run("one required int option", func(t *testing.T) {
		_, err := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			Required: true,
		})
		testUtil.NoError(t, err)
	})

	t.Run("one int option", func(t *testing.T) {
		_, err := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
		})
		testUtil.NoError(t, err)
	})

	t.Run("one required string option", func(t *testing.T) {
		_, err := New(Option{
			Key: "string",
			ValueType: "string",
			Description: "some string value",
			Required: true,
		})
		testUtil.NoError(t, err)
	})

	t.Run("one string option", func(t *testing.T) {
		_, err := New(Option{
			Key: "string",
			ValueType: "string",
			Description: "some string value",
			DefaultValue: "default value",
		})
		testUtil.NoError(t, err)
	})

	t.Run("one option with validator", func(t *testing.T) {
		_, err := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
			Validator: validator.IntWithin,
			ValidatorParam: []int{9, 11},
		})
		testUtil.NoError(t, err)
	})

	t.Run("invalid option (no key)", func(t *testing.T) {
		_, err := New(Option{
			Key: "",
			ValueType: "string",
			Description: "some string value",
			DefaultValue: "default value",
		})
		testUtil.WithError(t, err)
	})

	t.Run("invalid option (required and default value are specified)", func(t *testing.T) {
		_, err := New(Option{
			Key: "",
			ValueType: "string",
			Description: "some string value",
			Required: true,
			DefaultValue: "default value",
		})
		testUtil.WithError(t, err)
	})

	t.Run("invalid option with validator (validation failed)", func(t *testing.T) {
		_, err := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
			Validator: validator.IntSmallerThan,
			ValidatorParam: 10,
		})
		testUtil.NoError(t, err)
	})
}

func TestSetValue(t *testing.T) {
	t.Run("set value to a required int option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			Required: true,
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue(10)
		testUtil.NoError(t, setValueErr)
		testUtil.Match(t, true, opt.IsSet())

		var expected interface{} = 10
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("set value to an int option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue(10)
		testUtil.NoError(t, setValueErr)
		testUtil.Match(t, true, opt.IsSet())

		var expected interface{} = 10
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("set value to a required string option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "string",
			ValueType: "string",
			Description: "some string value",
			Required: true,
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue("some value")
		testUtil.NoError(t, setValueErr)
		testUtil.Match(t, true, opt.IsSet())

		var expected interface{} = "some value"
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("set value to an string option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "string",
			ValueType: "string",
			Description: "some string value",
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue("some value")
		testUtil.NoError(t, setValueErr)
		testUtil.Match(t, true, opt.IsSet())

		var expected interface{} = "some value"
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})
}

func TestIsSet(t *testing.T) {
	t.Run("no set option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
		})
		testUtil.NoError(t, newErr)
		testUtil.Match(t, false, opt.IsSet())
	})

	t.Run("no set option with default value", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
		})
		testUtil.NoError(t, newErr)
		testUtil.Match(t, false, opt.IsSet())
	})

	t.Run("set option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue(10)
		testUtil.NoError(t, setValueErr)
		testUtil.Match(t, true, opt.IsSet())
	})
}

func TestGetValue(t *testing.T) {
	t.Run("no set option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
		})
		testUtil.NoError(t, newErr)

		var expected interface{}
		expected = nil
		actual, getValueErr := opt.GetValue()
		testUtil.WithError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("no set option with default value", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
		})
		testUtil.NoError(t, newErr)

		var expected interface{}
		expected = 10
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("set option", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue(10)
		testUtil.NoError(t, setValueErr)

		var expected interface{}
		expected = 10
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})

	t.Run("set option with default value", func(t *testing.T) {
		opt, newErr := New(Option{
			Key: "int",
			ValueType: "int",
			Description: "some int value",
			DefaultValue: 10,
		})
		testUtil.NoError(t, newErr)

		setValueErr := opt.SetValue(20)
		testUtil.NoError(t, setValueErr)

		var expected interface{}
		expected = 20
		actual, getValueErr := opt.GetValue()
		testUtil.NoError(t, getValueErr)
		testUtil.Match(t, expected, actual)
	})
}
