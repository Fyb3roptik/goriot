package goriot

import (
	"fmt"
)

//League represents a League of Legends league
type League struct {
	Entries       []LeagueItem `json:"entries"`
	Name          string       `json:"name"`
	ParticipantId string       `json:"participantId"`
	Queue         string       `json:"queue"`
	Tier          string       `json:"tier"`
}

//LeagueItem is an entry in a League. It represents a player or team
type LeagueItem struct {
	IsFreshBlood     bool       `json:"isFreshBlood"`
	IsHotStreak      bool       `json:"isHotStreak"`
	IsInactive       bool       `json:"isInactive"`
	IsVeteran        bool       `json:"isVeteran"`
	LastPlayed       int64      `json:"lastPlayed"`
	LeagueName       string     `json:"leagueName"`
	LeaguePoints     int        `json:"leaguePoints"`
	MiniSeries       MiniSeries `json:"miniSeries"`
	PlayerOrTeamID   string     `json:"playerOrTeamId"`
	PlayerOrTeamName string     `json:"playerOrTeamName"`
	QueueType        string     `json:"queueType"`
	Rank             string     `json:"rank"`
	Tier             string     `json:"tier"`
	Wins             int        `json:"wins"`
}

//MiniSeries shows if a player is in their Series and how far they are within it
type MiniSeries struct {
	Losses               int    `json:"losses"`
	Progress             string `json:"progress"`
	Target               int    `json:"target"`
	TimeLeftToPlayMillis int64  `json:"timeLeftToPlayMillis"`
	Wins                 int    `json:"wins"`
}

//LeagueBySummoner retrieves the league of the supplied summonerID from Riot Games API.
//It returns a League and any errors that occured from the server
//The global API key must be set before use
func LeagueBySummoner(region string, summonerID int64) (league []League, err error) {
	if !IsKeySet() {
		return league, ErrAPIKeyNotSet
	}
	args := "api_key=" + apikey
	url := fmt.Sprintf("%v/lol/%v/v2.3/league/by-summoner/%d?%v", BaseURL, region, summonerID, args)
	err = requestAndUnmarshal(url, &league)
	if err != nil {
		return
	}
	return league, err
}
