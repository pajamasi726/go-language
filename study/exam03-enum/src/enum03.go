package main

func main()  {

	const (
		_ = iota + 1
		A
		B
		_
		C
	)

	println("A : ",A, " B : ",B, " C : ",C)
}
