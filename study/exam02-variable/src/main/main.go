package main

import "fmt"

func main(){
	// go language 에서는 변수를 선언하고 사용하지 않으면 compile 시 에러가 난다
	var str = "hello world"
	fmt.Println(str)

	str = "안녕하세요 go language 공부 시작 합니다"
	fmt.Println("문자열출력 : ",str)
	// output : 문자열출력 :  안녕하세요 go language 공부 시작 합니다

	var b = true
	fmt.Println("boolean : ",b)

}