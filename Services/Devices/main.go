package main

import mqtt "github.com/eclipse/paho.mqtt.golang"

var client1 mqtt.Client
var client2 mqtt.Client

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
	client1 = createClient()
	client2 = createClient2()
	token2 := client2.Connect()
	if token2.Wait() && token2.Error() != nil {
		panic(token2.Error())
	}
	token1 := client1.Connect()
	if token1.Wait() && token1.Error() != nil {
		panic(token1.Error())
	}
	go subSensor(client1)
	a.Run(":8080")
}
