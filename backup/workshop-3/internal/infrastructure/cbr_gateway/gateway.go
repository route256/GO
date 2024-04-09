package cbr_gateway

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"gitlab.ozon.dev/go/classroom-4/teachers/homework/internal/domain"
	"golang.org/x/text/encoding/charmap"
)

type Gateway struct {
	client *http.Client
}

func New() *Gateway {
	return &Gateway{
		client: http.DefaultClient,
	}
}

// https://www.cbr-xml-daily.ru/#json

func (gate *Gateway) FetchRatesOn(ctx context.Context, date time.Time) ([]domain.Rate, error) {
	// Чтобы не зависнуть навсегда
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	log.Printf("start receiving exchange rates on %s", date.Format("02/01/2006"))
	url := fmt.Sprintf("https://www.cbr-xml-daily.ru/daily_eng_utf8.xml?date_req=%s", date.Format("02/01/2006"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := gate.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get a list of currencies on the date %s", date.Format("02/01/2006"))
	}

	defer resp.Body.Close()

	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var cbrRates Rates
	if err = d.Decode(&cbrRates); err != nil {
		return nil, err
	}

	rates := make([]domain.Rate, len(cbrRates.Currencies))
	for _, rate := range cbrRates.Currencies {
		rates = append(rates, domain.Rate{
			Code:     rate.CharCode,
			Original: strings.Replace(rate.Value, ",", ".", 1),
			Nominal:  rate.Nominal,
		})
	}

	return rates, nil
}
