package shikihotest

import (
	"github.com/mattak/shikiho/pkg/shikiho"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type UnitTestContext struct{}

func (UnitTestContext) setup() {
	os.RemoveAll("/tmp/test")
}

func (UnitTestContext) tearDown() {
	os.RemoveAll("/tmp/test")
}

func TestParseInt(t *testing.T) {
	context := UnitTestContext{}
	context.setup()
	defer context.tearDown()

	assert.Equal(t, 1, shikiho.ParseInt("1"))
	assert.Equal(t, 1234, shikiho.ParseInt("1,234"))
	assert.Equal(t, 0, shikiho.ParseInt("―"))
	assert.Equal(t, 0, shikiho.ParseInt("‥"))
}

func TestParseFloat(t *testing.T) {
	context := UnitTestContext{}
	context.setup()
	defer context.tearDown()

	assert.Equal(t, 1.0, shikiho.ParseFloat("1"))
	assert.Equal(t, 1000.0, shikiho.ParseFloat("1,000"))
	assert.Equal(t, 1.0, shikiho.ParseFloat("1記"))
	assert.Equal(t, 1.0, shikiho.ParseFloat("1特"))
	assert.Equal(t, 1.0, shikiho.ParseFloat("1*"))
	assert.Equal(t, 2.0, shikiho.ParseFloat("2〜9"))
	assert.Equal(t, 0.0, shikiho.ParseFloat("‥"))
}

func TestParseJapanMoneyUnit(t *testing.T) {
	context := UnitTestContext{}
	context.setup()
	defer context.tearDown()

	assert.Equal(t, 0, shikiho.ParseJapanMoneyUnit("--"))
	assert.Equal(t, 1234, shikiho.ParseJapanMoneyUnit("1,234億円"))
}
