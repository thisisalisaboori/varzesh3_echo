package Appservice

import (
	Dto "varzesh3/Dto"
	"varzesh3/Repository"
)

func GetAllCountries() []Dto.CountryDto {
	dResults := Repository.GetAllCountries()
	result := make([]Dto.CountryDto, 0)
	for _, item := range dResults {
		//fmt.Printf("len=%d\t%d\t%s\t%d\n", len(dResults), item.Id, item.Title, item.Varzesh3Id)
		result = append(result, item.ToDto())
	}
	return result
}
