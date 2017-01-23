package models

//Request is a struct for received JSON
type Request struct {
	Site       []string `json:"site"`
	SearchText string   `json:"searchText"`
}
