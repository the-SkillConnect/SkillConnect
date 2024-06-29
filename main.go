package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"

	api "github.com/the-SkillConnect/SkillConnect/api/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

func main() {
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	instance := db.New(dbInstance)
	api.SetDbInstance(instance) 

	// Create the schema
	schemaConfig := graphql.SchemaConfig{
		Query:    api.RootQuery,
		Mutation: api.RootMutation,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		var input api.Input
		if err := ctx.QueryParser(&input); err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString("Cannot parse query parameters: " + err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  input.Query,
			OperationName:  input.OperationName,
			VariableValues: input.Variables,
		})

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		var input api.Input
		if err := ctx.BodyParser(&input); err != nil {
			return ctx.
				Status(fiber.StatusInternalServerError).
				SendString("Cannot parse body: " + err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  input.Query,
			OperationName:  input.OperationName,
			VariableValues: input.Variables,
		})

		ctx.Set("Content-Type", "application/graphql-response+json")
		return ctx.JSON(result)
	})

	log.Fatal(app.Listen(":8585"))
}
