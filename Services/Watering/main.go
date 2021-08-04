package main

func main() {
	a := App{}
	a.Init("postgres",
		"123456",
		"testdb")
	a.Run(":8010")

}
