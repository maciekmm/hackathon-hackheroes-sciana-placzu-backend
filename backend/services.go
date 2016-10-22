package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	rw.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(rw)
	err = enc.Encode(services)

	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
	}
}

type shallowService struct {
	Name               string   `json:"name" db:"name"`
	AverageWaitingTime float64  `json:"waiting_time,omitempty" db:"average_reliable_waiting_time_in_days"`
	Categories         string   `json:"-" db:"categories"`
	Cats               []string `json:"cats"`
}

var servicesCache []byte

//initServices caches services because of a pretty time consuming query
func initServices() {
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -currentTime.Day()+1)
	services := []*shallowService{}
	err := connection.Select(&services, "SELECT name, GROUP_CONCAT(DISTINCT category ORDER BY category SEPARATOR ';') as categories, GREATEST(AVG(TIMESTAMPDIFF(DAY, NOW(),first_available_date)),0) as average_reliable_waiting_time_in_days FROM services WHERE date_updated>=DATE(?) GROUP BY name", currentTime.AddDate(0, -2, 0).Format("2006-01-02"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, service := range services {
		service.Cats = strings.Split(service.Categories, ";")
	}
	servicesCache, err = json.Marshal(services)
	if err != nil {
		log.Fatalln(err)
	}
}

func servicesEndpoint(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(servicesCache)
}

func searchEndpoint(rw http.ResponseWriter, req *http.Request) {
	var err error
	currentTime := time.Now()
	currentTime = currentTime.AddDate(0, 0, -currentTime.Day()+1)

	name := req.URL.Query().Get("name")
	if name == "" {
		rw.WriteHeader(500)
		rw.Write([]byte("name parameter cannot be empty"))
		return
	}
	voivodeship := req.URL.Query().Get("voivodeship")
	category := req.URL.Query().Get("category")

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

	parameters := []interface{}{name}

	services := []Service{}

	buildQuery := "SELECT * FROM `services` WHERE `name`=? "
	if voivodeship != "" {
		buildQuery = buildQuery + "AND `voivodeship` LIKE ? "
		parameters = append(parameters, voivodeship)
	}
	if category != "" {
		buildQuery = buildQuery + "AND `category`=? "
		parameters = append(parameters, category)
	}
	buildQuery = buildQuery + " AND date_updated>=DATE(?) AND date_inserted>=DATE(?) ORDER BY `first_available_date` ASC LIMIT " + strconv.Itoa(limit)
	parameters = append(parameters, currentTime.AddDate(0, -2, 0).Format("2006-01-02"), currentTime.Format("2006-01-02"))

	err = connection.Select(&services, buildQuery, parameters...)
	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
		return
	}

	enc := json.NewEncoder(rw)
	err = enc.Encode(services)

	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
	}
}

func serviceEndpoint(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("id parameter cannot be empty"))
		return
	}
	service := &Service{}
	err = connection.Get(service, "SELECT * FROM services WHERE id=?", id)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("error occured while processing request"))
		return
	}
	enc := json.NewEncoder(rw)
	err = enc.Encode(service)

	if err != nil {
		rw.WriteHeader(500)
		log.Println(err)
	}
}
