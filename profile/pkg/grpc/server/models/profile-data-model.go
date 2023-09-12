package models

type ProfileData struct {
	Id int64 `json:"id,omitempty"`

	Address string `json:"address,omitempty"`

	Name string `json:"name,omitempty"`
}
