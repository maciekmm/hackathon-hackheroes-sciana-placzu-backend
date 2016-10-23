![hack-heroes hacathon](../assets/hackheroes-logo.png)

# Backend of Ściana Płaczu [sciana.placzu.pl](https://sciana.placzu.pl)

## Endpoints

### /top

Displays services with the longest waiting time.

#### Query parameter

- **limit**[1-50] - amount of results to return, default: 5

#### Example:

[https://sciana.placzu.pl/top?limit=1](https://sciana.placzu.pl/top?limit=1)

```json
[
  {
    "id": 94565,
    "name": "ENDOPROTEZOPLASTYKA STAWU KOLANOWEGO",
    "provider": {
      "name": "SAMODZIELNY PUBLICZNY WOJEWÓDZKI SZPITAL CHIRURGII URAZOWEJ IM.DR J. DAABA",
      "voivodeship": "ŚLĄSKIE"
    },
    "category": "STABLE",
    "cell": {
      "name": "ODDZIAŁ II URAZOWO-ORTOPEDYCZNY Z PODODDZIAŁEM ARTROSKOPII I CHIRURGII KOLANA",
      "city": "PIEKARY ŚLĄSKIE",
      "address": "BYTOMSKA  62",
      "phone": "323934323"
    },
    "waiting": 5288,
    "removed": 15,
    "average_waiting_time": 2225,
    "first_available_date": "2040-08-27",
    "date_prepared": "2016-09-23",
    "date_updated": "2016-08-01",
    "date_inserted": "2016-10-22 18:23:22"
  }
]
```

### /services 

Lists all services.

#### Example:
[https://sciana.placzu.pl/services](https://sciana.placzu.pl/services)

```json
[
  {
    "name": "AMBULATORYJNA STACJA DIALIZ",
    "cats": [
      "STABLE",
      "URGENT"
    ]
  },
  {
    "name": "BADANIA GENETYCZNE",
    "waiting_time": 25.4348, //in days
    "cats": [
      "STABLE",
      "URGENT"
    ]
  },
  {
    "name": "BADANIA MEDYCYNY NUKLEARNEJ",
    "waiting_time": 7.9224,
    "cats": [
      "STABLE",
      "URGENT"
    ]
  },
  {
    "name": "DZIAŁ (PRACOWNIA) FIZJOTERAPII",
    "waiting_time": 84.7462,
    "cats": [
      "STABLE",
      "URGENT"
    ]
  },
  {
    "name": "DZIAŁ (PRACOWNIA) FIZJOTERAPII DLA DZIECI",
    "waiting_time": 32.2361,
    "cats": [
      "STABLE",
      "URGENT"
    ]
  },
  {
    "name": "DZIAŁ (PRACOWNIA) FIZYKOTERAPII",
    "waiting_time": 3.75,
    "cats": [
      "STABLE",
      "URGENT"
    ]
  }
]
```

### /stats

Displays statistics about the aggregated data.

#### Example

[https://sciana.placzu.pl/stats](https://sciana.placzu.pl/stats)

```json
{
  "records": 141037,
  "services": 387,
  "providers": 14216
}
```

### /provider

#### Query parameter

- **provider** - provider's name

#### Example

[https://sciana.placzu.pl/provider?provider=SAMODZIELNY%20PUBLICZNY%20WOJEW%C3%93DZKI%20SZPITAL%20CHIRURGII%20URAZOWEJ%20IM.DR%20J.%20DAABA](https://sciana.placzu.pl/provider?provider=SAMODZIELNY%20PUBLICZNY%20WOJEW%C3%93DZKI%20SZPITAL%20CHIRURGII%20URAZOWEJ%20IM.DR%20J.%20DAABA)

```json
{
  "provider": {
    "name": "SAMODZIELNY PUBLICZNY WOJEWÓDZKI SZPITAL CHIRURGII URAZOWEJ IM.DR J. DAABA",
    "voivodeship": "ŚLĄSKIE"
  },
  "cells": [
    {
      "name": "DZIAŁ (PRACOWNIA) FIZJOTERAPII",
      "city": "KOCHCICE",
      "address": "ZAMKOWA 1",
      "phone": "343533631"
    },
    {
      "name": "ODDZIAŁ I URAZOWO-ORTOPEDYCZNY MĘSKI",
      "city": "PIEKARY ŚLĄSKIE",
      "address": "BYTOMSKA  62",
      "phone": "323934317"
    }
	 //...
  ]
}

```

### /cell

Returns all services in a specific cell

#### Query parameters

- **provider** - provider's name
- **cell** - cell's name

#### Example:

[http://sciana.placzu.pl/cell?provider=SAMODZIELNY%20ZESP%C3%93%C5%81%20PUBLICZNYCH%20ZAK%C5%81AD%C3%93W%20OPIEKI%20ZDROWOTNEJ%20W%20OLE%C5%9ANICY&cell=PORADNIA%20STOMATOLOGICZNA](http://sciana.placzu.pl/cell?provider=SAMODZIELNY%20ZESP%C3%93%C5%81%20PUBLICZNYCH%20ZAK%C5%81AD%C3%93W%20OPIEKI%20ZDROWOTNEJ%20W%20OLE%C5%9ANICY&cell=PORADNIA%20STOMATOLOGICZNA)

```json
[
  {
    "id": 24,
    "name": "PORADNIA STOMATOLOGICZNA",
    "category": "STABLE",
    "waiting": 7,
    "removed": 3,
    "average_waiting_time": 26,
    "first_available_date": "2016-10-18",
    "date_prepared": "2016-09-23",
    "date_updated": "2016-08-01",
    "date_inserted": "2016-10-22 18:16:56"
  },
  {
    "id": 5,
    "name": "LECZENIE PROTETYCZNE",
    "category": "URGENT",
    "waiting": 0,
    "removed": 0,
    "average_waiting_time": 0,
    "first_available_date": "2016-09-26",
    "date_prepared": "2016-09-23",
    "date_updated": "2016-08-01",
    "date_inserted": "2016-10-22 18:16:56"
  },
  {
    "id": 6,
    "name": "LECZENIE PROTETYCZNE",
    "category": "STABLE",
    "waiting": 0,
    "removed": 0,
    "average_waiting_time": 0,
    "first_available_date": "2016-09-26",
    "date_prepared": "2016-09-23",
    "date_updated": "2016-08-01",
    "date_inserted": "2016-10-22 18:16:56"
  }
]

```

### /search

Returns services matching specific criteria

#### Query parameters

- **name** - service's name
- **voivodeship** - voivodeship
- **category**[stable,unstable,undefined] - category

### /service

Returns information about a particular

#### Query parameters

- **id** - service's id

#### Example:

[http://sciana.placzu.pl/service?id=24](http://sciana.placzu.pl/service?id=24)
```json
{
  "id": 24,
  "name": "PORADNIA STOMATOLOGICZNA",
  "provider": {
    "name": "SAMODZIELNY ZESPÓŁ PUBLICZNYCH ZAKŁADÓW OPIEKI ZDROWOTNEJ W OLEŚNICY",
    "voivodeship": "DOLNOŚLĄSKIE"
  },
  "category": "STABLE",
  "cell": {
    "name": "PORADNIA STOMATOLOGICZNA",
    "city": "OLEŚNICA",
    "address": "LUDWIKOWSKA 10",
    "phone": "717982837"
  },
  "waiting": 7,
  "removed": 3,
  "average_waiting_time": 26,
  "first_available_date": "2016-10-18",
  "date_prepared": "2016-09-23",
  "date_updated": "2016-08-01",
  "date_inserted": "2016-10-22 18:16:56"
}

```

## Building

```bash
go get ./... #install dependencies
go build #build
```

## Running

The application requires either ``config.json`` to be present in the runtime directory or *-import configfile.json* flag specified.

```json
{
	"host": "",
	"username": "root",
	"password": "PtMXrJjuSXrd5jfMLY",
	"database": "scianaplaczu"
}
```

### Importing

In order to import the data, go to [http://kolejki.nfz.gov.pl/Informator/PobierzDane/Index/](http://kolejki.nfz.gov.pl/Informator/PobierzDane/Index/) and download all archives. Extract them to a single folder, **DON'T CHANGE FILE NAMES**

execute ``./backend -import ./folder-with-extracted-xlsx-files`` and wait:
