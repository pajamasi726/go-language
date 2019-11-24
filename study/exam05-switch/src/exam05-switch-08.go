package main

func main()  {

	switch e:="go";e{
	case "java":
		println("java")
		fallthrough
	case "go":
		println("go")
		fallthrough
	case "c":
		println("c")
	case "c#":
		println("c#")
		fallthrough
	case "python":
		println("python")
	}
}