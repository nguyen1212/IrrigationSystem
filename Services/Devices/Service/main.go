package main

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

	subChannel = make(chan DeviceMsg)
)

type Message struct {
	UserID string `json:"userId"`
	PlotID string `json:"plotId"`
	Name   string `json:"name"`
	Data   string `json:"data"`
}

type DeviceMsg struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
	Unit string `json:"unit"`
}

func handleWSConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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
	go subSensor(client)
	write(ws)
}

func read(ws *websocket.Conn, client mqtt.Client) {
	for {
		var msg Message
		var adaMsg *DeviceMsg
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if msg.Data == "1" {
			adaMsg = &DeviceMsg{
				Id:   "11",
				Name: "RELAY",
				Data: "0",
				Unit: "",
			}
		} else if msg.Data == "0" {
			adaMsg = &DeviceMsg{
				Id:   "11",
				Name: "RELAY",
				Data: "1",
				Unit: "",
			}
		}
		publish(client, adaMsg)
	}
}

func write(ws *websocket.Conn) {
	// adaMsg := <-subChannel
	for {
		adaMsg := <-subChannel
		sendingMsg := &Message{
			UserID: "0", //handle latter
			PlotID: "1", //handle latter
			Name:   adaMsg.Name,
			Data:   fmt.Sprint(adaMsg.Data),
		}
		err := ws.WriteJSON(sendingMsg)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

//Create a new client that connect to Adafruit via MQTT (Port 1883)
func createClient() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://io.adafruit.com:1883")
	opts.SetClientID("Demo")
	// Username of Adafruit account -> For demo only
	opts.SetUsername("MDPSmartFarm")
	// Key of Adafruit account: Get from MyKey tab -> For demo only
	opts.SetPassword("aio_hTES03bizhmCK7SiCw8A071FH1j9")
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
	subChannel <- sensorMsg

}

func publish(client mqtt.Client, msg *DeviceMsg) {
	topic := "MDPSmartFarm/feeds/pumprelay"
	data, _ := json.Marshal(msg)
	token := client.Publish(topic, 1, true, data)
	token.Wait()
}

func main() {
	// Create new Websocket to client
	http.HandleFunc("/devices/ws", handleWSConnections)

	fmt.Println("Server listening on port 8080")
	for {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
