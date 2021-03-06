package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	"encoding/json"
	"net/http"
	_ "strconv"

	"SmartFarm/Services/Devices/data"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
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

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/devices/ws", handleWSConnections).Methods("GET")
	a.Router.
		HandleFunc("/devices/data/log", data.HandlePlotDeviceData).
		Methods("POST", "OPTIONS")
	a.Router.
		HandleFunc("/devices/water/log", data.HandlePlotWater).
		Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/devices/water/on", HandleTurnOn).Methods("GET")
	a.Router.HandleFunc("/devices/water/off", HandleTurnOff).Methods("GET")
}

func HandleTurnOn(w http.ResponseWriter, r *http.Request) {
	println("Turn Relay On")
	adaMsg := &DeviceMsg{
		Id:   "11",
		Name: "RELAY",
		Data: "1",
		Unit: "",
	}
	publish(client2, adaMsg)
}

func HandleTurnOff(w http.ResponseWriter, r *http.Request) {
	println("Turn Relay Off")
	adaMsg := &DeviceMsg{
		Id:   "11",
		Name: "RELAY",
		Data: "0",
		Unit: "",
	}
	publish(client2, adaMsg)
}

func (a *App) Init(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

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

	log.Println("Websocket formed!")
	defer ws.Close()

	go read(ws, client1)
	write(ws)

}

func read(ws *websocket.Conn, client mqtt.Client) {
	for {
		var msg Message
		var adaMsg *DeviceMsg
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			ws.Close()
			client1.Disconnect(250)
			client2.Disconnect(250)
			log.Printf("%v", err)
			return
		}
		if msg.Data == "1" {
			log.Println("Force Turn Off")
			adaMsg = &DeviceMsg{
				Id:   "11",
				Name: "RELAY",
				Data: "0",
				Unit: "",
			}
		} else if msg.Data == "0" {
			log.Println("Force Turn On")
			adaMsg = &DeviceMsg{
				Id:   "11",
				Name: "RELAY",
				Data: "1",
				Unit: "",
			}
		}
		publish(client2, adaMsg)
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
	opts.SetUsername("CSE_BBC")
	// Key of Adafruit account: Get from MyKey tab -> For demo only
	opts.SetPassword("")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return mqtt.NewClient(opts)
}

func createClient2() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://io.adafruit.com:1883")
	opts.SetClientID("Demo2")
	// Username of Adafruit account -> For demo only
	//opts.SetUsername("CSE_BBC1")
	opts.SetUsername("CSE_BBC1")
	// Key of Adafruit account: Get from MyKey tab -> For demo only
	opts.SetPassword("")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	println("Client 2 ", opts.Username)
	return mqtt.NewClient(opts)
}

func subSensor(client mqtt.Client) {
	topic := make(map[string]byte)
	topic["CSE_BBC/feeds/bk-iot-soil"] = 2
	topic["CSE_BBC/feeds/bk-iot-temp-humid"] = 2
	client.SubscribeMultiple(topic, HandleReceivedMsg)
	//token.Wait()
	for topicName := range topic {
		fmt.Printf("Subscribed to topic %s", topicName)
	}
}

func HandleReceivedMsg(client mqtt.Client, msg mqtt.Message) {
	var sensorMsg DeviceMsg
	println("New Data Received")
	if err := json.Unmarshal(msg.Payload(), &sensorMsg); err != nil {
		log.Panic(err)
	}
	var url string
	if sensorMsg.Name == "SOIL" {
		url = "http://127.0.0.1:8010/api/soil"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg.Payload()))
		req.Header.Set("Content-Type", "application/json")
		checkError(err)
		sender := &http.Client{}
		response, err := sender.Do(req)
		checkError(err)

		defer response.Body.Close()
	}
	subChannel <- sensorMsg
}

func publish(client mqtt.Client, msg *DeviceMsg) {
	topic := "CSE_BBC1/feeds/bk-iot-relay"
	data, _ := json.Marshal(msg)
	token := client.Publish(topic, 1, true, data)
	token.Wait()
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
