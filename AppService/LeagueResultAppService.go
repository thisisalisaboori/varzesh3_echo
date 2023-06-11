package Appservice

import (
	"varzesh3/Domin"
	"varzesh3/Repository"
)

func GetLeagueResult(id int) []Domin.LeagueDetail {
	return Repository.GetLeagueResult(id)
}
