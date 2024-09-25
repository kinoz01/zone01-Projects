package server

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
	LOCATIONS    []string
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

// ApiLinks represents the JSON links structure
type ApiLinks struct {
	Artist         string   `json:"artist"`
	Locations      string   `json:"locations"`
	Dates          string   `json:"dates"`
	Relations      string   `json:"relations"`
	SerpApi        string   `json:"serp"`
	OtherLinks     []string `json:"others"`
	ErrorPage      string   `json:"error"`
	Home           string   `json:"home"`
	ReplacedImages []string `json:"replacedImages"`
	AllLocations   string   `json:"allLocations"`
}
