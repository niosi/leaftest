package main

import (
    "fmt"
    "sort"
)

func main() {
	testParam("H","E","LL","O","Hello")
}

func testParam(value...string){
	sort.Strings(value)
	fmt.Println(value)
	v := sort.SearchStrings(value,"9")
	fmt.Println(v)
}