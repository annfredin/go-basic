package loop

import (
	"fmt"
)

func init() {
	fmt.Println("==== START LOOP ===== ")
	start()
	fmt.Println("==== END LOOP ===== ")
	fmt.Println()
}

func start(){
	forLoop()
	forRangeLoop()
}


func forLoop(){
	for i := 0; i <= 5; i++ {
		if i%3 == 0 {
            continue
        }
        fmt.Printf("%v ", i)
    }

	//while fashion
	i := 0
	fmt.Println()
    for i <= 5 {
        fmt.Printf("%v ", i)

		
		
		if i >= 3 {
            break //break statement
        }
        i++
    }
}

func forRangeLoop(){
	/*	range loop =====>
		1. array or slice
		2. string
		3. maps
		4. channel 
	*/
	fmt.Println()

	letters := []string{"a", "b", "c"}
	for i, letter := range letters {
        fmt.Printf("Index: %d Value:%s\n", i, letter)
    }
	fmt.Println()
	}