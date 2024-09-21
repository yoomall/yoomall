package result_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"lazyfury.github.com/yoomall-server/core/result"
)

func TestResult(t *testing.T) {
	assert := assert.New(t)

	// ok must no err
	val := result.Ok(1)
	assert.Equal(1, val.Value)
	assert.Nil(val.Error)

	// Err str
	val2 := result.Err[any](errors.New("err"))
	assert.Equal("err", val2.Error.Error())

	//Err must err
	val3 := result.Err[any](nil)
	assert.Equal(true, val3.IsErr())
}
