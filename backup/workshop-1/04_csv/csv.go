package main

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
	res, err := http.Get("http://192.168.5.110:8080/storage/temp/geoip.csv.bz2")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("wrong status code:", res.StatusCode)
	}

	fmt.Println(countZones(res.Body, "RU"))
}

// startIpNum, endIpNum,     country,  region,  city,  postalCode,  latitude,  longitude, dmaCode, areaCode
// 1.0.0.0,    1.7.255.255,  "AU",     "",      "",    "",          -27.0000,  133.0000,   ,
// 1.9.0.0,    1.9.255.255,  "MY",     "",      "",    "",          2.5000,    112.5000,   ,
// 1.10.10.0,  1.10.10.255,  "AU",     "",      "",    "",          -27.0000,  133.0000,   ,

func countZones(reader io.Reader, country string) (int, error) {
	bzipReader := bzip2.NewReader(reader)
	csvReader := csv.NewReader(bzipReader)
	counter := 0
	rowId := 0
	for {
		rowId++
		if rowId%1_000_000 == 0 {
			log.Println("processed", rowId, "rows")
		}
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, errors.Wrap(err, "reading csv")
		}
		if len(row) < 3 {
			continue
		}

		if row[2] == country {
			counter++
		}
	}
	log.Println("total rows:", rowId)
	return counter, nil
}
