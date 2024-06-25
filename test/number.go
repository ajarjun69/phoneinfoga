package test

import "github.com/sundowndev/phoneinfoga/v2/lib/number"

func NewFakeUSNumber() *number.Number {
	n, _ := number.NewNumber("+91.8848978073")
	return n
}
