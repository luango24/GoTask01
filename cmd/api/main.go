package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

type Joke struct {
	ID          string   `json:"id"`
	CurrentJoke string   `json:"joke"`
	Categories  []string `json:"categories"`
}

func main() {
	http.HandleFunc("/", Task)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//Combine two existing webservices in a single response to the user
func Task(w http.ResponseWriter, r *http.Request) {
	response01, err := http.Get("https://names.mcquay.me/api/v0/")

	if err != nil {
		log.Fatal(err)
	}

	var user User
	byteUser, err := ioutil.ReadAll(response01.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteUser, &user)

	response02, err := http.Get("http://api.icndb.com/jokes/random?firstName=" + user.Name + "&lastName=" + user.LastName + "&limitTo=nerdy")

	if err != nil {
		log.Fatal(err)
	}

	byteJokeResponse, err := ioutil.ReadAll(response02.Body)

	if err != nil {
		log.Fatal(err)
	}

	var jokeResponse JokeResponse
	json.Unmarshal(byteJokeResponse, &jokeResponse)

	fmt.Fprint(w, jokeResponse.Value.CurrentJoke)
}
