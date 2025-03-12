package main

import (
	"fmt"

	"github.com/MarianCostea/LearnGo/blob/main/exportValues/file2"
)

var ExportedValue = "Hello World"

func main() {
	fmt.Println(ExportedValue)
	fmt.Println(file2.AnotherExportedValue)
}
