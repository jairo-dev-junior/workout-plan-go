package client

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var url string = "https://exercisedb.p.rapidapi.com/exercises/bodyPart/"
var RAPID_API_KEY_HEADER = "x-rapidapi-key"
var RAPID_API_HOST_HEADER = "x-rapidapi-host"

type Exercise struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	BodyPart         string   `json:"bodyPart"`
	Equipment        string   `json:"equipment"`
	Target           string   `json:"target"`
	SecondaryMuscles []string `json:"secondaryMuscles"`
	Instructions     []string `json:"instructions"`
	GifUrl           string   `json:"gifUrl"`
}

func getResponse(method string, url string) []byte {
	url += "?limit=4&offset=0"

	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add(RAPID_API_KEY_HEADER, os.Getenv("RAPID_API_KEY"))
	req.Header.Add(RAPID_API_HOST_HEADER, os.Getenv("RAPID_API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body

}

func GetExercise(bodyPart string) []Exercise {
	getUrl := url + bodyPart
	var response = []Exercise{}
	json.Unmarshal(getResponse("GET", getUrl), &response)
	return response
}
