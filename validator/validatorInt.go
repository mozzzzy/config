package validator

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

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */
func IntBiggerThan(val, min interface{}) error {
	intVal, intValOk := val.(int)
	if !intValOk {
		return errors.New(fmt.Sprintf("Specified value %v is not int type.", val))
	}
	intMin, intMinOk := min.(int)
	if !intMinOk {
		return errors.New(fmt.Sprintf("Specified mininum %v is not int type.", min))
	}
	if intVal < intMin {
		return errors.New(fmt.Sprintf("%v < %v.", intVal, intMin))
	}
	return nil
}

func IntSmallerThan(val, max interface{}) error {
	intVal, intValOk := val.(int)
	if !intValOk {
		return errors.New(fmt.Sprintf("Specified value %v is not int type.", val))
	}
	intMax, intMaxOk := max.(int)
	if !intMaxOk {
		return errors.New(fmt.Sprintf("Specified max %v is not int type.", max))
	}
	if intVal > intMax {
		return errors.New(fmt.Sprintf("%v > %v.", intVal, intMax))
	}
	return nil
}

func IntWithin(val interface{}, minMax interface{}) error {
	intVal, intValOk := val.(int)
	if !intValOk {
		return errors.New(fmt.Sprintf("Specified value %v is not int type.", val))
	}
	intAryMinMax, intAryMinMaxOk := minMax.([]int)
	if !intAryMinMaxOk {
		return errors.New(
			fmt.Sprintf("Specified value %v should be an array of min and max.", minMax))
	}
	if len(intAryMinMax) < 2 {
		return errors.New(fmt.Sprintf("Can't get min or max value from %v.", intAryMinMax))
	}
	if BiggerErr := IntBiggerThan(intVal, intAryMinMax[0]); BiggerErr != nil {
		return BiggerErr
	}
	if SmallerErr := IntSmallerThan(intVal, intAryMinMax[1]); SmallerErr != nil {
		return SmallerErr
	}
	return nil
}
