package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	r := gin.Default()
	r.GET("/", GinTask)
	err := r.Run(":5000")

	if err != nil {
		log.Fatal(err)
	}
}

//Combine two existing webservices in a single response to the user using gin
func GinTask(c *gin.Context) {

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

	//We are taking the values as coming from the service without any encode or decode operation
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

	//We are returning json content
	c.JSON(200, gin.H{
		"result": jokeResponse.Value.CurrentJoke,
	})
}
