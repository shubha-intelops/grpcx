package models

type Bill struct {
	Id int64 `json:"id,omitempty"`

	Ammount int `json:"ammount,omitempty"`

	Name int `json:"name,omitempty"`
}
