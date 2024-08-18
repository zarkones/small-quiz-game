package config

import "os"

// Environment variables.s
var (
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
)
