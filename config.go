package main

import (
	"os"
)

var (
	AccessToken = os.Getenv("ACCESS_TOKEN")
	Port        = os.Getenv("PORT")
)

func init() {
	if AccessToken == "" {
		AccessToken = ""
	}
	if Port == "" {
		Port = "3000"
	}
}
