package model

type Address struct {
	Value string `json:"value"`
	Lat   string `json:"lat"`
	Lng   string `json:"lon"`
}

type SearchIn struct {
	Query string `json:"query"`
}

type SearchOut struct {
	Addresses []*Address `json:"addresses"`
	Status    int        `json:"status"`
}

type GeocodeIn struct {
	Lat string `json:"lat"`
	Lng string `json:"lon"`
}

type GeocodeOut struct {
	Addresses []*Address `json:"addresses"`
	Status    int        `json:"status"`
}
