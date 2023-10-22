package packageone

import "fmt"

var PackageVar = "Some_PackageVar_value"

func PrintMe(varOne, varTwo, varThree string) {
	fmt.Println(varOne, varTwo, varThree)
}
