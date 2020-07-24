package main

import (
     "os"
     "strconv"
)

func main() {
     a := App{}
     portString := os.Getenv("APP_DB_PORT")
     port, _ := strconv.Atoi(portString)
     a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"),
		port)

     a.Run(":8010")
}