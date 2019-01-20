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

type characterActivities struct {
	Response struct {
		Activities []struct {
			Period          time.Time       `json:"Period"`
			ActivityDetails activityDetails `json:"ActivityDetails"`
			Values          struct {
				Assists                 statValue `json:"Assists"`
				Score                   statValue `json:"Score"`
				Kills                   statValue `json:"Kills"`
				AverageScorePerKill     statValue `json:"AverageScorePerKill"`
				Deaths                  statValue `json:"Deaths"`
				AverageScorePerLife     statValue `json:"AverageScorePerLife"`
				Completed               statValue `json:"Completed"`
				OpponentsDefeated       statValue `json:"OpponentsDefeated"`
				Efficiency              statValue `json:"Efficiency"`
				KillsDeathsRatio        statValue `json:"KillsDeathsRatio"`
				KillsDeathsAssists      statValue `json:"KillsDeathsAssists"`
				ActivityDurationSeconds statValue `json:"ActivityDurationSeconds"`
				CompletionReason        statValue `json:"CompletionReason"`
				FireteamID              statValue `json:"FireteamID"`
				StartSeconds            statValue `json:"StartSeconds"`
				TimePlayedSeconds       statValue `json:"TimePlayedSeconds"`
				PlayerCount             statValue `json:"PlayerCount"`
				TeamScore               statValue `json:"TeamScore"`
			} `json:"Values"`
		} `json:"Activities"`
	} `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type activityReport struct {
	Response        PGCR   `json:"Response"`
	ErrorCode       int    `json:"ErrorCode"`
	ThrottleSeconds int    `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type PGCR struct {
	Period          time.Time       `json:"period" bson:"Period"`
	ActivityDetails activityDetails `json:"activityDetails" bson:"ActivityDetails"`
	Entries         []struct {
		Standing int `json:"standing" bson:"Standing"`
		Score    struct {
			Basic basicValue `json:"basic" bson:"Basic"`
		} `json:"score" bson:"Score"`
		Player struct {
			DestinyUserInfo   Player `json:"destinyUserInfo" bson:"DestinyUserInfo"`
			CharacterClass    string `json:"characterClass" bson:"CharacterClass"`
			ClassHash         int    `json:"classHash" bson:"ClassHash"`
			RaceHash          int64  `json:"raceHash" bson:"RaceHash"`
			GenderHash        int64  `json:"genderHash" bson:"GenderHash"`
			CharacterLevel    int    `json:"characterLevel" bson:"CharacterLevel"`
			LightLevel        int    `json:"lightLevel" bson:"LightLevel"`
			BungieNetUserInfo struct {
				IconPath       string `json:"iconPath" bson:"IconPath"`
				MembershipType int    `json:"membershipType" bson:"MembershipType"`
				MembershipID   string `json:"membershipId" bson:"MembershipID"`
				DisplayName    string `json:"displayName" bson:"DisplayName"`
			} `json:"bungieNetUserInfo" bson:"BungieNetUserInfo"`
			EmblemHash int64 `json:"emblemHash" bson:"EmblemHash"`
		} `json:"player" bson:"Player"`
		CharacterID string `json:"characterId" bson:"CharacterID"`
		Values      struct {
			Assists struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"assists" bson:"Assists"`
			Completed struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"completed" bson:"Completed"`
			Deaths struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"deaths" bson:"Deaths"`
			Kills struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"kills" bson:"Kills"`
			OpponentsDefeated struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"opponentsDefeated" bson:"OpponentsDefeated"`
			Efficiency struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"efficiency" bson:"Efficiency"`
			KillsDeathsRatio struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"killsDeathsRatio" bson:"KillsDeathsRatio"`
			KillsDeathsAssists struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"killsDeathsAssists" bson:"KillsDeathsAssists"`
			Score struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"score" bson:"Score"`
			ActivityDurationSeconds struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"activityDurationSeconds" bson:"ActivityDurationSeconds"`
			CompletionReason struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"completionReason" bson:"CompletionReason"`
			FireteamID struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"fireteamId" bson:"FireteamID"`
			StartSeconds struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"startSeconds" bson:"StartSeconds"`
			TimePlayedSeconds struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"timePlayedSeconds" bson:"TimePlayedSeconds"`
			PlayerCount struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"playerCount" bson:"PlayerCount"`
			TeamScore struct {
				Basic basicValue `json:"basic" bson:"Basic"`
			} `json:"teamScore" bson:"TeamScore"`
		} `json:"values" bson:"Values"`
		Extended struct {
			Weapons []struct {
				ReferenceID int64 `json:"referenceId" bson:"ReferenceId"`
				Values      struct {
					UniqueWeaponKills struct {
						Basic basicValue `json:"basic" bson:"Basic"`
					} `json:"uniqueWeaponKills" bson:"UniqueWeaponKills"`
					UniqueWeaponPrecisionKills struct {
						Basic basicValue `json:"basic" bson:"Basic"`
					} `json:"uniqueWeaponPrecisionKills" bson:"UniqueWeaponPrecisionKills"`
					UniqueWeaponKillsPrecisionKills struct {
						Basic basicValue `json:"basic" bson:"Basic"`
					} `json:"uniqueWeaponKillsPrecisionKills" bson:"UniqueWeaponKillsPrecisionKills"`
				} `json:"values" bson:"Values"`
			} `json:"weapons" bson:"Weapons"`
			Values struct {
				PrecisionKills struct {
					Basic basicValue `json:"basic" bson:"Basic"`
				} `json:"precisionKills" bson:"PrecisionKills"`
				WeaponKillsGrenade struct {
					Basic basicValue `json:"basic" bson:"Basic"`
				} `json:"weaponKillsGrenade" bson:"WeaponKillsGrenade"`
				WeaponKillsMelee struct {
					Basic basicValue `json:"basic" bson:"Basic"`
				} `json:"weaponKillsMelee" bson:"WeaponKillsMelee"`
				WeaponKillsSuper struct {
					Basic basicValue `json:"basic" bson:"Basic"`
				} `json:"weaponKillsSuper" bson:"WeaponKillsSuper"`
				WeaponKillsAbility struct {
					Basic basicValue `json:"basic" bson:"Basic"`
				} `json:"weaponKillsAbility" bson:"WeaponKillsAbility"`
			} `json:"values" bson:"Values"`
		} `json:"extended" bson:"Extended"`
	} `json:"entries" bson:"Entries"`
	Teams []interface{} `json:"teams" bson:"Teams"`
}

type activityDetails struct {
	ReferenceID          int64  `json:"referenceId" bson:"ReferenceID"`
	DirectorActivityHash int64  `json:"directorActivityHash" bson:"DirectorActivityHash"`
	InstanceID           string `json:"instanceId" bson:"InstanceID"`
	Mode                 int    `json:"mode" bson:"Mode"`
	Modes                []int  `json:"modes" bson:"Modes"`
	IsPrivate            bool   `json:"isPrivate" bson:"IsPrivate"`
}

type statValue struct {
	StatID     string     `json:"statId" bson:"StatID"`
	Basic      basicValue `json:"basic" bson:"Basic"`
	Pga        basicValue `json:"pga,omitempty" bson:"Pga,omitempty"`
	Weighted   basicValue `json:"weighted,omitempty" bson:"Weighted,omitempty"`
	ActivityID string     `json:"activityId,omitempty" bson:"ActivityID,omitempty"`
}

type basicValue struct {
	Value        float64 `json:"value" bson:"Value"`
	DisplayValue string  `json:"displayValue" bson:"DisplayValue"`
}

type hashedActivityDetails struct {
	Response        HashedActivityDetails `json:"Response"`
	ErrorCode       int                   `json:"ErrorCode"`
	ThrottleSeconds int                   `json:"ThrottleSeconds"`
	ErrorStatus     string                `json:"ErrorStatus"`
	Message         string                `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type HashedActivityDetails struct {
	DisplayProperties struct {
		Description string `json:"description" bson:"Description,omitempty"`
		Name        string `json:"name" bson:"Name,omitempty"`
		Icon        string `json:"icon" bson:"Icon,omitempty"`
		HasIcon     bool   `json:"hasIcon" bson:"HasIcon,omitempty"`
	} `json:"displayProperties" bson:"DisplayProperties,omitempty"`
	ReleaseIcon           string        `json:"releaseIcon" bson:"ReleaseIcon,omitempty"`
	ReleaseTime           int           `json:"releaseTime" bson:"ReleaseTime,omitempty"`
	ActivityLevel         int           `json:"activityLevel" bson:"ActivityLevel,omitempty"`
	CompletionUnlockHash  int           `json:"completionUnlockHash" bson:"CompletionUnlockHash,omitempty"`
	ActivityLightLevel    int           `json:"activityLightLevel" bson:"ActivityLightLevel,omitempty"`
	DestinationHash       int64         `json:"destinationHash" bson:"DestinationHash,omitempty"`
	PlaceHash             int64         `json:"placeHash" bson:"PlaceHash,omitempty"`
	ActivityTypeHash      int           `json:"activityTypeHash" bson:"ActivityTypeHash,omitempty"`
	Tier                  int           `json:"tier" bson:"Tier,omitempty"`
	PgcrImage             string        `json:"pgcrImage" bson:"PGCRImage,omitempty"`
	Rewards               []interface{} `json:"rewards" bson:"Rewards,omitempty"`
	Modifiers             []interface{} `json:"modifiers" bson:"Modifiers,omitempty"`
	IsPlaylist            bool          `json:"isPlaylist" bson:"IsPlaylist,omitempty"`
	Challenges            []interface{} `json:"challenges" bson:"Challenges,omitempty"`
	OptionalUnlockStrings []interface{} `json:"optionalUnlockStrings" bson:"OptionalUnlockStrings,omitempty"`
	InheritFromFreeRoam   bool          `json:"inheritFromFreeRoam" bson:"InheritFromFreeRoam,omitempty"`
	SuppressOtherRewards  bool          `json:"suppressOtherRewards" bson:"SuppressOtherRewards,omitempty"`
	PlaylistItems         []interface{} `json:"playlistItems" bson:"PlaylistItems,omitempty"`
	ActivityGraphList     []struct {
		ActivityGraphHash int64 `json:"activityGraphHash" bson:"ActivityGraphHash,omitempty"`
	} `json:"activityGraphList" bson:"ActivityGraphList,omitempty"`
	Matchmaking struct {
		IsMatchmade          bool `json:"isMatchmade" bson:"IsMatchmade,omitempty"`
		MinParty             int  `json:"minParty" bson:"MinParty,omitempty"`
		MaxParty             int  `json:"maxParty" bson:"MaxParty,omitempty"`
		MaxPlayers           int  `json:"maxPlayers" bson:"MaxPlayers,omitempty"`
		RequiresGuardianOath bool `json:"requiresGuardianOath" bson:"RequiresGuardianOath,omitempty"`
	} `json:"matchmaking" bson:"Matchmaking,omitempty"`
	DirectActivityModeHash int   `json:"directActivityModeHash" bson:"DirectActivityModeHash,omitempty"`
	DirectActivityModeType int   `json:"directActivityModeType" bson:"DirectActivityModeType,omitempty"`
	ActivityModeHashes     []int `json:"activityModeHashes" bson:"ActivityModeHashes,omitempty"`
	ActivityModeTypes      []int `json:"activityModeTypes" bson:"ActivityModeTypes,omitempty"`
	IsPvP                  bool  `json:"isPvP" bson:"IsPvP,omitempty"`
	Hash                   int64 `json:"hash" bson:"Hash,omitempty"`
	Index                  int   `json:"index" bson:"Index,omitempty"`
	Redacted               bool  `json:"redacted" bson:"Redacted,omitempty"`
	Blacklisted            bool  `json:"blacklisted" bson:"Blacklisted,omitempty"`
}

func GetActivities(memberID string, characterID string, lastInstanceID string) []PGCR {
	page := 0
	activities := []string{}
	retval := []PGCR{}
	consumed := false

	for !consumed {
		url := fmt.Sprintf("https://bungie.net/Platform/Destiny2/%s/Account/%s/Character/%s/Stats/Activities?count=%d&page=%d", config.MembershipType, url.QueryEscape(memberID), url.QueryEscape(characterID), config.ActivityBatchSize, page)
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
		var record characterActivities
		if err := json.Unmarshal([]byte(string(body)), &record); err != nil {
			log.Printf("Unmarshal error: %v", err)
			return nil
		}

		if len(record.Response.Activities) == 0 {
			consumed = true
			break
		}

		for _, activity := range record.Response.Activities {
			if activity.ActivityDetails.InstanceID == lastInstanceID {
				consumed = true
				break
			}

			activities = append(activities, activity.ActivityDetails.InstanceID)
		}

		page = page + 1
	}

	for cnt, activity := range activities {
		fmt.Printf("Getting activity %s (%d of %d)\r\n", activity, cnt+1, len(activities))
		pgcr, err := GetPGCR(activity)
		if err != nil {
			log.Printf("Error getting PGCR: %v", err)
			return nil
		}
		retval = append(retval, pgcr)
	}

	return retval
}

func GetPGCR(InstanceID string) (PGCR, error) {
	url := fmt.Sprintf("https://bungie.net/Platform/Destiny2/Stats/PostGameCarnageReport/%s", InstanceID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", config.APIKey)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return PGCR{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return PGCR{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Body: ", err)
		return PGCR{}, err
	}

	// Fill the record with the data from the JSON
	var record activityReport
	if err := json.Unmarshal([]byte(string(body)), &record); err != nil {
		log.Printf("Unmarshal error: %v", err)
		return PGCR{}, err
	}

	return record.Response, nil
}
