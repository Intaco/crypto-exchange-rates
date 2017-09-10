package updater

import (
	"encoding/json"
	"github.com/intaco/crypto-exchange-rates/db"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func InitUpdater() {
	go actualize()          //1st run initially
	actualizePeriodically() //later every 5 minutes
}

const CRYPTOCOMPAREURL = "https://min-api.cryptocompare.com/data/pricemulti"

func actualize() (db.Currencies, error) {
	log.Println("actulizing currencies started...")
	currencies := db.Currencies{}
	currencies.InitTypes()
	u, err := url.Parse(CRYPTOCOMPAREURL)
	if err != nil {
		log.Printf("Failed to parse CryptoCompare base url! Error: %v", err)
		return currencies, err
	}
	q := u.Query()
	q.Add("fsyms", "BTC,ETH,LTC,XMR,ETC,DASH,MAID,REP,XEM")
	q.Add("tsyms", "USD,EUR,RUR")
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("Could not access website %s. Error: %v", CRYPTOCOMPAREURL, err)
		return currencies, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read answer! Error: %v", err)
		return currencies, err

	}
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		log.Printf("Failed to unmarshal response body! Error: %v", err)
		return currencies, err

	}
	log.Println("actulizing successful...")
	db.WriteCurrencies(currencies)

	return currencies, nil
}

const INTERVAL_PERIOD time.Duration = 5 * time.Minute

func actualizePeriodically() {
	for range time.NewTicker(INTERVAL_PERIOD).C {
		go actualize()
	}
}
