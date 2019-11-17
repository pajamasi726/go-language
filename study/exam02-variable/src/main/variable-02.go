package main

import "fmt"

func main()  {

	// var ( )  scope 로 변수를 선언 할 수 있다.
	var(
		name string = "홍길동"
		age int32
		isMan bool = true
	)

	age = 34
	fmt.Println("name : ",name," age : ",age," isMan : ",isMan)


	const PROFILE = "Student"
	const (
		PATH = "/v1/user"
	)


}
