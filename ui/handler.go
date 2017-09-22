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
	From time.Time `json:"from"` //TODO use
	To   time.Time `json:"to"` //TODO use for custom search
	Duration string `json:"duration"`
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
			http.Error(rw, "Failed to read request body!", http.StatusBadRequest)
			return
		}
		m := moments{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			log.Printf("Failed to unmarshal request body! Error: %v\n", err)
			http.Error(rw, "Failed to unmarshal request body!", http.StatusBadRequest)
			return
		}
		currs, err := db.RetrieveCurrencies(m.Duration)

		if err != nil {
			log.Printf("An error found at retrieving currencies! Error: %v\n", err)
			http.NotFound(rw, req)
			return
		}
		currenciesJson, err := json.Marshal(currs)
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
		break
	default:
		http.Error(rw, "Bad request, try POST!", http.StatusBadRequest)
	}
}
