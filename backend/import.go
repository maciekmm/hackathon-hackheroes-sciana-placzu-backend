package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

var voivodeships map[string]string = map[string]string{
	"01": "DOLNOŚLĄSKIE",
	"02": "KUJAWSKO-POMORSKIE",
	"03": "LUBELSKIE",
	"04": "LUBUSKIE",
	"05": "ŁÓDZKIE",
	"06": "MAŁOPOLSKIE",
	"07": "MAZOWIECKIE",
	"08": "OPOLSKIE",
	"09": "PODKARPACKIE",
	"10": "PODLASKIE",
	"11": "POMORSKIE",
	"12": "ŚLĄSKIE",
	"13": "ŚWIĘTOKRZYSKIE",
	"14": "WARMIŃSKO-MAZURSKIE",
	"15": "WIELKOPOLSKIE",
	"16": "ZACHODNIOPOMORSKIE",
}

func beautifyPhone(phone string) string {
	replacer := strings.NewReplacer(" ", "", "-", "", "(", "", ")", "", ",", ";")
	phones := strings.Split(replacer.Replace(phone), ";")
	for i, phone := range phones {
		phones[i] = strings.TrimLeft(strings.TrimLeft(phone, "0"), "+48")
	}
	return strings.Join(phones, ";")
}

func startImport(folder string) error {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}

	//iterate through files to import
	for _, info := range files {
		fileName := info.Name()
		//voivodeship is included in filename
		voivodeship := voivodeships[strings.Split(fileName, "_")[0]]
		//open excel file
		xlFile, err := xlsx.OpenFile(path.Join(folder, fileName))
		if err != nil {
			log.Fatalln(err)
			return err
		}
		stmt, _ := connection.Prepare("INSERT INTO services(`voivodeship`,`name`,`category`,`city`,`address`,`phone`,`provider_name`,`cell`,`waiting`,`removed`,`average_waiting_time`,`first_available_date`,`date_prepared`,`date_updated`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		defer stmt.Close()
		for _, row := range xlFile.Sheets[0].Rows[2:] { //omit header rows
			service := Service{
				Provider: &Provider{},
				Cell:     &Cell{},
				Name:     strings.TrimSpace(row.Cells[0].Value),
			}
			service.ProviderName = strings.TrimSpace(row.Cells[2].Value)
			service.Voivodeship = voivodeship
			service.Cell.Cell = strings.TrimSpace(row.Cells[3].Value)

			if strings.Contains(row.Cells[1].Value, "pilny") {
				service.Category = Urgent
			} else if strings.Contains(row.Cells[1].Value, "stabilny") {
				service.Category = Stable
			} else {
				service.Category = Undefined
			}

			//
			addressParts := strings.Split(row.Cells[4].Value, "\n")
			service.City = strings.TrimSpace(addressParts[0])
			service.Address = strings.TrimSpace(addressParts[1])
			service.Phone = beautifyPhone(strings.TrimSpace(addressParts[2]))

			parsed := strings.Replace(row.Cells[5].Value, " ", "", -1)
			service.Waiting, err = strconv.Atoi(parsed)
			//if can't parse set 0
			if err != nil {
				service.Waiting = 0
			}

			service.Removed, err = strconv.Atoi(strings.Replace(row.Cells[6].Value, " ", "", -1))
			//if can't parse set 0
			if err != nil {
				service.Removed = 0
			}

			service.AverageWaitingTime, err = strconv.Atoi(strings.TrimSpace(row.Cells[7].Value))
			if err != nil {
				service.AverageWaitingTime = 0
			}

			rawDate, err := row.Cells[8].FormattedValue()
			if err != nil {
				fmt.Println("error occured while parsing date: " + row.Cells[8].Value)
			} else {
				date := strings.Split(strings.Replace(rawDate, " ", "", -1), "/")
				if len(date) == 2 {
					service.DateUpdated = fmt.Sprintf("%s-%s-%s", date[1], date[0], "01")
				}
			}

			rawDate, err = row.Cells[9].FormattedValue()
			if err != nil {
				fmt.Println("error occured while parsing date, skipping: " + row.Cells[9].Value)
				continue
			} else {
				date := strings.Split(rawDate, "-")
				if len(date) == 3 {
					day, _ := strconv.Atoi(date[1])
					service.FirstAvailableDate = fmt.Sprintf("20%s-%s-%02d", date[2], date[0], day)
				}
			}
			rawDate, err = row.Cells[10].FormattedValue()
			if err != nil {
				fmt.Println("error occured while parsing date: " + row.Cells[10].Value)
			} else {
				date := strings.Split(rawDate, "-")
				if len(date) == 3 {
					day, _ := strconv.Atoi(date[1])
					service.DatePrepared = fmt.Sprintf("20%s-%s-%02d", date[2], date[0], day)
				}
			}
			_, err = stmt.Exec(voivodeship, service.Name, string(service.Category), service.City, service.Address, service.Phone, service.ProviderName, service.Cell.Cell, service.Waiting, service.Removed, service.AverageWaitingTime, service.FirstAvailableDate, service.DatePrepared, service.DateUpdated)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
