package main

import (
	"log"

	"p03_kanban_board/server"

	_ "github.com/joho/godotenv/autoload"
)

// @title Hacktiv8 KanbanBoard API
// @version 1.0
// @description This is API for completing final project 3 hacktiv8
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
