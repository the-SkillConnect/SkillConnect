package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	gql "github.com/the-SkillConnect/SkillConnect/api/graphql"
)

// GraphQLHandler handles the GraphQL queries and mutations
func GraphQLHandler(schema graphql.Schema) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var input gql.Input
		if err := ctx.BodyParser(&input); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Cannot parse request body: " + err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  input.Query,
			OperationName:  input.OperationName,
			VariableValues: input.Variables,
		})

		if len(result.Errors) > 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(result.Errors)
		}

		return ctx.JSON(result)
	}
}
