package fizzbuzz

import "errors"

var (
	ErrorInt1NotValidNumber      = errors.New("int1 is not valid number")
	ErrorInt2NotValidNumber      = errors.New("int2 is not valid number")
	ErrorInt1CanNotBeEqualToZero = errors.New("int1 can't be equal to 0")
	ErrorInt2CanNotBeEqualToZero = errors.New("int2 can't be equal to 0")
	ErrorLimitNotValidNumber     = errors.New("limit is not valid number")
	ErrorLimitNegativeValue      = errors.New("limit can't be negative")
	ErrorStr1CanNotBeEmpty       = errors.New("str1 can't be empty")
	ErrorStr2CanNotBeEmpty       = errors.New("str2 can't be empty")
)
