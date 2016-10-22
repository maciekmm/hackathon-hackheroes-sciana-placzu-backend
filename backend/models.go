package main

import "database/sql/driver"

type Category string

var (
	Urgent    Category = "URGENT"
	Stable    Category = "STABLE"
	Undefined Category = "UNDEFINED"
)

func (c *Category) Scan(value interface{}) error {
	if value == nil {
		*c = Undefined
		return nil
	}
	*c = Category(value.([]uint8))
	return nil
}

func (c Category) Value() (driver.Value, error) {
	return string(c), nil
}

type Cell struct {
	Cell    string `db:"cell" json:"name,omitempty"`
	City    string `db:"city" json:"city,omitempty"`
	Address string `db:"address" json:"address,omitempty"`
	Phone   string `db:"phone" json:"phone,omitempty"`
}

type Provider struct {
	ProviderName string `db:"provider_name" json:"name,omitempty"`
	Voivodeship  string `db:"voivodeship" json:"voivodeship,omitempty"`
}

type Service struct {
	ID                 int    `db:"id" json:"id"`
	Name               string `db:"name" json:"name,omitempty"`
	*Provider          `json:"provider,omitempty"`
	Category           Category `db:"category" json:"category,omitempty"`
	*Cell              `json:"cell,omitempty"`
	Waiting            int    `db:"waiting" json:"waiting"`
	Removed            int    `db:"removed" json:"removed"`
	AverageWaitingTime int    `db:"average_waiting_time" json:"average_waiting_time"`
	FirstAvailableDate string `db:"first_available_date" json:"first_available_date,omitempty"`
	DatePrepared       string `db:"date_prepared" json:"date_prepared,omitempty"`
	DateUpdated        string `db:"date_updated" json:"date_updated,omitempty"`
	DateInserted       string `db:"date_inserted" json:"date_inserted"`
}
