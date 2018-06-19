package main

type Player struct {
	ID int `json:"id" bson:"_id"`
	Points int `json:"points" bson:"points"`
}

type Tournament struct {
	ID int `json:"id" bson:"_id"`
	Deposit int `json:"deposit" bson:"deposit"`
}
