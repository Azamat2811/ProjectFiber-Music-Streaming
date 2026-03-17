package models

type Artists struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Genre             string `json:"genre"`
	Monthly_listeners int    `json:"monthly_listeners"`
}
