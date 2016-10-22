package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

const maxLimit int = 50

//SELECT name, GROUP_CONCAT(DISTINCT category ORDER BY category SEPARATOR ';'), FLOOR(AVG(average_waiting_time)) FROM services WHERE date_updated>=2016-08-01 GROUP BY name
func topEndpoint(rw http.ResponseWriter, req *http.Request) {
	var err error
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -currentTime.Day()+1)

	limit := 5

	if req.URL.Query().Get("limit") != "" {
		limit, err = strconv.Atoi(req.URL.Query().Get("limit"))
		if err != nil {
			rw.WriteHeader(400)
			rw.Write([]byte(err.Error()))
			return
		}
		if limit > maxLimit {
			rw.WriteHeader(400)
			rw.Write([]byte("Limit exceeded the maximum limit of 50"))
			return
		}
	}

	services := &[]Service{}
	err = connection.Select(services, "SELECT * FROM `services` WHERE `date_updated`>=DATE(?) AND `date_inserted`>=DATE(?) ORDER BY `first_available_date` DESC LIMIT "+strconv.Itoa(limit), currentTime.AddDate(0, -2, 0).Format("2006-01-02"), currentTime.Format("2006-01-02"))
	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
		return
	}
	enc := json.NewEncoder(rw)
	enc.Encode(services)
}

func servicesEndpoint(rw http.ResponseWriter, req *http.Request) {
	services := &[]string{}
	err := connection.Select(services, "SELECT DISTINCT name FROM services ORDER BY name")
	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
		return
	}
	enc := json.NewEncoder(rw)
	enc.Encode(services)
}
