package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type plugStatus struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Body       plugStatusBody `json:"body"`
}

type plugStatusBody struct {
	DeviceId         string  `json:"deviceId"`
	DeviceType       string  `json:"deviceType"`
	HubDeviceId      string  `json:"hubDeviceId"`
	Power            string  `json:"power"`
	Voltage          float64 `json:"voltage"`
	Weight           float64 `json:"weight"`
	ElectricityOfDay float64 `json:"electricityOfDay"`
	ElectricCurrent  float64 `json:"electricCurrent"`
}

func getToken() string {
	token := os.Getenv("SWITCHBOT_TOKEN")
	if token == "" {
		log.Fatal("SWITCHBOT_TOKEN is not set")
	}
	return token
}

func retrievePlugStatus(deviceId string) (*plugStatusBody, error) {
	url := fmt.Sprintf("https://api.switch-bot.com/v1.0/devices/%s/status", deviceId)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", getToken())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	status := plugStatus{}
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	fmt.Println(status)

	return &status.Body, nil
}
