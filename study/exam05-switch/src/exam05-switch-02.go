package main

func main()  {

	switch b:=28;{
	case b < 0 :
		println(b,"는 음수")
	case b > 0 :
		println(b,"는 양수")
	case b == 0 :
		println(b,"는 0")
	}
}