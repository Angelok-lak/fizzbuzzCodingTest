package fizzbuzz

import (
	"strconv"
)

// FizzbuzzStruct
type FizzbuzzStruct struct {
	Int1  string `json:"int1" binding:"required"`
	Int2  string `json:"int2" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
	Limit string `json:"limit" binding:"required"`
}

// fizzbuzzStruct
type fizzbuzzStruct struct {
	Int1  int64
	Int2  int64
	Str1  string
	Str2  string
	Limit int64
}

func CheckParams(s FizzbuzzStruct) (*fizzbuzzStruct, error) {
	// check valid parameter
	fizzbuzzParams := new(fizzbuzzStruct)

	// check int1
	tmp, err := strconv.ParseInt(s.Int1, 10, 64)
	if err != nil {
		return nil, ErrorInt1NotValidNumber
	} else if tmp == 0 {
		return nil, ErrorInt1CanNotBeEqualToZero
	}
	fizzbuzzParams.Int1 = tmp

	// check int2
	tmp, err = strconv.ParseInt(s.Int2, 10, 64)
	if err != nil {
		return nil, ErrorInt2NotValidNumber
	} else if tmp == 0 {
		return nil, ErrorInt2CanNotBeEqualToZero
	}
	fizzbuzzParams.Int2 = tmp

	// check str1
	if s.Str1 == "" {
		return nil, ErrorStr1CanNotBeEmpty
	}
	fizzbuzzParams.Str1 = s.Str1

	// check str2
	if s.Str2 == "" {
		return nil, ErrorStr2CanNotBeEmpty
	}
	fizzbuzzParams.Str2 = s.Str2

	// check limit
	tmp, err = strconv.ParseInt(s.Limit, 10, 64)
	if err != nil {
		return nil, ErrorLimitNotValidNumber
	} else if tmp < 1 {
		return nil, ErrorLimitNegativeValue
	}
	fizzbuzzParams.Limit = tmp

	return fizzbuzzParams, nil
}

func Fizzbuzz(s *fizzbuzzStruct) []string {
	// initialise variables
	var i int64
	ret := make([]string, s.Limit)
	fizzbuzzIndex := s.Int1 * s.Int2
	fizzbuzzStr := s.Str1 + s.Str2

	// process values and assemble the string
	for i = 1; i <= s.Limit; i++ {
		switch {
		case i%fizzbuzzIndex == 0:
			ret[i-1] = fizzbuzzStr
		case i%s.Int1 == 0:
			ret[i-1] = s.Str1
		case i%s.Int2 == 0:
			ret[i-1] = s.Str2
		default:
			ret[i-1] = strconv.FormatInt(i, 10)
		}
	}
	return ret
}
