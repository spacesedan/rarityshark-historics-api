package models

import "time"

type TokenTransaction struct {
	TransactionHash string    `bson:"transactionHash" json:"transactionHash"`
	From            string    `bson:"from" json:"from"`
	To              string    `bson:"to" json:"to"`
	BlockNumber     int       `bson:"blockNumber" json:"blockNumber"`
	TokenID         int64     `bson:"tokenId" json:"tokenId"`
	Time            time.Time `bson:"time" json:"time"`
	Price           float32   `bson:"price" json:"price"`
	Marketplace     string    `bson:"marketPlace" json:"marketplace"`
	EtherscanLink   string    `bson:"etherscanLink" json:"etherscanLink"`
}
