package main

import (
	"fmt"

	_ "module.com/basic/loop"
	_ "module.com/basic/variable"
)

func main(){
	defer test()
	defer test1()
	fmt.Println("Welcome to go-basic demo!")

}

func test() {
    fmt.Println("In Defer 1")
}

func test1() {
    fmt.Println("In Defer 2")
}