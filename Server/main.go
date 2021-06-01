package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var testClients = make(map[string]bool)
var upgrader = websocket.Upgrader{}

// var broadcast = make(chan forceMode)

func main() {
	http.HandleFunc("/devices/ws", handleConnection)
	http.HandleFunc("/api/thumbnail", thumbnailHandler)
	http.HandleFunc("/api/homepage", homePageHandler)

	fs := http.FileServer(http.Dir("../Frontend/dist"))
	http.Handle("/", fs)

	// go sendInfo()

	fmt.Println("Server listening on port 8081")
	log.Panic(
		http.ListenAndServe(":8081", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

type thumbnailRequest struct {
	Url string `json:"url"`
}
type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

type forceMode struct {
	UserId string `json:"userId"`
	PlotId string `json:"plotId"`
	State  string `json:"state"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	// Try to decode the request into the thumbnailRequest struct.
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a struct with the parameters needed to call the ScreenshotAPI.
	apiRequest := screenshotAPIRequest{
		Token:          "QPGG181-8DM429R-KQB7PGD-1ZXEFF2",
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	// Convert the struct to a JSON string.
	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	// Create a HTTP request.
	req, err := http.NewRequest("POST", "https://shot.screenshotapi.net/screenshot", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request.
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	// Tell Go to close the response at the end of the function.
	defer response.Body.Close()

	// Read the raw response into a Go struct.
	type screenshotAPIResponse struct {
		Screenshot string `json:"screenshot"`
	}
	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	// Pass back the screenshot URL to the frontend.
	_, err = fmt.Fprintf(w, `{ "screenShotUrl": "%s" }`, apiResponse.Screenshot)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil) // upgrade to websocket connection
	checkError(err)

	defer ws.Close()

	clients[ws] = true
	for {
		var msg forceMode
		err := ws.ReadJSON(&msg) // decode json to msg variable
		checkError(err)
		fmt.Println(msg.PlotId)
	}
}
