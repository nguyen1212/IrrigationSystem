package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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
func HandlePlotWater(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	type WaterRequest struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}
	var deviceDataReq WaterRequest

	err := json.NewDecoder(r.Body).Decode(&deviceDataReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://io.adafruit.com/api/v2/CSE_BBC1/feeds/%s/data?start_time=%s&end_time=%s&include=value,created_at", "bk-iot-relay", deviceDataReq.StartDate, deviceDataReq.EndDate)
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
		Data []string `json:"Data"`
		Avg  string   `json:"avg"`
	}
	var apiResponse []map[string]string
	var listValue [][]string
	er := json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(er)
	tmp := &responseData{}
	var newList [][]string
	layout := "2006-01-02T15:04:05Z"
	count2 := 0
	for i := len(apiResponse) - 1; i >= 0; i-- {
		err := json.Unmarshal([]byte(apiResponse[i]["value"]), tmp)
		if err != nil {
			continue
		}
		dateTime := apiResponse[i]["created_at"]
		if len(tmp.Data) == 0 {
			continue
		}
		listValue = append(listValue, []string{dateTime, tmp.Data})
		if i == len(apiResponse)-1 {
			newList = append(newList, []string{dateTime[0:10], "0"})
		} else {
			if dateTime[0:10] != newList[count2][0] {
				newList = append(newList, []string{dateTime[0:10], "0"})
				count2++
			}
		}

	}
	var sumDay float64
	sumDay = 0
	count := 0
	var newArr []string
	for i := 0; i < len(listValue)-1; i++ {
		realdateoff, _ := time.Parse(layout, listValue[i+1][0])
		realdateon, _ := time.Parse(layout, listValue[i][0])
		timeDif := float64(realdateoff.Sub(realdateon) / time.Second)
		amount := timeDif * 0.2
		str := fmt.Sprintf("%d-%02d-%02d",
			realdateon.Year(), realdateon.Month(), realdateon.Day())
		if str == newList[count][0] {
			tempAmount, _ := strconv.ParseFloat(newList[count][1], 64)
			sumDay += amount
			newList[count][1] = fmt.Sprintf("%f", tempAmount+amount)
		} else {
			count += 1
		}
		i++
	}
	println(len(newList))
	for i := 0; i < len(newList); i++ {
		println(newList[i][0])
		if i > 0 {
			time1, _ := time.Parse("2006-01-02", newList[i-1][0])
			time2, _ := time.Parse("2006-01-02", newList[i][0])
			dif := int(time2.Sub(time1) / time.Hour)
			if dif/24 > 1 {
				for j := 0; j < dif-1; j++ {
					newArr = append(newArr, "0")
					println("run 0, ", i)
				}
			} else {
				newArr = append(newArr, newList[i][1])
				println("run ", i)
			}
		} else {
			newArr = append(newArr, newList[i][1])
		}
	}
	// for i := 0; i < arrlen; i++ {
	// 	if i > 0 {
	// 		time1, _ := time.Parse("2006-01-02", newList[c-1][0])
	// 		time2, _ := time.Parse("2006-01-02", newList[c][0])
	// 		arrlen := int64((time2.Sub(time1) * time.Hour) / 24)
	// 		if arrlen > 1 {
	// 			newArr = append(newArr, "0")
	// 		}
	// 	}
	// 	newArr = append(newArr, newList[i][1])
	// }
	jsonmsg := dataLog{newArr, fmt.Sprintf("%.2f", sumDay/float64(count))}
	jsonString, _ := json.Marshal(jsonmsg)
	w.Write(jsonString)
}

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
