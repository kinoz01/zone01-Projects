package server

var Logs string

type Ports struct {
	Port    string
	ApiPort string
}

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type MyAPI struct {
	Image          string              `json:"image"`
	MembersImages  map[string]string   `json:"membersImages"`
	LocationImages map[string][]string `json:"datesLocations"`
	YoutubeURL     []string            `json:"youtubeUrl"`
}

type ArtistDetails struct {
	Artist    Artist
	Locations Locations
	Dates     Dates
	Relations Relations
	MyAPI     MyAPI
}
