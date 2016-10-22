package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var cachedStats []byte

func initStats() {
	stats := struct {
		Records   int `db:"records" json:"records"`
		Servuices int `db:"services" json:"services"`
		Providers int `db:"providers" json:"providers"`
	}{}

	err := connection.Get(&stats, "SELECT COUNT(*) as records, COUNT(DISTINCT name) as services, COUNT(DISTINCT provider_name) as providers FROM `services`")
	if err != nil {
		log.Fatalln("Caching statistics failed during db access ", err)
	}

	cachedStats, err = json.Marshal(stats)
	if err != nil {
		log.Fatalln("Caching statistics failed during marshalling ", err)
	}
}

func statsEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(cachedStats)
}
