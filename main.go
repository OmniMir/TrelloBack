package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Trello = "https://api.trello.com/1/"

type Configuration struct {
	ApiKey            string `json:"api_key"`
	ApiToken          string `json:"api_token"`
	DestinationFolder string `json:"destination_folder"`
}

type Organizations struct {
	ID          string   `json:"id"`
	Name string   `json:"displayName"`
}

type Boards struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Lists struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Cards struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Card struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Description string `json:"desc"`
	Labels string `json:"idLabels"`
	URL string `json:"shortUrl"`
	LastActivity string `json:"dateLastActivity"`

}


func main() {
	//Reading configuration file and getting authorization keys
	configBlob, _ := ioutil.ReadFile("./config.json")
	var myConfig Configuration
	json.Unmarshal(configBlob, &myConfig)
	echo("Configuration OK", myConfig.ApiToken)
	myAuth := "?key=" + myConfig.ApiKey + "&token=" + myConfig.ApiToken

	//Getting all organizations of user
	originOrganizations := getResponse("members/me/organizations/", myAuth)
	var myOrgs []Organizations
	json.Unmarshal(originOrganizations, &myOrgs)
	fmt.Println(myOrgs)
	fmt.Println("")



	//Getting all boards of organization
	originBoards := getResponse("organizations/" + myOrgs[0].ID + "/boards", myAuth)
	var myBoards []Boards
	json.Unmarshal(originBoards, &myBoards)
	fmt.Println(myBoards)
	fmt.Println("")

	//Getting all lists of board
	originLists := getResponse("boards/" + myBoards[1].ID + "/lists", myAuth)
	var mylists []Lists
	json.Unmarshal(originLists, &mylists)
	fmt.Println(mylists)
	fmt.Println("")

	//Getting all cards of list
	originCards := getResponse("list/" + mylists[2].ID +"/cards", myAuth)
	var myCards []Cards
	json.Unmarshal(originCards, &myCards)
	//fmt.Println(myCards)
	fmt.Println(myCards)
	fmt.Println("")

	//Getting all of card

	originCard := getResponse("cards/" + myCards[3].ID, myAuth)
	var myCard Card
	json.Unmarshal(originCard, &myCard)
	fmt.Println(myCard)
	fmt.Println("")

}

func getResponse(request string, authorization string) []byte {
	response, err := http.Get(Trello + request + authorization)
	echoErr(err)
	responseJSON, err := ioutil.ReadAll(response.Body)
	echoErr(err)
	return responseJSON
}

//Simple print to CLI
func echo(message string, condition string) {
	if condition != "" {
		fmt.Println(message)
	}
}

func echoErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
