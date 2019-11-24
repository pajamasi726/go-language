package main

func main()  {

	switch i,j := 20,30;{
	case i < j:
		println("i < j")
	case i == j:
		println("i == j")
	case i > j:
		println("i가 j보다 큼")
	}
}
