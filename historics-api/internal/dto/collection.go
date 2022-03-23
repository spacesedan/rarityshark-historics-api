package dto

type Collection struct {
	Collection string `json:"collection" binding:"required"`
}
