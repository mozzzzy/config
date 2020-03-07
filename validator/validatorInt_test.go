package validator

/*
 * Module Dependencies
 */

import (
	"testing"
	"github.com/mozzzzy/testUtil"
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

func TestIntBiggerThan(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		err := IntBiggerThan(10, 10)
		testUtil.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := IntBiggerThan(10, 11)
		testUtil.WithError(t, err)
	})
}

func TestIntSmallerThan(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		err := IntSmallerThan(10, 10)
		testUtil.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := IntSmallerThan(10, 9)
		testUtil.WithError(t, err)
	})
}

func TestIntWithin(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		err := IntWithin(10, []int{9, 10})
		testUtil.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		err := IntWithin(10, []int{8, 9})
		testUtil.WithError(t, err)
	})
}
