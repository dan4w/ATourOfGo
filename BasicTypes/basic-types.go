/**
 * BasicTypes
 * Page 11/17
 * Basic types in Go
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 5:57 PM
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z,z)
}
