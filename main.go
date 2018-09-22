package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const Trello = "https://api.trello.com/1/"

type Configuration struct {
	ApiKey            string `json:"api_key"`
	ApiToken          string `json:"api_token"`
	DestinationFolder string `json:"destination_folder"`
}

type Organizations struct {
	ID   string `json:"id"`
	Name string `json:"displayName"`
}

type BoardsListsCards struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Card struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"desc"`
	Labels       string `json:"idLabels"`
	URL          string `json:"shortUrl"`
	LastActivity string `json:"dateLastActivity"`
}

type Attachments struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"name"`
}

func main() {
	// args --full for full backup --zip for archive backup

	//Reading configuration file and getting authorization keys
	configBlob, _ := ioutil.ReadFile("./config.json")
	var myConfig Configuration
	json.Unmarshal(configBlob, &myConfig)
	echo("Configuration OK", myConfig.ApiToken)
	myAuth := "?key=" + myConfig.ApiKey + "&token=" + myConfig.ApiToken

	//Last backup has found

	//Getting all organizations of user
	originOrganizations := getResponse("members/me/organizations/", myAuth)
	var myOrgs []Organizations
	json.Unmarshal(originOrganizations, &myOrgs)

	//Getting all boards of organization
	originBoards := getResponse("organizations/"+myOrgs[0].ID+"/boards", myAuth)
	var myBoards []BoardsListsCards
	json.Unmarshal(originBoards, &myBoards)

	//Getting all lists of board
	originLists := getResponse("boards/"+myBoards[1].ID+"/lists", myAuth)
	var mylists []BoardsListsCards
	json.Unmarshal(originLists, &mylists)

	//Getting labels of board

	//Getting all cards of list
	originCards := getResponse("list/"+mylists[2].ID+"/cards", myAuth)
	var myCards []BoardsListsCards
	json.Unmarshal(originCards, &myCards)
	//fmt.Println(myCards)

	for i := range myCards {

		//Getting all of card
		originCard := getResponse("cards/"+myCards[i].ID, myAuth)
		var myCard Card
		json.Unmarshal(originCard, &myCard)

		//Getting card attachments
		originAttachments := getResponse("cards/"+myCards[i].ID+"/attachments", myAuth)
		var myAttachments []Attachments
		json.Unmarshal(originAttachments, &myAttachments)

		//TODO incremental backup
		//Making path and filename
		backupFolder := myConfig.DestinationFolder + "Trello 00-00-0000\\"
		boardFolder := myOrgs[0].Name + "__" + myBoards[1].Name + "\\"
		cardFile := mylists[2].Name + "__" + myCards[i].Name
		extension := ".json"
		//Making all needed folders
		fileFolder := backupFolder + boardFolder
		os.MkdirAll(fileFolder, 0644)
		filename := fileFolder + cardFile + extension
		//Converting cards and attachments to text
		backupCard, _ := json.Marshal(myCard)
		backupAttachments, _ := json.Marshal(myAttachments)
		//Adding new line before attachments
		backupAttachmentsNewLine := append([]byte("\r\n"), backupAttachments...)
		backupFile := append(backupCard, backupAttachmentsNewLine...)
		//Writting file of card
		echo(filename, filename)
		ioutil.WriteFile(filename, backupFile, 0644)


	}

	//TODO zipping backup
}

func getResponse(request string, authorization string) []byte {
	response, err := http.Get(Trello + request + authorization)
	check(err)
	responseJSON, err := ioutil.ReadAll(response.Body)
	check(err)
	return responseJSON
}

//Simple print to CLI
func echo(message string, condition string) {
	if condition != "" {
		fmt.Println(message)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
}
