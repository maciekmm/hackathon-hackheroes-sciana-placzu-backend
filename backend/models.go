package main

type Category string

var (
	Urgent    Category = "URGENT"
	Stable    Category = "STABLE"
	Undefined Category = "UNDEFINED"
)

type Service struct {
	ID                 int      `sql:"id" json:"id"`
	Voivodeship        string   `sql:"voivodeship" json:"voivodeship"`
	Name               string   `sql:"name" json:"name"`
	Category           Category `sql:"category" json:"category"`
	City               string   `sql:"city" json:"city"`
	Address            string   `sql:"address" json:"address"`
	Phone              string   `sql:"phone" json:"phone"`
	ProviderName       string   `sql:"provider_name" json:"provider_name"`
	Cell               string   `sql:"cell" json:"cell"`
	Waiting            int      `sql:"waiting" json:"waiting"`
	Removed            int      `sql:"removed" json:"removed"`
	AverageWaitingTime int      `sql:"average_waiting_time" json:"average_waiting_time"`
	FirstAvailableDate string   `sql:"first_available_date" json:"first_available_date"`
	DatePrepared       string   `sql:"date_prepared" json:"date_prepared"`
	DateUpdated        string   `sql:"date_updated" json:"date_updated"`
}
