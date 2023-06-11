package Repository

import (
	"fmt"
	"varzesh3/Domin"
	Unitofwork "varzesh3/UnitOFWork"
)

func GetHistory(id int) []Domin.LeagueHistory {
	db := Unitofwork.GetDb()
	defer db.Close()
	rows, er := db.Query("SELECT id, country_id, period	FROM league_histories where country_id=$1 ;", id)
	defer rows.Close()
	if er != nil {
		fmt.Print(er.Error())
		return nil
	}
	obj := Domin.LeagueHistory{}
	result := make([]Domin.LeagueHistory, 0)
	for rows.Next() {
		er = rows.Scan(&obj.Id, &obj.CountryId, &obj.Period)
		if er != nil && obj.Id > 0 {
			continue
		}
		result = append(result, obj)
	}
	return result
}
