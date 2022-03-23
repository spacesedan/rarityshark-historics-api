package dto

type GetHistorics struct {
	Collection string `json:"collection" bson:"collection"`
	Limit      int64  `json:"limit" bson:"limit"`
}
