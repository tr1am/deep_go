package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	pointer := unsafe.Pointer(&number)
	result := new(uint32)
	lt := int(unsafe.Sizeof(number))
	
	for i := 0; i < lt; i++ {
		byteValue := *(*uint8)(unsafe.Add(pointer, i))
		*(*uint8)(unsafe.Add((unsafe.Pointer(result)), lt-i-1)) = byteValue
	}
	return *result
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
