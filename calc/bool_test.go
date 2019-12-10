package calc

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unary_Bool(t *testing.T) {
	assert.True(t, BoolSolve("!0"), "Bool operation fails")
	assert.False(t, BoolSolve("!1"), "Bool operation fails")
	assert.False(t, BoolSolve("!(0 | 1)"), "Bool operation fails")
	assert.True(t, BoolSolve("!(0 | 0)"), "Bool operation fails")
	assert.True(t, BoolSolve("(1+1+1+1 > 3)"), "Bool operation fails")
	assert.False(t, BoolSolve("!(1+1+1+1 > 3)"), "Bool operation fails")
	assert.True(t, BoolSolve("!0 & !0"), "Bool operation fails")
	assert.False(t, BoolSolve("!0 & !1"), "Bool operation fails")
	assert.False(t, BoolSolve("!0 & !1)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & !0)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1 | !(0 & 0 & 0)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1 | !(1 & 0 & 0)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1 | !(0 & 1 & 0)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1 | !(0 & 0 & 1)"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1 | !(!0 & !0 & 1)"), "Bool operation fails")
	assert.False(t, BoolSolve("1 & 1 | !(!0 & !0 & !1)"), "Bool operation fails")
}

func Test_And_Bool(t *testing.T) {
	assert.False(t, BoolSolve("0 & 0"), "Bool operation fails")
	assert.False(t, BoolSolve("1 & 0"), "Bool operation fails")
	assert.False(t, BoolSolve("0 & 1"), "Bool operation fails")
	assert.True(t, BoolSolve("1 & 1"), "Bool operation fails")
}

func Test_Or_Bool(t *testing.T) {
	assert.False(t, BoolSolve("0 | 0"), "Bool operation fails")
	assert.True(t, BoolSolve("1 | 0"), "Bool operation fails")
	assert.True(t, BoolSolve("0 | 1"), "Bool operation fails")
	assert.True(t, BoolSolve("1 | 1"), "Bool operation fails")
}

func Test_Parens_Bool(t *testing.T) {
	assert.False(t, BoolSolve("(0|0)&(1)"), "Bool operation fails")
	assert.True(t, BoolSolve("(1|1)&(1)"), "Bool operation fails")
	assert.True(t, BoolSolve("(1|1)&(1)"), "Bool operation fails")
}

func Test_Complex_Bool(t *testing.T) {
	assert.True(t, BoolSolve("1&0&(1&1&1)|1&1|1"), "Bool operation fails")
	assert.True(t, BoolSolve("1&0&(1&1&1)|(1&1|1)"), "Bool operation fails")
	assert.False(t, BoolSolve("1&0&(1&1&1)|1&1|1&0"), "Bool operation fails")
	assert.False(t, BoolSolve("1&0&(1&1&1)|(1&1|1)&(1&0)"), "Bool operation fails")
}

func Test_Threshold_Bool(t *testing.T) {
	assert.True(t, BoolSolve("1+1+1+1 > 3"), "Bool operation fails")
	assert.True(t, BoolSolve("0+1+1+1 > 2)"), "Bool operation fails")
	assert.True(t, BoolSolve("0 | 0+1+1+1 > 2)"), "Bool operation fails")
	assert.False(t, BoolSolve("0 | 0+1+1+0 > 2)"), "Bool operation fails")
	assert.False(t, BoolSolve("0+1+1+0 > 2 | 0+1+1+0 > 2"), "Bool operation fails")
	assert.True(t, BoolSolve("0+1+1+0 > 2 | 0+1+1+1 > 2"), "Bool operation fails")
	assert.True(t, BoolSolve("0+1+1+0 > 2 | 0+1+1+0 > 2 | 1"), "Bool operation fails")
	assert.True(t, BoolSolve("1 | 0+1+1+0 > 2 | 0+1+1+0 > 2"), "Bool operation fails")
}

func Test_Decode1_Bool(t *testing.T) {
	//A couple of tests for my specific use case
	keys := make(map[string]string)
	keys["P"] = "1"
	keys["C"] = "1"
	keys["B"] = "1"
	keys["T1"] = "1"
	keys["T2"] = "1"
	keys["T3"] = "1"
	keys["T4"] = "1"

	rule := "P & C & (T1 + T2 + T4 + T4 > 3)"
	testString := decode(keys, rule)
	assert.True(t, BoolSolve(testString), "Bool operation fails")
}
func Test_Decode2_Bool(t *testing.T) {
	//A couple of tests for my specific use case
	keys := make(map[string]string)
	keys["P"] = "1"
	keys["C"] = "1"
	keys["B"] = "1"
	keys["T1"] = "1"
	keys["T2"] = "1"
	keys["T3"] = "1"
	keys["T4"] = "0"

	rule := "P & C & (T1 + T2 + T3 + T4 > 3)"

	testString := decode(keys, rule)
	assert.False(t, BoolSolve(testString), "Bool operation fails")
}

func decode(keys map[string]string, rule string) string {
	for key, _ := range keys {
		rule = strings.ReplaceAll(rule, key, keys[key])
	}
	return rule
}
