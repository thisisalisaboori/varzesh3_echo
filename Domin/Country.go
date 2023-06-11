package Domin

import "varzesh3/Dto"

type Country struct {
	Id         int    `json:"id"`
	Title      string `json:"leage_title"`
	Varzesh3Id int    `json:"varzesh3id"`
}

func (T Country) ToDto() Dto.CountryDto {
	newItem := Dto.CountryDto{Title: T.Title, Varzesh3Id: T.Varzesh3Id, Id: T.Id}
	return newItem
}
