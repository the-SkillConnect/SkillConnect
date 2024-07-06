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
					Resolve: resolver.ResolveGetUsersIdentity,
				},
				// New profile queries
				"userProfile": &graphql.Field{
					Type: UserProfileType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetUserProfile,
				},
				// Project queries
				"project": &graphql.Field{
					Type: ProjectType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetProjectByID,
				},
				// Comment queries
				"comment": &graphql.Field{
					Type: graphql.NewList(CommentType),
					Args: graphql.FieldConfigArgument{
						"project_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetCommentByID,
				},
				// AssignProject queries
				"assignedUsersByProject": &graphql.Field{
					Type: graphql.NewList(AssignProjectType),
					Args: graphql.FieldConfigArgument{
						"project_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetAssignedUsersByProjectID,
				},
				"assignedProjectsByUser": &graphql.Field{
					Type: graphql.NewList(AssignProjectType),
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetAssignedProjectsByUserID,
				},
				// Category queries
				"category": &graphql.Field{
					Type: CategoryType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveGetCategory,
				},
				"categories": &graphql.Field{
					Type:    graphql.NewList(CategoryType),
					Resolve: resolver.ResolveGetCategories,
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
				// Project mutations
				"insertProject": &graphql.Field{
					Type: ProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertProjectInputType)},
					},
					Resolve: resolver.ResolveInsertProject,
				},
				"updateProject": &graphql.Field{
					Type: ProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(UpdateProjectInputType)},
					},
					Resolve: resolver.ResolveUpdateProject,
				},
				"deleteProject": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteProject,
				},
				// AssignProject mutations
				"insertAssignProject": &graphql.Field{
					Type: AssignProjectType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertAssignProjectInputType)},
					},
					Resolve: resolver.ResolveInsertAssignProject,
				},
				"deleteAssignProject": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(DeleteAssignProjectInputType)},
					},
					Resolve: resolver.ResolveDeleteAssignProject,
				},
				// Comment mutations
				"insertComment": &graphql.Field{
					Type: CommentType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertCommentInputType)},
					},
					Resolve: resolver.ResolveInsertComment,
				},
				"deleteComment": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteComment,
				},
				"insertCategory": &graphql.Field{
					Type: CategoryType,
					Args: graphql.FieldConfigArgument{
						"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(InsertCategoryInputType)},
					},
					Resolve: resolver.ResolveInsertCategory,
				},
				"deleteCategory": &graphql.Field{
					Type: graphql.Boolean,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					},
					Resolve: resolver.ResolveDeleteCategory,
				},
			},
		}),
	})
}
