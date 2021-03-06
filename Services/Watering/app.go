package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var (
	Auto = AutoPreset{
		Id:       1,
		Name:     "",
		Moisture: 30,
		Note:     "",
	}
	Repeat   = RepeatPreset{}
	Interval = IntervalPreset{}
	Mode     = "AUTO"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init(user, password, dbname string) {
	connectString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.InitRoutes()

}

func (a *App) InitRoutes() {
	a.Router.HandleFunc("/api/soil", HandleSoil).Methods("POST")
	//a.Router.HandleFunc("/api/temp-hum", HandleTemp)
	a.Router.HandleFunc("/api/preset/auto", HandleAutoPreset).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/api/preset/repeat", HandleRepeat)
	a.Router.HandleFunc("/api/preset/interval", HandleInterval)
}

func (a *App) Run(address string) {
	println("Running on port 8010")
	log.Fatal(http.ListenAndServe(address, a.Router))
}

type SensorMsg struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
	Unit string `json:"unit"`
}

func HandleAutoPreset(w http.ResponseWriter, r *http.Request) {
	var auto AutoPreset
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode("OK")
	if err := json.NewDecoder(r.Body).Decode(&auto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Auto = auto
	Mode = "AUTO"
}

func HandleInterval(w http.ResponseWriter, r *http.Request) {
	var interval IntervalPreset
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode("OK")
	if err := json.NewDecoder(r.Body).Decode(&interval); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Interval = interval
	Mode = "INTERVAL"
	str := strings.Split(Interval.Interval, "")
	data, _ := strconv.Atoi(str[0])
	for Mode == "INTERVAL" {
		IntervalMode(strconv.Atoi(Interval.Amount))

		timer1 := time.NewTimer(time.Duration(data*60) * time.Second)
		<-timer1.C
		println("Time to water!")
	}
}

func IntervalMode(amount int, _ error) {
	url := "http://127.0.0.1:8080/devices/water/on"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	checkError(err)
	sender := &http.Client{}
	_, err = sender.Do(req)
	checkError(err)

	time.Sleep(time.Duration(float64(amount)/(0.2)) * time.Second)

	url = "http://127.0.0.1:8080/devices/water/off"
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	checkError(err)
	sender = &http.Client{}
	_, err = sender.Do(req)
	checkError(err)
}

func HandleRepeat(w http.ResponseWriter, r *http.Request) {
	var repeat RepeatPreset
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode("OK")
	if err := json.NewDecoder(r.Body).Decode(&repeat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Repeat = repeat
	Mode = "REPEAT"
	for Mode == "REPEAT" {
		t := time.Now()
		layout := "15:04:00"
		timeset, err := time.Parse(layout, Repeat.Time)
		checkError(err)
		if t.Hour() == timeset.Hour() && t.Minute() == timeset.Minute() {
			IntervalMode(strconv.Atoi(Repeat.Amount))
		}
		time.Sleep(65 * time.Second)
	}
}

func HandleSoil(w http.ResponseWriter, r *http.Request) {
	println("New Data Received")
	var sensorData SensorMsg
	err := json.NewDecoder(r.Body).Decode(&sensorData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AutoMode(sensorData.Data) // run when receiving any soil data
}

// "id": 1,
// "name": "Rosie Preset",
// "moistureThreshold": 30,
// "notes": "roses are red violets are blue"
type AutoPreset struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Moisture float64 `json:"moistureThreshold"`
	Note     string  `json:"notes"`
}

type IntervalPreset struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	Interval string `json:"interval"`
}

type RepeatPreset struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	Amount string `json:"amount"`
	Repeat string `json:"repeat"`
}

func AutoMode(input string) {
	data, _ := strconv.ParseFloat(input, 64)
	if Mode == "AUTO" {
		if data <= Auto.Moisture {
			url := "http://127.0.0.1:8080/devices/water/on"
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Set("Content-Type", "application/json")
			checkError(err)
			sender := &http.Client{}
			_, err = sender.Do(req)

			checkError(err)

			time.Sleep(6 * time.Second)

			url = "http://127.0.0.1:8080/devices/water/off"
			req, err = http.NewRequest("GET", url, nil)
			req.Header.Set("Content-Type", "application/json")
			checkError(err)
			sender = &http.Client{}
			_, err = sender.Do(req)
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
