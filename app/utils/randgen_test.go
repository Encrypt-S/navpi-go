package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// General us
func Test_GenerateRandomBytes(t *testing.T) {

	l := 10
	b1, _ := GenerateRandomBytes(l)
	b2, _ := GenerateRandomBytes(l)

	assert.NotEqual(t, b1, b2)
	assert.Equal(t, l, len(b1))
	assert.Equal(t, l, len(b2))

}

// test the byte length generation up to 5000 bytes
func Test_GenerateRandomBytes_len(t *testing.T) {

	stopTestLen := 5000

	b1, _ := GenerateRandomBytes(1)

	for i := 1; i <= stopTestLen; i++ {

		b1, _ = GenerateRandomBytes(i)
		assert.Equal(t, i, len(b1))

	}

}

// General us
func Test_GenerateRandomString(t *testing.T) {

	l := 10
	s1, _ := GenerateRandomString(l)
	s2, _ := GenerateRandomString(l)

	assert.NotEqual(t, s1, s2)

}
