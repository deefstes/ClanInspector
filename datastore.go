package main

import (
	"fmt"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func RetrieveMembers() error {
	// First get a list of all members in db
	collectionMembers := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Members")
	collectionCharacters := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Characters")
	var dbPlayers []Player
	err := collectionMembers.Find(bson.M{"Enabled": true}).All(&dbPlayers)
	if err != nil {
		fmt.Printf("Error reading members: %s\r\n", err.Error())
		return err
	}

	// Get list of Members from API
	apiPlayers := GetMembers()

	// Disable players in DB that are no longer in clan
	for _, player := range dbPlayers {
		if !ContainsMember(apiPlayers, player.MembershipID) {
			fmt.Printf("Disabling member: %s (%s)\r\n", player.DisplayName, player.MembershipID)
			err = collectionMembers.Update(
				bson.M{"MembershipID": player.MembershipID},
				bson.M{"$set": bson.M{"Enabled": false}},
			)
			if err != nil {
				fmt.Printf("Error disabling member: %s\r\n", err.Error())
			}
			err = collectionCharacters.Update(
				bson.M{"MembershipID": player.MembershipID},
				bson.M{"$set": bson.M{"Enabled": false}},
			)
			if err != nil {
				fmt.Printf("Error disabling characters: %s\r\n", err.Error())
			}
		}
	}

	// Upsert members that are not in DB (upsert because a member might already be in the DB but disabled after having left the clan)
	for _, player := range apiPlayers {
		if !ContainsMember(dbPlayers, player.MembershipID) {
			player.Enabled = true
			colQuerier := bson.M{"MembershipID": player.MembershipID}
			_, err = collectionMembers.Upsert(colQuerier, player)
			if err != nil {
				fmt.Printf("Error inserting member: %s\r\n", err.Error())
				return err
			}
			dbPlayers = append(dbPlayers, player)
			fmt.Printf("New Member: %s (%s)\r\n", player.DisplayName, player.MembershipID)
		}

		// Find  all characters for member
		characters := GetCharacters(player.MembershipID)

		// Upsert characters
		//charactersCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Characters")
		for _, character := range characters {
			colQuerier := bson.M{"CharacterID": character.CharacterID}
			record := bson.M{"$set": bson.M{
				"MembershipID":   player.MembershipID,
				"Race":           character.Race,
				"Gender":         character.Gender,
				"Class":          character.Class,
				"DateLastPlayed": character.DateLastPlayed,
				"Enabled":        true,
			}}
			//_, err = charactersCollection.Upsert(colQuerier, record)
			_, err = collectionCharacters.Upsert(colQuerier, record)
			if err != nil {
				fmt.Printf("Error inserting character: %s\r\n", err.Error())
				return err
			}
			fmt.Printf("Character updated: %s %s %s for %s\r\n", Gender(character.Gender), Race(character.Race), Class(character.Class), player.DisplayName)
		}
	}

	return nil
}

func RetrieveActivities() error {
	// Get all characters from DB
	c := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Characters")
	activitiesCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")

	var characters []Character
	err := c.Find(bson.M{"Enabled": true}).All(&characters)
	if err != nil {
		fmt.Printf("Error reading characters: %s\r\n", err.Error())
		return err
	}

	// Iterate through characters and check if last retrieved activity is old enough to warrant retrieving activities
	for cnt, character := range characters {
		if int(character.DateLastPlayed.Sub(character.LastRetrievedDate).Hours()) > config.ActivityAgeCutoff {
			fmt.Printf("Retrieving activities for %s (%s) [%d of %d]\r\n", Class(character.Class), character.CharacterID, cnt, len(characters))
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

func RetrievePlayersStats() error {
	collectionMembers := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Members")
	collectionStats := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("PlayerStats")
	//collectionHashedActivities := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("HashedActivities")

	var dbPlayers []Player
	err := collectionMembers.Find(bson.M{"Enabled": true}).All(&dbPlayers)
	if err != nil {
		fmt.Printf("Error reading players: %s\r\n", err.Error())
		return err
	}

	var dbStats PlayerStats
	err = collectionStats.Find(bson.M{}).Sort("-BatchID").One(&dbStats)
	if err != nil {
		fmt.Printf("Error finding BatchID: %s\r\n", err.Error())
		//return err
	}

	var dbHashedActivities []HashedActivityDetails
	err = collectionStats.Find(bson.M{}).All(&dbHashedActivities)
	if err != nil {
		fmt.Printf("Error reading hashed activities: %s\r\n", err.Error())
		return err
	}

	batchID := dbStats.BatchID + 1
	if batchID == 1 {
		batchID = 106
	}
	fmt.Printf("BatchID: %d\r\n", batchID)

	// Iterate through players
	for cnt, player := range dbPlayers {
		fmt.Printf("Player %d/%d - %s (%s)... ", cnt+1, len(dbPlayers), player.DisplayName, player.MembershipID)
		stats, err := GetMemberStats(player.MembershipID)
		if err != nil {
			fmt.Printf("%s\r\n", err.Error)
		} else {
			newStats := PlayerStats{
				BatchID:      batchID,
				BatchTime:    time.Now(),
				MembershipID: player.MembershipID,
				MemberName:   player.DisplayName,
				PvE: PveStats{
					SecondsPlayed:          stats.MergedAllCharacters.Results.AllPvE.AllTime.SecondsPlayed.Basic.Value,
					Kills:                  stats.MergedAllCharacters.Results.AllPvE.AllTime.Kills.Basic.Value,
					Assists:                stats.MergedAllCharacters.Results.AllPvE.AllTime.Assists.Basic.Value,
					Deaths:                 stats.MergedAllCharacters.Results.AllPvE.AllTime.Deaths.Basic.Value,
					AverageKillDistance:    stats.MergedAllCharacters.Results.AllPvE.AllTime.AverageKillDistance.Basic.Value,
					AverageDeathDistance:   stats.MergedAllCharacters.Results.AllPvE.AllTime.AverageDeathDistance.Basic.Value,
					LongestKillDistance:    stats.MergedAllCharacters.Results.AllPvE.AllTime.LongestKillDistance.Basic.Value,
					KDRatio:                stats.MergedAllCharacters.Results.AllPvE.AllTime.KillsDeathsRatio.Basic.Value,
					PrecisionKills:         stats.MergedAllCharacters.Results.AllPvE.AllTime.PrecisionKills.Basic.Value,
					ResurrectionsPerformed: stats.MergedAllCharacters.Results.AllPvE.AllTime.ResurrectionsPerformed.Basic.Value,
					Suicides:               stats.MergedAllCharacters.Results.AllPvE.AllTime.Suicides.Basic.Value,
					WeaponKills: WeaponStats{
						AutoRifle:       stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsAutoRifle.Basic.Value,
						BeamRifle:       stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsBeamRifle.Basic.Value,
						Bow:             stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsBow.Basic.Value,
						FusionRifle:     stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsFusionRifle.Basic.Value,
						HandCannon:      stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsHandCannon.Basic.Value,
						TraceRifle:      stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsTraceRifle.Basic.Value,
						PulseRifle:      stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsPulseRifle.Basic.Value,
						RocketLauncher:  stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsRocketLauncher.Basic.Value,
						ScoutRifle:      stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsScoutRifle.Basic.Value,
						Shotgun:         stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsShotgun.Basic.Value,
						Sniper:          stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsSniper.Basic.Value,
						Submachinegun:   stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsSubmachinegun.Basic.Value,
						Relic:           stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsRelic.Basic.Value,
						SideArm:         stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsSideArm.Basic.Value,
						Sword:           stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsSword.Basic.Value,
						Ability:         stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsAbility.Basic.Value,
						Grenade:         stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsGrenade.Basic.Value,
						GrenadeLauncher: stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsGrenadeLauncher.Basic.Value,
						Super:           stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsSuper.Basic.Value,
						Melee:           stats.MergedAllCharacters.Results.AllPvE.AllTime.WeaponKillsMelee.Basic.Value,
					},
					OrbsDropped:        stats.MergedAllCharacters.Results.AllPvE.AllTime.OrbsDropped.Basic.Value,
					PublicEvents:       stats.MergedAllCharacters.Results.AllPvE.AllTime.PublicEventsCompleted.Basic.Value,
					HeroicPublicEvents: stats.MergedAllCharacters.Results.AllPvE.AllTime.HeroicPublicEventsCompleted.Basic.Value,
					Adventures:         stats.MergedAllCharacters.Results.AllPvE.AllTime.AdventuresCompleted.Basic.Value,
				},
				PvP: PvpStats{
					SecondsPlayed:          stats.MergedAllCharacters.Results.AllPvP.AllTime.SecondsPlayed.Basic.Value,
					Kills:                  stats.MergedAllCharacters.Results.AllPvP.AllTime.Kills.Basic.Value,
					Assists:                stats.MergedAllCharacters.Results.AllPvP.AllTime.Assists.Basic.Value,
					Deaths:                 stats.MergedAllCharacters.Results.AllPvP.AllTime.Deaths.Basic.Value,
					AverageKillDistance:    stats.MergedAllCharacters.Results.AllPvP.AllTime.AverageKillDistance.Basic.Value,
					AverageDeathDistance:   stats.MergedAllCharacters.Results.AllPvP.AllTime.AverageDeathDistance.Basic.Value,
					LongestKillDistance:    stats.MergedAllCharacters.Results.AllPvP.AllTime.LongestKillDistance.Basic.Value,
					KDRatio:                stats.MergedAllCharacters.Results.AllPvP.AllTime.KillsDeathsRatio.Basic.Value,
					PrecisionKills:         stats.MergedAllCharacters.Results.AllPvP.AllTime.PrecisionKills.Basic.Value,
					ResurrectionsPerformed: stats.MergedAllCharacters.Results.AllPvP.AllTime.ResurrectionsPerformed.Basic.Value,
					Suicides:               stats.MergedAllCharacters.Results.AllPvP.AllTime.Suicides.Basic.Value,
					WeaponKills: WeaponStats{
						AutoRifle:       stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsAutoRifle.Basic.Value,
						BeamRifle:       stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsBeamRifle.Basic.Value,
						Bow:             stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsBow.Basic.Value,
						FusionRifle:     stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsFusionRifle.Basic.Value,
						HandCannon:      stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsHandCannon.Basic.Value,
						TraceRifle:      stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsTraceRifle.Basic.Value,
						PulseRifle:      stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsPulseRifle.Basic.Value,
						RocketLauncher:  stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsRocketLauncher.Basic.Value,
						ScoutRifle:      stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsScoutRifle.Basic.Value,
						Shotgun:         stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsShotgun.Basic.Value,
						Sniper:          stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsSniper.Basic.Value,
						Submachinegun:   stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsSubmachinegun.Basic.Value,
						Relic:           stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsRelic.Basic.Value,
						SideArm:         stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsSideArm.Basic.Value,
						Sword:           stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsSword.Basic.Value,
						Ability:         stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsAbility.Basic.Value,
						Grenade:         stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsGrenade.Basic.Value,
						GrenadeLauncher: stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsGrenadeLauncher.Basic.Value,
						Super:           stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsSuper.Basic.Value,
						Melee:           stats.MergedAllCharacters.Results.AllPvP.AllTime.WeaponKillsMelee.Basic.Value,
					},
					OrbsDropped:         stats.MergedAllCharacters.Results.AllPvP.AllTime.OrbsDropped.Basic.Value,
					BestSingleGameKills: stats.MergedAllCharacters.Results.AllPvP.AllTime.BestSingleGameKills.Basic.Value,
					LongestKillSpree:    stats.MergedAllCharacters.Results.AllPvP.AllTime.LongestKillSpree.Basic.Value,
				},
			}

			err := collectionStats.Insert(newStats)
			if err != nil {
				fmt.Printf("Error inserting stats: %s\r\n", err.Error())
			}

			fmt.Printf("Stats retrieved\r\n")
		}

		//fmt.Println(StructToJSON(stats))
		//fmt.Println()
		//fmt.Println(dbStats)
	}

	return nil
}

func RetrievePlayersAggregateStats() error {
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

func FixActivities() {
	activitiesCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")

	var activities []PGCR
	activitiesFound := true
	totalActivities, err := activitiesCollection.Find(bson.M{"Entries.CharacterID": bson.M{"$exists": false}}).Count()
	if err != nil {
		fmt.Printf("Error counting activities: %s\r\n", err.Error())
		return
	}
	fmt.Printf("%d incorrect activities found\r\n", totalActivities)
	majorCnt := 0
	for activitiesFound {
		err = activitiesCollection.Find(bson.M{"Entries.CharacterID": bson.M{"$exists": false}}).Limit(100).All(&activities)
		if err != nil {
			fmt.Printf("Error reading activities: %s\r\n", err.Error())
			return
		}

		for minorCnt, activity := range activities {
			fmt.Printf("Fixing activity %d of %d (%s)... ", majorCnt*100+minorCnt+1, totalActivities, activity.ActivityDetails.InstanceID)
			err = FixActivity(activity.ActivityDetails.InstanceID)
			if err == nil {
				fmt.Printf("done\r\n")
			} else {
				fmt.Printf("error: %s\r\n", err.Error())
			}
		}

		majorCnt = majorCnt + 1
	}
}

func FixActivity(InstanceID string) error {
	activitiesCollection := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")

	activity, err := GetPGCR(InstanceID)
	if err != nil {
		return err
	}

	colQuerier := bson.M{"ActivityDetails.InstanceID": InstanceID}
	err = activitiesCollection.Update(colQuerier, activity)
	if err != nil {
		return err
	}

	return nil
}

func WhoPlaysWithWho(startDate time.Time, endDate time.Time, postfix string, duplicates bool) {
	f, err := os.Create(fmt.Sprintf("ClanInspector%s_%s.json", config.ClanID, postfix))
	if err != nil {
		fmt.Printf("Error opening file: %s", err.Error())
		return
	}

	// First get a list of all members in db
	collectionMembers := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Members")
	collectionActivities := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")
	var dbPlayers []Player
	err = collectionMembers.Find(bson.M{"Enabled": true}).All(&dbPlayers)
	if err != nil {
		fmt.Printf("Error reading members: %s\r\n", err.Error())
		return
	}

	f.WriteString("{\r\n\t\"nodes\": [\r\n")
	for _, player := range dbPlayers {
		f.WriteString(fmt.Sprintf("\t\t{\"id\": \"%s\", \"group\": 1},\r\n", player.DisplayName))
	}
	f.WriteString("\t],\r\n\t\"links\": [\r\n")

	cnt := 0
	for i1 := 0; i1 < len(dbPlayers); i1++ {
		startPos := i1 + 1
		if duplicates {
			startPos = 0
		}
		for i2 := startPos; i2 < len(dbPlayers); i2++ {
			// db.Activities.find({"Entries.Player.DestinyUserInfo.MembershipID": {$all: [dbPlayers[i1].MembershipID,dbPlayers[i2].MembershipID]}})
			cnt, _ = collectionActivities.Find(
				bson.M{
					"Entries.Player.DestinyUserInfo.MembershipID": bson.M{
						"$all": []interface{}{
							dbPlayers[i1].MembershipID,
							dbPlayers[i2].MembershipID,
						},
					},
					"Period": bson.M{
						"$gt": startDate,
						"$lt": endDate,
					},
				},
			).Count()
			if cnt > 0 {
				fmt.Printf("%s,%s,%d\r\n", dbPlayers[i1].DisplayName, dbPlayers[i2].DisplayName, cnt)
				f.WriteString(fmt.Sprintf("\t\t{\"source\": \"%s\", \"target\": \"%s\", \"value\": %d},\r\n", dbPlayers[i1].DisplayName, dbPlayers[i2].DisplayName, cnt))
			}
		}
	}

	f.WriteString("\t]\r\n}")
}

func WhoPlaysWhen(startDate time.Time, endDate time.Time, postfix string) {
	f, err := os.Create(fmt.Sprintf("ClanInspector%s_%s.tsv", config.ClanID, postfix))
	if err != nil {
		fmt.Printf("Error opening file: %s", err.Error())
		return
	}

	// First get a list of all members in db
	c1 := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Members")
	c2 := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")
	var dbPlayers []Player
	err = c1.Find(bson.M{}).All(&dbPlayers)
	if err != nil {
		fmt.Printf("Error reading members: %s\r\n", err.Error())
		return
	}

	f.WriteString("player\thour\tvalue\r\n")

	for _, player := range dbPlayers {
		if !player.Enabled {
			continue
		}
		var dbActivities []PGCR
		//timeSlots := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var timeSlots [24]int
		c2.Find(
			bson.M{
				"Entries.Player.DestinyUserInfo.MembershipID": player.MembershipID,
				"Period": bson.M{
					"$gt": startDate,
					"$lt": endDate,
				},
			},
		).All(&dbActivities)
		timezone, _ := time.LoadLocation("Europe/London")

		for i, activity := range dbActivities {
			fmt.Printf("Activity %d of %d for %s\r\n", i+1, len(dbActivities), player.DisplayName)
			duration := 0
			for _, participant := range activity.Entries {
				if participant.Player.DestinyUserInfo.MembershipID == player.MembershipID {
					duration = int(participant.Values.TimePlayedSeconds.Basic.Value) / 60
				}
			}
			for slot := activity.Period.In(timezone).Hour(); slot < (activity.Period.In(timezone).Hour() + duration); slot++ {
				timeSlots[slot%24]++
			}
		}

		for i, slot := range timeSlots {
			//if slot > 0 {
			f.WriteString(fmt.Sprintf("%s\t%d\t%d\r\n", player.DisplayName, i, slot))
			//}
		}
	}
}

func TestActivity(InstanceID string) {
	c := mongoSession.DB(fmt.Sprintf("ClanInspector%s", config.ClanID)).C("Activities")
	var activity PGCR
	err := c.Find(bson.M{"ActivityDetails.InstanceID": InstanceID}).One(&activity)
	if err != nil {
		fmt.Printf("Error finding activity: %s\r\n", err.Error())
	}
	timezone, _ := time.LoadLocation("Europe/London")

	fmt.Printf("Activity %s starting hour: %d\r\n", activity.ActivityDetails.InstanceID, activity.Period.In(timezone).Hour())

	for slot := activity.Period.In(timezone).Hour(); slot < (activity.Period.In(timezone).Hour() + 6); slot++ {
		fmt.Printf("%d ", slot%24)
	}
}
