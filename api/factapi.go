package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// json blueprint for a cat fact
type CatFact struct {
	Fact string `json:"fact"`
	Time string `json:"time"`
}

// json blueprint for an array of cat facts
type Facts struct {
	Facts []CatFact `json:"facts"`
}

// Obtains a cat fact using the API url
func GetCatFact(client *http.Client, url string) (CatFact, error) {
	catFact := new(CatFact)
	response, getErr := client.Get(url)
	if getErr != nil {
		return *catFact, getErr
	}

	defer response.Body.Close()
	catFact.Time = time.Now().Format(time.RFC1123) // time format: day, date month year hh:mm:ss
	decodeErr := json.NewDecoder(response.Body).Decode(catFact)

	return *catFact, decodeErr
}
