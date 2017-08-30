package updater

import (
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"time"
)

type Answer struct {
	a string
}

func ABC() {
	actualize()
}
type Currencies struct {
	BTC Price 
	ETH Price
	LTC Price
	XMR Price
	ETC Price
	DASH Price
	MAID Price
	REP Price
	XEM Price
}
type Price struct {
	USD float32
	EUR float32
	RUR float32
}

const CRYPTOCOMPAREURL = "https://min-api.cryptocompare.com/data/pricemulti"

func actualize() (Currencies, error) {
	var currencies Currencies
	
	u, err := url.Parse(CRYPTOCOMPAREURL)
	if err != nil {
		log.Fatalf("Failed to parse CSE base url! Error: %s", err)
		return currencies, err
	}
	q := u.Query()
	q.Add("fsyms", "BTC,ETH,LTC,XMR,ETC,DASH,MAID,REP,XEM")
	q.Add("tsyms", "USD,EUR,RUR")
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("Could not access website")
		return currencies, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read answer! Error: %s", err)
		return currencies, err
	
	}
	err = json.Unmarshal(body, &currencies)
		if err != nil {
		log.Fatalf("Failed to unmarshal response body! Error: %s", err)
		return currencies, err
	
	}
	
	return currencies, nil
}

const INTERVAL_PERIOD time.Duration = 10 * time.Minute

