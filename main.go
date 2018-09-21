package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	)

const Trello  = "https://api.trello.com/1/"

type Configuration struct {
	ApiKey            string `json:"api_key"`
	ApiToken          string `json:"api_token"`
	DestinationFolder string `json:"destination_folder"`
}

type Organizations struct {
	ID            string `json:"id"`
	DisplayName          string `json:"displayName"`
	BoardsIDs []string `json:"idBoards"`
}

type Boards struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
}


func main() {
	//Reading configuration file
	configBlob, _ := ioutil.ReadFile("./config.json")
	var myConfig Configuration
	json.Unmarshal(configBlob, &myConfig)

	//Getting all organizations of user
	authorization := "?key=" + myConfig.ApiKey + "&token=" + myConfig.ApiToken
	response, _ := http.Get(Trello + "members/me/organizations/" + authorization)
	responseJSON, _ := ioutil.ReadAll(response.Body)
	var myOrgs []Organizations
	json.Unmarshal(responseJSON, &myOrgs)


	fmt.Println(myOrgs[0])

	fmt.Println(myOrgs[0].BoardsIDs[1])
	response1, _ := http.Get(Trello + "boards/" + myOrgs[0].BoardsIDs[1] + authorization)
	responseJSON1, _ := ioutil.ReadAll(response1.Body)
	fmt.Println(responseJSON1)
	var myBoards Boards
	json.Unmarshal(responseJSON1, &myBoards)
	fmt.Println(myBoards)


















}
