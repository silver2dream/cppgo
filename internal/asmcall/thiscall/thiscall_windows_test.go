package thiscall

import (
	"testing"

	"github.com/lsegal/cppgo/internal/asmcall/asmcalltest"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	v, e := Call(asmcalltest.GetStdcallF0Addr(), 0)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(42), v)

	v, e = Call(asmcalltest.GetStdcallF1Addr(), 0, 16)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(17), v)

	v, e = Call(asmcalltest.GetStdcallF2Addr(), 0, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(2), v)

	v, e = Call(asmcalltest.GetStdcallF3Addr(), 0, 4, 2, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetStdcallF4Addr(), 0, 16, 2, 4, 2)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(1), v)

	v, e = Call(asmcalltest.GetStdcallF5Addr(), 0, 99, 99, 99, 99, 12)
	assert.NoError(t, e)
	assert.Equal(t, uintptr(12), v)
}
