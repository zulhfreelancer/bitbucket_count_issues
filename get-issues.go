package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func getIssues(token string, repoName string) issuesResponse {
	teamName := viper.GetString("TEAM_NAME")
	url := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%v/%v/issues", teamName, repoName)
	req, _ := http.NewRequest("GET", url, nil)

	accessToken := fmt.Sprintf("Bearer %v", token)
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.bitbucket.org")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var ir issuesResponse
	err := json.NewDecoder(res.Body).Decode(&ir)
	if err != nil {
		log.Fatalf("Error decoding API response: %v", err)
	}

	return ir
}
