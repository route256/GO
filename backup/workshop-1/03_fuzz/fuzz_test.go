package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	assert.Equal(t, "123", reverseString("321"))
	assert.Equal(t, "", reverseString(""))
}

func Fuzz_ParseLine(f *testing.F) {
	f.Add("123")
	f.Add("test string")
	f.Add("")

	f.Fuzz(func(t *testing.T, data string) {
		reverseString(data)

		// assert.Equal(t,
		// 	utf8.ValidString(data),
		// 	utf8.ValidString(reverseString(data)),
		// )
	})
}
