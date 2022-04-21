package variable

import (
	"fmt"
)

func init() {
	fmt.Println("==== START VARIABLE ===== ")
	startVariable()
	fmt.Println("==== END VARIABLE ===== ")
	fmt.Println()
}

func startVariable(){
	v1()
	v2()
}


func v1(){
	//all the variable will have default value assigned internally in go.
	var a int //single dec
	var x, y int = 8, 9  //multi dec
	var (
		z bool
		s string = "TEST"
	)

	//short declaration
	name,year := "Fredin", 2022

	fmt.Println("V1 -> INT Val ", a, x, y,z,s,name,year)
}

func v2(){
    fmt.Println("====>")

	var t = 123      //Type Inferred will be int
    var u = "circle" //Type Inferred will be string
    var v = 5.6      //Type Inferred will be float64
    var w = true     //Type Inferred will be bool
    var x = 'a'      //Type Inferred will be rune
    var y = 3 + 5i   //Type Inferred will be complex128

    fmt.Printf("Type: %T Value: %v\n", t, t)
    fmt.Printf("Type: %T Value: %v\n", u, u)
    fmt.Printf("Type: %T Value: %v\n", v, v)
    fmt.Printf("Type: %T Value: %v\n", w, w)
    fmt.Printf("Type: %T Value: %v\n", x, x)
    fmt.Printf("Type: %T Value: %v\n", y, y)
    fmt.Println()
}

