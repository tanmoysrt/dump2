package main

type LogRequest struct {
	Event     string  `json:"event" query:"event"`
	ID        string  `json:"id" query:"id"`
	Key       string  `json:"key" query:"key"`
	Lon       float64 `json:"lon" query:"lon"`
	Lat       float64 `json:"lat" query:"lat"`
	Timestamp int64   `json:"timestamp" query:"timestamp"`
}

type IgnitionLog struct {
	Event     string `json:"event" query:"event"`
	ID        string `json:"id" query:"id"`
	Key       string `json:"key" query:"key"`
	Enabled   bool   `json:"enabled"`
	Timestamp int64  `json:"timestamp"`
}
