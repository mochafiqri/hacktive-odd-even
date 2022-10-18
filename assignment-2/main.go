package main

import (
	"fmt"
	"html/template"
	"math/rand"
)
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = "index.html"
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var wind = rand.Intn(20-3) + 3
		var water = rand.Intn(10-3) + 3
		var statusWind = ""
		var statusWater = ""

		if water < 5 {
			statusWater = "aman"
		} else if water >= 6 && water <= 8 {
			statusWater = "siaga"
		} else {
			statusWater = "bahaya"
		}

		if wind < 6 {
			statusWind = "aman"
		} else if wind >= 7 && wind <= 15 {
			statusWind = "siaga"
		} else {
			statusWind = "bahaya"
		}

		var data = map[string]interface{}{
			"name":         "Wind & Water Status",
			"wind":         wind,
			"water":        water,
			"wind_status":  statusWind,
			"water_status": statusWater,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
