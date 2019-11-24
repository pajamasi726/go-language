package main

func main()  {
	a := 10-17

	switch {
	case a < 0 :
		println(a,"는 음수")
	case a == 0 :
		println(a,"는 0")
	case a > 0 :
		println(a,"는 양수")
	}
}
