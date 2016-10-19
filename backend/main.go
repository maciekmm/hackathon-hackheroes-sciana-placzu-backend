package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type databaseCredentials struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var connection *sql.DB

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
	flag.Parse()

	if *configPath == "" {
		log.Fatalln("Configuration file path must be specified")
	}

	config, err := configFromFile(*configPath)
	if err != nil {
		log.Fatalln("Could not read config file ", err)
	}
	fmt.Println(fmt.Sprintf("%s:%s@%s/%s?parseTime=true", config.Username, config.Password, config.Host, config.Database))
	connection, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", config.Username, config.Password, config.Host, config.Database))
	if err != nil {
		log.Fatalln("Could create database connection ", err)
	}

	err = connection.Ping()
	if err != nil {
		log.Fatalln("Could not connect to database ", err)
	}

	if *importPath != "" {
		fmt.Println(startImport(*importPath))
	}
}
