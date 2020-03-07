package main

/*
 * Module Dependencies
 */
import (
	"fmt"

	"github.com/mozzzzy/config/json/config"
	"github.com/mozzzzy/config/json/configOption"
	"github.com/mozzzzy/config/validator"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func main() {
	var conf config.Config
	if err := conf.AddOptions([]configOption.Option{
		{
			Key:         "string",
			ValueType:   "string",
			Description: "some string.",
			Required:    true,
		},
		{
			Key:            "int",
			ValueType:      "int",
			Description:    "some int.",
			Required:       true,
			Validator:      validator.IntBiggerThan,
			ValidatorParam: 1,
		},
		{
			Key:          "array",
			ValueType:    "array",
			Description:  "some array.",
			DefaultValue: []interface{}{1, 2, 3},
		},
		{
			Key:         "object",
			ValueType:   "object",
			Description: "some object.",
			Required:    true,
		},
		{
			Key:         "object.string",
			ValueType:   "string",
			Description: "some string.",
		},
		{
			Key:          "object.int",
			ValueType:    "int",
			Description:  "some integer.",
			DefaultValue: 1,
		},
		{
			Key:          "object.array",
			ValueType:    "array",
			Description:  "some array.",
			DefaultValue: []interface{}{1, 2, 3},
		},
	}); err != nil {
		fmt.Println(err.Error())
		return
	}

	configPath := "./config.json"
	if err := conf.Parse(configPath); err != nil {
		fmt.Printf("Failed to parse %v. %v\n", configPath, err)
		fmt.Printf("Following options are avairable.\n\n%v\n", conf)
		return
	}

	if integer, err := conf.GetInt("int"); err != nil {
		fmt.Println(err)
		fmt.Println(conf)
	} else {
		fmt.Println(integer)
	}

	if str, err := conf.GetString("string"); err != nil {
		fmt.Println(err)
		fmt.Println(conf)
	} else {
		fmt.Println(str)
	}

	if intArray, err := conf.GetIntArray("array"); err != nil {
		fmt.Println(err)
		fmt.Println(conf)
	} else {
		fmt.Println(intArray)
	}

	fmt.Println(conf)
}
