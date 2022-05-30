package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CheckParams failed tests suite
func TestCheckParamsFailedEmptyStruct(t *testing.T) {
	s := FizzbuzzStruct{}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: empty struct should be failed")
	}
	assert.Equal(t, err, ErrorInt1NotValidNumber, ErrorInt1NotValidNumber)
	assert.Equal(t, err.Error(), "int1 is not valid number", "int1 is not valid number")
}

func TestCheckParamsFailedInt1(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "",
		Int2: "1",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int1 is not valid number should be failed")
	}
	assert.Equal(t, err, ErrorInt1NotValidNumber, ErrorInt1NotValidNumber)
	assert.Equal(t, err.Error(), "int1 is not valid number", "int1 is not valid number")
}

func TestCheckParamsFailedZeroValueInt1(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "0",
		Int2: "1",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int1 can't be equal to 0 should be failed")
	}
	assert.Equal(t, err, ErrorInt1CanNotBeEqualToZero, ErrorInt1CanNotBeEqualToZero)
	assert.Equal(t, err.Error(), "int1 can't be equal to 0", "int1 can't be equal to 0")
}

func TestCheckParamsFailedStringValueInt1(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "abc",
		Int2: "1",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int1 is not valid number should be failed")
	}
	assert.Equal(t, err, ErrorInt1NotValidNumber, ErrorInt1NotValidNumber)
	assert.Equal(t, err.Error(), "int1 is not valid number", "int1 is not valid number")
}

func TestCheckParamsFailedInt2(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "1",
		Int2: "",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int2 is not valid number should be failed")
	}
	assert.Equal(t, err, ErrorInt2NotValidNumber, ErrorInt2NotValidNumber)
	assert.Equal(t, err.Error(), "int2 is not valid number", "int2 is not valid number")
}

func TestCheckParamsFailedZeroValueInt2(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "1",
		Int2: "0",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int2 can't be equal to 0 should be failed")
	}
	assert.Equal(t, err, ErrorInt2CanNotBeEqualToZero, ErrorInt2CanNotBeEqualToZero)
	assert.Equal(t, err.Error(), "int2 can't be equal to 0", "int2 can't be equal to 0")
}

func TestCheckParamsFailedStringValueInt2(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "1",
		Int2: "abc",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: int2 is not valid number should be failed")
	}
	assert.Equal(t, err, ErrorInt2NotValidNumber, ErrorInt2NotValidNumber)
	assert.Equal(t, err.Error(), "int2 is not valid number", "int2 is not valid number")
}

func TestCheckParamsFailedStr1(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "1",
		Int2: "1",
		Str1: "",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: str1 can't be empty should be failed")
	}
	assert.Equal(t, err, ErrorStr1CanNotBeEmpty, ErrorStr1CanNotBeEmpty)
	assert.Equal(t, err.Error(), "str1 can't be empty", "str1 can't be empty")
}

func TestCheckParamsFailedStr2(t *testing.T) {
	s := FizzbuzzStruct{
		Int1: "1",
		Int2: "1",
		Str1: "fizz",
		Str2: "",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: str2 can't be empty should be failed")
	}
	assert.Equal(t, err, ErrorStr2CanNotBeEmpty, ErrorStr2CanNotBeEmpty)
	assert.Equal(t, err.Error(), "str2 can't be empty", "str2 can't be empty")
}

func TestCheckParamsFailedLimitNegativeValue(t *testing.T) {
	s := FizzbuzzStruct{
		Int1:  "1",
		Int2:  "1",
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: "-1",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: limit can't be negative should be failed")
	}
	assert.Equal(t, err, ErrorLimitNegativeValue, ErrorLimitNegativeValue)
	assert.Equal(t, err.Error(), "limit can't be negative", "limit can't be negative")
}

func TestCheckParamsFailedLimitNotValidNumber(t *testing.T) {
	s := FizzbuzzStruct{
		Int1:  "1",
		Int2:  "1",
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: "abc",
	}
	_, err := CheckParams(s)
	if err == nil {
		t.Error("CheckParams: limit is not valid number")
	}
	assert.Equal(t, err, ErrorLimitNotValidNumber, ErrorLimitNotValidNumber)
	assert.Equal(t, err.Error(), "limit is not valid number", "limit is not valid number")
}

// CheckParams success tests suite
func TestCheckParamsSuccess(t *testing.T) {
	s := FizzbuzzStruct{
		Int1:  "1",
		Int2:  "1",
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: "100",
	}
	expected := &fizzbuzzStruct{
		Int1:  1,
		Int2:  1,
		Str1:  "fizz",
		Str2:  "buzz",
		Limit: 100,
	}
	ret, err := CheckParams(s)
	if err != nil {
		t.Error("CheckParams: should be success")
	}
	assert.Equal(t, expected, ret, "the return value should be equal to expected")
}
