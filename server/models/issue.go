package models

type Issue struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Reporter    string   `json:"reporter"`
	Watchers    []string `json:"watchers"`
}
