/**
 * TypeConversion
 * Page 13/17
 * Type conversions.
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 6:09 PM
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import (
	"math"
	"fmt"
)

func main(){
	var x , y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,f,z)
}
