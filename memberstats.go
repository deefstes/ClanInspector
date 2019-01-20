package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type memberStats struct {
	Response        MemberStats `json:"Response"`
	ErrorCode       int         `json:"ErrorCode"`
	ThrottleSeconds int         `json:"ThrottleSeconds"`
	ErrorStatus     string      `json:"ErrorStatus"`
	Message         string      `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type MemberStats struct {
	MergedDeletedCharacters struct {
		Results struct {
		} `json:"results"`
		Merged struct {
		} `json:"merged"`
	} `json:"mergedDeletedCharacters"`
	MergedAllCharacters struct {
		Results struct {
			AllPvE struct {
				AllTime struct {
					ActivitiesCleared             statValue `json:"activitiesCleared"`
					ActivitiesEntered             statValue `json:"activitiesEntered"`
					Assists                       statValue `json:"assists"`
					TotalDeathDistance            statValue `json:"totalDeathDistance"`
					AverageDeathDistance          statValue `json:"averageDeathDistance"`
					TotalKillDistance             statValue `json:"totalKillDistance"`
					Kills                         statValue `json:"kills"`
					AverageKillDistance           statValue `json:"averageKillDistance"`
					SecondsPlayed                 statValue `json:"secondsPlayed"`
					Deaths                        statValue `json:"deaths"`
					AverageLifespan               statValue `json:"averageLifespan"`
					BestSingleGameKills           statValue `json:"bestSingleGameKills"`
					BestSingleGameScore           statValue `json:"bestSingleGameScore"`
					OpponentsDefeated             statValue `json:"opponentsDefeated"`
					Efficiency                    statValue `json:"efficiency"`
					KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
					KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
					ObjectivesCompleted           statValue `json:"objectivesCompleted"`
					PrecisionKills                statValue `json:"precisionKills"`
					ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
					ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
					Score                         statValue `json:"score"`
					HeroicPublicEventsCompleted   statValue `json:"heroicPublicEventsCompleted"`
					AdventuresCompleted           statValue `json:"adventuresCompleted"`
					Suicides                      statValue `json:"suicides"`
					WeaponKillsBow                statValue `json:"weaponKillsBow"`
					WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
					WeaponKillsBeamRifle          statValue `json:"weaponKillsBeamRifle"`
					WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
					WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
					WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
					WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
					WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
					WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
					WeaponKillsSword              statValue `json:"weaponKillsSword"`
					WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
					WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
					WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
					WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
					WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
					WeaponBestType                statValue `json:"weaponBestType"`
					AllParticipantsCount          statValue `json:"allParticipantsCount"`
					AllParticipantsScore          statValue `json:"allParticipantsScore"`
					AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
					LongestKillSpree              statValue `json:"longestKillSpree"`
					LongestSingleLife             statValue `json:"longestSingleLife"`
					MostPrecisionKills            statValue `json:"mostPrecisionKills"`
					OrbsDropped                   statValue `json:"orbsDropped"`
					OrbsGathered                  statValue `json:"orbsGathered"`
					PublicEventsCompleted         statValue `json:"publicEventsCompleted"`
					RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
					TeamScore                     statValue `json:"teamScore"`
					TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
					FastestCompletionMs           statValue `json:"fastestCompletionMs"`
					LongestKillDistance           statValue `json:"longestKillDistance"`
					HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
					HighestLightLevel             statValue `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvE"`
			AllPvP struct {
				AllTime struct {
					ActivitiesEntered             statValue `json:"activitiesEntered"`
					ActivitiesWon                 statValue `json:"activitiesWon"`
					Assists                       statValue `json:"assists"`
					TotalDeathDistance            statValue `json:"totalDeathDistance"`
					AverageDeathDistance          statValue `json:"averageDeathDistance"`
					TotalKillDistance             statValue `json:"totalKillDistance"`
					Kills                         statValue `json:"kills"`
					AverageKillDistance           statValue `json:"averageKillDistance"`
					SecondsPlayed                 statValue `json:"secondsPlayed"`
					Deaths                        statValue `json:"deaths"`
					AverageLifespan               statValue `json:"averageLifespan"`
					Score                         statValue `json:"score"`
					AverageScorePerKill           statValue `json:"averageScorePerKill"`
					AverageScorePerLife           statValue `json:"averageScorePerLife"`
					BestSingleGameKills           statValue `json:"bestSingleGameKills"`
					BestSingleGameScore           statValue `json:"bestSingleGameScore"`
					OpponentsDefeated             statValue `json:"opponentsDefeated"`
					Efficiency                    statValue `json:"efficiency"`
					KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
					KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
					ObjectivesCompleted           statValue `json:"objectivesCompleted"`
					PrecisionKills                statValue `json:"precisionKills"`
					ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
					ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
					Suicides                      statValue `json:"suicides"`
					WeaponKillsBow                statValue `json:"weaponKillsBow"`
					WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
					WeaponKillsBeamRifle          statValue `json:"weaponKillsBeamRifle"`
					WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
					WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
					WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
					WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
					WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
					WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
					WeaponKillsSword              statValue `json:"weaponKillsSword"`
					WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
					WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
					WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
					WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
					WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
					WeaponBestType                statValue `json:"weaponBestType"`
					WinLossRatio                  statValue `json:"winLossRatio"`
					AllParticipantsCount          statValue `json:"allParticipantsCount"`
					AllParticipantsScore          statValue `json:"allParticipantsScore"`
					AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
					LongestKillSpree              statValue `json:"longestKillSpree"`
					LongestSingleLife             statValue `json:"longestSingleLife"`
					MostPrecisionKills            statValue `json:"mostPrecisionKills"`
					OrbsDropped                   statValue `json:"orbsDropped"`
					OrbsGathered                  statValue `json:"orbsGathered"`
					RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
					TeamScore                     statValue `json:"teamScore"`
					TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
					CombatRating                  statValue `json:"combatRating"`
					FastestCompletionMs           statValue `json:"fastestCompletionMs"`
					LongestKillDistance           statValue `json:"longestKillDistance"`
					HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
					HighestLightLevel             statValue `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvP"`
		} `json:"results"`
		Merged struct {
			AllTime struct {
				ActivitiesCleared             statValue `json:"activitiesCleared"`
				ActivitiesEntered             statValue `json:"activitiesEntered"`
				Assists                       statValue `json:"assists"`
				TotalDeathDistance            statValue `json:"totalDeathDistance"`
				AverageDeathDistance          statValue `json:"averageDeathDistance"`
				TotalKillDistance             statValue `json:"totalKillDistance"`
				Kills                         statValue `json:"kills"`
				AverageKillDistance           statValue `json:"averageKillDistance"`
				SecondsPlayed                 statValue `json:"secondsPlayed"`
				Deaths                        statValue `json:"deaths"`
				AverageLifespan               statValue `json:"averageLifespan"`
				BestSingleGameKills           statValue `json:"bestSingleGameKills"`
				BestSingleGameScore           statValue `json:"bestSingleGameScore"`
				OpponentsDefeated             statValue `json:"opponentsDefeated"`
				Efficiency                    statValue `json:"efficiency"`
				KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
				KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
				ObjectivesCompleted           statValue `json:"objectivesCompleted"`
				PrecisionKills                statValue `json:"precisionKills"`
				ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
				ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
				Score                         statValue `json:"score"`
				HeroicPublicEventsCompleted   statValue `json:"heroicPublicEventsCompleted"`
				AdventuresCompleted           statValue `json:"adventuresCompleted"`
				Suicides                      statValue `json:"suicides"`
				WeaponKillsBow                statValue `json:"weaponKillsBow"`
				WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
				WeaponKillsBeamRifle          statValue `json:"weaponKillsBeamRifle"`
				WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
				WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
				WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
				WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
				WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
				WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
				WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
				WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
				WeaponKillsSword              statValue `json:"weaponKillsSword"`
				WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
				WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
				WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
				WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
				WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
				WeaponBestType                statValue `json:"weaponBestType"`
				AllParticipantsCount          statValue `json:"allParticipantsCount"`
				AllParticipantsScore          statValue `json:"allParticipantsScore"`
				AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
				LongestKillSpree              statValue `json:"longestKillSpree"`
				LongestSingleLife             statValue `json:"longestSingleLife"`
				MostPrecisionKills            statValue `json:"mostPrecisionKills"`
				OrbsDropped                   statValue `json:"orbsDropped"`
				OrbsGathered                  statValue `json:"orbsGathered"`
				PublicEventsCompleted         statValue `json:"publicEventsCompleted"`
				RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
				TeamScore                     statValue `json:"teamScore"`
				TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
				FastestCompletionMs           statValue `json:"fastestCompletionMs"`
				LongestKillDistance           statValue `json:"longestKillDistance"`
				HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
				HighestLightLevel             statValue `json:"highestLightLevel"`
				ActivitiesWon                 statValue `json:"activitiesWon"`
				AverageScorePerKill           statValue `json:"averageScorePerKill"`
				AverageScorePerLife           statValue `json:"averageScorePerLife"`
				WinLossRatio                  statValue `json:"winLossRatio"`
				CombatRating                  statValue `json:"combatRating"`
			} `json:"allTime"`
		} `json:"merged"`
	} `json:"mergedAllCharacters"`
	Characters []struct {
		CharacterID string `json:"characterId"`
		Deleted     bool   `json:"deleted"`
		Results     struct {
			AllPvP struct {
				AllTime struct {
					ActivitiesEntered             statValue `json:"activitiesEntered"`
					ActivitiesWon                 statValue `json:"activitiesWon"`
					Assists                       statValue `json:"assists"`
					TotalDeathDistance            statValue `json:"totalDeathDistance"`
					AverageDeathDistance          statValue `json:"averageDeathDistance"`
					TotalKillDistance             statValue `json:"totalKillDistance"`
					Kills                         statValue `json:"kills"`
					AverageKillDistance           statValue `json:"averageKillDistance"`
					SecondsPlayed                 statValue `json:"secondsPlayed"`
					Deaths                        statValue `json:"deaths"`
					AverageLifespan               statValue `json:"averageLifespan"`
					Score                         statValue `json:"score"`
					AverageScorePerKill           statValue `json:"averageScorePerKill"`
					AverageScorePerLife           statValue `json:"averageScorePerLife"`
					BestSingleGameKills           statValue `json:"bestSingleGameKills"`
					BestSingleGameScore           statValue `json:"bestSingleGameScore"`
					OpponentsDefeated             statValue `json:"opponentsDefeated"`
					Efficiency                    statValue `json:"efficiency"`
					KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
					KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
					ObjectivesCompleted           statValue `json:"objectivesCompleted"`
					PrecisionKills                statValue `json:"precisionKills"`
					ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
					ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
					Suicides                      statValue `json:"suicides"`
					WeaponKillsBow                statValue `json:"weaponKillsBow"`
					WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
					WeaponKillsBeamRifle          statValue `json:"weaponKillsBeamRifle"`
					WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
					WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
					WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
					WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
					WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
					WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
					WeaponKillsSword              statValue `json:"weaponKillsSword"`
					WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
					WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
					WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
					WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
					WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
					WeaponBestType                statValue `json:"weaponBestType"`
					WinLossRatio                  statValue `json:"winLossRatio"`
					AllParticipantsCount          statValue `json:"allParticipantsCount"`
					AllParticipantsScore          statValue `json:"allParticipantsScore"`
					AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
					LongestKillSpree              statValue `json:"longestKillSpree"`
					LongestSingleLife             statValue `json:"longestSingleLife"`
					MostPrecisionKills            statValue `json:"mostPrecisionKills"`
					OrbsDropped                   statValue `json:"orbsDropped"`
					OrbsGathered                  statValue `json:"orbsGathered"`
					RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
					TeamScore                     statValue `json:"teamScore"`
					TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
					CombatRating                  statValue `json:"combatRating"`
					FastestCompletionMs           statValue `json:"fastestCompletionMs"`
					LongestKillDistance           statValue `json:"longestKillDistance"`
					HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
					HighestLightLevel             statValue `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvP"`
			AllPvE struct {
				AllTime struct {
					ActivitiesCleared             statValue `json:"activitiesCleared"`
					ActivitiesEntered             statValue `json:"activitiesEntered"`
					Assists                       statValue `json:"assists"`
					TotalDeathDistance            statValue `json:"totalDeathDistance"`
					AverageDeathDistance          statValue `json:"averageDeathDistance"`
					TotalKillDistance             statValue `json:"totalKillDistance"`
					Kills                         statValue `json:"kills"`
					AverageKillDistance           statValue `json:"averageKillDistance"`
					SecondsPlayed                 statValue `json:"secondsPlayed"`
					Deaths                        statValue `json:"deaths"`
					AverageLifespan               statValue `json:"averageLifespan"`
					BestSingleGameKills           statValue `json:"bestSingleGameKills"`
					BestSingleGameScore           statValue `json:"bestSingleGameScore"`
					OpponentsDefeated             statValue `json:"opponentsDefeated"`
					Efficiency                    statValue `json:"efficiency"`
					KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
					KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
					ObjectivesCompleted           statValue `json:"objectivesCompleted"`
					PrecisionKills                statValue `json:"precisionKills"`
					ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
					ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
					Score                         statValue `json:"score"`
					HeroicPublicEventsCompleted   statValue `json:"heroicPublicEventsCompleted"`
					AdventuresCompleted           statValue `json:"adventuresCompleted"`
					Suicides                      statValue `json:"suicides"`
					WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
					WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
					WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
					WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
					WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
					WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
					WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
					WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
					WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
					WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
					WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
					WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
					WeaponKillsSword              statValue `json:"weaponKillsSword"`
					WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
					WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
					WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
					WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
					WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
					WeaponBestType                statValue `json:"weaponBestType"`
					AllParticipantsCount          statValue `json:"allParticipantsCount"`
					AllParticipantsScore          statValue `json:"allParticipantsScore"`
					AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
					LongestKillSpree              statValue `json:"longestKillSpree"`
					LongestSingleLife             statValue `json:"longestSingleLife"`
					MostPrecisionKills            statValue `json:"mostPrecisionKills"`
					OrbsDropped                   statValue `json:"orbsDropped"`
					OrbsGathered                  statValue `json:"orbsGathered"`
					PublicEventsCompleted         statValue `json:"publicEventsCompleted"`
					RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
					TeamScore                     statValue `json:"teamScore"`
					TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
					FastestCompletionMs           statValue `json:"fastestCompletionMs"`
					LongestKillDistance           statValue `json:"longestKillDistance"`
					HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
					HighestLightLevel             statValue `json:"highestLightLevel"`
				} `json:"allTime"`
			} `json:"allPvE"`
		} `json:"results"`
		Merged struct {
			AllTime struct {
				ActivitiesCleared             statValue `json:"activitiesCleared"`
				ActivitiesEntered             statValue `json:"activitiesEntered"`
				Assists                       statValue `json:"assists"`
				TotalDeathDistance            statValue `json:"totalDeathDistance"`
				AverageDeathDistance          statValue `json:"averageDeathDistance"`
				TotalKillDistance             statValue `json:"totalKillDistance"`
				Kills                         statValue `json:"kills"`
				AverageKillDistance           statValue `json:"averageKillDistance"`
				SecondsPlayed                 statValue `json:"secondsPlayed"`
				Deaths                        statValue `json:"deaths"`
				AverageLifespan               statValue `json:"averageLifespan"`
				BestSingleGameKills           statValue `json:"bestSingleGameKills"`
				BestSingleGameScore           statValue `json:"bestSingleGameScore"`
				OpponentsDefeated             statValue `json:"opponentsDefeated"`
				Efficiency                    statValue `json:"efficiency"`
				KillsDeathsRatio              statValue `json:"killsDeathsRatio"`
				KillsDeathsAssists            statValue `json:"killsDeathsAssists"`
				ObjectivesCompleted           statValue `json:"objectivesCompleted"`
				PrecisionKills                statValue `json:"precisionKills"`
				ResurrectionsPerformed        statValue `json:"resurrectionsPerformed"`
				ResurrectionsReceived         statValue `json:"resurrectionsReceived"`
				Score                         statValue `json:"score"`
				HeroicPublicEventsCompleted   statValue `json:"heroicPublicEventsCompleted"`
				AdventuresCompleted           statValue `json:"adventuresCompleted"`
				Suicides                      statValue `json:"suicides"`
				WeaponKillsAutoRifle          statValue `json:"weaponKillsAutoRifle"`
				WeaponKillsFusionRifle        statValue `json:"weaponKillsFusionRifle"`
				WeaponKillsHandCannon         statValue `json:"weaponKillsHandCannon"`
				WeaponKillsTraceRifle         statValue `json:"weaponKillsTraceRifle"`
				WeaponKillsPulseRifle         statValue `json:"weaponKillsPulseRifle"`
				WeaponKillsRocketLauncher     statValue `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle         statValue `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun            statValue `json:"weaponKillsShotgun"`
				WeaponKillsSniper             statValue `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun      statValue `json:"weaponKillsSubmachinegun"`
				WeaponKillsRelic              statValue `json:"weaponKillsRelic"`
				WeaponKillsSideArm            statValue `json:"weaponKillsSideArm"`
				WeaponKillsSword              statValue `json:"weaponKillsSword"`
				WeaponKillsAbility            statValue `json:"weaponKillsAbility"`
				WeaponKillsGrenade            statValue `json:"weaponKillsGrenade"`
				WeaponKillsGrenadeLauncher    statValue `json:"weaponKillsGrenadeLauncher"`
				WeaponKillsSuper              statValue `json:"weaponKillsSuper"`
				WeaponKillsMelee              statValue `json:"weaponKillsMelee"`
				WeaponBestType                statValue `json:"weaponBestType"`
				AllParticipantsCount          statValue `json:"allParticipantsCount"`
				AllParticipantsScore          statValue `json:"allParticipantsScore"`
				AllParticipantsTimePlayed     statValue `json:"allParticipantsTimePlayed"`
				LongestKillSpree              statValue `json:"longestKillSpree"`
				LongestSingleLife             statValue `json:"longestSingleLife"`
				MostPrecisionKills            statValue `json:"mostPrecisionKills"`
				OrbsDropped                   statValue `json:"orbsDropped"`
				OrbsGathered                  statValue `json:"orbsGathered"`
				PublicEventsCompleted         statValue `json:"publicEventsCompleted"`
				RemainingTimeAfterQuitSeconds statValue `json:"remainingTimeAfterQuitSeconds"`
				TeamScore                     statValue `json:"teamScore"`
				TotalActivityDurationSeconds  statValue `json:"totalActivityDurationSeconds"`
				FastestCompletionMs           statValue `json:"fastestCompletionMs"`
				LongestKillDistance           statValue `json:"longestKillDistance"`
				HighestCharacterLevel         statValue `json:"highestCharacterLevel"`
				HighestLightLevel             statValue `json:"highestLightLevel"`
				ActivitiesWon                 statValue `json:"activitiesWon"`
				AverageScorePerKill           statValue `json:"averageScorePerKill"`
				AverageScorePerLife           statValue `json:"averageScorePerLife"`
				WinLossRatio                  statValue `json:"winLossRatio"`
				CombatRating                  statValue `json:"combatRating"`
			} `json:"allTime"`
		} `json:"merged"`
	} `json:"characters"`
}

type characterAggregateStats struct {
	Response        CharacterAgregateStats `json:"Response"`
	ErrorCode       int                    `json:"ErrorCode"`
	ThrottleSeconds int                    `json:"ThrottleSeconds"`
	ErrorStatus     string                 `json:"ErrorStatus"`
	Message         string                 `json:"Message"`
	MessageData     struct {
	} `json:"MessageData"`
}

type CharacterAgregateStats struct {
	Activities []struct {
		ActivityHash int64 `json:"activityHash"`
		Values       struct {
			FastestCompletionMsForActivity statValue `json:"fastestCompletionMsForActivity"`
			ActivityCompletions            statValue `json:"activityCompletions"`
			ActivityDeaths                 statValue `json:"activityDeaths"`
			ActivityKills                  statValue `json:"activityKills"`
			ActivitySecondsPlayed          statValue `json:"activitySecondsPlayed"`
			ActivityWins                   statValue `json:"activityWins"`
			ActivityGoalsMissed            statValue `json:"activityGoalsMissed"`
			ActivitySpecialActions         statValue `json:"activitySpecialActions"`
			ActivityBestGoalsHit           statValue `json:"activityBestGoalsHit"`
			ActivityGoalsHit               statValue `json:"activityGoalsHit"`
			ActivitySpecialScore           statValue `json:"activitySpecialScore"`
			ActivityBestSingleGameScore    statValue `json:"activityBestSingleGameScore"`
			ActivityKillsDeathsRatio       statValue `json:"activityKillsDeathsRatio"`
			ActivityAssists                statValue `json:"activityAssists"`
			ActivityKillsDeathsAssists     statValue `json:"activityKillsDeathsAssists"`
			ActivityPrecisionKills         statValue `json:"activityPrecisionKills"`
		} `json:"values"`
	} `json:"activities"`
}

type PlayerStats struct {
	BatchID      int       `bson:"BatchID"`
	BatchTime    time.Time `bson:"BatchTime"`
	MembershipID string    `bson:"MembershipID"`
	MemberName   string    `bson:"MemberName"`
	PvE          PveStats  `bson:"PvE,omitempty"`
	PvP          PvpStats  `bson:"PvP,omitempty"`
}

type WeaponStats struct {
	AutoRifle       float64 `bson:"AutoRifle,omitempty"`
	BeamRifle       float64 `bson:"BeamRifle,omitempty"`
	Bow             float64 `bson:"Bow,omitempty"`
	FusionRifle     float64 `bson:"FusionRifle,omitempty"`
	HandCannon      float64 `bson:"HandCannon,omitempty"`
	TraceRifle      float64 `bson:"TraceRifle,omitempty"`
	PulseRifle      float64 `bson:"PulseRifle,omitempty"`
	RocketLauncher  float64 `bson:"RocketLauncher,omitempty"`
	ScoutRifle      float64 `bson:"ScoutRifle,omitempty"`
	Shotgun         float64 `bson:"Shotgun,omitempty"`
	Sniper          float64 `bson:"Sniper,omitempty"`
	Submachinegun   float64 `bson:"Submachinegun,omitempty"`
	Relic           float64 `bson:"Relic,omitempty"`
	SideArm         float64 `bson:"SideArm,omitempty"`
	Sword           float64 `bson:"Sword,omitempty"`
	Ability         float64 `bson:"Ability,omitempty"`
	Grenade         float64 `bson:"Grenade,omitempty"`
	GrenadeLauncher float64 `bson:"GrenadeLauncher,omitempty"`
	Super           float64 `bson:"Super,omitempty"`
	Melee           float64 `bson:"Melee,omitempty"`
}

type PveStats struct {
	SecondsPlayed          float64     `bson:"SecondsPlayed,omitempty"`
	Kills                  float64     `bson:"Kills,omitempty"`
	Assists                float64     `bson:"Assists,omitempty"`
	Deaths                 float64     `bson:"Deaths,omitempty"`
	AverageKillDistance    float64     `bson:"AverageKillDistance,omitempty"`
	AverageDeathDistance   float64     `bson:"AverageDeathDistance,omitempty"`
	LongestKillDistance    float64     `bson:"LongestKillDistance,omitempty"`
	KDRatio                float64     `bson:"KDRatio,omitempty"`
	PrecisionKills         float64     `bson:"PrecisionKills,omitempty"`
	ResurrectionsPerformed float64     `bson:"ResurrectionsPerformed,omitempty"`
	Suicides               float64     `bson:"Suicides,omitempty"`
	WeaponKills            WeaponStats `bson:"WeaponKills,omitempty"`
	OrbsDropped            float64     `bson:"OrbsDropped,omitempty"`
	PublicEvents           float64     `bson:"PublicEvents,omitempty"`
	HeroicPublicEvents     float64     `bson:"HeroicPublicEvents,omitempty"`
	Adventures             float64     `bson:"Adventures,omitempty"`
	Raids                  float64     `bson:"Raids,omitempty"`
	PrestigeRaids          float64     `bson:"PrestigeRaids,omitempty"`
	Nightfalls             float64     `bson:"Nightfalls,omitempty"`
	PrestigeNightfalls     float64     `bson:"PrestigeNightfalls,omitempty"`
	NightfallHighScore     float64     `bson:"NightfallHighScore,omitempty"`
	Strikes                float64     `bson:"Strikes,omitempty"`
	HeroicStrikes          float64     `bson:"HeroicStrikes,omitempty"`
}

type PvpStats struct {
	SecondsPlayed          float64     `bson:"SecondsPlayed,omitempty"`
	Kills                  float64     `bson:"Kills,omitempty"`
	Assists                float64     `bson:"Assists,omitempty"`
	Deaths                 float64     `bson:"Deaths,omitempty"`
	AverageKillDistance    float64     `bson:"AverageKillDistance,omitempty"`
	AverageDeathDistance   float64     `bson:"AverageDeathDistance,omitempty"`
	LongestKillDistance    float64     `bson:"LongestKillDistance,omitempty"`
	KDRatio                float64     `bson:"KDRatio,omitempty"`
	PrecisionKills         float64     `bson:"PrecisionKills,omitempty"`
	ResurrectionsPerformed float64     `bson:"ResurrectionsPerformed,omitempty"`
	Suicides               float64     `bson:"Suicides,omitempty"`
	WeaponKills            WeaponStats `bson:"WeaponKills,omitempty"`
	OrbsDropped            float64     `bson:"OrbsDropped,omitempty"`
	BestSingleGameKills    float64     `bson:"BestSingleGameKills,omitempty"`
	LongestKillSpree       float64     `bson:"LongestKillSpree,omitempty"`
}

func GetMemberStats(memberID string) (*MemberStats, error) {
	url := fmt.Sprintf("https://bungie.net/Platform/Destiny2/%s/Account/%s/Stats", config.MembershipType, url.QueryEscape(memberID))
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", config.APIKey)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return &MemberStats{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return &MemberStats{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Body: ", err)
		return &MemberStats{}, err
	}

	// Fill the record with the data from the JSON
	var record memberStats
	if err := json.Unmarshal([]byte(string(body)), &record); err != nil {
		log.Printf("Unmarshal error: %v", err)
		return &MemberStats{}, err
	}

	if record.ErrorCode != 1 {
		return &MemberStats{}, errors.New(record.Message)
	}

	return &record.Response, nil
}
