package models

import "time"

type Historic struct {
	FloorPrice FloorPrice `bson:"floorPrice" json:"floorPrice"`
	Listings   Listing    `bson:"listings" json:"listings"`
	Stats      Stat       `bson:"stats" json:"stats"`
	Date       time.Time  `bson:"date" json:"date"`
}

type FloorPrice struct {
	Price float64 `bson:"price" json:"price"`
}

type Listing struct {
	Listings int `bson:"listings" json:"listings"`
}

type Stat struct {
	OneDayVolume          float64 `bson:"one_day_volume" json:"oneDayVolume"`
	OneDayChange          float64 `bson:"one_day_change" json:"oneDayChange"`
	OneDaySales           float64 `bson:"one_day_sales" json:"oneDaySales"`
	OneDayAveragePrice    float64 `bson:"one_day_average_price" json:"oneDayAveragePrice"`
	SevenDayVolume        float64 `bson:"seven_day_volume" json:"sevenDayVolume"`
	SevenDayChange        float64 `bson:"seven_day_change" json:"sevenDayChange"`
	SevenDaySales         float64 `bson:"seven_day_sales" json:"sevenDaySales"`
	SevenDayAveragePrice  float64 `bson:"seven_day_average_price" json:"sevenDayAveragePrice"`
	ThirtyDayVolume       float64 `bson:"thirty_day_volume" json:"thirtyDayVolume"`
	ThirtyDayChange       float64 `bson:"thirty_day_change" json:"thirtyDayChange"`
	ThirtyDaySales        float64 `bson:"thirty_day_sales" json:"thirtyDaySales"`
	ThirtyDayAveragePrice float64 `bson:"thirty_day_average_price" json:"thirtyDayAveragePrice"`
	TotalVolume           float64 `bson:"total_volume" json:"totalVolume"`
	TotalSales            float64 `bson:"total_sales" json:"totalSales"`
	TotalSupply           float64 `bson:"total_supply" json:"totalSupply"`
	Count                 float64 `bson:"count" json:"count"`
	NumOwners             int     `bson:"num_owners" json:"numOwners"`
	AveragePrice          float64 `bson:"average_price" json:"averagePrice"`
	NumReports            int     `bson:"num_reports" json:"numReports"`
	MarketCap             float64 `bson:"market_cap" json:"marketCap"`
	FloorPrice            float64 `bson:"floor_price" json:"floorPrice"`
}
