package main

func main()  {

	// shot 선언
	shortVar := 3
	println(shortVar)

	shortVar = 10
	println(shortVar)

	a:=2
	b:=21
	println(b%a)
	// if 안에서 할당이 가능
	if c:=b%a; c==0  {
		println("짝수")
	}
}
