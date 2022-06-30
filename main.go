package main

import (
	database "lab-jarkom-be/Database"
)

func main() {
	db := database.Open()
	if db != nil {
		println("nice")
	}
}
