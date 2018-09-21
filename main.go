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
	//Reading configuration file and getting authorization keys
	configBlob, _ := ioutil.ReadFile("./config.json")
	var myConfig Configuration
	json.Unmarshal(configBlob, &myConfig)
	myAuth := "?key=" + myConfig.ApiKey + "&token=" + myConfig.ApiToken

	//Getting all organizations of user
	originOrganizations := getResponse("members/me/organizations/", myAuth)
	var myOrgs []Organizations
	json.Unmarshal(originOrganizations, &myOrgs)


	fmt.Println(myOrgs[0])

	fmt.Println(myOrgs[0].BoardsIDs[1])

	//Getting all boards of organization
	originBoards := getResponse("boards/" + myOrgs[0].BoardsIDs[1], myAuth)
	var myBoards Boards
	json.Unmarshal(originBoards, &myBoards)
	fmt.Println(myBoards)














}

func getResponse(request string, authorization string) []byte {
	response, _ := http.Get(Trello + request + authorization)

	responseJSON, _ := ioutil.ReadAll(response.Body)

	return responseJSON
}
