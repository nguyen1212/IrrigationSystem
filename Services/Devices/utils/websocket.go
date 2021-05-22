package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

var (
	messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}

	connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	connectLostHandler mqtt.ConnectionLostHandler = func(clietn mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Message struct {
	UserID string `json:"user"`
	PlotID string `json:"plot"`
	State  string `json:"state"`
}

type DeviceMsg struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Data int    `json:"data"`
	Unit string `json:"unit"`
}

func handleWSConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	log.Println("Websocket formed!")

	//Create a new connection to Adafruit Server
	client := createClient()

	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go read(ws, client)
	write(ws, client)
}

func read(ws *websocket.Conn, client mqtt.Client) {
	for {
		var msg Message
		var adaMsg *DeviceMsg
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("error: %v", err)
			continue
		}
		if msg.State == "ON" {
			adaMsg = &DeviceMsg{
				Id:   1,
				Name: "Demo",
				Data: 0,
				Unit: "",
			}
		} else if msg.State == "OFF" {
			adaMsg = &DeviceMsg{
				Id:   1,
				Name: "Demo",
				Data: 1,
				Unit: "",
			}
		}
		publish(client, adaMsg)
	}
}

func write(ws *websocket.Conn, client mqtt.Client) {
	// for {
	// 	subSensor(client)
	// 	err := ws.WriteJSON()
	// 	if err != nil {
	// 		log.Printf("error: %v", err)
	// 	}
	// }
}

//Create a new client that connect to Adafruit via MQTT (Port 1883)
func createClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://io.adafruit.com:1883")
	opts.SetClientID("Demo")
	// Username of Adafruit account -> For demo only
	opts.SetUsername("MDPSmartFarm")
	// Key of Adafruit account: Get from MyKey tab -> For demo only
	opts.SetPassword("aio_TaBK62eP4aIDchqWz7oGt39utSVT")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return mqtt.NewClient(opts)
}

func subSensor(client mqtt.Client) {
	topic := make(map[string]byte)
	topic["MDPSmartFarm/feeds/soilmoist"] = 2
	topic["MDPSmartFarm/feeds/temphumid"] = 2
	token := client.SubscribeMultiple(topic, HandleReceivedMsg)
	token.Wait()
	for topicName := range topic {
		fmt.Println("Subscribed to topic %s", topicName)
	}
}

func HandleReceivedMsg(client mqtt.Client, msg mqtt.Message) {
	var sensorMsg DeviceMsg
	json.Unmarshal(msg.Payload(), &sensorMsg)
}

func publish(client mqtt.Client, msg *DeviceMsg) {
	topic := "MDPSmartFarm/feeds/pumprelay"
	data, _ := json.Marshal(msg)
	token := client.Publish(topic, 1, true, data)
	token.Wait()
}
