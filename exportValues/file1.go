package main

import (
	"fmt"

	"./file2"
)

var ExportedValue = "Hello World"

func main() {
	fmt.Println(ExportedValue)
	fmt.Println(file2.AnotherExportedValue)
}
