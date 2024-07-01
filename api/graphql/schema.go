package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/the-SkillConnect/SkillConnect/db"
)

func NewSchema(dbInstance db.Querier) (graphql.Schema, error) {
	resolver := NewResolver(dbInstance)

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				// Existing user queries
				"userIdentity": &graphql.Field{
					Type: UserIdentityType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetUserByID,
				},
				"usersIdentity": &graphql.Field{
					Type:    graphql.NewList(UserIdentityType),
					Resolve: resolver.ResolveGetUsers,
				},
				// New profile queries
				"userProfile": &graphql.Field{
					Type: UserProfileType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetUserProfile,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				// Existing user mutations
				"insertUserIdentity": &graphql.Field{
					Type: UserIdentityType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertUserIdentityInputType)},
					},
					Resolve: resolver.ResolveInsertUser,
				},
				"deleteUserIdentity": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteUser,
				},
				"updateUserIdentity": &graphql.Field{
					Type: UserIdentityType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateUserIdentityInputType)},
					},
					Resolve: resolver.ResolveUpdateUser,
				},
				// New profile mutations
				"insertUserProfile": &graphql.Field{
					Type: UserProfileType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertUserProfileInputType)},
					},
					Resolve: resolver.ResolveInsertUserProfile,
				},
				"updateUserProfile": &graphql.Field{
					Type: UserProfileType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateUserProfileInputType)},
					},
					Resolve: resolver.ResolveUpdateUserProfile,
				},
				"deleteUserProfile": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteUserProfile,
				},
			},
		}),
	})
}
