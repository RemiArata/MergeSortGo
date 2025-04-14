package main

import (
	"MergeSortGo/helpers"
	"MergeSortGo/ms"
	"fmt"
)

func main() {
	fmt.Println("hello world")

	data := helpers.GenerateData(true, 10)
	newData := ms.MergeSort(data)
	fmt.Println(newData)
}
