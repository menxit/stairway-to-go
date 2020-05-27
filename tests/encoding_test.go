package tests

import (
	"testing"

	"github.com/menxit/stairway-to-go/encoding"
	"github.com/stretchr/testify/assert"
)

func Test_encoding_DaASCIIABase64(t *testing.T) {
	msg := "hello, world"
	dovreiEssereHelloWorldInBase64 := encoding.DaASCIIABase64(msg)
	// don't trust, verify https://www.base64encode.org/
	assert.Equal(t, dovreiEssereHelloWorldInBase64, "aGVsbG8sIHdvcmxk")
}
