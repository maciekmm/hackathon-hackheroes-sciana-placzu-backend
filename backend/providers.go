package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type providerData struct {
	Provider
	Cell
}

type providerOut struct {
	Provider Provider `json:"provider"`
	Cells    []Cell   `json:"cells"`
}

func providerEndpoint(rw http.ResponseWriter, req *http.Request) {
	services := []providerData{}
	cells := []Cell{}
	providerName := req.URL.Query().Get("provider")
	if providerName == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("provider parameter cannot be empty"))
		return
	}
	err := connection.Select(&services, "SELECT provider_name,voivodeship,cell,phone,city,address FROM `services` WHERE provider_name=? GROUP BY cell", providerName)
	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
		return
	}

	if len(services) == 0 {
		rw.WriteHeader(404)
		return
	}

	for _, service := range services {
		cells = append(cells, service.Cell)
	}

	enc := json.NewEncoder(rw)
	enc.Encode(&providerOut{Provider: services[0].Provider, Cells: cells})
}

type cellOut struct {
	*Cell
	Provider *Provider  `json:"provider"`
	Services []*Service `json:"services"`
}

func cellEndpoint(rw http.ResponseWriter, req *http.Request) {
	providerName := req.URL.Query().Get("provider")
	if providerName == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("provider parameter cannot be empty"))
		return
	}
	cellName := req.URL.Query().Get("cell")
	if cellName == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("cell parameter cannot be empty"))
		return
	}
	services := []*Service{}

	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -currentTime.Day()+1)

	err := connection.Select(&services, "SELECT * FROM `services` WHERE `cell`=? AND `provider_name`=? AND `date_updated`>=DATE(?) AND `date_inserted`>=DATE(?) ORDER BY `first_available_date` DESC", cellName, providerName, currentTime.AddDate(0, -2, 0).Format("2006-01-02"), currentTime.Format("2006-01-02"))
	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
		return
	}

	if len(services) == 0 {
		rw.WriteHeader(404)
		return
	}
	out := &cellOut{Cell: services[0].Cell, Provider: services[0].Provider}
	for _, service := range services {
		service.Cell = nil
		service.Provider = nil
	}
	out.Services = services
	enc := json.NewEncoder(rw)
	enc.Encode(out)
}
