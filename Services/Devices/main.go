package main

import mqtt "github.com/eclipse/paho.mqtt.golang"

var client mqtt.Client

func main() {

	//db, _ := dbConnect()
	// checkError(err)
	// fmt.Println("Database successfully connected")
	// defer db.Close()

	a := App{}
	a.Init(
		"postgres",
		"123456",
		"testdb")
	client = createClient()
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	a.Run(":8080")
}
