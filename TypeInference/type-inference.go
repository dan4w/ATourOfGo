/**
 * TypeInference
 * Page 14/17
 * When declaring a variable without specifying the type, the type is inferred by what is on the right hand side.
 * If the right hand side contains a numeric constant, the type can be either int, float64 or complex128
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 6:16 PM
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

func main(){
	v:= 42 // change me
	fmt.Printf("v is a type %T\n", v)
}