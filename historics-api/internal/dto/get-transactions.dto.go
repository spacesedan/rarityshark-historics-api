package dto

type GetTransaction struct {
	Collection string `json:"collection" bson:"collection"`
	Limit      int64  `json:"limit" bson:"limit"`
	Page       int64  `json:"page" bson:"page"`
}
