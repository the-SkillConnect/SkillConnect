package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/the-SkillConnect/SkillConnect/api/graphql"
	"github.com/the-SkillConnect/SkillConnect/api/handlers"
	"github.com/the-SkillConnect/SkillConnect/db"
)

func main() {

	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.CreateTables(dbInstance); err != nil {
		log.Fatal(err)
	}
	instance := db.New(dbInstance)

	schema, err := graphql.NewSchema(instance)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	app := fiber.New()
	app.Post("/graphql", handlers.GraphQLHandler(schema))

	log.Fatal(app.Listen(os.Getenv("HTTP_LISTEN_ADDRESS")))
}
