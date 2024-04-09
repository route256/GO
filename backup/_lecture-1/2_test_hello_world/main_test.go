package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	require.Equal(t, hello(), "Hello, 世界")
}
