package goquestion

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestIota(t *testing.T) {

	fmt.Println(second)
	// 1

	convey.Convey("iota", t, func() {
		convey.So(second, convey.ShouldEqual, 1)
	})
}

// below references: go by exam
// link:<https://golangbyexample.com/iota-in-golang/>

// There will be no increment if there is a empty line or a commented line
// 空行或者注释,不会自增
func TestIota2(t *testing.T) {
	const (
		a = iota

		b
		//comment
		c
	)
	fmt.Println(a, b, c)
	// 0 1 2
}

// Iota value will reset and again start with zero if the const keyword is used again
// 出现的const,iota值被重置为0
func TestIota3(t *testing.T) {
	const (
		a = iota
		b
	)
	const (
		c = iota
	)
	fmt.Println(a, b, c)
	// 0 1 0
}

// iota increment can be skipped using a blank identifier
// 使用 _ 匿名标识可以跳过某个自增值
func TestIota4(t *testing.T) {
	const (
		a = iota
		_
		b
		c
	)
	fmt.Println(a, b, c)
	// 0 2 3
}

// iota expressions – iota allows expressions which can be used to set any value for the constant
// 支持运算符
func TestIota5(t *testing.T) {
	const (
		a = iota
		b = iota + 4
		c = iota * 4
	)
	fmt.Println(a, b, c)
	// 0 5 8
	/*
		The first-time iota value was zero, hence the output is zero.
		On the next line iota value is 1 hence the output is 1+4=5.
		On the next line, iota value is 2 hence output 2*4=8
	*/
}

// iota can also start from non-zero number- iota expressions can also be used to start iota from any number
// 支持从任意值开始
func TestIota6(t *testing.T) {
	const (
		a = iota + 10
		b
		c
	)
	fmt.Println(a, b, c)
	// 10 11 12
}

// IOTA provides an automated way to create a enum in Golang.
// 支持枚举
func TestIota7(t *testing.T) {
	type Size uint8
	const (
		small Size = iota
		medium
		large
		extraLarge
	)
	fmt.Println(small, medium, large, extraLarge)
	// 0 1 2 3
}
