package Repository

import (
	"varzesh3/Domin"
	Unitofwork "varzesh3/UnitOFWork"
)

func GetLeagueResult(id int) []Domin.LeagueDetail {
	db := Unitofwork.GetDb()
	defer db.Close()
	rows, er := db.Query("SELECT id, av, eq, league_id, loss, play, point, rank, team, win"+
		" FROM league_details where league_id=$1 order by rank;", id)
	defer rows.Close()
	if er != nil {
		return nil
	}
	result := make([]Domin.LeagueDetail, 0)
	obj := Domin.LeagueDetail{}
	for rows.Next() {
		er = rows.Scan(&obj.Id, &obj.Av, &obj.Draw, &obj.LeagueId, &obj.Loss,
			&obj.Play, &obj.Point, &obj.Rank, &obj.Team, &obj.Win)
		if er != nil {
			return nil
		}
		result = append(result, obj)
	}
	return result
}
