package models

type Bill struct {
	Id int64 `json:"id,omitempty"`

	Amount int32 `json:"amount,omitempty"`

	Name string `json:"name,omitempty"`
}
