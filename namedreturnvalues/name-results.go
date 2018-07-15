/**
 * namedreturnvalues
 * Page 7/17
 * Go's return values may be name. If so, they are treated as variables defined at the top of the functions
 * Prof says this is bad and messy. This is called "naked" returns.
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 5:13 PM ----2018-07-14T02:59:28-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

func split(sum int) (x,y int){
	x = sum*4/9
	y= sum -x
	return
}

func main() {
	fmt.Println(split(17))
}