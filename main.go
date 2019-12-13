package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	accessToken := getAPIToken()
	totalIssuesAcrossTeam := 0

	file, err := os.Open("./repositories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		repoName := scanner.Text()
		issues := getIssues(accessToken, repoName)

		issuesLength := issues.Size
		totalIssuesAcrossTeam = totalIssuesAcrossTeam + issuesLength
		fmt.Printf("%v: %v\n", repoName, issuesLength)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total issues across team:", totalIssuesAcrossTeam)
}
