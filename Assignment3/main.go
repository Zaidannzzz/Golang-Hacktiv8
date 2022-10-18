package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type StatusData struct {
	Status struct {
		Water int `json: "water"`
		Wind int `json: "wind"`
	} `json: "status"`
}

func AutoReloadJSON()  {
	for {
		min := 1
		max := 100
		wind := rand.Intn(max-min)+min
		water := rand.Intn(max-min)+min

		data := StatusData{}
		data.Status.Wind = wind
		data.Status.Water = water

		jsonData, err := json.Marshal(data)

		if err != nil {
			log.Fatal("Error while marshalling Data: ", err.Error())
		}
		err = ioutil.WriteFile("data.json", jsonData, 0644)

		if err != nil {
			log.Fatal("Error while writing Data to data.json file", err.Error())
		}
		time.Sleep(10 * time.Second)
	}
}

func AutoReloadWeb(w http.ResponseWriter, r *http.Request)  {
	fileData, err := ioutil.ReadFile("data.json")

	if err != nil {
		log.Fatal("Error while reading data from data.json file", err.Error())
	}

	var StatusData StatusData
	err = json.Unmarshal(fileData, &StatusData)
	if err != nil {
		log.Fatal("Error while UnMarshalling from data.json file", err.Error())
	}

	waterval := StatusData.Status.Water
	windval := StatusData.Status.Wind

	var (
		waterStatus string
		windStatus string
 	)

	WaterValue := strconv.Itoa(waterval)
	WindValue := strconv.Itoa(windval)

	switch {
	case waterval <=5:
		waterStatus = "Aman"
	case waterval >=6 && waterval <=8:
		waterStatus = "Siaga"
	case waterval >8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Water Value not defined"
	}

	switch {
	case windval <=6:
		windStatus = "Aman"
	case windval >=7 && windval <=15:
		windStatus = "Siaga"
	case windval >15:
		windStatus = "Bahaya"
	default:
		windStatus = "Wind Value not defined"
	}

	data := map[string]string{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterValue":  WaterValue,
		"windValue":   WindValue,
	}

	tmplt, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatal("Error while parsing HTML:", err.Error())
	}

	tmplt.Execute(w, data)
}

func main() {
	go AutoReloadJSON()
	http.HandleFunc("/", AutoReloadWeb)
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))
	fmt.Println("Listening at port ", ":4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}