package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DB struct {
	Router   *mux.Router
	Database *sql.DB
}
type DeviceData struct {
	Name     string `json:"Name"`
	Id       string `json:"Id"`
	UserId   string `json:"UserId"`
	PlotId   string `json:"PlotId"`
	Type     string `json:"Type"`
	Position string `json:"Position"`
}

type DeviceAPIRequest struct {
	UserId string `json:"UserId"`
	PlotId string `json:"PlotId"`
}

func (db *DB) SetupRouter() {
	db.Router.
		HandleFunc("/devices/data", db.handleDeviceData).
		Methods("POST", "OPTIONS")
		// Path("/devices/data").

	db.Router.
		HandleFunc("/", db.handleDeviceData)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func checkQueryError(err error) {
	if err != nil {
		log.Fatal("Database Select Failed")
	}
}

func checkScanError(err error) {
	if err != nil {
		log.Fatal("Cannot Scan Data")
	}
}

func (db *DB) handleDeviceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	var decoded DeviceAPIRequest
	var deviceData []DeviceData
	var tmp DeviceData

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	results, err := db.Database.Query("CALL retrieve_device_data(?,?);", decoded.UserId, decoded.PlotId)
	checkQueryError(err)
	defer results.Close()

	for results.Next() {
		err = results.Scan(&tmp.Id, &tmp.Name, &tmp.UserId, &tmp.PlotId, &tmp.Type, &tmp.Position)
		checkScanError(err)
		deviceData = append(deviceData, tmp)
	}
	jsonString, err := json.Marshal(deviceData)
	checkError(err)

	// buff := bytes.NewBuffer(jsonString)
	fmt.Println("Successfully Handle Request")
	w.Write(jsonString)
}
