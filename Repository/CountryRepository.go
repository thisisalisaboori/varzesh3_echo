package Repository

import (
	"fmt"
	"varzesh3/Domin"
	Unitofwork "varzesh3/UnitOFWork"
)

func GetAllCountries() []Domin.Country {
	db := Unitofwork.GetDb()
	defer db.Close()
	rows, er := db.Query("SELECT id, leage_title, varzesh3id FROM countries; ")
	defer rows.Close()
	if er != nil {
		return nil
	}
	obj := Domin.Country{}
	result := make([]Domin.Country, 0)
	for rows.Next() {

		er = rows.Scan(&obj.Id, &obj.Title, &obj.Varzesh3Id)
		if er != nil && obj.Id > 0 {
			fmt.Println(er.Error())
		} else {
			result = append(result, obj)
		}

	}

	return result

}
