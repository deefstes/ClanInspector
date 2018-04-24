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

type memberChars struct {
	Response struct {
		Characters struct {
			Data    map[int64]character `json:"data"`
			Privacy int                 `json:"privacy"`
		} `json:"characters"`
		ItemComponents struct {
		} `json:"itemComponents"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type character struct {
	MembershipID             string        `json:"membershipId"`
	MembershipType           int           `json:"membershipType"`
	CharacterID              string        `json:"characterId"`
	DateLastPlayed           time.Time     `json:"dateLastPlayed"`
	MinutesPlayedThisSession string        `json:"minutesPlayedThisSession"`
	MinutesPlayedTotal       string        `json:"minutesPlayedTotal"`
	Light                    int           `json:"light"`
	Stats                    map[int64]int `json:"stats"`
	RaceHash                 int64         `json:"raceHash"`
	GenderHash               int64         `json:"genderHash"`
	ClassHash                int64         `json:"classHash"`
	RaceType                 int           `json:"raceType"`
	ClassType                int           `json:"classType"`
	GenderType               int           `json:"genderType"`
	EmblemPath               string        `json:"emblemPath"`
	EmblemBackgroundPath     string        `json:"emblemBackgroundPath"`
	EmblemHash               int64         `json:"emblemHash"`
	EmblemColor              struct {
		Red   int `json:"red"`
		Green int `json:"green"`
		Blue  int `json:"blue"`
		Alpha int `json:"alpha"`
	} `json:"emblemColor"`
	LevelProgression struct {
		ProgressionHash     int `json:"progressionHash"`
		DailyProgress       int `json:"dailyProgress"`
		DailyLimit          int `json:"dailyLimit"`
		WeeklyProgress      int `json:"weeklyProgress"`
		WeeklyLimit         int `json:"weeklyLimit"`
		CurrentProgress     int `json:"currentProgress"`
		Level               int `json:"level"`
		LevelCap            int `json:"levelCap"`
		StepIndex           int `json:"stepIndex"`
		ProgressToNextLevel int `json:"progressToNextLevel"`
		NextLevelAt         int `json:"nextLevelAt"`
	} `json:"levelProgression"`
	BaseCharacterLevel int     `json:"baseCharacterLevel"`
	PercentToNextLevel float32 `json:"percentToNextLevel"`
}

type Character struct {
	MembershipID          string    `json:"MembershipID" bson:"MembershipID"`
	CharacterID           string    `json:"CharacterID" bson:"CharacterID"`
	Race                  int       `json:"Race" bson:"Race"`
	Gender                int       `json:"Gender" bson:"Gender"`
	Class                 int       `json:"Class" bson:"Class"`
	LastRetrievedActivity string    `json:"LastRetrievedActivity" bson:"LastRetrievedActivity"`
	LastRetrievedDate     time.Time `json:"LastRetrievedDate" bson:"LastRetrievedDate"`
	DateLastPlayed        time.Time `json:"DateLastPlayed" bson:"DateLastPlayed"`
}

func GetCharacters(memberID string) []Character {
	url := fmt.Sprintf("https://bungie.net/Platform/Destiny2/%s/Profile/%s?components=200", config.MembershipType, url.QueryEscape(memberID))
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
	var record memberChars
	if err := json.Unmarshal([]byte(string(body)), &record); err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	characters := []Character{}
	for _, character := range record.Response.Characters.Data {
		//fmt.Printf("%s\r\n", member.DestinyUserInfo.DisplayName)
		characters = append(
			characters,
			Character{
				MembershipID:   memberID,
				CharacterID:    character.CharacterID,
				Race:           character.RaceType,
				Gender:         character.GenderType,
				Class:          character.ClassType,
				DateLastPlayed: character.DateLastPlayed,
			},
		)
	}

	return characters
}

func Race(raceID int) string {
	switch raceID {
	case 0:
		return "Human"
	case 1:
		return "Awoken"
	case 2:
		return "Exo"
	}

	return "Unknown Race"
}

func Gender(genderID int) string {
	switch genderID {
	case 0:
		return "Male"
	case 1:
		return "Female"
	}

	return "Genderless"
}

func Class(classID int) string {
	switch classID {
	case 0:
		return "Titan"
	case 1:
		return "Hunter"
	case 2:
		return "Warlock"
	}

	return "Unknown Class"
}
