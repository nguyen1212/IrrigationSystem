package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
}

type DeviceListResponse struct {
	SoilDevices  []DeviceInfoResponse `json:"soilDevices"`
	HumidDevices []DeviceInfoResponse `json:"humidDevices"`
	TempDevices  []DeviceInfoResponse `json:"tempDevices"`
}

type DeviceDataRequest struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	DeviceName string `json:"deviceName"`
	Type       string `json:"type"`
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
// 	if deviceListReq.PlotId == "" {
// 		return
// 	}
// 	results, err := db.Database.Query("CALL retrieve_device_data(?);", deviceListReq.PlotId)
// 	checkError(err)
// 	defer results.Close()

// 	for results.Next() {
// 		err = results.Scan(&tmp.Id, &tmp.Name, &tmp.Type)
// 		checkError(err)
// 		deviceListRes = append(deviceListRes, tmp)
// 	}
// 	jsonString, err := json.Marshal(deviceListRes)
// 	checkError(err)

// 	// buff := bytes.NewBuffer(jsonString)
// 	fmt.Println("Successfully Handle Request")
// 	w.Write(jsonString)
// }

func HandlePlotDeviceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	var deviceDataReq DeviceDataRequest

	err := json.NewDecoder(r.Body).Decode(&deviceDataReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var feedKey string
	if deviceDataReq.Type == "soil" {
		feedKey = "bk-iot-soil"
	} else {
		feedKey = "bk-iot-temp-humid"
	}
	url := fmt.Sprintf("https://io.adafruit.com/api/v2/CSE_BBC/feeds/%s/data?start_time=%s&end_time=%s&include=value,created_at", feedKey, deviceDataReq.Start_time, deviceDataReq.End_time)
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
		Data    [][]string `json:"Data"`
		Avg     string     `json:"avg"`
		Max     string     `json:"max"`
		Min     string     `json:"min"`
		MinTime string     `json:"minTime"`
		MaxTime string     `json:"maxTime"`
	}
	var apiResponse []map[string]string
	var listValue [][]string
	er := json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(er)

	tmp := &responseData{}
	var min, max, sum, count float64
	var maxTime, minTime string
	for i := len(apiResponse) - 1; i >= 0; i-- {
		err := json.Unmarshal([]byte(apiResponse[i]["value"]), tmp)
		if err != nil {
			continue
		}
		if tmp.Name == deviceDataReq.DeviceName {
			var tmpData []string
			dateTime := strings.Split(apiResponse[i]["created_at"], "T")
			if deviceDataReq.Type != "soil" {
				tmpData = strings.Split(tmp.Data, "-")
				if deviceDataReq.Type == "temp" {
					tmp.Data = tmpData[0]
				}
				if deviceDataReq.Type == "humid" {
					tmp.Data = tmpData[1]
				}
			}
			if len(tmp.Data) == 0 {
				continue
			}
			if n, err := strconv.Atoi(tmp.Data); err == nil && (n > 100 || n < -100) {
				continue
			}
			floatData, _ := strconv.ParseFloat(tmp.Data, 64)
			sum += floatData
			count++
			if max < floatData {
				max = floatData
				maxTime = dateTime[0] + " " + dateTime[1] + " -0700"
			}
			if count == 1 || min > floatData {
				min = floatData
				minTime = dateTime[0] + " " + dateTime[1] + " -0700"
			}
			listValue = append(listValue, []string{dateTime[0] + " " + dateTime[1] + " -0700", tmp.Data})
		}
	}
	jsonmsg := dataLog{listValue, fmt.Sprintf("%.2f", sum/count), fmt.Sprintf("%.2f", max), fmt.Sprintf("%.2f", min), minTime, maxTime}
	jsonString, _ := json.Marshal(jsonmsg)
	w.Write(jsonString)
}
