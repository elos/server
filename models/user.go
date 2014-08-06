package models

import (
	"encoding/json"

	"time"
)

type User struct {
	Id          int `json:"id"`
	AccessToken string
	CreatedAt   time.Time `json:"created_at"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.MarshalIndent(*u, "", "    ")
}
