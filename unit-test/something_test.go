package somethingPackage

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func testSomething(t *testing.T) {
	v := SomethingMethod()

	assert.Equal(t, "somethingValue", v)
}