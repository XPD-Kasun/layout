package models

type Product struct {
	Id          int
	Img         string
	Name        string
	Description string
	Price       float64
	Features    []string
}
