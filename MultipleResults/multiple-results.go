/**
 * MultipleResults
 * Page 6/17
 * A function can return any number of results.
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 5:00 PM ----2018-07-14T02:59:28-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T17:02:08-07:00
 */

package main

import "fmt"

func swap(x, y string) (string, string){
	return y, x
}

func main(){
	a, b := swap("hello", "world")
	fmt.Println(a,b)
}


