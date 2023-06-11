package Appservice

import (
	"varzesh3/Domin"
	"varzesh3/Repository"
)

func GetHistory(id int) []Domin.LeagueHistory {
	return Repository.GetHistory(id)
}
