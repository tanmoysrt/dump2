package main

type LogRequest struct {
	ID        string `json:"id" query:"id"`
	Key       string `json:"key" query:"key"`
	Lon       string `json:"lon" query:"lon"`
	Lat       string `json:"lat" query:"lat"`
	Timestamp string `json:"timestamp" query:"timestamp"`
}
