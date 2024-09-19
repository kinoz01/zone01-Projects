package apiserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ApiData struct {
	BandImages    []string            `json:"BandImages"`
	MembersImages []string            `json:"MembersImages"`
	YoutubeLinks  map[string][]string `json:"YoutubeLinks"`
}

// Initialise my ApiData struct fields.
func (data *ApiData) Initialise() {

	d, err := LoadDataFromFile("./API/apidata.json")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	data.BandImages = d.BandImages
	data.MembersImages = d.MembersImages
	data.YoutubeLinks = d.YoutubeLinks
}

func LoadDataFromFile(filename string) (*ApiData, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data ApiData
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
