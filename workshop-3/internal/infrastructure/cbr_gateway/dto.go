package cbr_gateway

type Rates struct {
	Currencies []*Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int64  `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}
