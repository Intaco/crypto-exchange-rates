package ui

import (
	"net/http"
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
	"github.com/intaco/crypto-exchange-rates/db"
)

type moments struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func Serve() {
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/", fs)
	http.HandleFunc("/api/data", handleMomentRequests)

	http.ListenAndServe(":3000", nil)

	log.Println("Listening to port 3000...")

}

func handleMomentRequests(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	switch req.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Failed to read request body! Error: %v\n", err)
			http.Error(rw, "", http.StatusBadRequest)
			return
		}
		m := moments{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			log.Printf("Failed to unmarshal request body! Error: %v\n", err)
			http.Error(rw, "", http.StatusBadRequest)
			return
		}
		currencies := db.RetrieveCurrencies(m.From, m.To)
		if len(currencies) == 0 {
			http.NotFound(rw, req)
			return
		}
		currenciesJson, err := json.Marshal(currencies)
		if err != nil {
			log.Printf("Failed to marshal retrieved currencies! Error: %v\n", err)
			http.Error(rw, "", http.StatusInternalServerError)
			return
		}
		status, err := rw.Write(currenciesJson)
		if err != nil {
			log.Printf("Failed to send currencies answer! Status: %v Error: %v\n", status, err)
			http.Error(rw, "", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(rw, "Bad request, try POST!", http.StatusBadRequest)
	}
}
