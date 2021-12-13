// ============================================================================
// = main.go																  =
// = 	Description: used to run the goconvert package functions			  =
// = 	Date: December 12, 2021												  =
// ============================================================================

package main

import (
	"fmt"
	n2w "goconvert/num2words"
)

func main() {
	fmt.Println(n2w.Num2Words(n2w.Pow(n2w.New(10), n2w.New(65)), true))
}