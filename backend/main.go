package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type databaseCredentials struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var connection *sqlx.DB

func configFromFile(file string) (*databaseCredentials, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	cfg := &databaseCredentials{}
	return cfg, json.Unmarshal(content, cfg)
}

func main() {
	importPath := flag.String("import", "", "folder (with xlsx files) to import data from")
	configPath := flag.String("config", "config.json", "path to database access configuration file")
	port := flag.Int("port", 2001, "port to bind the backend to")
	flag.Parse()

	if *configPath == "" {
		log.Fatalln("Configuration file path must be specified")
	}

	config, err := configFromFile(*configPath)
	if err != nil {
		log.Fatalln("Could not read config file ", err)
	}

	connection, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", config.Username, config.Password, config.Host, config.Database))
	if err != nil {
		log.Fatalln("Could create database connection ", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalln("Could not connect to database ", err)
	}

	//If there's stuff to import, do it
	if *importPath != "" {
		err = startImport(*importPath)
		if err != nil {
			log.Fatalln("Importing data failed ", err)
		}
	}

	if *port == 2001 && os.Getenv("PORT") != "" {
		if *port, err = strconv.Atoi(os.Getenv("PORT")); err != nil {
			log.Fatalln("Invalid port ", err)
		}
	}

	//cache services
	initServices()

	http.HandleFunc("/top", topEndpoint)
	http.HandleFunc("/services", servicesEndpoint)
	http.HandleFunc("/provider", providerEndpoint)
	http.HandleFunc("/cell", cellEndpoint)
	http.HandleFunc("/search", searchEndpoint)
	http.HandleFunc("/service", serviceEndpoint)

	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
