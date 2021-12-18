package main

import (
	"mvc/config"
	"mvc/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
