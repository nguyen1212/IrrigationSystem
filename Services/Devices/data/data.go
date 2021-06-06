package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
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

func (db *Data) SetupRouter() {
	db.Router.
		HandleFunc("/devices/data/info", db.handleDeviceDataMsg).
		Methods("POST", "OPTIONS")
		// Path("/devices/data").

	db.Router.
		HandleFunc("/devices/data/log", handlePlotDeviceData).
		Methods("POST", "OPTIONS")
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

func (db *Data) handleDeviceDataMsg(w http.ResponseWriter, r *http.Request) {
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
	if decoded.PlotId == "" {
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

func handlePlotDeviceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	req, err := http.NewRequest("GET", "https://io.adafruit.com/api/v2/MDPSmartFarm/feeds/soilmoist/data?start_time=2021-05-04T13:45:00Z&end_time=2021-06-04T13:47:00Z&include=value,created_at", nil)
	// req.Header.Set("Content-Type", "application/json")
	checkError(err)
	// Execute the HTTP request.
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	// Tell Go to close the response at the end of the function.
	defer response.Body.Close()

	// Read the raw response into a Go struct.
	type responseData struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Data string `json:"data"`
		Unit string `json:"unit"`
	}

	type dataLog struct {
		Data [][]string `json:"Data"`
	}
	var apiResponse []map[string]string
	var listValue [][]string
	er := json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(er)

	tmp := &responseData{}
	for i := len(apiResponse) - 1; i >= 0; i-- {
		err := json.Unmarshal([]byte(apiResponse[i]["value"]), tmp)
		checkError(err)
		if tmp.Name == "SOIL" {
			dateTime := strings.Split(apiResponse[i]["created_at"], "T")
			listValue = append(listValue, []string{dateTime[0] + " " + dateTime[1] + " -0700", tmp.Data})
			// listDateTime = append(listDateTime, )
		}
	}
	// fmt.Println(listDateTime[0])
	jsonmsg := dataLog{listValue}
	jsonString, _ := json.Marshal(jsonmsg)
	w.Write(jsonString)
}
