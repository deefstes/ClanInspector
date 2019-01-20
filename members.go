package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type clanMembers struct {
	Response struct {
		Results []struct {
			MemberType        int       `json:"memberType"`
			IsOnline          bool      `json:"isOnline"`
			GroupID           string    `json:"groupId"`
			DestinyUserInfo   Player    `json:"destinyUserInfo"`
			JoinDate          time.Time `json:"joinDate"`
			BungieNetUserInfo struct {
				SupplementalDisplayName string `json:"supplementalDisplayName"`
				IconPath                string `json:"iconPath"`
				MembershipType          int    `json:"membershipType"`
				MembershipID            string `json:"membershipId"`
				DisplayName             string `json:"displayName"`
			} `json:"bungieNetUserInfo,omitempty"`
		} `json:"results"`
		TotalResults int  `json:"totalResults"`
		HasMore      bool `json:"hasMore"`
		Query        struct {
			ItemsPerPage int `json:"itemsPerPage"`
			CurrentPage  int `json:"currentPage"`
		} `json:"query"`
		UseTotalResults bool `json:"useTotalResults"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type Player struct {
	IconPath       string `json:"iconPath" bson:"IconPath"`
	MembershipType int    `json:"membershipType" bson:"MembershipType"`
	MembershipID   string `json:"membershipId" bson:"MembershipID"`
	DisplayName    string `json:"displayName" bson:"DisplayName"`
	Enabled        bool   `json:"enabled" bson:"Enabled"`
}

func GetMembers() []Player {
	url := fmt.Sprintf("https://bungie.net/Platform/GroupV2/%s/Members/", url.QueryEscape(config.ClanID))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", config.APIKey)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Body: ", err)
		return nil
	}

	// Fill the record with the data from the JSON
	var record clanMembers
	if err := json.Unmarshal([]byte(string(body)), &record); err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	players := []Player{}
	for _, member := range record.Response.Results {
		//fmt.Printf("%s\r\n", member.DestinyUserInfo.DisplayName)
		players = append(players, member.DestinyUserInfo)
	}

	return players
}
