package iblclient

import (
	"encoding/json"
)

type Programme struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Title  string `json:"title"`
	Number int
}

type Channel struct {
	Id            string `json:"id"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	HasSchedule   bool   `json:"has_schedule"`
	MasterBrandId string `json:"master_brand_id"`
}

type ResponseWrapper struct {
	Programmes []Programme `json:"elements"`
	Channel    Channel     `json:"channel"`
	Count      int         `json:"count"`
}

type Response struct {
	Schema  string          `json:"schema"`
	Version string          `json:"version"`
	Data    ResponseWrapper `json:"channel_programmes"`
}

func ParseResponse(in string) Response {
	out := Response{}
	json.Unmarshal([]byte(in), &out)
	return out
}
