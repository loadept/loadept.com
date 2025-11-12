package main

import (
	"log"
	"os"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/infrastructure/database"
)

func main() {
	config.LoadEnviron()

	con, err := database.NewConnection()
	if err != nil {
		log.Printf("\033[31mError to create database connection: %v\033[0m\n", err)
		os.Exit(1)
	}
	defer con.Close()

	m, err := database.NewMigration(con.GetDB())
	if err != nil {
		log.Printf("\033[31mError to create new migration: %v\033[0m", err)
		os.Exit(1)
	}

	if err := m.RunMigrations(); err != nil {
		log.Printf("\033[31mError to run migrations: %v\033[0m\n", err)
		os.Exit(1)
	}
	log.Println("\033[32mAll migrations already\033[0m")
}
