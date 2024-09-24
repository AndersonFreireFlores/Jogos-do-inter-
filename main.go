package main

import (
	"net/http"
)

func main() {

	url := "https://api.football-data.org/v4/teams/6684/matches?status=SCHEDULED"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Auth-Token", "502cfc9d9bb54dcf88cfd8b78dd5344b")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
