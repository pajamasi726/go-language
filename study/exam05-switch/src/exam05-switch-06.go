package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)

	switch {
	case value == 0:
		println("0")
	case value > 0 :
		println("0보다 큼")
	case value < 0:
		println("0보다 작음")
	}

	fmt.Println(value)
}
