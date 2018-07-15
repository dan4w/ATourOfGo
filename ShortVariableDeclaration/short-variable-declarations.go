/**
 * ShortVariableDeclaration
 * Page 10/17
 * Using := to delare means we dont need to type var x =
 * := must be used inside a function, can't be used at the package level
 *
 * @Author: Daniel Forwood
 * @Date:   2018-07-14 5:28 PM ----2018-07-14T02:59:28-07:00
 * @Last modified by:   Daniel Forwood
 * @Last modified time: 2018-07-14T16:10:24-07:00
 */

package main

import "fmt"

func main() {
	var i, j int = 1,2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
