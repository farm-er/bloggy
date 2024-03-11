package request

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"tester/dummydata"
)

var url string = "http://localhost:4242/"

func SignUp() {

	user := dummydata.UserGenerator()

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(user)

	if err != nil {
		log.Fatal("error encoding data")
	}

	r, err := http.NewRequest("POST", url+"signup", &buf)

	if err != nil {
		log.Fatal("error in posting the request")
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
