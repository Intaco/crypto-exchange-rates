package updater

import (
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"time"
	"github.com/intaco/crypto-exchange-rates/db"
)

type Answer struct {
	a string
}

func InitUpdater() {
	go actualize(time.Now()) //1st run initially
	actualizePeriodically() //later every N minutes
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

func actualize(t time.Time) (Currencies, error) {
	log.Println("actulizing currencies...")
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
		log.Fatalf("Could not access website %s. Error: %s", CRYPTOCOMPAREURL, err)
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
	log.Println("actulizing successful...")
	db.Temp()

	return currencies, nil
}

const INTERVAL_PERIOD time.Duration = 3 * time.Minute

func actualizePeriodically(){
    for t := range time.NewTicker(INTERVAL_PERIOD).C {
        go actualize(t)
    }
}


