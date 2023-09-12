package models

type Name struct {
	Id int64 `json:"id,omitempty"`

	Address string `json:"address,omitempty"`

	Age int32 `json:"age,omitempty"`

	Email_address string `json:"emailAddress,omitempty"`

	Name string `json:"name,omitempty"`
}
