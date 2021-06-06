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

type PlotListRequest struct {
	UserId string `json:"UserId"`
}

type PlotInfoResponse struct {
	Name string `json:"PlotName"`
	Id   string `json:"Id"`
}

type DeviceListRequest struct {
	PlotId string `json:"PlotId"`
}

type DeviceInfoResponse struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
	Type string `json:"Type"`
}

type DeviceDataRequest struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	DeviceName string `json:"deviceName"`
	Type       string `json:"type"`
}

func (db *Data) SetupRouter() {
	db.Router.
		HandleFunc("/devices/data/log", handlePlotDeviceData).
		Methods("POST", "OPTIONS")
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// func (db *Data) handleDeviceDataMsg(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.WriteHeader(http.StatusOK)
// 	var deviceListReq DeviceListRequest
// 	var deviceListRes []DeviceInfoResponse
// 	var tmp DeviceInfoResponse

// 	err := json.NewDecoder(r.Body).Decode(&deviceListReq)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if decoded.PlotId == "" {
// 		return
// 	}
// 	results, err := db.Database.Query("CALL retrieve_device_data(?,?);", decoded.UserId, decoded.PlotId)
// 	checkQueryError(err)
// 	defer results.Close()

// 	for results.Next() {
// 		err = results.Scan(&tmp.Id, &tmp.Name, &tmp.UserId, &tmp.PlotId, &tmp.Type, &tmp.Position)
// 		checkScanError(err)
// 		deviceData = append(deviceData, tmp)
// 	}
// 	jsonString, err := json.Marshal(deviceData)
// 	checkError(err)

// 	// buff := bytes.NewBuffer(jsonString)
// 	fmt.Println("Successfully Handle Request")
// 	w.Write(jsonString)
// }

func handlePlotDeviceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var deviceDataReq DeviceDataRequest

	err := json.NewDecoder(r.Body).Decode(&deviceDataReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var feed string
	if deviceDataReq.Type == "soil" {
		feed = "soilmoist"
	} else {
		feed = "temphumid"
	}
	url := fmt.Sprintf("https://io.adafruit.com/api/v2/MDPSmartFarm/feeds/%s/data?start_time=%s&end_time=%s&include=value,created_at", feed, deviceDataReq.Start_time, deviceDataReq.End_time)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	checkError(err)
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()

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
		if tmp.Name == deviceDataReq.DeviceName {
			dateTime := strings.Split(apiResponse[i]["created_at"], "T")
			listValue = append(listValue, []string{dateTime[0] + " " + dateTime[1] + " -0700", tmp.Data})
		}
	}
	jsonmsg := dataLog{listValue}
	jsonString, _ := json.Marshal(jsonmsg)
	w.Write(jsonString)
}
