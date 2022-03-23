package models

type Meta struct {
	Collection    string `json:"collection"`
	Contract      string `json:"contract"`
	ContractType  string `json:"contractType"`
	StartingBlock int64  `json:"startingBlock"`
}
