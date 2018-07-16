/**
 * NumericConstants
 * Page 16/17
 *
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 7:55 PM
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

const (
	// Create a huge number by shifting 1 bit left 100 places
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10+1
}
func needFloat(x float64) float64 {
	return x*0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}