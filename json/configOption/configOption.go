package configOption

/*
 * Module Dependencies
 */

import (
	"errors"
	"fmt"
)

/*
 * Types
 */

type Option struct {
	Key       string
	Description    string
	ValueType      string
	DefaultValue   interface{}
	Value          interface{}
	Required       bool
	set            bool
	Validator      func(interface{}, interface{}) error
	ValidatorParam interface{}
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Package Private Functions
 */

func validateRule(opt Option) error {
	if opt.Key == "" {
		return errors.New("Key is required.")
	}
	if opt.Required && opt.DefaultValue != nil {
		return errors.New(
			fmt.Sprintf(
				"Required option %v can't be specified its default value.",
				opt.Key))
	}
	if opt.DefaultValue != nil {
		switch(opt.ValueType) {
			case "int":
				if val, ok := opt.DefaultValue.(int); !ok {
					return errors.New(fmt.Sprintf("Invalid int default value %v.", val))
				}
			case "string":
				if val, ok := opt.DefaultValue.(string); !ok {
					return errors.New(fmt.Sprintf("Invalid string default value %v.", val))
				}
		}
	}
	return nil
}

/*
 * Public Functions
 */

func (opt *Option) GetValue() (interface{}, error) {
	if !opt.set && opt.DefaultValue == nil {
		return nil, errors.New(
			fmt.Sprintf(
				"No value and no default value for %v are set.",
				opt.Key))
	}
	if !opt.set {
		return opt.DefaultValue, nil
	}
	return opt.Value, nil
}

func (opt Option)IsSet() bool {
	return opt.set
}

func New(opt Option) (*Option, error) {
	// Validate
	if err := validateRule(opt); err != nil {
		return nil, err
	}

	// Set value type to nil if it is not specified
	if opt.ValueType == "" {
		opt.ValueType = "nil"
	}

	return &opt, nil
}

func (opt *Option) SetValue(value interface{}) error {
	if value == nil {
		return errors.New("nil is invalid for SetValue func's param.")
	}
	switch opt.ValueType {
	case "":
	case "nil":
	case "array":
		ary, ok := value.([]interface{})
		if ok {
			opt.Value = ary
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is array ([]interface{}). "+
						"But specified value is %T.", value))
		}
	case "float64":
		flt64, ok := value.(float64)
		if ok {
			opt.Value = flt64
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is float64. "+
						"But specified value is %T.", value))
		}
	case "int":
		integer, ok := value.(int)
		if ok {
			opt.Value = integer
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is int. "+
						"But specified value is %T.", value))
		}
	case "string":
		str, ok := value.(string)
		if ok {
			opt.Value = str
		} else {
			return errors.New(
				fmt.Sprintf(
					"Failed to SetValue to option. "+
						"The ValueType is string. "+
						"But specified value is %T.", value))
		}
	}
	opt.set = true
	return nil
}

func (opt Option) Validate() error {
	// Required but not set
	if opt.Required && opt.set == false {
		return errors.New(
			fmt.Sprintf("Required option %v is not provided.", opt.Key))
	}

	// Execute validator
	if opt.Validator != nil {
		return opt.Validator(opt.Value, opt.ValidatorParam)
	}
	return nil
}
