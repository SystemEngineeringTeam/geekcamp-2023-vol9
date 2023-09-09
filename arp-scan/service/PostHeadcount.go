package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type RequestBody struct {
	Time      string `json:"time"`
	HeadCount int    `json:"headcount"`
}

func PostCount(id, count int, time string) {
	godotenv.Load(".env")

	requestBody := &RequestBody{
		Time:      time,
		HeadCount: count,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		panic("Error")
	}

	endpoint := fmt.Sprint(os.Getenv("SERVER_URL"), id)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	fmt.Println(string(byteArray))
}
