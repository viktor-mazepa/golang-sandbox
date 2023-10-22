package main

import "go-scope-project/packageone"

var myVar = "Some_myVar_value"

func main() {
	var blockVar = "Some_blockVar_value"
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}
