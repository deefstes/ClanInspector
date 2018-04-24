package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func RetrieveMembers() error {
	// First get a list of all members in db
	c := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Members")
	var dbPlayers []Player
	err := c.Find(bson.M{}).All(&dbPlayers)
	if err != nil {
		fmt.Printf("Error reading members: %s\r\n", err.Error())
		return err
	}

	// Get list of Members from API
	apiPlayers := GetMembers()

	// Add members that are not in DB
	for _, player := range apiPlayers {
		if !ContainsMember(dbPlayers, player.MembershipID) {
			err = c.Insert(player)
			if err != nil {
				fmt.Printf("Error inserting member: %s\r\n", err.Error())
				return err
			}
			dbPlayers = append(dbPlayers, player)
			fmt.Printf("New Member: %s (%s)\r\n", player.DisplayName, player.MembershipID)
		}

		// Find all characters for member
		characters := GetCharacters(player.MembershipID)

		// Upsert characters
		charactersCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Characters")
		for _, character := range characters {
			colQuerier := bson.M{"CharacterID": character.CharacterID}
			record := bson.M{"$set": bson.M{
				"MembershipID":   player.MembershipID,
				"Race":           character.Race,
				"Gender":         character.Gender,
				"Class":          character.Class,
				"DateLastPlayed": character.DateLastPlayed,
			}}
			_, err = charactersCollection.Upsert(colQuerier, record)
			if err != nil {
				fmt.Printf("Error inserting character: %s\r\n", err.Error())
				return err
			}
			fmt.Printf("Character updated: %s %s %s for %s\r\n", Gender(character.Gender), Race(character.Race), Class(character.Class), player.DisplayName)
		}
	}

	// TODO: Remove/disable members in db that are no longer in clan

	return nil
}

func RetrieveActivities() error {
	// Get all characters from DB
	c := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Characters")
	activitiesCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")

	var characters []Character
	err := c.Find(bson.M{}).All(&characters)
	if err != nil {
		fmt.Printf("Error reading characters: %s\r\n", err.Error())
		return err
	}

	// Iterate through characters and check if last retrieved activity is old enough to warrant retrieving activities
	for _, character := range characters {
		if int(character.DateLastPlayed.Sub(character.LastRetrievedDate).Hours()) > config.ActivityAgeCutoff {
			fmt.Printf("Retrieving activities for %s (%s)\r\n", Class(character.Class), character.CharacterID)
			activities := GetActivities(character.MembershipID, character.CharacterID, character.LastRetrievedActivity)
			lastActivityID := character.LastRetrievedActivity
			lastActivityDate := character.LastRetrievedDate
			if lastActivityID == "" && len(activities) > 0 {
				lastActivityID = activities[0].ActivityDetails.InstanceID
			}
			fmt.Printf("Found total of %d\r\n", len(activities))

			// Now iterate through activities and insert into DB if not already inserted
			retrievedCnt := 0
			for _, activity := range activities {
				err = activitiesCollection.Insert(activity)
				if err != nil {
					if !mgo.IsDup(err) {
						fmt.Printf("Error inserting activity: %s\r\n", err.Error())
						return err
					} else {
						retrievedCnt = retrievedCnt + 1
						fmt.Printf("Activity %s retrieved\r\n", activity.ActivityDetails.InstanceID)
					}
				}
				if activity.Period.After(lastActivityDate) {
					lastActivityDate = activity.Period
					lastActivityID = activity.ActivityDetails.InstanceID
				}
			}

			if retrievedCnt == 0 {
				lastActivityDate = character.DateLastPlayed
			}

			fmt.Printf("Retrieved total of %d\r\n", retrievedCnt)

			colQuerier := bson.M{"CharacterID": character.CharacterID}
			record := bson.M{"$set": bson.M{
				"LastRetrievedActivity": lastActivityID,
				"LastRetrievedDate":     lastActivityDate,
			}}
			err = c.Update(colQuerier, record)
			if err != nil {
				fmt.Printf("Error updating character: %s\r\n", err.Error())
				//return err
			}
		} else {
			fmt.Printf("No new activities for %s (%s)\r\n", Class(character.Class), character.CharacterID)
		}
	}
	return nil
}

func ContainsMember(baselist []Player, membershipID string) bool {
	for _, player := range baselist {
		if player.MembershipID == membershipID {
			return true
		}
	}

	return false
}

func ContainsCharacter(baselist []Character, characterID string) bool {
	for _, character := range baselist {
		if character.CharacterID == characterID {
			return true
		}
	}

	return false
}

func FixActivity(InstanceID string) {
	activitiesCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")

	var activity PGCR
	err := activitiesCollection.Find(bson.M{"ActivityDetails.InstanceID": InstanceID}).One(&activity)
	if err != nil {
		fmt.Printf("Error reading activity: %s\r\n", err.Error())
		return
	}

	activity, err = GetPGCR(InstanceID)
	if err != nil {
		fmt.Printf("Error getting PGCR: %v", err)
		return
	}

	colQuerier := bson.M{"ActivityDetails.InstanceID": InstanceID}
	err = activitiesCollection.Update(colQuerier, activity)
	if err != nil {
		fmt.Printf("Error updating activity: %s\r\n", err.Error())
		return
	}
}
